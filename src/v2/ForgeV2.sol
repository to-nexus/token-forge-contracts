// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.30;

import {EnumerableSet} from "@openzeppelin-contracts-5.3.0/utils/structs/EnumerableSet.sol";
import {NoncesUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/utils/NoncesUpgradeable.sol";
import {
    ERC721HolderUpgradeable
} from "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC721/utils/ERC721HolderUpgradeable.sol";
import {
    ERC1155HolderUpgradeable
} from "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC1155/utils/ERC1155HolderUpgradeable.sol";

import {SignatureChecker} from "@openzeppelin-contracts-5.3.0/utils/cryptography/SignatureChecker.sol";
import {IERC20, SafeERC20} from "@openzeppelin-contracts-5.3.0/token/ERC20/utils/SafeERC20.sol";
import {IERC721} from "@openzeppelin-contracts-5.3.0/token/ERC721/IERC721.sol";
import {IERC1155} from "@openzeppelin-contracts-5.3.0/token/ERC1155/IERC1155.sol";

import {IERC20Forge} from "../interfaces/IERC20Forge.sol";
import {IERC721Forge} from "../interfaces/IERC721Forge.sol";
import {IERC1155Forge} from "../interfaces/IERC1155Forge.sol";
import {TokenType, IForgeFactoryAlert} from "../interfaces/IForgeFactory.sol";
import {BaseForge} from "./BaseForge.sol";

abstract contract ERC20ForgeV2 is BaseForge {
    using EnumerableSet for EnumerableSet.AddressSet;
    using SafeERC20 for IERC20;

    error ERC20ForgeV2__InvalidAccountSignature(address account);

    bytes32 private constant ERC20_MINT_TYPE_HASH = keccak256(
        "ERC20Mint(address token,uint256 amount,address feeRecipient,uint256 feeBPS,uint256 nonce,uint256 deadline)"
    );
    bytes32 private constant ERC20_VALIDATOR_MINT_TYPE_HASH =
        keccak256("ValidatorERC20Mint(address recipient,bytes recipientSig)");

    bytes32 private constant ERC20_TRANSFER_TYPE_HASH = keccak256(
        "ERC20Transfer(address token,uint256 amount,address feeRecipient,uint256 feeBPS,uint256 nonce,uint256 deadline)"
    );
    bytes32 private constant ERC20_VALIDATOR_TRANSFER_TYPE_HASH =
        keccak256("ValidatorERC20Transfer(address recipient,bytes recipientSig)");

    bytes32 private constant ERC20_TRANSFER_FROM_TYPE_HASH = keccak256(
        "ERC20TransferFrom(address token,uint256 amount,address feeRecipient,uint256 feeBPS,uint256 nonce,uint256 deadline)"
    );
    bytes32 private constant ERC20_VALIDATOR_TRANSFER_FROM_TYPE_HASH =
        keccak256("ValidatorERC20TransferFrom(address from,bytes fromSig)");

    bytes32 private constant ERC20_BURN_TYPE_HASH = keccak256(
        "ERC20Burn(address token,uint256 amount,address feeRecipient,uint256 feeBPS,uint256 nonce,uint256 deadline)"
    );
    bytes32 private constant ERC20_VALIDATOR_BURN_TYPE_HASH =
        keccak256("ValidatorERC20Burn(address from,bytes fromSig)");

    function mintERC20(
        address recipient,
        address token,
        uint256 amount,
        address feeRecipient,
        uint256 feeBPS,
        uint256 deadline,
        bytes memory recipientSig,
        bytes memory validatorSig
    ) external checkDeadline(deadline) notBlacklisted(recipient) {
        uint256 nonce = _useNonce(recipient);
        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC20_MINT_TYPE_HASH, token, amount, feeRecipient, feeBPS, nonce, deadline));
            bytes32 recipientHash = _hashTypedDataV4(recipientStructHash);
            if (!SignatureChecker.isValidSignatureNow(recipient, recipientHash, recipientSig)) {
                revert ERC20ForgeV2__InvalidAccountSignature(recipient);
            }
        }
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC20_VALIDATOR_MINT_TYPE_HASH, recipient, keccak256(recipientSig)));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }
        (uint256 fee, uint256 value) = _erc20CalcFee(feeRecipient, feeBPS, amount);
        IERC20Forge(token).mint(recipient, value);
        if (fee != 0) {
            IERC20Forge(token).mint(feeRecipient, fee);
        }
        _alertMintToFactory(
            TokenType.ERC20, _calcUUID(recipient, nonce), token, abi.encode(recipient, amount, feeRecipient, fee)
        );
    }

    function transferERC20(
        address recipient,
        address token,
        uint256 amount,
        address feeRecipient,
        uint256 feeBPS,
        uint256 deadline,
        bytes memory recipientSig,
        bytes memory validatorSig
    ) external checkDeadline(deadline) notBlacklisted(recipient) {
        uint256 nonce = _useNonce(recipient);
        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC20_TRANSFER_TYPE_HASH, token, amount, feeRecipient, feeBPS, nonce, deadline));
            bytes32 recipientHash = _hashTypedDataV4(recipientStructHash);
            if (!SignatureChecker.isValidSignatureNow(recipient, recipientHash, recipientSig)) {
                revert ERC20ForgeV2__InvalidAccountSignature(recipient);
            }
        }
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC20_VALIDATOR_TRANSFER_TYPE_HASH, recipient, keccak256(recipientSig)));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }

        (uint256 fee, uint256 value) = _erc20CalcFee(feeRecipient, feeBPS, amount);
        IERC20(token).safeTransfer(recipient, value);
        if (fee != 0) {
            IERC20(token).safeTransfer(feeRecipient, fee);
        }
        _alertTransferToFactory(
            TokenType.ERC20, _calcUUID(recipient, nonce), token, abi.encode(recipient, amount, feeRecipient, fee)
        );
    }

    function transferFromERC20(
        address from,
        address token,
        uint256 amount,
        address feeRecipient,
        uint256 feeBPS,
        uint256 deadline,
        bytes memory fromSig,
        bytes memory validatorSig
    ) external checkDeadline(deadline) notBlacklisted(from) {
        _transferFromERC20(from, token, amount, feeRecipient, feeBPS, deadline, fromSig, validatorSig);
    }

    function transferFromERC20Permit(
        address from,
        address token,
        uint256 amount,
        address feeRecipient,
        uint256 feeBPS,
        uint256 deadline,
        bytes memory fromSig,
        bytes memory validatorSig,
        bytes memory permitSig
    ) external checkDeadline(deadline) notBlacklisted(from) erc20Permit(from, token, amount, deadline, permitSig) {
        _transferFromERC20(from, token, amount, feeRecipient, feeBPS, deadline, fromSig, validatorSig);
    }

    function _transferFromERC20(
        address from,
        address token,
        uint256 amount,
        address feeRecipient,
        uint256 feeBPS,
        uint256 deadline,
        bytes memory fromSig,
        bytes memory validatorSig
    ) private {
        uint256 nonce = _useNonce(from);
        {
            bytes32 fromStructHash = keccak256(
                abi.encode(ERC20_TRANSFER_FROM_TYPE_HASH, token, amount, feeRecipient, feeBPS, nonce, deadline)
            );
            bytes32 fromHash = _hashTypedDataV4(fromStructHash);
            if (!SignatureChecker.isValidSignatureNow(from, fromHash, fromSig)) {
                revert ERC20ForgeV2__InvalidAccountSignature(from);
            }
        }
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC20_VALIDATOR_TRANSFER_FROM_TYPE_HASH, from, keccak256(fromSig)));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }

        (uint256 fee,) = _erc20CalcFee(feeRecipient, feeBPS, amount);
        IERC20(token).safeTransferFrom(from, address(this), amount);
        if (fee != 0) {
            IERC20(token).safeTransfer(feeRecipient, fee);
        }
        _alertTransferFromToFactory(
            TokenType.ERC20, _calcUUID(from, nonce), token, abi.encode(from, amount, feeRecipient, fee)
        );
    }

    function burnERC20(
        address from,
        address token,
        uint256 amount,
        address feeRecipient,
        uint256 feeBPS,
        uint256 deadline,
        bytes memory fromSig,
        bytes memory validatorSig
    ) external checkDeadline(deadline) notBlacklisted(from) {
        _burnERC20(from, token, amount, feeRecipient, feeBPS, deadline, fromSig, validatorSig);
    }

    function burnERC20Permit(
        address from,
        address token,
        uint256 amount,
        address feeRecipient,
        uint256 feeBPS,
        uint256 deadline,
        bytes memory fromSig,
        bytes memory validatorSig,
        bytes memory permitSig
    ) external checkDeadline(deadline) notBlacklisted(from) erc20Permit(from, token, amount, deadline, permitSig) {
        _burnERC20(from, token, amount, feeRecipient, feeBPS, deadline, fromSig, validatorSig);
    }

    function _burnERC20(
        address from,
        address token,
        uint256 amount,
        address feeRecipient,
        uint256 feeBPS,
        uint256 deadline,
        bytes memory fromSig,
        bytes memory validatorSig
    ) private {
        uint256 nonce = _useNonce(from);
        {
            bytes32 fromStructHash =
                keccak256(abi.encode(ERC20_BURN_TYPE_HASH, token, amount, feeRecipient, feeBPS, nonce, deadline));
            bytes32 fromHash = _hashTypedDataV4(fromStructHash);
            if (!SignatureChecker.isValidSignatureNow(from, fromHash, fromSig)) {
                revert ERC20ForgeV2__InvalidAccountSignature(from);
            }
        }
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC20_VALIDATOR_BURN_TYPE_HASH, from, keccak256(fromSig)));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }

        (uint256 fee, uint256 value) = _erc20CalcFee(feeRecipient, feeBPS, amount);
        IERC20Forge(token).burnFrom(from, value);
        if (fee != 0) {
            IERC20(token).safeTransferFrom(from, feeRecipient, fee);
        }
        _alertBurnToFactory(TokenType.ERC20, _calcUUID(from, nonce), token, abi.encode(from, amount, feeRecipient, fee));
    }
}

abstract contract ERC721ForgeV2 is BaseForge, ERC721HolderUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;

    error ERC721ForgeV2__InvalidAccountSignature(address account);

    bytes32 private constant ERC721_MINT_TYPE_HASH =
        keccak256("ERC721Mint(address token,uint256 tokenID,bytes data,uint256 nonce,uint256 deadline)");
    bytes32 private constant ERC721_VALIDATOR_MINT_TYPE_HASH =
        keccak256("ValidatorERC721Mint(address recipient,bytes recipientSig)");

    bytes32 private constant ERC721_TRANSFER_TYPE_HASH =
        keccak256("ERC721Transfer(address token,uint256 tokenID,bytes data,uint256 nonce,uint256 deadline)");
    bytes32 private constant ERC721_VALIDATOR_TRANSFER_TYPE_HASH =
        keccak256("ValidatorERC721Transfer(address recipient,bytes recipientSig)");

    bytes32 private constant ERC721_BURN_TYPE_HASH =
        keccak256("ERC721Burn(address token,uint256 tokenID,uint256 nonce,uint256 deadline)");
    bytes32 private constant ERC721_VALIDATOR_BURN_TYPE_HASH =
        keccak256("ValidatorERC721Burn(address from,bytes fromSig)");

    function mintERC721(
        address recipient,
        address token,
        uint256 tokenID,
        bytes calldata data,
        uint256 deadline,
        bytes calldata recipientSig,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) notBlacklisted(recipient) {
        uint256 nonce = _useNonce(recipient);
        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC721_MINT_TYPE_HASH, token, tokenID, keccak256(data), nonce, deadline));
            bytes32 recipientHash = _hashTypedDataV4(recipientStructHash);
            if (!SignatureChecker.isValidSignatureNow(recipient, recipientHash, recipientSig)) {
                revert ERC721ForgeV2__InvalidAccountSignature(recipient);
            }
        }
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC721_VALIDATOR_MINT_TYPE_HASH, recipient, keccak256(recipientSig)));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }
        IERC721Forge(token).mint(recipient, tokenID, data);
        _alertMintToFactory(TokenType.ERC721, _calcUUID(recipient, nonce), token, abi.encode(recipient, tokenID));
    }

    function transferERC721(
        address recipient,
        address token,
        uint256 tokenID,
        bytes memory data,
        uint256 deadline,
        bytes calldata recipientSig,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) notBlacklisted(recipient) {
        uint256 nonce = _useNonce(recipient);
        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC721_TRANSFER_TYPE_HASH, token, tokenID, keccak256(data), nonce, deadline));
            bytes32 recipientHash = _hashTypedDataV4(recipientStructHash);
            if (!SignatureChecker.isValidSignatureNow(recipient, recipientHash, recipientSig)) {
                revert ERC721ForgeV2__InvalidAccountSignature(recipient);
            }
        }
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC721_VALIDATOR_TRANSFER_TYPE_HASH, recipient, keccak256(recipientSig)));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }
        IERC721(token).safeTransferFrom(address(this), recipient, tokenID, data);
        _alertTransferToFactory(TokenType.ERC721, _calcUUID(recipient, nonce), token, abi.encode(recipient, tokenID));
    }

    function burnERC721(
        address from,
        address token,
        uint256 tokenID,
        uint256 deadline,
        bytes calldata fromSig,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) notBlacklisted(from) {
        uint256 nonce = _useNonce(from);
        {
            bytes32 fromStructHash = keccak256(abi.encode(ERC721_BURN_TYPE_HASH, token, tokenID, nonce, deadline));
            bytes32 fromHash = _hashTypedDataV4(fromStructHash);
            if (!SignatureChecker.isValidSignatureNow(from, fromHash, fromSig)) {
                revert ERC721ForgeV2__InvalidAccountSignature(from);
            }
        }
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC721_VALIDATOR_BURN_TYPE_HASH, from, keccak256(fromSig)));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }
        IERC721Forge(token).burnFrom(from, tokenID);
        _alertBurnToFactory(TokenType.ERC721, _calcUUID(from, nonce), token, abi.encode(from, tokenID));
    }
}

abstract contract ERC1155ForgeV2 is BaseForge, ERC1155HolderUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;

    error ERC1155ForgeV2__InvalidAccountSignature(address account);

    bytes32 private constant ERC1155_MINT_TYPE_HASH = keccak256(
        "ERC1155Mint(address token,uint256 tokenID,uint256 amount,bytes data,uint256 nonce,uint256 deadline)"
    );
    bytes32 private constant ERC1155_VALIDATOR_MINT_TYPE_HASH =
        keccak256("ValidatorERC1155Mint(address recipient,bytes recipientSig)");
    bytes32 private constant ERC1155_TRANSFER_TYPE_HASH = keccak256(
        "ERC1155Transfer(address token,uint256 tokenID,uint256 amount,bytes data,uint256 nonce,uint256 deadline)"
    );
    bytes32 private constant ERC1155_VALIDATOR_TRANSFER_TYPE_HASH =
        keccak256("ValidatorERC1155Transfer(address recipient,bytes recipientSig)");
    bytes32 private constant ERC1155_BURN_TYPE_HASH =
        keccak256("ERC1155Burn(address token,uint256 tokenID,uint256 amount,uint256 nonce,uint256 deadline)");
    bytes32 private constant ERC1155_VALIDATOR_BURN_TYPE_HASH =
        keccak256("ValidatorERC1155Burn(address from,bytes fromSig)");

    bytes32 private constant ERC1155_BATCH_MINT_TYPE_HASH = keccak256(
        "ERC1155BatchMint(address token,uint256[] tokenIDs,uint256[] amounts,bytes data,uint256 nonce,uint256 deadline)"
    );
    bytes32 private constant ERC1155_VALIDATOR_BATCH_MINT_TYPE_HASH =
        keccak256("ValidatorERC1155BatchMint(address recipient,bytes recipientSig)");
    bytes32 private constant ERC1155_BATCH_TRANSFER_TYPE_HASH = keccak256(
        "ERC1155BatchTransfer(address token,uint256[] tokenIDs,uint256[] amounts,bytes data,uint256 nonce,uint256 deadline)"
    );
    bytes32 private constant ERC1155_VALIDATOR_BATCH_TRANSFER_TYPE_HASH =
        keccak256("ValidatorERC1155BatchTransfer(address recipient,bytes recipientSig)");
    bytes32 private constant ERC1155_BATCH_BURN_TYPE_HASH = keccak256(
        "ERC1155BatchBurn(address token,uint256[] tokenIDs,uint256[] amounts,uint256 nonce,uint256 deadline)"
    );
    bytes32 private constant ERC1155_VALIDATOR_BATCH_BURN_TYPE_HASH =
        keccak256("ValidatorERC1155BatchBurn(address from,bytes fromSig)");

    function mintERC1155(
        address recipient,
        address token,
        uint256 tokenID,
        uint256 amount,
        bytes memory data,
        uint256 deadline,
        bytes memory recipientSig,
        bytes memory validatorSig
    ) external checkDeadline(deadline) notBlacklisted(recipient) {
        uint256 nonce = _useNonce(recipient);
        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC1155_MINT_TYPE_HASH, token, tokenID, amount, keccak256(data), nonce, deadline));
            bytes32 recipientHash = _hashTypedDataV4(recipientStructHash);
            if (!SignatureChecker.isValidSignatureNow(recipient, recipientHash, recipientSig)) {
                revert ERC1155ForgeV2__InvalidAccountSignature(recipient);
            }
        }
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_MINT_TYPE_HASH, recipient, keccak256(recipientSig)));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }
        IERC1155Forge(token).mint(recipient, tokenID, amount, data);
        _alertMintToFactory(
            TokenType.ERC1155,
            _calcUUID(recipient, nonce),
            token,
            abi.encode(false, abi.encode(recipient, tokenID, amount))
        );
    }

    function transferERC1155(
        address recipient,
        address token,
        uint256 tokenID,
        uint256 amount,
        bytes memory data,
        uint256 deadline,
        bytes memory recipientSig,
        bytes memory validatorSig
    ) external checkDeadline(deadline) notBlacklisted(recipient) {
        uint256 nonce = _useNonce(recipient);
        {
            bytes32 recipientStructHash = keccak256(
                abi.encode(ERC1155_TRANSFER_TYPE_HASH, token, tokenID, amount, keccak256(data), nonce, deadline)
            );
            bytes32 recipientHash = _hashTypedDataV4(recipientStructHash);
            if (!SignatureChecker.isValidSignatureNow(recipient, recipientHash, recipientSig)) {
                revert ERC1155ForgeV2__InvalidAccountSignature(recipient);
            }
        }
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_TRANSFER_TYPE_HASH, recipient, keccak256(recipientSig)));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }
        IERC1155(token).safeTransferFrom(address(this), recipient, tokenID, amount, data);
        _alertTransferToFactory(
            TokenType.ERC1155,
            _calcUUID(recipient, nonce),
            token,
            abi.encode(false, abi.encode(recipient, tokenID, amount))
        );
    }

    function burnERC1155(
        address from,
        address token,
        uint256 tokenID,
        uint256 amount,
        uint256 deadline,
        bytes memory fromSig,
        bytes memory validatorSig
    ) external checkDeadline(deadline) notBlacklisted(from) {
        uint256 nonce = _useNonce(from);
        {
            bytes32 fromStructHash =
                keccak256(abi.encode(ERC1155_BURN_TYPE_HASH, token, tokenID, amount, nonce, deadline));
            bytes32 fromHash = _hashTypedDataV4(fromStructHash);
            if (!SignatureChecker.isValidSignatureNow(from, fromHash, fromSig)) {
                revert ERC1155ForgeV2__InvalidAccountSignature(from);
            }
        }
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_BURN_TYPE_HASH, from, keccak256(fromSig)));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }
        IERC1155Forge(token).burnFrom(from, tokenID, amount);
        _alertBurnToFactory(
            TokenType.ERC1155, _calcUUID(from, nonce), token, abi.encode(false, abi.encode(from, tokenID, amount))
        );
    }

    function mintERC1155Batch(
        address recipient,
        address token,
        uint256[] memory tokenIDs,
        uint256[] memory amounts,
        bytes memory data,
        uint256 deadline,
        bytes memory recipientSig,
        bytes memory validatorSig
    ) external checkDeadline(deadline) notBlacklisted(recipient) {
        uint256 nonce = _useNonce(recipient);
        {
            bytes32 recipientStructHash = keccak256(
                abi.encode(
                    ERC1155_BATCH_MINT_TYPE_HASH,
                    token,
                    keccak256(abi.encodePacked(tokenIDs)),
                    keccak256(abi.encodePacked(amounts)),
                    keccak256(data),
                    nonce,
                    deadline
                )
            );
            bytes32 recipientHash = _hashTypedDataV4(recipientStructHash);
            if (!SignatureChecker.isValidSignatureNow(recipient, recipientHash, recipientSig)) {
                revert ERC1155ForgeV2__InvalidAccountSignature(recipient);
            }
        }
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_BATCH_MINT_TYPE_HASH, recipient, keccak256(recipientSig)));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }
        IERC1155Forge(token).mintBatch(recipient, tokenIDs, amounts, data);
        _alertMintToFactory(
            TokenType.ERC1155,
            _calcUUID(recipient, nonce),
            token,
            abi.encode(true, abi.encode(recipient, tokenIDs, amounts))
        );
    }

    function transferERC1155Batch(
        address recipient,
        address token,
        uint256[] memory tokenIDs,
        uint256[] memory amounts,
        bytes memory data,
        uint256 deadline,
        bytes memory recipientSig,
        bytes memory validatorSig
    ) external checkDeadline(deadline) notBlacklisted(recipient) {
        uint256 nonce = _useNonce(recipient);
        {
            bytes32 recipientStructHash = keccak256(
                abi.encode(
                    ERC1155_BATCH_TRANSFER_TYPE_HASH,
                    token,
                    keccak256(abi.encodePacked(tokenIDs)),
                    keccak256(abi.encodePacked(amounts)),
                    keccak256(data),
                    nonce,
                    deadline
                )
            );
            bytes32 recipientHash = _hashTypedDataV4(recipientStructHash);
            if (!SignatureChecker.isValidSignatureNow(recipient, recipientHash, recipientSig)) {
                revert ERC1155ForgeV2__InvalidAccountSignature(recipient);
            }
        }
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_BATCH_TRANSFER_TYPE_HASH, recipient, keccak256(recipientSig)));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }
        IERC1155(token).safeBatchTransferFrom(address(this), recipient, tokenIDs, amounts, data);
        _alertTransferToFactory(
            TokenType.ERC1155,
            _calcUUID(recipient, nonce),
            token,
            abi.encode(true, abi.encode(recipient, tokenIDs, amounts))
        );
    }

    function burnERC1155Batch(
        address from,
        address token,
        uint256[] memory tokenIDs,
        uint256[] memory amounts,
        uint256 deadline,
        bytes memory fromSig,
        bytes memory validatorSig
    ) external checkDeadline(deadline) notBlacklisted(from) {
        uint256 nonce = _useNonce(from);
        {
            bytes32 fromStructHash = keccak256(
                abi.encode(
                    ERC1155_BATCH_BURN_TYPE_HASH,
                    token,
                    keccak256(abi.encodePacked(tokenIDs)),
                    keccak256(abi.encodePacked(amounts)),
                    nonce,
                    deadline
                )
            );
            bytes32 fromHash = _hashTypedDataV4(fromStructHash);
            if (!SignatureChecker.isValidSignatureNow(from, fromHash, fromSig)) {
                revert ERC1155ForgeV2__InvalidAccountSignature(from);
            }
        }
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_BATCH_BURN_TYPE_HASH, from, keccak256(fromSig)));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }
        IERC1155Forge(token).burnFromBatch(from, tokenIDs, amounts);
        _alertBurnToFactory(
            TokenType.ERC1155, _calcUUID(from, nonce), token, abi.encode(true, abi.encode(from, tokenIDs, amounts))
        );
    }
}

contract ForgeV2 is ERC20ForgeV2, ERC721ForgeV2, ERC1155ForgeV2 {
    // This contract combines the functionality of the three forge contracts
    // and can be used to mint, transfer, and burn ERC20, ERC721, and ERC1155 tokens.
    // It inherits from the individual forge contracts to provide a unified interface.

    }
