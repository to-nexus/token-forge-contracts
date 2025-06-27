// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {EnumerableSet} from "@openzeppelin-contracts-5.3.0/utils/structs/EnumerableSet.sol";
import {NoncesUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/utils/NoncesUpgradeable.sol";
import {ERC1155HolderUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC1155/utils/ERC1155HolderUpgradeable.sol";

import {ECDSA} from "@openzeppelin-contracts-5.3.0/utils/cryptography/ECDSA.sol";

import {IERC1155Forge} from "./interfaces/IERC1155Forge.sol";
import {TokenType, ITokenForgeFactoryAlert} from "./interfaces/ITokenForgeFactory.sol";
import {BaseForge} from "./BaseForge.sol";

abstract contract ERC1155MintForge is BaseForge {
    using EnumerableSet for EnumerableSet.AddressSet;
    using ECDSA for bytes32;

    error ERC1155MintForge__InvalidRecipientSignature(address recipient);

    bytes32 private constant ERC1155_MINT_TYPEHASH = keccak256(
        "ERC1155Mint(uint256 uuid, address recipient, address token, uint256 tokenID, uint256 amount, uint256 nonce, uint256 deadline)"
    );
    bytes32 private constant ERC1155_MINT_BATCH_TYPEHASH = keccak256(
        "ERC1155MintBatch(uint256 uuid, address recipient, address token, uint256[] tokenIDs, uint256[] amounts, uint256 nonce, uint256 deadline)"
    );

    bytes32 private constant ERC1155_MINTTO_TYPEHASH =
        keccak256("ERC1155MintTo(address token, uint256 tokenID, uint256 amount, uint256 nonce, uint256 deadline)");
    bytes32 private constant ERC1155_MINTTO_BATCH_TYPEHASH = keccak256(
        "ERC1155MintToBatch(address token, uint256[] tokenIDs, uint256[] amounts, uint256 nonce, uint256 deadline)"
    );
    bytes32 private constant ERC1155_VALIDATOR_MINTTO_TYPEHASH =
        keccak256("ValidatorERC1155MintTo(uint256 uuid, address recipient, bytes recipientSig)");
    bytes32 private constant ERC1155_VALIDATOR_MINTTO_BATCH_TYPEHASH =
        keccak256("ValidatorERC1155MintToBatch(uint256 uuid, address recipient, bytes recipientSig)");

    // recipient: msg.sender
    function mintERC1155(
        uint256 uuid,
        address token,
        uint256 tokenID,
        uint256 amount,
        bytes calldata data,
        uint256 deadline,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) {
        address recipient = _msgSender();
        {
            bytes32 structHash = keccak256(
                abi.encode(
                    ERC1155_MINT_TYPEHASH, uuid, recipient, token, tokenID, amount, _useNonce(recipient), deadline
                )
            );

            bytes32 hash = _hashTypedDataV4(structHash);
            _verifyValidatorSignature(hash, validatorSig);
        }

        IERC1155Forge(token).mint(recipient, tokenID, amount, data);
        _alertMintToFactory(TokenType.ERC1155, uuid, token, abi.encode(false, abi.encode(recipient, tokenID, amount)));
    }

    function mintERC1155Batch(
        uint256 uuid,
        address token,
        uint256[] memory tokenIDs,
        uint256[] memory amounts,
        bytes memory data,
        uint256 deadline,
        bytes memory validatorSig
    ) external checkDeadline(deadline) {
        address recipient = _msgSender();
        bytes32 structHash = keccak256(
            abi.encode(
                ERC1155_MINT_BATCH_TYPEHASH, uuid, recipient, token, tokenIDs, amounts, _useNonce(recipient), deadline
            )
        );

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC1155Forge(token).mintBatch(recipient, tokenIDs, amounts, data);
        _alertMintToFactory(TokenType.ERC1155, uuid, token, abi.encode(true, abi.encode(recipient, tokenIDs, amounts)));
    }

    function mintERC1155To(
        uint256 uuid,
        address recipient,
        address token,
        uint256 tokenID,
        uint256 amount,
        bytes memory data,
        uint256 deadline,
        bytes memory recipientSig,
        bytes memory validatorSig
    ) external checkDeadline(deadline) {
        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC1155_MINTTO_TYPEHASH, token, tokenID, amount, _useNonce(recipient), deadline));
            bytes32 recipientHash = _hashTypedDataV4(recipientStructHash);
            address recipientSigner = ECDSA.recover(recipientHash, recipientSig);
            if (recipientSigner != recipient) revert ERC1155MintForge__InvalidRecipientSignature(recipient);
        }
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_MINTTO_TYPEHASH, uuid, recipient, recipientSig));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }
        IERC1155Forge(token).mint(recipient, tokenID, amount, data);
        _alertMintToFactory(TokenType.ERC1155, uuid, token, abi.encode(false, abi.encode(recipient, tokenID, amount)));
    }

    function mintERC1155ToBatch(
        uint256 uuid,
        address recipient,
        address token,
        uint256[] memory tokenIDs,
        uint256[] memory amounts,
        bytes memory data,
        uint256 deadline,
        bytes memory recipientSig,
        bytes memory validatorSig
    ) external checkDeadline(deadline) {
        {
            bytes32 recipientStructHash = keccak256(
                abi.encode(ERC1155_MINTTO_BATCH_TYPEHASH, token, tokenIDs, amounts, _useNonce(recipient), deadline)
            );
            bytes32 recipientHash = _hashTypedDataV4(recipientStructHash);
            address recipientSigner = ECDSA.recover(recipientHash, recipientSig);
            if (recipientSigner != recipient) revert ERC1155MintForge__InvalidRecipientSignature(recipient);
        }
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_MINTTO_BATCH_TYPEHASH, uuid, recipient, recipientSig));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }
        IERC1155Forge(token).mintBatch(recipient, tokenIDs, amounts, data);
        _alertMintToFactory(TokenType.ERC1155, uuid, token, abi.encode(true, abi.encode(recipient, tokenIDs, amounts)));
    }
}

abstract contract ERC1155TransferForge is BaseForge, ERC1155HolderUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;
    using ECDSA for bytes32;

    error ERC1155TransferForge__InvalidRecipientSignature(address recipient);

    bytes32 private constant ERC1155_TRANSFER_TYPEHASH = keccak256(
        "ERC1155Transfer(uint256 uuid, address recipient, address token, uint256 tokenID, uint256 amount, uint256 nonce, uint256 deadline)"
    );
    bytes32 private constant ERC1155_TRANSFER_BATCH_TYPEHASH = keccak256(
        "ERC1155TransferBatch(uint256 uuid, address recipient, address token, uint256[] tokenIDs, uint256[] amounts, uint256 nonce, uint256 deadline)"
    );

    bytes32 private constant ERC1155_TRANSFERTO_TYPEHASH =
        keccak256("ERC1155TransferTo(address token, uint256 tokenID, uint256 amount, uint256 nonce, uint256 deadline)");
    bytes32 private constant ERC1155_TRANSFERTO_BATCH_TYPEHASH = keccak256(
        "ERC1155TransferToBatch(address token, uint256[] tokenIDs, uint256[] amounts, uint256 nonce, uint256 deadline)"
    );
    bytes32 private constant ERC1155_VALIDATOR_TRANSFERTO_TYPEHASH =
        keccak256("ValidatorERC1155TransferTo(uint256 uuid, address recipient, bytes recipientSig)");
    bytes32 private constant ERC1155_VALIDATOR_TRANSFERTO_BATCH_TYPEHASH =
        keccak256("ValidatorERC1155TransferToBatch(uint256 uuid, address recipient, bytes recipientSig)");

    // recipient: msg.sender
    function transferERC1155(
        uint256 uuid,
        address token,
        uint256 tokenID,
        uint256 amount,
        bytes calldata data,
        uint256 deadline,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) {
        address recipient = _msgSender();
        {
            bytes32 structHash = keccak256(
                abi.encode(
                    ERC1155_TRANSFER_TYPEHASH, uuid, recipient, token, tokenID, amount, _useNonce(recipient), deadline
                )
            );

            bytes32 hash = _hashTypedDataV4(structHash);
            _verifyValidatorSignature(hash, validatorSig);
        }

        IERC1155Forge(token).safeTransferFrom(address(this), recipient, tokenID, amount, data);
        _alertTransferToFactory(
            TokenType.ERC1155, uuid, token, abi.encode(false, abi.encode(recipient, tokenID, amount))
        );
    }

    function transferERC1155Batch(
        uint256 uuid,
        address token,
        uint256[] memory tokenIDs,
        uint256[] memory amounts,
        bytes memory data,
        uint256 deadline,
        bytes memory validatorSig
    ) external checkDeadline(deadline) {
        address recipient = _msgSender();
        bytes32 structHash = keccak256(
            abi.encode(
                ERC1155_TRANSFER_BATCH_TYPEHASH,
                uuid,
                recipient,
                token,
                tokenIDs,
                amounts,
                _useNonce(recipient),
                deadline
            )
        );

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC1155Forge(token).safeBatchTransferFrom(address(this), recipient, tokenIDs, amounts, data);
        _alertTransferToFactory(
            TokenType.ERC1155, uuid, token, abi.encode(true, abi.encode(recipient, tokenIDs, amounts))
        );
    }

    function transferERC1155To(
        uint256 uuid,
        address recipient,
        address token,
        uint256 tokenID,
        uint256 amount,
        bytes memory data,
        uint256 deadline,
        bytes memory recipientSig,
        bytes memory validatorSig
    ) external checkDeadline(deadline) {
        {
            bytes32 recipientStructHash = keccak256(
                abi.encode(ERC1155_TRANSFERTO_TYPEHASH, token, tokenID, amount, _useNonce(recipient), deadline)
            );
            bytes32 recipientHash = _hashTypedDataV4(recipientStructHash);
            address recipientSigner = ECDSA.recover(recipientHash, recipientSig);
            if (recipientSigner != recipient) revert ERC1155TransferForge__InvalidRecipientSignature(recipient);
        }
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_TRANSFERTO_TYPEHASH, uuid, recipient, recipientSig));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }
        IERC1155Forge(token).safeTransferFrom(address(this), recipient, tokenID, amount, data);
        _alertTransferToFactory(
            TokenType.ERC1155, uuid, token, abi.encode(false, abi.encode(recipient, tokenID, amount))
        );
    }

    function transferERC1155ToBatch(
        uint256 uuid,
        address recipient,
        address token,
        uint256[] memory tokenIDs,
        uint256[] memory amounts,
        bytes memory data,
        uint256 deadline,
        bytes memory recipientSig,
        bytes memory validatorSig
    ) external checkDeadline(deadline) {
        {
            bytes32 recipientStructHash = keccak256(
                abi.encode(ERC1155_TRANSFERTO_BATCH_TYPEHASH, token, tokenIDs, amounts, _useNonce(recipient), deadline)
            );
            bytes32 recipientHash = _hashTypedDataV4(recipientStructHash);
            address recipientSigner = ECDSA.recover(recipientHash, recipientSig);
            if (recipientSigner != recipient) revert ERC1155TransferForge__InvalidRecipientSignature(recipient);
        }
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_TRANSFERTO_BATCH_TYPEHASH, uuid, recipient, recipientSig));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }
        IERC1155Forge(token).safeBatchTransferFrom(address(this), recipient, tokenIDs, amounts, data);
        _alertTransferToFactory(
            TokenType.ERC1155, uuid, token, abi.encode(true, abi.encode(recipient, tokenIDs, amounts))
        );
    }
}

abstract contract ERC1155BurnForge is BaseForge {
    using EnumerableSet for EnumerableSet.AddressSet;
    using ECDSA for bytes32;

    error ERC1155BurnForge__InvalidFromSignature(address from);

    bytes32 private constant ERC1155_BURN_TYPEHASH = keccak256(
        "ERC1155Burn(uint256 uuid, address from, address token, uint256 tokenID, uint256 amount, uint256 nonce, uint256 deadline)"
    );
    bytes32 private constant ERC1155_BURN_BATCH_TYPEHASH = keccak256(
        "ERC1155BurnBatch(uint256 uuid, address from, address token, uint256[] tokenIDs, uint256[] amounts, uint256 nonce, uint256 deadline)"
    );
    bytes32 private constant ERC1155_BURNFROM_TYPEHASH =
        keccak256("ERC1155BurnFrom(address token, uint256 tokenID, uint256 amount, uint256 nonce, uint256 deadline)");
    bytes32 private constant ERC1155_BURNFROM_BATCH_TYPEHASH = keccak256(
        "ERC1155BurnFromBatch(address token, uint256[] tokenIDs, uint256[] amounts, uint256 nonce, uint256 deadline)"
    );
    bytes32 private constant ERC1155_VALIDATOR_BURNFROM_TYPEHASH =
        keccak256("ValidatorERC1155BurnFrom(uint256 uuid, address from, bytes fromSig)");
    bytes32 private constant ERC1155_VALIDATOR_BURNFROM_BATCH_TYPEHASH =
        keccak256("ValidatorERC1155BurnFromBatch(uint256 uuid, address from, bytes fromSig)");

    // from: msg.sender
    function burnERC1155(
        uint256 uuid,
        address token,
        uint256 tokenID,
        uint256 amount,
        uint256 deadline,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) {
        address from = _msgSender();
        bytes32 structHash =
            keccak256(abi.encode(ERC1155_BURN_TYPEHASH, uuid, from, token, tokenID, amount, _useNonce(from), deadline));

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC1155Forge(token).burnFrom(from, tokenID, amount);
        _alertBurnToFactory(TokenType.ERC1155, uuid, token, abi.encode(false, abi.encode(from, tokenID, amount)));
    }

    function burnERC1155Batch(
        uint256 uuid,
        address token,
        uint256[] memory tokenIDs,
        uint256[] memory amounts,
        uint256 deadline,
        bytes memory validatorSig
    ) external checkDeadline(deadline) {
        address from = _msgSender();
        bytes32 structHash = keccak256(
            abi.encode(ERC1155_BURN_BATCH_TYPEHASH, uuid, from, token, tokenIDs, amounts, _useNonce(from), deadline)
        );

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC1155Forge(token).burnFromBatch(from, tokenIDs, amounts);
        _alertBurnToFactory(TokenType.ERC1155, uuid, token, abi.encode(true, abi.encode(from, tokenIDs, amounts)));
    }

    function burnERC1155From(
        uint256 uuid,
        address from,
        address token,
        uint256 tokenID,
        uint256 amount,
        uint256 deadline,
        bytes memory fromSig,
        bytes memory validatorSig
    ) external checkDeadline(deadline) {
        {
            bytes32 fromStructHash =
                keccak256(abi.encode(ERC1155_BURNFROM_TYPEHASH, token, tokenID, amount, _useNonce(from), deadline));
            bytes32 fromHash = _hashTypedDataV4(fromStructHash);
            address fromSigner = ECDSA.recover(fromHash, fromSig);
            if (fromSigner != from) revert ERC1155BurnForge__InvalidFromSignature(from);
        }
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_BURNFROM_TYPEHASH, uuid, from, fromSig));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }
        IERC1155Forge(token).burnFrom(from, tokenID, amount);
        _alertBurnToFactory(TokenType.ERC1155, uuid, token, abi.encode(false, abi.encode(from, tokenID, amount)));
    }

    function burnERC1155FromBatch(
        uint256 uuid,
        address from,
        address token,
        uint256[] memory tokenIDs,
        uint256[] memory amounts,
        uint256 deadline,
        bytes memory fromSig,
        bytes memory validatorSig
    ) external checkDeadline(deadline) {
        {
            bytes32 fromStructHash = keccak256(
                abi.encode(ERC1155_BURNFROM_BATCH_TYPEHASH, token, tokenIDs, amounts, _useNonce(from), deadline)
            );
            bytes32 fromHash = _hashTypedDataV4(fromStructHash);
            address fromSigner = ECDSA.recover(fromHash, fromSig);
            if (fromSigner != from) revert ERC1155BurnForge__InvalidFromSignature(from);
        }
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_BURNFROM_BATCH_TYPEHASH, uuid, from, fromSig));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }
        IERC1155Forge(token).burnFromBatch(from, tokenIDs, amounts);
        _alertBurnToFactory(TokenType.ERC1155, uuid, token, abi.encode(true, abi.encode(from, tokenIDs, amounts)));
    }
}

contract ERC1155ForgeFacet is ERC1155MintForge, ERC1155TransferForge, ERC1155BurnForge {}
