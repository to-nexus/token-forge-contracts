// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {EnumerableSet} from "@openzeppelin-contracts-5.3.0/utils/structs/EnumerableSet.sol";
import {NoncesUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/utils/NoncesUpgradeable.sol";
import {ERC721HolderUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC721/utils/ERC721HolderUpgradeable.sol";
import {ERC1155HolderUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC1155/utils/ERC1155HolderUpgradeable.sol";

import {ECDSA} from "@openzeppelin-contracts-5.3.0/utils/cryptography/ECDSA.sol";
import {IERC20, SafeERC20} from "@openzeppelin-contracts-5.3.0/token/ERC20/utils/SafeERC20.sol";
import {IERC721} from "@openzeppelin-contracts-5.3.0/token/ERC721/IERC721.sol";
import {IERC1155} from "@openzeppelin-contracts-5.3.0/token/ERC1155/IERC1155.sol";

import {IERC20Forge} from "../interfaces/IERC20Forge.sol";
import {IERC721Forge} from "../interfaces/IERC721Forge.sol";
import {IERC1155Forge} from "../interfaces/IERC1155Forge.sol";
import {TokenType, IForgeFactoryAlert} from "../interfaces/IForgeFactory.sol";
import {BaseForge} from "./BaseForge.sol";

abstract contract ERC20ForgeV3 is BaseForge {
    using EnumerableSet for EnumerableSet.AddressSet;
    using ECDSA for bytes32;
    using SafeERC20 for IERC20;

    bytes32 private constant ERC20_MINT_TYPE_HASH =
        keccak256("ERC20Mint(address recipient,address token,uint256 amount,uint256 nonce,uint256 deadline)");
    bytes32 private constant ERC20_TRANSFER_TYPE_HASH =
        keccak256("ERC20Transfer(address recipient,address token,uint256 amount,uint256 nonce,uint256 deadline)");
    bytes32 private constant ERC20_TRANSFER_FROM_TYPE_HASH =
        keccak256("ERC20TransferFrom(address from,address token,uint256 amount,uint256 nonce,uint256 deadline)");
    bytes32 private constant ERC20_BURN_TYPE_HASH =
        keccak256("ERC20Burn(address from,address token,uint256 amount,uint256 nonce,uint256 deadline)");

    function mintERC20(address recipient, address token, uint256 amount, uint256 deadline, bytes calldata validatorSig)
        external
        checkDeadline(deadline)
    {
        uint256 nonce = _useNonce(recipient);
        bytes32 structHash = keccak256(abi.encode(ERC20_MINT_TYPE_HASH, recipient, token, amount, nonce, deadline));

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC20Forge(token).mint(recipient, amount);
        _alertMintToFactory(TokenType.ERC20, _calcUUID(recipient, nonce), token, abi.encode(recipient, amount));
    }

    function transferERC20(
        address recipient,
        address token,
        uint256 amount,
        uint256 deadline,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) {
        uint256 nonce = _useNonce(recipient);
        bytes32 structHash = keccak256(abi.encode(ERC20_TRANSFER_TYPE_HASH, recipient, token, amount, nonce, deadline));

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC20(token).safeTransfer(recipient, amount);
        _alertTransferToFactory(TokenType.ERC20, _calcUUID(recipient, nonce), token, abi.encode(recipient, amount));
    }

    function transferFromERC20Permit(
        address from,
        address token,
        uint256 amount,
        uint256 deadline,
        bytes calldata validatorSig,
        bytes memory permitSig
    ) external checkDeadline(deadline) erc20Permit(from, token, amount, deadline, permitSig) {
        _transferFromERC20(from, token, amount, deadline, validatorSig);
    }

    function _transferFromERC20(
        address from,
        address token,
        uint256 amount,
        uint256 deadline,
        bytes calldata validatorSig
    ) private {
        uint256 nonce = _useNonce(from);
        bytes32 structHash = keccak256(abi.encode(ERC20_TRANSFER_FROM_TYPE_HASH, from, token, amount, nonce, deadline));

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC20(token).safeTransferFrom(from, address(this), amount);
        _alertTransferFromToFactory(TokenType.ERC20, _calcUUID(from, nonce), token, abi.encode(from, amount));
    }

    function burnERC20Permit(
        address from,
        address token,
        uint256 amount,
        uint256 deadline,
        bytes calldata validatorSig,
        bytes memory permitSig
    ) external checkDeadline(deadline) erc20Permit(from, token, amount, deadline, permitSig) {
        _burnERC20(from, token, amount, deadline, validatorSig);
    }

    function _burnERC20(address from, address token, uint256 amount, uint256 deadline, bytes calldata validatorSig)
        private
    {
        uint256 nonce = _useNonce(from);
        bytes32 structHash = keccak256(abi.encode(ERC20_BURN_TYPE_HASH, from, token, amount, nonce, deadline));

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC20Forge(token).burnFrom(from, amount);
        _alertBurnToFactory(TokenType.ERC20, _calcUUID(from, nonce), token, abi.encode(from, amount));
    }
}

abstract contract ERC721ForgeV3 is BaseForge, ERC721HolderUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;
    using ECDSA for bytes32;

    error ERC721MintForge__InvalidRecipientSignature(address recipient);

    bytes32 private constant ERC721_MINT_TYPE_HASH =
        keccak256("ERC721Mint(address recipient,address token,uint256 tokenID,uint256 nonce,uint256 deadline)");
    bytes32 private constant ERC721_TRANSFER_TYPE_HASH =
        keccak256("ERC721Transfer(address recipient,address token,uint256 tokenID,uint256 nonce,uint256 deadline)");

    function mintERC721(
        address recipient,
        address token,
        uint256 tokenID,
        uint256 deadline,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) {
        uint256 nonce = _useNonce(recipient);
        bytes32 structHash = keccak256(abi.encode(ERC721_MINT_TYPE_HASH, recipient, token, tokenID, nonce, deadline));

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC721Forge(token).mint(recipient, tokenID);
        _alertMintToFactory(TokenType.ERC721, _calcUUID(recipient, nonce), token, abi.encode(recipient, tokenID));
    }

    function transferERC721(
        address recipient,
        address token,
        uint256 tokenID,
        uint256 deadline,
        bytes calldata validatorSig,
        bytes calldata data
    ) external checkDeadline(deadline) {
        uint256 nonce = _useNonce(recipient);
        bytes32 structHash =
            keccak256(abi.encode(ERC721_TRANSFER_TYPE_HASH, recipient, token, tokenID, nonce, deadline));

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC721(token).safeTransferFrom(address(this), recipient, tokenID, data);
        _alertTransferToFactory(TokenType.ERC721, _calcUUID(recipient, nonce), token, abi.encode(recipient, tokenID));
    }
}

abstract contract ERC1155ForgeV3 is BaseForge, ERC1155HolderUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;
    using ECDSA for bytes32;

    bytes32 private constant ERC1155_MINT_TYPE_HASH = keccak256(
        "ERC1155Mint(address recipient,address token,uint256 tokenID,uint256 amount,uint256 nonce,uint256 deadline)"
    );
    bytes32 private constant ERC1155_TRANSFER_TYPE_HASH = keccak256(
        "ERC1155Transfer(address recipient,address token,uint256 tokenID,uint256 amount,uint256 nonce,uint256 deadline)"
    );

    bytes32 private constant ERC1155_BATCH_MINT_TYPE_HASH = keccak256(
        "ERC1155BatchMint(address recipient,address token,uint256[] tokenIDs,uint256[] amounts,uint256 nonce,uint256 deadline)"
    );
    bytes32 private constant ERC1155_BATCH_TRANSFER_TYPE_HASH = keccak256(
        "ERC1155BatchTransfer(address recipient,address token,uint256[] tokenIDs,uint256[] amounts,uint256 nonce,uint256 deadline)"
    );

    function mintERC1155(
        address recipient,
        address token,
        uint256 tokenID,
        uint256 amount,
        uint256 deadline,
        bytes memory validatorSig,
        bytes memory data
    ) external checkDeadline(deadline) {
        uint256 nonce = _useNonce(recipient);
        {
            bytes32 structHash =
                keccak256(abi.encode(ERC1155_MINT_TYPE_HASH, recipient, token, tokenID, amount, nonce, deadline));

            bytes32 hash = _hashTypedDataV4(structHash);
            _verifyValidatorSignature(hash, validatorSig);
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
        uint256 deadline,
        bytes memory validatorSig,
        bytes memory data
    ) external checkDeadline(deadline) {
        uint256 nonce = _useNonce(recipient);
        {
            bytes32 structHash =
                keccak256(abi.encode(ERC1155_TRANSFER_TYPE_HASH, recipient, token, tokenID, amount, nonce, deadline));

            bytes32 hash = _hashTypedDataV4(structHash);
            _verifyValidatorSignature(hash, validatorSig);
        }

        IERC1155(token).safeTransferFrom(address(this), recipient, tokenID, amount, data);
        _alertTransferToFactory(
            TokenType.ERC1155,
            _calcUUID(recipient, nonce),
            token,
            abi.encode(false, abi.encode(recipient, tokenID, amount))
        );
    }

    function mintERC1155Batch(
        address recipient,
        address token,
        uint256[] memory tokenIDs,
        uint256[] memory amounts,
        uint256 deadline,
        bytes memory validatorSig,
        bytes memory data
    ) external checkDeadline(deadline) {
        uint256 nonce = _useNonce(recipient);
        bytes32 structHash = keccak256(
            abi.encode(
                ERC1155_BATCH_MINT_TYPE_HASH,
                recipient,
                token,
                keccak256(abi.encodePacked(tokenIDs)),
                keccak256(abi.encodePacked(amounts)),
                nonce,
                deadline
            )
        );

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

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
        uint256 deadline,
        bytes memory validatorSig,
        bytes memory data
    ) external checkDeadline(deadline) {
        uint256 nonce = _useNonce(recipient);
        bytes32 structHash = keccak256(
            abi.encode(
                ERC1155_BATCH_TRANSFER_TYPE_HASH,
                recipient,
                token,
                keccak256(abi.encodePacked(tokenIDs)),
                keccak256(abi.encodePacked(amounts)),
                nonce,
                deadline
            )
        );

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC1155(token).safeBatchTransferFrom(address(this), recipient, tokenIDs, amounts, data);
        _alertTransferToFactory(
            TokenType.ERC1155,
            _calcUUID(recipient, nonce),
            token,
            abi.encode(true, abi.encode(recipient, tokenIDs, amounts))
        );
    }
}

contract ForgeV3 is ERC20ForgeV3, ERC721ForgeV3, ERC1155ForgeV3 {
// This contract serves as a version marker for the Forge contracts.
// It does not contain any additional logic or state.
// Future versions can inherit from this contract to maintain compatibility.
}
