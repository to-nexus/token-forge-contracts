// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {EnumerableSet} from "@openzeppelin-contracts-5.3.0/utils/structs/EnumerableSet.sol";
import {NoncesUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/utils/NoncesUpgradeable.sol";
import {ERC721HolderUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC721/utils/ERC721HolderUpgradeable.sol";
import {ERC1155HolderUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC1155/utils/ERC1155HolderUpgradeable.sol";

import {ECDSA} from "@openzeppelin-contracts-5.3.0/utils/cryptography/ECDSA.sol";
import {SafeERC20} from "@openzeppelin-contracts-5.3.0/token/ERC20/utils/SafeERC20.sol";

import {IERC20, IERC20Forge} from "./interfaces/IERC20Forge.sol";
import {IERC721Forge} from "./interfaces/IERC721Forge.sol";
import {IERC1155Forge} from "./interfaces/IERC1155Forge.sol";
import {TokenType, ITokenForgeFactoryAlert} from "./interfaces/ITokenForgeFactory.sol";
import {BaseForge} from "./BaseForge.sol";

abstract contract ERC20ForgeV1 is BaseForge {
    using EnumerableSet for EnumerableSet.AddressSet;
    using ECDSA for bytes32;
    using SafeERC20 for IERC20;

    bytes32 private constant ERC20_MINT_TYPE_HASH = keccak256(
        "ERC20Mint(uint256 uuid, address recipient, address token, uint256 amount, uint256 nonce, uint256 deadline)"
    );
    bytes32 private constant ERC20_TRANSFER_TYPE_HASH = keccak256(
        "ERC20Transfer(uint256 uuid, address recipient, address token, uint256 amount, uint256 nonce, uint256 deadline)"
    );
    bytes32 private constant ERC20_BURN_TYPE_HASH = keccak256(
        "ERC20Burn(uint256 uuid, address from, address token, uint256 amount, uint256 nonce, uint256 deadline)"
    );

    function mintERC20(uint256 uuid, address token, uint256 amount, uint256 deadline, bytes calldata validatorSig)
        external
        checkDeadline(deadline)
    {
        address recipient = _msgSender();
        bytes32 structHash =
            keccak256(abi.encode(ERC20_MINT_TYPE_HASH, uuid, recipient, token, amount, _useNonce(recipient), deadline));

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC20Forge(token).mint(recipient, amount);
        _alertMintToFactory(TokenType.ERC20, uuid, token, abi.encode(recipient, amount));
    }

    function transferERC20(uint256 uuid, address token, uint256 amount, uint256 deadline, bytes calldata validatorSig)
        external
        checkDeadline(deadline)
    {
        address recipient = _msgSender();
        bytes32 structHash = keccak256(
            abi.encode(ERC20_TRANSFER_TYPE_HASH, uuid, recipient, token, amount, _useNonce(recipient), deadline)
        );

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC20(token).safeTransfer(recipient, amount);
        _alertTransferToFactory(TokenType.ERC20, uuid, token, abi.encode(recipient, amount));
    }

    // recipient: msg.sender
    function burnERC20(uint256 uuid, address token, uint256 amount, uint256 deadline, bytes calldata validatorSig)
        external
        checkDeadline(deadline)
    {
        address from = _msgSender();
        bytes32 structHash =
            keccak256(abi.encode(ERC20_BURN_TYPE_HASH, uuid, from, token, amount, _useNonce(from), deadline));

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC20Forge(token).burnFrom(from, amount);
        _alertBurnToFactory(TokenType.ERC20, uuid, token, abi.encode(from, amount));
    }
}

abstract contract ERC721ForgeV1 is BaseForge, ERC721HolderUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;
    using ECDSA for bytes32;

    error ERC721MintForge__InvalidRecipientSignature(address recipient);

    bytes32 private constant ERC721_MINT_TYPE_HASH = keccak256(
        "ERC721Mint(uint256 uuid, address recipient, address token, uint256 tokenID, uint256 nonce, uint256 deadline)"
    );
    bytes32 private constant ERC721_TRANSFER_TYPE_HASH = keccak256(
        "ERC721Transfer(uint256 uuid, address recipient, address token, uint256 tokenID, uint256 nonce, uint256 deadline)"
    );
    bytes32 private constant ERC721_BURN_TYPE_HASH = keccak256(
        "ERC721Burn(uint256 uuid, address from, address token, uint256 tokenID, uint256 nonce, uint256 deadline)"
    );

    function mintERC721(uint256 uuid, address token, uint256 tokenID, uint256 deadline, bytes calldata validatorSig)
        external
        checkDeadline(deadline)
    {
        address recipient = _msgSender();
        bytes32 structHash = keccak256(
            abi.encode(ERC721_MINT_TYPE_HASH, uuid, recipient, token, tokenID, _useNonce(recipient), deadline)
        );

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC721Forge(token).mint(recipient, tokenID);
        _alertMintToFactory(TokenType.ERC721, uuid, token, abi.encode(recipient, tokenID));
    }

    function transferERC721(
        uint256 uuid,
        address token,
        uint256 tokenID,
        uint256 deadline,
        bytes calldata validatorSig,
        bytes calldata data
    ) external checkDeadline(deadline) {
        address recipient = _msgSender();
        bytes32 structHash = keccak256(
            abi.encode(ERC721_TRANSFER_TYPE_HASH, uuid, recipient, token, tokenID, _useNonce(recipient), deadline)
        );

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC721Forge(token).safeTransferFrom(address(this), recipient, tokenID, data);
        _alertMintToFactory(TokenType.ERC721, uuid, token, abi.encode(recipient, tokenID));
    }

    function burnERC721(uint256 uuid, address token, uint256 tokenID, uint256 deadline, bytes calldata validatorSig)
        external
        checkDeadline(deadline)
    {
        address from = _msgSender();
        bytes32 structHash =
            keccak256(abi.encode(ERC721_BURN_TYPE_HASH, uuid, from, token, tokenID, _useNonce(from), deadline));

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC721Forge(token).burnFrom(from, tokenID);
        _alertBurnToFactory(TokenType.ERC721, uuid, token, abi.encode(from, tokenID));
    }
}

abstract contract ERC1155ForgeV1 is BaseForge, ERC1155HolderUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;
    using ECDSA for bytes32;

    bytes32 private constant ERC1155_MINT_TYPE_HASH = keccak256(
        "ERC1155Mint(uint256 uuid, address recipient, address token, uint256 tokenID, uint256 amount, uint256 nonce, uint256 deadline)"
    );
    bytes32 private constant ERC1155_TRANSFER_TYPE_HASH = keccak256(
        "ERC1155Transfer(uint256 uuid, address recipient, address token, uint256 tokenID, uint256 amount, uint256 nonce, uint256 deadline)"
    );
    bytes32 private constant ERC1155_BURN_TYPE_HASH = keccak256(
        "ERC1155Burn(uint256 uuid, address from, address token, uint256 tokenID, uint256 amount, uint256 nonce, uint256 deadline)"
    );

    bytes32 private constant ERC1155_BATCH_MINT_TYPE_HASH = keccak256(
        "ERC1155BatchMint(uint256 uuid, address recipient, address token, uint256[] tokenIDs, uint256[] amounts, uint256 nonce, uint256 deadline)"
    );
    bytes32 private constant ERC1155_BATCH_TRANSFER_TYPE_HASH = keccak256(
        "ERC1155BatchTransfer(uint256 uuid, address recipient, address token, uint256[] tokenIDs, uint256[] amounts, uint256 nonce, uint256 deadline)"
    );
    bytes32 private constant ERC1155_BATCH_BURN_TYPE_HASH = keccak256(
        "ERC1155BatchBurn(uint256 uuid, address from, address token, uint256[] tokenIDs, uint256[] amounts, uint256 nonce, uint256 deadline)"
    );

    function mintERC1155(
        uint256 uuid,
        address token,
        uint256 tokenID,
        uint256 amount,
        uint256 deadline,
        bytes calldata validatorSig,
        bytes calldata data
    ) external checkDeadline(deadline) {
        address recipient = _msgSender();
        {
            bytes32 structHash = keccak256(
                abi.encode(
                    ERC1155_MINT_TYPE_HASH, uuid, recipient, token, tokenID, amount, _useNonce(recipient), deadline
                )
            );

            bytes32 hash = _hashTypedDataV4(structHash);
            _verifyValidatorSignature(hash, validatorSig);
        }

        IERC1155Forge(token).mint(recipient, tokenID, amount, data);
        _alertMintToFactory(TokenType.ERC1155, uuid, token, abi.encode(false, abi.encode(recipient, tokenID, amount)));
    }

    function transferERC1155(
        uint256 uuid,
        address token,
        uint256 tokenID,
        uint256 amount,
        uint256 deadline,
        bytes calldata validatorSig,
        bytes calldata data
    ) external checkDeadline(deadline) {
        address recipient = _msgSender();
        {
            bytes32 structHash = keccak256(
                abi.encode(
                    ERC1155_TRANSFER_TYPE_HASH, uuid, recipient, token, tokenID, amount, _useNonce(recipient), deadline
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
            keccak256(abi.encode(ERC1155_BURN_TYPE_HASH, uuid, from, token, tokenID, amount, _useNonce(from), deadline));

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC1155Forge(token).burnFrom(from, tokenID, amount);
        _alertBurnToFactory(TokenType.ERC1155, uuid, token, abi.encode(false, abi.encode(from, tokenID, amount)));
    }

    function mintERC1155Batch(
        uint256 uuid,
        address token,
        uint256[] memory tokenIDs,
        uint256[] memory amounts,
        uint256 deadline,
        bytes memory validatorSig,
        bytes memory data
    ) external checkDeadline(deadline) {
        address recipient = _msgSender();
        bytes32 structHash = keccak256(
            abi.encode(
                ERC1155_BATCH_MINT_TYPE_HASH, uuid, recipient, token, tokenIDs, amounts, _useNonce(recipient), deadline
            )
        );

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC1155Forge(token).mintBatch(recipient, tokenIDs, amounts, data);
        _alertMintToFactory(TokenType.ERC1155, uuid, token, abi.encode(true, abi.encode(recipient, tokenIDs, amounts)));
    }

    function transferERC1155Batch(
        uint256 uuid,
        address token,
        uint256[] memory tokenIDs,
        uint256[] memory amounts,
        uint256 deadline,
        bytes memory validatorSig,
        bytes memory data
    ) external checkDeadline(deadline) {
        address recipient = _msgSender();
        bytes32 structHash = keccak256(
            abi.encode(
                ERC1155_BATCH_TRANSFER_TYPE_HASH,
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
            abi.encode(ERC1155_BATCH_BURN_TYPE_HASH, uuid, from, token, tokenIDs, amounts, _useNonce(from), deadline)
        );

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC1155Forge(token).burnFromBatch(from, tokenIDs, amounts);
        _alertBurnToFactory(TokenType.ERC1155, uuid, token, abi.encode(true, abi.encode(from, tokenIDs, amounts)));
    }
}

contract ForgeV1 is ERC20ForgeV1, ERC721ForgeV1, ERC1155ForgeV1 {
// This contract serves as a version marker for the Forge contracts.
// It does not contain any additional logic or state.
// Future versions can inherit from this contract to maintain compatibility.
}
