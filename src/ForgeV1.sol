// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {EnumerableSet} from "@openzeppelin-contracts-5.3.0/utils/structs/EnumerableSet.sol";
import {NoncesUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/utils/NoncesUpgradeable.sol";
import {ERC721HolderUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC721/utils/ERC721HolderUpgradeable.sol";
import {ERC1155HolderUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC1155/utils/ERC1155HolderUpgradeable.sol";

import {IERC20, SafeERC20} from "@openzeppelin-contracts-5.3.0/token/ERC20/utils/SafeERC20.sol";
import {IERC721} from "@openzeppelin-contracts-5.3.0/token/ERC721/IERC721.sol";
import {IERC1155} from "@openzeppelin-contracts-5.3.0/token/ERC1155/IERC1155.sol";

import {IERC20Forge} from "./interfaces/IERC20Forge.sol";
import {IERC721Forge} from "./interfaces/IERC721Forge.sol";
import {IERC1155Forge} from "./interfaces/IERC1155Forge.sol";
import {TokenType, IForgeFactoryAlert} from "./interfaces/IForgeFactory.sol";
import {BaseForge} from "./BaseForge.sol";

abstract contract ERC20ForgeV1 is BaseForge {
    using EnumerableSet for EnumerableSet.AddressSet;
    using SafeERC20 for IERC20;

    bytes32 private constant ERC20_MINT_TYPE_HASH = keccak256(
        "ERC20Mint(address recipient,address token,uint256 amount,address feeRecipient,uint256 feeBPS,uint256 nonce,uint256 deadline)"
    );
    bytes32 private constant ERC20_TRANSFER_TYPE_HASH = keccak256(
        "ERC20Transfer(address recipient,address token,uint256 amount,address feeRecipient,uint256 feeBPS,uint256 nonce,uint256 deadline)"
    );
    bytes32 private constant ERC20_TRANSFER_FROM_TYPE_HASH = keccak256(
        "ERC20TransferFrom(address from,address token,uint256 amount,address feeRecipient,uint256 feeBPS,uint256 nonce,uint256 deadline)"
    );
    bytes32 private constant ERC20_BURN_TYPE_HASH = keccak256(
        "ERC20Burn(address from,address token,uint256 amount,address feeRecipient,uint256 feeBPS,uint256 nonce,uint256 deadline)"
    );

    function mintERC20(
        address token,
        uint256 amount,
        address feeRecipient,
        uint256 feeBPS,
        uint256 deadline,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) {
        address recipient = _msgSender();
        uint256 nonce = _useNonce(recipient);
        {
            bytes32 structHash = keccak256(
                abi.encode(ERC20_MINT_TYPE_HASH, recipient, token, amount, feeRecipient, feeBPS, nonce, deadline)
            );

            bytes32 hash = _hashTypedDataV4(structHash);
            _verifyValidatorSignature(hash, validatorSig);
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
        address token,
        uint256 amount,
        address feeRecipient,
        uint256 feeBPS,
        uint256 deadline,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) {
        address recipient = _msgSender();
        uint256 nonce = _useNonce(recipient);
        {
            bytes32 structHash = keccak256(
                abi.encode(ERC20_TRANSFER_TYPE_HASH, recipient, token, amount, feeRecipient, feeBPS, nonce, deadline)
            );

            bytes32 hash = _hashTypedDataV4(structHash);
            _verifyValidatorSignature(hash, validatorSig);
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
        address token,
        uint256 amount,
        address feeRecipient,
        uint256 feeBPS,
        uint256 deadline,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) {
        _transferFromERC20(token, amount, feeRecipient, feeBPS, deadline, validatorSig);
    }

    function transferFromERC20Permit(
        address token,
        uint256 amount,
        address feeRecipient,
        uint256 feeBPS,
        uint256 deadline,
        bytes calldata validatorSig,
        bytes memory permitSig
    ) external checkDeadline(deadline) erc20Permit(_msgSender(), token, amount, deadline, permitSig) {
        _transferFromERC20(token, amount, feeRecipient, feeBPS, deadline, validatorSig);
    }

    function _transferFromERC20(
        address token,
        uint256 amount,
        address feeRecipient,
        uint256 feeBPS,
        uint256 deadline,
        bytes calldata validatorSig
    ) private {
        address from = _msgSender();
        uint256 nonce = _useNonce(from);
        {
            bytes32 structHash = keccak256(
                abi.encode(ERC20_TRANSFER_FROM_TYPE_HASH, from, token, amount, feeRecipient, feeBPS, nonce, deadline)
            );

            bytes32 hash = _hashTypedDataV4(structHash);
            _verifyValidatorSignature(hash, validatorSig);
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

    // recipient: msg.sender
    function burnERC20(
        address token,
        uint256 amount,
        address feeRecipient,
        uint256 feeBPS,
        uint256 deadline,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) {
        _burnERC20(token, amount, feeRecipient, feeBPS, deadline, validatorSig);
    }

    // recipient: msg.sender
    function burnERC20Permit(
        address token,
        uint256 amount,
        address feeRecipient,
        uint256 feeBPS,
        uint256 deadline,
        bytes calldata validatorSig,
        bytes memory permitSig
    ) external checkDeadline(deadline) erc20Permit(_msgSender(), token, amount, deadline, permitSig) {
        _burnERC20(token, amount, feeRecipient, feeBPS, deadline, validatorSig);
    }

    function _burnERC20(
        address token,
        uint256 amount,
        address feeRecipient,
        uint256 feeBPS,
        uint256 deadline,
        bytes calldata validatorSig
    ) private {
        address from = _msgSender();
        uint256 nonce = _useNonce(from);
        {
            bytes32 structHash =
                keccak256(abi.encode(ERC20_BURN_TYPE_HASH, from, token, amount, feeRecipient, feeBPS, nonce, deadline));

            bytes32 hash = _hashTypedDataV4(structHash);
            _verifyValidatorSignature(hash, validatorSig);
        }

        (uint256 fee, uint256 value) = _erc20CalcFee(feeRecipient, feeBPS, amount);
        IERC20Forge(token).burnFrom(from, value);
        if (fee != 0) {
            IERC20(token).safeTransferFrom(from, feeRecipient, fee);
        }
        _alertBurnToFactory(TokenType.ERC20, _calcUUID(from, nonce), token, abi.encode(from, amount, feeRecipient, fee));
    }
}

abstract contract ERC721ForgeV1 is BaseForge, ERC721HolderUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;

    error ERC721MintForge__InvalidRecipientSignature(address recipient);

    bytes32 private constant ERC721_MINT_TYPE_HASH = keccak256(
        "ERC721Mint(address recipient,address token,uint256 tokenID,bytes data,uint256 nonce,uint256 deadline)"
    );
    bytes32 private constant ERC721_TRANSFER_TYPE_HASH = keccak256(
        "ERC721Transfer(address recipient,address token,uint256 tokenID,bytes data,uint256 nonce,uint256 deadline)"
    );
    bytes32 private constant ERC721_BURN_TYPE_HASH =
        keccak256("ERC721Burn(address from,address token,uint256 tokenID,uint256 nonce,uint256 deadline)");

    function mintERC721(
        address token,
        uint256 tokenID,
        bytes calldata data,
        uint256 deadline,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) {
        address recipient = _msgSender();
        uint256 nonce = _useNonce(recipient);
        bytes32 structHash =
            keccak256(abi.encode(ERC721_MINT_TYPE_HASH, recipient, token, tokenID, keccak256(data), nonce, deadline));

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC721Forge(token).mint(recipient, tokenID, data);
        _alertMintToFactory(TokenType.ERC721, _calcUUID(recipient, nonce), token, abi.encode(recipient, tokenID));
    }

    function transferERC721(
        address token,
        uint256 tokenID,
        bytes calldata data,
        uint256 deadline,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) {
        address recipient = _msgSender();
        uint256 nonce = _useNonce(recipient);
        {
            bytes32 structHash = keccak256(
                abi.encode(ERC721_TRANSFER_TYPE_HASH, recipient, token, tokenID, keccak256(data), nonce, deadline)
            );

            bytes32 hash = _hashTypedDataV4(structHash);
            _verifyValidatorSignature(hash, validatorSig);
        }
        IERC721(token).safeTransferFrom(address(this), recipient, tokenID, data);
        _alertTransferToFactory(TokenType.ERC721, _calcUUID(recipient, nonce), token, abi.encode(recipient, tokenID));
    }

    function burnERC721(address token, uint256 tokenID, uint256 deadline, bytes calldata validatorSig)
        external
        checkDeadline(deadline)
    {
        address from = _msgSender();
        uint256 nonce = _useNonce(from);
        bytes32 structHash = keccak256(abi.encode(ERC721_BURN_TYPE_HASH, from, token, tokenID, nonce, deadline));

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC721Forge(token).burnFrom(from, tokenID);
        _alertBurnToFactory(TokenType.ERC721, _calcUUID(from, nonce), token, abi.encode(from, tokenID));
    }
}

abstract contract ERC1155ForgeV1 is BaseForge, ERC1155HolderUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;

    bytes32 private constant ERC1155_MINT_TYPE_HASH = keccak256(
        "ERC1155Mint(address recipient,address token,uint256 tokenID,uint256 amount,bytes data,uint256 nonce,uint256 deadline)"
    );
    bytes32 private constant ERC1155_TRANSFER_TYPE_HASH = keccak256(
        "ERC1155Transfer(address recipient,address token,uint256 tokenID,uint256 amount,bytes data,uint256 nonce,uint256 deadline)"
    );
    bytes32 private constant ERC1155_BURN_TYPE_HASH = keccak256(
        "ERC1155Burn(address from,address token,uint256 tokenID,uint256 amount,uint256 nonce,uint256 deadline)"
    );

    bytes32 private constant ERC1155_BATCH_MINT_TYPE_HASH = keccak256(
        "ERC1155BatchMint(address recipient,address token,uint256[] tokenIDs,uint256[] amounts,bytes data,uint256 nonce,uint256 deadline)"
    );
    bytes32 private constant ERC1155_BATCH_TRANSFER_TYPE_HASH = keccak256(
        "ERC1155BatchTransfer(address recipient,address token,uint256[] tokenIDs,uint256[] amounts,bytes data,uint256 nonce,uint256 deadline)"
    );
    bytes32 private constant ERC1155_BATCH_BURN_TYPE_HASH = keccak256(
        "ERC1155BatchBurn(address from,address token,uint256[] tokenIDs,uint256[] amounts,uint256 nonce,uint256 deadline)"
    );

    function mintERC1155(
        address token,
        uint256 tokenID,
        uint256 amount,
        bytes calldata data,
        uint256 deadline,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) {
        address recipient = _msgSender();
        uint256 nonce = _useNonce(recipient);
        {
            bytes32 structHash = keccak256(
                abi.encode(ERC1155_MINT_TYPE_HASH, recipient, token, tokenID, amount, keccak256(data), nonce, deadline)
            );

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
        address token,
        uint256 tokenID,
        uint256 amount,
        bytes calldata data,
        uint256 deadline,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) {
        address recipient = _msgSender();
        uint256 nonce = _useNonce(recipient);
        {
            bytes32 structHash = keccak256(
                abi.encode(
                    ERC1155_TRANSFER_TYPE_HASH, recipient, token, tokenID, amount, keccak256(data), nonce, deadline
                )
            );

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

    function burnERC1155(address token, uint256 tokenID, uint256 amount, uint256 deadline, bytes calldata validatorSig)
        external
        checkDeadline(deadline)
    {
        address from = _msgSender();
        uint256 nonce = _useNonce(from);
        {
            bytes32 structHash =
                keccak256(abi.encode(ERC1155_BURN_TYPE_HASH, from, token, tokenID, amount, nonce, deadline));

            bytes32 hash = _hashTypedDataV4(structHash);
            _verifyValidatorSignature(hash, validatorSig);
        }
        IERC1155Forge(token).burnFrom(from, tokenID, amount);
        _alertBurnToFactory(
            TokenType.ERC1155, _calcUUID(from, nonce), token, abi.encode(false, abi.encode(from, tokenID, amount))
        );
    }

    function mintERC1155Batch(
        address token,
        uint256[] memory tokenIDs,
        uint256[] memory amounts,
        bytes calldata data,
        uint256 deadline,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) {
        address recipient = _msgSender();
        uint256 nonce = _useNonce(recipient);
        {
            bytes32 structHash = keccak256(
                abi.encode(
                    ERC1155_BATCH_MINT_TYPE_HASH,
                    recipient,
                    token,
                    keccak256(abi.encodePacked(tokenIDs)),
                    keccak256(abi.encodePacked(amounts)),
                    keccak256(data),
                    nonce,
                    deadline
                )
            );

            bytes32 hash = _hashTypedDataV4(structHash);
            _verifyValidatorSignature(hash, validatorSig);
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
        address token,
        uint256[] memory tokenIDs,
        uint256[] memory amounts,
        bytes calldata data,
        uint256 deadline,
        bytes calldata validatorSig
    ) external checkDeadline(deadline) {
        address recipient = _msgSender();
        uint256 nonce = _useNonce(recipient);
        {
            bytes32 structHash = keccak256(
                abi.encode(
                    ERC1155_BATCH_TRANSFER_TYPE_HASH,
                    recipient,
                    token,
                    keccak256(abi.encodePacked(tokenIDs)),
                    keccak256(abi.encodePacked(amounts)),
                    keccak256(data),
                    nonce,
                    deadline
                )
            );

            bytes32 hash = _hashTypedDataV4(structHash);
            _verifyValidatorSignature(hash, validatorSig);
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
        address token,
        uint256[] memory tokenIDs,
        uint256[] memory amounts,
        uint256 deadline,
        bytes memory validatorSig
    ) external checkDeadline(deadline) {
        address from = _msgSender();
        uint256 nonce = _useNonce(from);
        bytes32 structHash = keccak256(
            abi.encode(
                ERC1155_BATCH_BURN_TYPE_HASH,
                from,
                token,
                keccak256(abi.encodePacked(tokenIDs)),
                keccak256(abi.encodePacked(amounts)),
                nonce,
                deadline
            )
        );

        bytes32 hash = _hashTypedDataV4(structHash);
        _verifyValidatorSignature(hash, validatorSig);

        IERC1155Forge(token).burnFromBatch(from, tokenIDs, amounts);
        _alertBurnToFactory(
            TokenType.ERC1155, _calcUUID(from, nonce), token, abi.encode(true, abi.encode(from, tokenIDs, amounts))
        );
    }
}

contract ForgeV1 is ERC20ForgeV1, ERC721ForgeV1, ERC1155ForgeV1 {
// This contract serves as a version marker for the Forge contracts.
// It does not contain any additional logic or state.
// Future versions can inherit from this contract to maintain compatibility.
}
