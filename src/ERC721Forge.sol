// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {EnumerableSet} from "@openzeppelin-contracts-5.3.0/utils/structs/EnumerableSet.sol";
import {NoncesUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/utils/NoncesUpgradeable.sol";
import {ERC721HolderUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC721/utils/ERC721HolderUpgradeable.sol";

import {ECDSA} from "@openzeppelin-contracts-5.3.0/utils/cryptography/ECDSA.sol";

import {IERC721Forge} from "./interfaces/IERC721Forge.sol";
import {TokenType, ITokenForgeFactoryAlert} from "./interfaces/ITokenForgeFactory.sol";
import {BaseForge} from "./BaseForge.sol";

abstract contract ERC721MintForge is BaseForge {
    using EnumerableSet for EnumerableSet.AddressSet;
    using ECDSA for bytes32;

    error ERC721MintForge__InvalidRecipientSignature(address recipient);

    bytes32 private constant ERC721_MINT_TYPEHASH = keccak256(
        "ERC721Mint(uint256 uuid, address recipient, address token, uint256 tokenID, uint256 nonce, uint256 deadline)"
    );
    bytes32 private constant ERC721_MINTTO_TYPEHASH =
        keccak256("ERC721MintTo(address token, uint256 tokenID, uint256 nonce, uint256 deadline)");
    bytes32 private constant ERC721_VALIDATOR_MINTTO_TYPEHASH =
        keccak256("ValidatorERC721MintTo(uint256 uuid, address recipient, bytes recipientSig)");

    // recipient: msg.sender
    function mintERC721(uint256 uuid, address token, uint256 tokenID, uint256 deadline, bytes calldata validatorSig)
        external
        checkDeadline(deadline)
    {
        address recipient = _msgSender();
        bytes32 structHash =
            keccak256(abi.encode(ERC721_MINT_TYPEHASH, uuid, recipient, token, tokenID, _useNonce(recipient), deadline));

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC721Forge(token).mint(recipient, tokenID);
        _alertMintToFactory(TokenType.ERC721, uuid, token, abi.encode(recipient, tokenID));
    }

    function mintERC721To(
        uint256 uuid,
        address recipient,
        address token,
        uint256 tokenID,
        uint256 deadline,
        bytes calldata recipientSig,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) {
        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC721_MINTTO_TYPEHASH, token, tokenID, _useNonce(recipient), deadline));
            bytes32 recipientHash = _hashTypedDataV4(recipientStructHash);
            address recipientSigner = ECDSA.recover(recipientHash, recipientSig);
            if (recipientSigner != recipient) revert ERC721MintForge__InvalidRecipientSignature(recipient);
        }
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC721_VALIDATOR_MINTTO_TYPEHASH, uuid, recipient, recipientSig));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }
        IERC721Forge(token).mint(recipient, tokenID);
        _alertMintToFactory(TokenType.ERC721, uuid, token, abi.encode(recipient, tokenID));
    }
}

abstract contract ERC721TransferForge is BaseForge, ERC721HolderUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;
    using ECDSA for bytes32;

    error ERC721TransferForge__InvalidRecipientSignature(address recipient);

    bytes32 private constant ERC721_TRANSFER_TYPEHASH = keccak256(
        "ERC721Transfer(uint256 uuid, address recipient, address token, uint256 tokenID, uint256 nonce, uint256 deadline)"
    );
    bytes32 private constant ERC721_TRANSFERTO_TYPEHASH =
        keccak256("ERC721TransferTo(address token, uint256 tokenID, uint256 nonce, uint256 deadline)");
    bytes32 private constant ERC721_VALIDATOR_TRANSFERTO_TYPEHASH =
        keccak256("ValidatorERC721TransferTo(uint256 uuid, address recipient, bytes recipientSig)");

    // recipient: msg.sender
    function transferERC721(
        uint256 uuid,
        address token,
        uint256 tokenID,
        bytes calldata data,
        uint256 deadline,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) {
        address recipient = _msgSender();
        bytes32 structHash = keccak256(
            abi.encode(ERC721_TRANSFER_TYPEHASH, uuid, recipient, token, tokenID, _useNonce(recipient), deadline)
        );

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC721Forge(token).safeTransferFrom(address(this), recipient, tokenID, data);
        _alertMintToFactory(TokenType.ERC721, uuid, token, abi.encode(recipient, tokenID));
    }

    function transferERC721To(
        uint256 uuid,
        address recipient,
        address token,
        uint256 tokenID,
        bytes memory data,
        uint256 deadline,
        bytes calldata recipientSig,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) {
        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC721_TRANSFERTO_TYPEHASH, token, tokenID, _useNonce(recipient), deadline));
            bytes32 recipientHash = _hashTypedDataV4(recipientStructHash);
            address recipientSigner = ECDSA.recover(recipientHash, recipientSig);
            if (recipientSigner != recipient) revert ERC721TransferForge__InvalidRecipientSignature(recipient);
        }
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC721_VALIDATOR_TRANSFERTO_TYPEHASH, uuid, recipient, recipientSig));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }
        IERC721Forge(token).safeTransferFrom(address(this), recipient, tokenID, data);
        _alertMintToFactory(TokenType.ERC721, uuid, token, abi.encode(recipient, tokenID));
    }
}

abstract contract ERC721BurnForge is BaseForge {
    using EnumerableSet for EnumerableSet.AddressSet;
    using ECDSA for bytes32;

    error ERC721BurnForge__InvalidFromSignature(address from);

    bytes32 private constant ERC721_BURN_TYPEHASH = keccak256(
        "ERC721Burn(uint256 uuid, address from, address token, uint256 tokenID, uint256 nonce, uint256 deadline)"
    );
    bytes32 private constant ERC721_BURNFROM_TYPEHASH =
        keccak256("ERC721BurnFrom(address token, uint256 tokenID, uint256 nonce, uint256 deadline)");
    bytes32 private constant ERC721_VALIDATOR_BURNFROM_TYPEHASH =
        keccak256("ValidatorERC721BurnFrom(uint256 uuid, address from, bytes fromSig)");

    // from: msg.sender
    function burnERC721(uint256 uuid, address token, uint256 tokenID, uint256 deadline, bytes calldata validatorSig)
        external
        checkDeadline(deadline)
    {
        address from = _msgSender();
        bytes32 structHash =
            keccak256(abi.encode(ERC721_BURN_TYPEHASH, uuid, from, token, tokenID, _useNonce(from), deadline));

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC721Forge(token).burnFrom(from, tokenID);
        _alertBurnToFactory(TokenType.ERC721, uuid, token, abi.encode(from, tokenID));
    }

    function burnERC721From(
        uint256 uuid,
        address from,
        address token,
        uint256 tokenID,
        uint256 deadline,
        bytes calldata fromSig,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) {
        {
            bytes32 fromStructHash =
                keccak256(abi.encode(ERC721_BURNFROM_TYPEHASH, token, tokenID, _useNonce(from), deadline));
            bytes32 fromHash = _hashTypedDataV4(fromStructHash);
            address fromSigner = ECDSA.recover(fromHash, fromSig);
            if (fromSigner != from) revert ERC721BurnForge__InvalidFromSignature(from);
        }
        {
            bytes32 validatorStructHash = keccak256(abi.encode(ERC721_VALIDATOR_BURNFROM_TYPEHASH, uuid, from, fromSig));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }
        IERC721Forge(token).burnFrom(from, tokenID);
        _alertBurnToFactory(TokenType.ERC721, uuid, token, abi.encode(from, tokenID));
    }
}

contract ERC721ForgeFacet is ERC721MintForge, ERC721TransferForge, ERC721BurnForge {}
