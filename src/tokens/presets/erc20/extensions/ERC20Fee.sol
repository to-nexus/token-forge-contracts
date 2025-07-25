// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC20Base} from "../ERC20Base.sol";
import {EnumerableMap} from "@openzeppelin/contracts/utils/structs/EnumerableMap.sol";

abstract contract ERC20Fee is ERC20Base {
    using EnumerableMap for EnumerableMap.AddressToUintMap;

    error ERC20Fee__InvalidFeeBPS(uint256 feeBPS);

    event FeeRecipientUpdated(address indexed oldRecipient, address indexed newRecipient);
    event FeeBPSUpdated(uint256 oldBPS, uint256 newBPS);
    event ForgeFeeBPSUpdated(address indexed forge, uint256 oldBPS, uint256 newBPS);
    event ForgeFeeBPSRemoved(address indexed forge, uint256 oldBPS);
    event FeeCollected(address indexed caller, address indexed recipient, uint256 amount);

    /// @custom:storage-location erc7201:cross.storage.forge.erc20.fee
    struct ERC20FeeStorage {
        address feeRecipient;
        uint96 feeBPS; // BPS(1 / 10000) feeBPS < 10000 - 전역 기본 수수료율 (fallback)
        EnumerableMap.AddressToUintMap forgeFeeBPS; // forge별 개별 수수료율
    }

    // keccak256(abi.encode(uint256(keccak256("cross.storage.forge.erc20.fee")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant ERC20FeeStorageLocation =
        0xd0b546179495f4880d81845138f4b5012ee24fd5441f32d121fce6837dc28d00;

    function _getERC20FeeStorage() private pure returns (ERC20FeeStorage storage $) {
        assembly {
            $.slot := ERC20FeeStorageLocation
        }
    }

    function __ERC20Fee_init(address _feeRecipient, uint256 _feeBPS) internal onlyInitializing {
        if (_feeRecipient == address(0)) revert TokenBase__NullInput("feeRecipient");
        if (_feeBPS > 10000) revert ERC20Fee__InvalidFeeBPS(_feeBPS);

        ERC20FeeStorage storage $ = _getERC20FeeStorage();
        $.feeRecipient = _feeRecipient;
        $.feeBPS = uint96(_feeBPS);

        emit FeeRecipientUpdated(address(0), _feeRecipient);
        emit FeeBPSUpdated(0, _feeBPS);
    }

    function feeBPS(address forge) public view returns (uint256) {
        ERC20FeeStorage storage $ = _getERC20FeeStorage();
        (bool exists, uint256 forgeFee) = $.forgeFeeBPS.tryGet(forge);
        return exists ? forgeFee : $.feeBPS;
    }

    function feeRecipient() public view returns (address) {
        return _getERC20FeeStorage().feeRecipient;
    }

    function getAllForgeFeeBPS() external view returns (address[] memory forges, uint256[] memory fees) {
        EnumerableMap.AddressToUintMap storage forgeFeeBPS = _getERC20FeeStorage().forgeFeeBPS;
        uint256 length = forgeFeeBPS.length();

        forges = new address[](length);
        fees = new uint256[](length);

        for (uint256 i = 0; i < length; i++) {
            (forges[i], fees[i]) = forgeFeeBPS.at(i);
        }
    }

    function setFeeBPS(uint256 _feeBPS) external onlyOwner {
        if (_feeBPS > 10000) revert ERC20Fee__InvalidFeeBPS(_feeBPS);

        ERC20FeeStorage storage $ = _getERC20FeeStorage();
        emit FeeBPSUpdated($.feeBPS, _feeBPS);
        $.feeBPS = uint96(_feeBPS);
    }

    function setFeeRecipient(address _feeRecipient) external onlyOwner {
        if (_feeRecipient == address(0)) revert TokenBase__NullInput("feeRecipient");

        ERC20FeeStorage storage $ = _getERC20FeeStorage();
        emit FeeRecipientUpdated($.feeRecipient, _feeRecipient);
        $.feeRecipient = _feeRecipient;
    }

    function setForgeFeeBPS(address forge, uint256 _feeBPS) external onlyOwner {
        if (forge == address(0)) revert TokenBase__NullInput("forge");
        if (_feeBPS > 10000) revert ERC20Fee__InvalidFeeBPS(_feeBPS);

        EnumerableMap.AddressToUintMap storage forgeFeeBPS = _getERC20FeeStorage().forgeFeeBPS;
        (bool exists, uint256 oldBPS) = forgeFeeBPS.tryGet(forge);
        uint256 oldFee = exists ? oldBPS : 0;

        forgeFeeBPS.set(forge, _feeBPS);
        emit ForgeFeeBPSUpdated(forge, oldFee, _feeBPS);
    }

    function removeForgeFeeBPS(address forge) external onlyOwner {
        EnumerableMap.AddressToUintMap storage forgeFeeBPS = _getERC20FeeStorage().forgeFeeBPS;
        (bool exists, uint256 oldBPS) = forgeFeeBPS.tryGet(forge);

        if (exists) {
            forgeFeeBPS.remove(forge);
            emit ForgeFeeBPSRemoved(forge, oldBPS);
        }
    }
}
