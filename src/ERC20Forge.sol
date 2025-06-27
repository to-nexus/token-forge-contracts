// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {EnumerableSet} from "@openzeppelin-contracts-5.3.0/utils/structs/EnumerableSet.sol";
import {NoncesUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/utils/NoncesUpgradeable.sol";

import {ECDSA} from "@openzeppelin-contracts-5.3.0/utils/cryptography/ECDSA.sol";
import {SafeERC20} from "@openzeppelin-contracts-5.3.0/token/ERC20/utils/SafeERC20.sol";

import {IERC20, IERC20Forge} from "./interfaces/IERC20Forge.sol";
import {TokenType, ITokenForgeFactoryAlert} from "./interfaces/ITokenForgeFactory.sol";
import {BaseForge} from "./BaseForge.sol";

abstract contract ERC20MintForge is BaseForge {
    using EnumerableSet for EnumerableSet.AddressSet;
    using ECDSA for bytes32;

    error ERC20MintForge__InvalidRecipientSignature(address recipient);

    bytes32 private constant ERC20_MINT_TYPEHASH = keccak256(
        "ERC20Mint(uint256 uuid, address recipient, address token, uint256 amount, uint256 nonce, uint256 deadline)"
    );
    bytes32 private constant ERC20_MINTTO_TYPEHASH =
        keccak256("ERC20MintTo(address token, uint256 amount, uint256 nonce, uint256 deadline)");
    bytes32 private constant ERC20_VALIDATOR_MINTTO_TYPEHASH =
        keccak256("ValidatorERC20MintTo(uint256 uuid, address recipient, bytes recipientSig)");

    // recipient: msg.sender
    function mintERC20(uint256 uuid, address token, uint256 amount, uint256 deadline, bytes calldata validatorSig)
        external
        checkDeadline(deadline)
    {
        address recipient = _msgSender();
        bytes32 structHash =
            keccak256(abi.encode(ERC20_MINT_TYPEHASH, uuid, recipient, token, amount, _useNonce(recipient), deadline));

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC20Forge(token).mint(recipient, amount);
        _alertMintToFactory(TokenType.ERC20, uuid, token, abi.encode(recipient, amount));
    }

    function mintERC20To(
        uint256 uuid,
        address recipient,
        address token,
        uint256 amount,
        uint256 deadline,
        bytes calldata recipientSig,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) {
        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC20_MINTTO_TYPEHASH, token, amount, _useNonce(recipient), deadline));
            bytes32 recipientHash = _hashTypedDataV4(recipientStructHash);
            address recipientSigner = ECDSA.recover(recipientHash, recipientSig);
            if (recipientSigner != recipient) revert ERC20MintForge__InvalidRecipientSignature(recipient);
        }
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC20_VALIDATOR_MINTTO_TYPEHASH, uuid, recipient, recipientSig));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }
        IERC20Forge(token).mint(recipient, amount);
        _alertMintToFactory(TokenType.ERC20, uuid, token, abi.encode(recipient, amount));
    }
}

abstract contract ERC20TransferForge is BaseForge {
    using EnumerableSet for EnumerableSet.AddressSet;
    using ECDSA for bytes32;
    using SafeERC20 for IERC20;

    error ERC20TransferForge__InvalidRecipientSignature(address recipient);

    bytes32 private constant ERC20_TRANSFER_TYPEHASH = keccak256(
        "ERC20Transfer(uint256 uuid, address recipient, address token, uint256 amount, uint256 nonce, uint256 deadline)"
    );
    bytes32 private constant ERC20_TRANSFERTO_TYPEHASH =
        keccak256("ERC20TransferTo(address token, uint256 amount, uint256 nonce, uint256 deadline)");
    bytes32 private constant ERC20_VALIDATOR_TRANSFERTO_TYPEHASH =
        keccak256("ValidatorERC20TransferTo(uint256 uuid, address recipient, bytes recipientSig)");

    // recipient: msg.sender
    function transferERC20(uint256 uuid, address token, uint256 amount, uint256 deadline, bytes calldata validatorSig)
        external
        checkDeadline(deadline)
    {
        address recipient = _msgSender();
        bytes32 structHash = keccak256(
            abi.encode(ERC20_TRANSFER_TYPEHASH, uuid, recipient, token, amount, _useNonce(recipient), deadline)
        );

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC20(token).safeTransfer(recipient, amount);
        _alertTransferToFactory(TokenType.ERC20, uuid, token, abi.encode(recipient, amount));
    }

    function transferERC20To(
        uint256 uuid,
        address recipient,
        address token,
        uint256 amount,
        uint256 deadline,
        bytes calldata recipientSig,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) {
        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC20_TRANSFERTO_TYPEHASH, token, amount, _useNonce(recipient), deadline));
            bytes32 recipientHash = _hashTypedDataV4(recipientStructHash);
            address recipientSigner = ECDSA.recover(recipientHash, recipientSig);
            if (recipientSigner != recipient) revert ERC20TransferForge__InvalidRecipientSignature(recipient);
        }
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC20_VALIDATOR_TRANSFERTO_TYPEHASH, uuid, recipient, recipientSig));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }
        IERC20(token).safeTransfer(recipient, amount);
        _alertTransferToFactory(TokenType.ERC20, uuid, token, abi.encode(recipient, amount));
    }
}

abstract contract ERC20BurnForge is BaseForge {
    using EnumerableSet for EnumerableSet.AddressSet;
    using ECDSA for bytes32;

    error ERC20BurnForge__InvalidFromSignature(address from);

    bytes32 private constant ERC20_BURN_TYPEHASH = keccak256(
        "ERC20Burn(uint256 uuid, address from, address token, uint256 amount, uint256 nonce, uint256 deadline)"
    );
    bytes32 private constant ERC20_BURNFROM_TYPEHASH =
        keccak256("ERC20BurnFrom(address token, uint256 amount, uint256 nonce, uint256 deadline)");
    bytes32 private constant ERC20_VALIDATOR_BURNFROM_TYPEHASH =
        keccak256("ValidatorERC20BurnFrom(uint256 uuid, address from, bytes fromSig)");

    // recipient: msg.sender
    function burnERC20(uint256 uuid, address token, uint256 amount, uint256 deadline, bytes calldata validatorSig)
        external
        checkDeadline(deadline)
    {
        address from = _msgSender();
        bytes32 structHash =
            keccak256(abi.encode(ERC20_BURN_TYPEHASH, uuid, from, token, amount, _useNonce(from), deadline));

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC20Forge(token).burnFrom(from, amount);
        _alertBurnToFactory(TokenType.ERC20, uuid, token, abi.encode(from, amount));
    }

    function burnERC20From(
        uint256 uuid,
        address from,
        address token,
        uint256 amount,
        uint256 deadline,
        bytes calldata fromSig,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) {
        {
            bytes32 fromStructHash =
                keccak256(abi.encode(ERC20_BURNFROM_TYPEHASH, token, amount, _useNonce(from), deadline));
            bytes32 fromHash = _hashTypedDataV4(fromStructHash);
            address fromSigner = ECDSA.recover(fromHash, fromSig);
            if (fromSigner != from) revert ERC20BurnForge__InvalidFromSignature(from);
        }
        {
            bytes32 validatorStructHash = keccak256(abi.encode(ERC20_VALIDATOR_BURNFROM_TYPEHASH, uuid, from, fromSig));
            bytes32 validatorHash = _hashTypedDataV4(validatorStructHash);
            _verifyValidatorSignature(validatorHash, validatorSig);
        }
        IERC20Forge(token).burnFrom(from, amount);
        _alertBurnToFactory(TokenType.ERC20, uuid, token, abi.encode(from, amount));
    }
}

contract ERC20ForgeFacet is ERC20MintForge, ERC20TransferForge, ERC20BurnForge {}
