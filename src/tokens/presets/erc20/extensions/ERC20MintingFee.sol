// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC20Base} from "../ERC20Base.sol";

abstract contract ERC20MintingFee is ERC20Base {
    error ERC20MintingFee__InvalidFeeBPS(uint256 feeBPS);
    error ERC20MintingFee__InvalidFeeRecipient(address recipient);

    event FeeRecipientUpdated(address indexed oldRecipient, address indexed newRecipient);
    event FeeBPSUpdated(uint256 oldBPS, uint256 newBPS);
    event MintingFeeCollected(address indexed recipient, uint256 amount);

    /// @custom:storage-location erc7201:cross.storage.forge.erc20.mintingfee
    struct ERC20MintingFeeStorage {
        address feeRecipient;
        uint96 feeBPS; // BPS(1 / 10000) feeBPS < 10000
    }

    // keccak256(abi.encode(uint256(keccak256("cross.storage.forge.erc20.mintingfee")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant ERC20MintingFeeStorageLocation =
        0x9e07bc4d06789dd0b871d60d502656ec22e58628182c2156c52dcfc63440ac00;

    function _getERC20MintingFeeStorage() private pure returns (ERC20MintingFeeStorage storage $) {
        assembly {
            $.slot := ERC20MintingFeeStorageLocation
        }
    }

    function __ERC20MintingFee_init(address feeRecipient, uint256 feeBPS) internal onlyInitializing {
        if (feeRecipient == address(0)) revert ERC20MintingFee__InvalidFeeRecipient(feeRecipient);
        if (feeBPS > 10000) revert ERC20MintingFee__InvalidFeeBPS(feeBPS);

        ERC20MintingFeeStorage storage $ = _getERC20MintingFeeStorage();
        $.feeRecipient = feeRecipient;
        $.feeBPS = uint96(feeBPS);

        emit FeeRecipientUpdated(address(0), feeRecipient);
        emit FeeBPSUpdated(0, feeBPS);
    }

    function mint(address to, uint256 amount) public virtual override {
        ERC20MintingFeeStorage storage $ = _getERC20MintingFeeStorage();
        if ($.feeBPS == 0) {
            super.mint(to, amount);
        } else {
            uint256 fee = (amount * $.feeBPS) / 10000;
            if (fee != 0) {
                amount -= fee;
                super.mint($.feeRecipient, fee);
                emit MintingFeeCollected($.feeRecipient, fee);
            }
            super.mint(to, amount);
        }
    }

    function mintingFeeBPS() external view returns (uint256) {
        return _getERC20MintingFeeStorage().feeBPS;
    }

    function mintingFeeRecipient() external view returns (address) {
        return _getERC20MintingFeeStorage().feeRecipient;
    }

    function setMintingFeeBPS(uint256 _feeBPS) external onlyOwner {
        if (_feeBPS > 10000) revert ERC20MintingFee__InvalidFeeBPS(_feeBPS);

        ERC20MintingFeeStorage storage $ = _getERC20MintingFeeStorage();
        emit FeeBPSUpdated($.feeBPS, _feeBPS);
        $.feeBPS = uint96(_feeBPS);
    }

    function setMintingFeeRecipient(address _feeRecipient) external onlyOwner {
        if (_feeRecipient == address(0)) revert ERC20MintingFee__InvalidFeeRecipient(_feeRecipient);

        ERC20MintingFeeStorage storage $ = _getERC20MintingFeeStorage();
        emit FeeRecipientUpdated($.feeRecipient, _feeRecipient);
        $.feeRecipient = _feeRecipient;
    }
}
