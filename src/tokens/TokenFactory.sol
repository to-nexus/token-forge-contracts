// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {AccessControlUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/access/AccessControlUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/proxy/utils/UUPSUpgradeable.sol";
import {EnumerableSet} from "@openzeppelin-contracts-5.3.0/utils/structs/EnumerableSet.sol";
import {Create2} from "@openzeppelin-contracts-5.3.0/utils/Create2.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts-5.3.0/proxy/ERC1967/ERC1967Proxy.sol";

import {IERC165} from "@openzeppelin-contracts-5.3.0/interfaces/IERC165.sol";
import {IERC20Forge} from "../interfaces/IERC20Forge.sol";
import {IERC721Forge} from "../interfaces/IERC721Forge.sol";
import {IERC1155Forge} from "../interfaces/IERC1155Forge.sol";

import {ITokenFactory} from "../interfaces/ITokenFactory.sol";
import {IForgeFactoryForgeByService} from "../interfaces/IForgeFactory.sol";

contract TokenFactory is ITokenFactory, AccessControlUpgradeable, UUPSUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;

    error TokenFactory__ZeroAddress(bytes32 field);
    error TokenFactory__InvalidLogic(TokenType, address);
    error TokenFactory__DeployFailed(TokenType, address);

    bytes32 public constant MANAGER_ROLE = keccak256("MANAGER");

    /// @custom:storage-location erc7201:cross.storage.forge.TokenFactory
    struct TokenFactoryStorage {
        EnumerableSet.AddressSet erc20Impls;
        EnumerableSet.AddressSet erc721Impls;
        EnumerableSet.AddressSet erc1155Impls;
    }

    // keccak256(abi.encode(uint256(keccak256("cross.storage.forge.TokenFactory")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant TOKEN_FACTORY_STORAGE_LOCATION =
        0xd9ce7ba49891d0b3422c928b675902668a780038fd864217ee35722e549a9f00;

    function _getTokenFactoryStorage() private pure returns (TokenFactoryStorage storage $) {
        assembly {
            $.slot := TOKEN_FACTORY_STORAGE_LOCATION
        }
    }

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(
        address owner,
        address[] calldata erc20Impls,
        address[] calldata erc721Impls,
        address[] calldata erc1155Impls
    ) external initializer {
        __AccessControl_init();
        __UUPSUpgradeable_init();
        __TokenFactory_init(owner, erc20Impls, erc721Impls, erc1155Impls);
    }

    function __TokenFactory_init(
        address owner,
        address[] calldata erc20Impls,
        address[] calldata erc721Impls,
        address[] calldata erc1155Impls
    ) private onlyInitializing {
        if (owner == address(0)) revert TokenFactory__ZeroAddress("owner");

        _grantRole(MANAGER_ROLE, owner);
        _grantRole(DEFAULT_ADMIN_ROLE, owner);

        TokenFactoryStorage storage $ = _getTokenFactoryStorage();
        _setImpls(TokenType.ERC20, type(IERC20Forge).interfaceId, $.erc20Impls, erc20Impls, true);
        _setImpls(TokenType.ERC721, type(IERC721Forge).interfaceId, $.erc721Impls, erc721Impls, true);
        _setImpls(TokenType.ERC1155, type(IERC1155Forge).interfaceId, $.erc1155Impls, erc1155Impls, true);
    }

    function getPresetLogics(TokenType tokenType) external view returns (address[] memory) {
        TokenFactoryStorage storage $ = _getTokenFactoryStorage();
        if (tokenType == TokenType.ERC20) {
            return $.erc20Impls.values();
        } else if (tokenType == TokenType.ERC721) {
            return $.erc721Impls.values();
        } else {
            return $.erc1155Impls.values();
        }
    }

    function deployERC20(
        address owner,
        address manager,
        string memory name,
        string memory symbol,
        uint8 decimals,
        uint256 initialSupply,
        address initialRecipient,
        bytes memory data,
        address logic
    ) external onlyRole(MANAGER_ROLE) returns (address token) {
        TokenFactoryStorage storage $ = _getTokenFactoryStorage();
        {
            if (!$.erc20Impls.contains(logic)) {
                revert TokenFactory__InvalidLogic(TokenType.ERC20, logic);
            }
            token = address(
                new ERC1967Proxy(
                    logic,
                    abi.encodeCall(
                        IERC20Forge.initialize,
                        (owner, manager, name, symbol, decimals, initialSupply, initialRecipient, data)
                    )
                )
            );
        }
        if (token == address(0)) revert TokenFactory__DeployFailed(TokenType.ERC20, logic);

        emit TokenDeployed(owner, TokenType.ERC20, token, logic);
    }

    function deployERC721(
        address owner,
        address manager,
        string memory name,
        string memory symbol,
        string memory baseTokenURI,
        bytes memory data,
        address logic
    ) external onlyRole(MANAGER_ROLE) returns (address token) {
        TokenFactoryStorage storage $ = _getTokenFactoryStorage();
        {
            if (!$.erc721Impls.contains(logic)) {
                revert TokenFactory__InvalidLogic(TokenType.ERC721, logic);
            }
            token = address(
                new ERC1967Proxy(
                    logic, abi.encodeCall(IERC721Forge.initialize, (owner, manager, name, symbol, baseTokenURI, data))
                )
            );
        }
        if (token == address(0)) revert TokenFactory__DeployFailed(TokenType.ERC721, logic);

        emit TokenDeployed(owner, TokenType.ERC721, token, logic);
    }

    function deployERC1155(address owner, address manager, string memory uri, bytes memory data, address logic)
        external
        onlyRole(MANAGER_ROLE)
        returns (address token)
    {
        TokenFactoryStorage storage $ = _getTokenFactoryStorage();
        {
            if (!$.erc1155Impls.contains(logic)) {
                revert TokenFactory__InvalidLogic(TokenType.ERC1155, logic);
            }
            token =
                address(new ERC1967Proxy(logic, abi.encodeCall(IERC1155Forge.initialize, (owner, manager, uri, data))));
        }
        if (token == address(0)) revert TokenFactory__DeployFailed(TokenType.ERC1155, logic);

        emit TokenDeployed(owner, TokenType.ERC1155, token, logic);
    }

    function setPresetLogics(TokenType tokenType, address[] calldata logicAddress, bool add)
        external
        onlyRole(MANAGER_ROLE)
    {
        EnumerableSet.AddressSet storage _impls;
        bytes4 interfaceId;
        if (tokenType == TokenType.ERC20) {
            (_impls, interfaceId) = (_getTokenFactoryStorage().erc20Impls, type(IERC20Forge).interfaceId);
        } else if (tokenType == TokenType.ERC721) {
            (_impls, interfaceId) = (_getTokenFactoryStorage().erc721Impls, type(IERC721Forge).interfaceId);
        } else {
            (_impls, interfaceId) = (_getTokenFactoryStorage().erc1155Impls, type(IERC1155Forge).interfaceId);
        }
        _setImpls(tokenType, interfaceId, _impls, logicAddress, add);
    }

    function _setImpls(
        TokenType tokeType,
        bytes4 interfaceId,
        EnumerableSet.AddressSet storage _impls,
        address[] calldata impls,
        bool add
    ) private {
        uint256 length = impls.length;
        unchecked {
            if (add) {
                for (uint256 i = 0; i < length; ++i) {
                    address impl = impls[i];
                    if (impl == address(0)) {
                        revert TokenFactory__ZeroAddress("impls");
                    }
                    if (!IERC165(impl).supportsInterface(interfaceId)) {
                        revert TokenFactory__InvalidLogic(tokeType, impl);
                    }
                    if (_impls.add(impl)) {
                        emit PresetLogicSet(tokeType, impl);
                    }
                }
            } else {
                for (uint256 i = 0; i < length; ++i) {
                    address impl = impls[i];
                    if (_impls.remove(impl)) {
                        emit PresetLogicRemoved(tokeType, impl);
                    }
                }
            }
        }
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE) {}
}
