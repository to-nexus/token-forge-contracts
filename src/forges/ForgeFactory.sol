// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {IDiamondCut} from "diamond-3-hardhat-1.0.0/interfaces/IDiamondCut.sol";
import {Diamond} from "diamond-3-hardhat-1.0.0/Diamond.sol";
import {AccessControlUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/access/AccessControlUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/proxy/utils/UUPSUpgradeable.sol";

import {ECDSA} from "@openzeppelin-contracts-5.3.0/utils/cryptography/ECDSA.sol";
import {EnumerableMap} from "@openzeppelin-contracts-5.3.0/utils/structs/EnumerableMap.sol";
import {Create2} from "@openzeppelin-contracts-5.3.0/utils/Create2.sol";

import {TokenType, IForgeFactoryAlert} from "../interfaces/IForgeFactory.sol";
import {IDefaultDiamondCut} from "../interfaces/IDefaultDiamondCut.sol";
import {ForgeProxyCode} from "./ForgeProxy.sol";

contract ForgeFactory is IForgeFactoryAlert, AccessControlUpgradeable, UUPSUpgradeable {
    using ECDSA for bytes32;
    using EnumerableMap for EnumerableMap.Bytes32ToAddressMap;

    error TokenForgeFactory__InvalidData(bytes32 field);
    error TokenForgeFactory__AlreadyUsedService(bytes32 service);
    error TokenForgeFactory__CallerIsNotForge(address caller);
    error TokenForgeFactory__CallerIsPausedForge(address caller);
    error TokenForgeFactory__ServiceNotFound(bytes32 service);
    error TokenForge__InvalidTokenType();

    event ERC20Minted(bytes32 indexed service, uint256 indexed uuid, address indexed to, address token, uint256 amount);
    event ERC721Minted(
        bytes32 indexed service, uint256 indexed uuid, address indexed to, address token, uint256 tokenID
    );
    event ERC1155Minted(
        bytes32 indexed service,
        uint256 indexed uuid,
        address indexed to,
        address token,
        uint256 tokenID,
        uint256 amount
    );

    event ERC20Transferred(
        bytes32 indexed service, uint256 indexed uuid, address indexed to, address token, uint256 amount
    );
    event ERC20TransferredFrom(
        bytes32 indexed service, uint256 indexed uuid, address indexed from, address token, uint256 amount
    );
    event ERC721Transferred(
        bytes32 indexed service, uint256 indexed uuid, address indexed to, address token, uint256 tokenID
    );
    event ERC1155Transferred(
        bytes32 indexed service,
        uint256 indexed uuid,
        address indexed to,
        address token,
        uint256 tokenID,
        uint256 amount
    );

    event ERC20Burned(
        bytes32 indexed service, uint256 indexed uuid, address indexed from, address token, uint256 amount
    );
    event ERC721Burned(
        bytes32 indexed service, uint256 indexed uuid, address indexed from, address token, uint256 tokenID
    );
    event ERC1155Burned(
        bytes32 indexed service,
        uint256 indexed uuid,
        address indexed from,
        address token,
        uint256 tokenID,
        uint256 amount
    );

    event ERC20FeeCollected(uint256 indexed uuid, address indexed feeRecipient, address indexed token, uint256 fee);

    event ServiceRegistered(bytes32 indexed service, address indexed forge);
    event ServiceUnregistered(bytes32 indexed service);
    event ServicePaused(bytes32 indexed service, bool paused);

    bytes32 private constant MANAGER_ROLE = keccak256("MANAGER");

    /// @custom:storage-location erc7201:cross.storage.forge.ForgeFactory
    struct ForgeFactoryStorage {
        ForgeProxyCode _forgeProxyCode;
        address _diamondImpl;
        address _baseImpl;
        EnumerableMap.Bytes32ToAddressMap _serviceToForge;
        mapping(address forge => bytes32) _forgeToService;
        mapping(bytes32 service => bool) _isRunning;
    }

    // keccak256(abi.encode(uint256(keccak256("cross.storage.forge.ForgeFactory")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant FORGE_FACTORY_STORAGE_LOCATION =
        0x11d8116c90363a5c9c6e1d5a32c41784fb521cabceaeed518cf2f20332b24e00;

    function _getForgeFactoryStorage() private pure returns (ForgeFactoryStorage storage $) {
        assembly {
            $.slot := FORGE_FACTORY_STORAGE_LOCATION
        }
    }

    function initialize(address _owner, address _forgeProxyCode, address _diamondImpl, address _baseImpl)
        external
        initializer
    {
        __AccessControl_init();
        __UUPSUpgradeable_init();
        __ForgeFactory_init(_owner, ForgeProxyCode(_forgeProxyCode), _diamondImpl, _baseImpl);
    }

    function __ForgeFactory_init(
        address _owner,
        ForgeProxyCode _forgeProxyCode,
        address _diamondImpl,
        address _baseImpl
    ) private onlyInitializing {
        // check initial data
        if (_owner == address(0)) revert TokenForgeFactory__InvalidData("owner");
        if (address(_forgeProxyCode) == address(0)) revert TokenForgeFactory__InvalidData("forgeProxyCode");
        if (_diamondImpl == address(0)) revert TokenForgeFactory__InvalidData("diamondImpl");
        if (_baseImpl == address(0)) revert TokenForgeFactory__InvalidData("baseImpl");

        // check contracts
        if (_forgeProxyCode.code().length == 0) {
            revert TokenForgeFactory__InvalidData("forgeProxyCode code");
        }
        if (_diamondImpl != IDefaultDiamondCut(_diamondImpl).defaultDiamondFacetCut().facetAddress) {
            revert TokenForgeFactory__InvalidData("diamondImpl facet");
        }
        if (_baseImpl != IDefaultDiamondCut(_baseImpl).defaultDiamondFacetCut().facetAddress) {
            revert TokenForgeFactory__InvalidData("baseImpl facet");
        }

        ForgeFactoryStorage storage $ = _getForgeFactoryStorage();
        $._forgeProxyCode = _forgeProxyCode;
        $._diamondImpl = _diamondImpl;
        $._baseImpl = _baseImpl;

        _setRoleAdmin(MANAGER_ROLE, DEFAULT_ADMIN_ROLE);
        _grantRole(DEFAULT_ADMIN_ROLE, _owner);
        _grantRole(MANAGER_ROLE, _owner);
    }

    ////////////////////
    // View functions //
    ////////////////////

    function allForges() external view returns (bytes32[] memory, address[] memory) {
        ForgeFactoryStorage storage $ = _getForgeFactoryStorage();

        uint256 length = $._serviceToForge.length();
        bytes32[] memory services = new bytes32[](length);
        address[] memory forges = new address[](length);
        unchecked {
            for (uint256 i = 0; i < length; ++i) {
                (services[i], forges[i]) = $._serviceToForge.at(i);
            }
        }
        return (services, forges);
    }

    function isRunningForge(address forge) external view returns (bytes32 service, bool running) {
        ForgeFactoryStorage storage $ = _getForgeFactoryStorage();
        service = $._forgeToService[forge];
        if (service != bytes32(0)) {
            running = $._isRunning[service];
        }
    }

    function forgeByService(bytes32 service) external view returns (address forge, bool running) {
        ForgeFactoryStorage storage $ = _getForgeFactoryStorage();
        bool ok;
        (ok, forge) = $._serviceToForge.tryGet(service);
        if (ok) running = $._isRunning[service];
    }

    function lengthAllForges() external view returns (uint256) {
        return _getForgeFactoryStorage()._serviceToForge.length();
    }

    function forgeByIndex(uint256 index) external view returns (bytes32 service, address forge, bool running) {
        ForgeFactoryStorage storage $ = _getForgeFactoryStorage();
        (service, forge) = $._serviceToForge.at(index);
        running = $._isRunning[service];
    }

    /////////////////////
    // Alert functions //
    /////////////////////

    function alertMint(TokenType tokenType, uint256 uuid, address token, bytes calldata data) external {
        bytes32 service = _checkRunningService(_getForgeFactoryStorage(), msg.sender);
        if (tokenType == TokenType.ERC20) {
            (address to, uint256 amount, address feeRecipient, uint256 fee) =
                abi.decode(data, (address, uint256, address, uint256));
            emit ERC20Minted(service, uuid, to, token, amount);
            if (fee != 0) {
                emit ERC20FeeCollected(uuid, feeRecipient, token, fee);
            }
        } else if (tokenType == TokenType.ERC721) {
            (address to, uint256 tokenID) = abi.decode(data, (address, uint256));
            emit ERC721Minted(service, uuid, to, token, tokenID);
        } else if (tokenType == TokenType.ERC1155) {
            (bool isBatch, bytes memory data1155) = abi.decode(data, (bool, bytes));
            if (isBatch) {
                unchecked {
                    (address to, uint256[] memory tokenIDs, uint256[] memory amounts) =
                        abi.decode(data1155, (address, uint256[], uint256[]));
                    uint256 length = tokenIDs.length;
                    for (uint256 i = 0; i < length; ++i) {
                        emit ERC1155Minted(service, uuid, to, token, tokenIDs[i], amounts[i]);
                    }
                }
            } else {
                (address to, uint256 tokenID, uint256 amount) = abi.decode(data1155, (address, uint256, uint256));
                emit ERC1155Minted(service, uuid, to, token, tokenID, amount);
            }
        } else {
            revert TokenForge__InvalidTokenType();
        }
    }

    function alertTransfer(TokenType tokenType, uint256 uuid, address token, bytes calldata data) external {
        bytes32 service = _checkRunningService(_getForgeFactoryStorage(), msg.sender);
        if (tokenType == TokenType.ERC20) {
            (address to, uint256 amount, address feeRecipient, uint256 fee) =
                abi.decode(data, (address, uint256, address, uint256));
            emit ERC20Transferred(service, uuid, to, token, amount);
            if (fee != 0) {
                emit ERC20FeeCollected(uuid, feeRecipient, token, fee);
            }
        } else if (tokenType == TokenType.ERC721) {
            (address to, uint256 tokenID) = abi.decode(data, (address, uint256));
            emit ERC721Transferred(service, uuid, to, token, tokenID);
        } else if (tokenType == TokenType.ERC1155) {
            (bool isBatch, bytes memory data1155) = abi.decode(data, (bool, bytes));
            if (isBatch) {
                unchecked {
                    (address to, uint256[] memory tokenIDs, uint256[] memory amounts) =
                        abi.decode(data1155, (address, uint256[], uint256[]));
                    uint256 length = tokenIDs.length;
                    for (uint256 i = 0; i < length; ++i) {
                        emit ERC1155Transferred(service, uuid, to, token, tokenIDs[i], amounts[i]);
                    }
                }
            } else {
                (address to, uint256 tokenID, uint256 amount) = abi.decode(data1155, (address, uint256, uint256));
                emit ERC1155Transferred(service, uuid, to, token, tokenID, amount);
            }
        } else {
            revert TokenForge__InvalidTokenType();
        }
    }

    function alertTransferFrom(TokenType tokenType, uint256 uuid, address token, bytes calldata data) external {
        bytes32 service = _checkRunningService(_getForgeFactoryStorage(), msg.sender);
        if (tokenType == TokenType.ERC20) {
            (address from, uint256 amount, address feeRecipient, uint256 fee) =
                abi.decode(data, (address, uint256, address, uint256));
            emit ERC20TransferredFrom(service, uuid, from, token, amount);
            if (fee != 0) {
                emit ERC20FeeCollected(uuid, feeRecipient, token, fee);
            }
        } else {
            revert TokenForge__InvalidTokenType();
        }
    }

    function alertBurn(TokenType tokenType, uint256 uuid, address token, bytes calldata data) external {
        bytes32 service = _checkRunningService(_getForgeFactoryStorage(), msg.sender);
        if (tokenType == TokenType.ERC20) {
            (address from, uint256 amount, address feeRecipient, uint256 fee) =
                abi.decode(data, (address, uint256, address, uint256));
            emit ERC20Burned(service, uuid, from, token, amount);
            if (fee != 0) {
                emit ERC20FeeCollected(uuid, feeRecipient, token, fee);
            }
        } else if (tokenType == TokenType.ERC721) {
            (address from, uint256 tokenID) = abi.decode(data, (address, uint256));
            emit ERC721Burned(service, uuid, from, token, tokenID);
        } else if (tokenType == TokenType.ERC1155) {
            (bool isBatch, bytes memory data1155) = abi.decode(data, (bool, bytes));
            if (isBatch) {
                unchecked {
                    (address from, uint256[] memory tokenIDs, uint256[] memory amounts) =
                        abi.decode(data1155, (address, uint256[], uint256[]));
                    uint256 length = tokenIDs.length;
                    for (uint256 i = 0; i < length; ++i) {
                        emit ERC1155Burned(service, uuid, from, token, tokenIDs[i], amounts[i]);
                    }
                }
            } else {
                (address from, uint256 tokenID, uint256 amount) = abi.decode(data1155, (address, uint256, uint256));
                emit ERC1155Burned(service, uuid, from, token, tokenID, amount);
            }
        } else {
            revert TokenForge__InvalidTokenType();
        }
    }

    function _checkRunningService(ForgeFactoryStorage storage $, address forge) private view returns (bytes32) {
        bytes32 service = $._forgeToService[forge];
        if (service == bytes32(0)) {
            revert TokenForgeFactory__CallerIsNotForge(forge);
        }
        if (!$._isRunning[service]) {
            revert TokenForgeFactory__CallerIsPausedForge(forge);
        }
        return service;
    }

    ////////////////////
    // Role functions //
    ////////////////////

    function addService(address owner, address validator, bytes32 service, IDiamondCut.FacetCut[] calldata addCuts)
        external
        onlyRole(MANAGER_ROLE)
        returns (address forge)
    {
        if (owner == address(0)) revert TokenForgeFactory__InvalidData("owner");
        if (validator == address(0)) revert TokenForgeFactory__InvalidData("validator");
        if (service == bytes32(0)) revert TokenForgeFactory__InvalidData("service");

        ForgeFactoryStorage storage $ = _getForgeFactoryStorage();
        if ($._serviceToForge.contains(service)) revert TokenForgeFactory__AlreadyUsedService(service);

        // deploy forge proxy
        forge = Create2.deploy(
            0,
            service, // salt is service name
            abi.encodePacked(
                $._forgeProxyCode.code(), abi.encode(owner, validator, addCuts, service, $._diamondImpl, $._baseImpl)
            )
        );
        if (forge == address(0)) revert TokenForgeFactory__InvalidData("forge deploy");

        // enroll service&forge to storage
        $._serviceToForge.set(service, forge);
        $._forgeToService[forge] = service;
        $._isRunning[service] = true;
        emit ServiceRegistered(service, forge);
    }

    function removeService(bytes32 service) external onlyRole(MANAGER_ROLE) {
        ForgeFactoryStorage storage $ = _getForgeFactoryStorage();

        // check if service exists
        (bool ok, address forge) = $._serviceToForge.tryGet(service);
        if (!ok) revert TokenForgeFactory__ServiceNotFound(service);

        // remove data
        $._serviceToForge.remove(service);
        delete $._forgeToService[forge];
        delete $._isRunning[service];
        emit ServiceUnregistered(service);
    }

    function pauseService(bytes32 service, bool paused) external onlyRole(MANAGER_ROLE) {
        ForgeFactoryStorage storage $ = _getForgeFactoryStorage();
        if (!$._serviceToForge.contains(service)) revert TokenForgeFactory__ServiceNotFound(service);
        if ($._isRunning[service] == paused) {
            if (paused) delete $._isRunning[service];
            else $._isRunning[service] = true;

            emit ServicePaused(service, paused);
        }
    }

    function _authorizeUpgrade(address) internal override onlyRole(DEFAULT_ADMIN_ROLE) {}
}
