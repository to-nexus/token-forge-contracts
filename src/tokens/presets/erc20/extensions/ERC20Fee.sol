// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC20Base} from "../ERC20Base.sol";

abstract contract ERC20Fee is ERC20Base {
    error ERC20Fee__InvalidFeeBPS(uint256 feeBPS);

    event FeeRecipientUpdated(address indexed oldRecipient, address indexed newRecipient);
    event FeeBPSUpdated(uint256 oldBPS, uint256 newBPS);
    event MintingFeeCollected(address indexed recipient, uint256 amount);

    /// @custom:storage-location erc7201:cross.storage.forge.erc20.fee
    struct ERC20FeeStorage {
        address feeRecipient;
        uint96 feeBPS; // BPS(1 / 10000) feeBPS < 10000
    }

    // keccak256(abi.encode(uint256(keccak256("cross.storage.forge.erc20.fee")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant ERC20FeeStorageLocation =
        0x9e07bc4d06789dd0b871d60d502656ec22e58628182c2156c52dcfc63440ac00;

    function _getERC20FeeStorage() private pure returns (ERC20FeeStorage storage $) {
        assembly {
            $.slot := ERC20FeeStorageLocation
        }
    }

    function __ERC20Fee_init(address feeRecipient, uint256 feeBPS) internal onlyInitializing {
        if (feeRecipient == address(0)) revert ERC20Base__NullInput("feeRecipient");
        if (feeBPS > 10000) revert ERC20Fee__InvalidFeeBPS(feeBPS);

        ERC20FeeStorage storage $ = _getERC20FeeStorage();
        $.feeRecipient = feeRecipient;
        $.feeBPS = uint96(feeBPS);

        emit FeeRecipientUpdated(address(0), feeRecipient);
        emit FeeBPSUpdated(0, feeBPS);
    }

    function feeBPS() public view returns (uint256) {
        return _getERC20FeeStorage().feeBPS;
    }

    function feeRecipient() public view returns (address) {
        return _getERC20FeeStorage().feeRecipient;
    }

    function setFeeBPS(uint256 _feeBPS) external onlyOwner {
        if (_feeBPS > 10000) revert ERC20Fee__InvalidFeeBPS(_feeBPS);

        ERC20FeeStorage storage $ = _getERC20FeeStorage();
        emit FeeBPSUpdated($.feeBPS, _feeBPS);
        $.feeBPS = uint96(_feeBPS);
    }

    function setFeeRecipient(address _feeRecipient) external onlyOwner {
        if (_feeRecipient == address(0)) revert ERC20Base__NullInput("feeRecipient");

        ERC20FeeStorage storage $ = _getERC20FeeStorage();
        emit FeeRecipientUpdated($.feeRecipient, _feeRecipient);
        $.feeRecipient = _feeRecipient;
    }
}
