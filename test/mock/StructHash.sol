// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

// V1
// ERC20
bytes32 constant ERC20_MINT_TYPE_HASH_V1 = keccak256(
    "ERC20Mint(address recipient,address token,uint256 amount,address feeRecipient,uint256 feeBPS,uint256 nonce,uint256 deadline)"
);
bytes32 constant ERC20_TRANSFER_TYPE_HASH_V1 = keccak256(
    "ERC20Transfer(address recipient,address token,uint256 amount,address feeRecipient,uint256 feeBPS,uint256 nonce,uint256 deadline)"
);
bytes32 constant ERC20_TRANSFER_FROM_TYPE_HASH_V1 = keccak256(
    "ERC20TransferFrom(address from,address token,uint256 amount,address feeRecipient,uint256 feeBPS,uint256 nonce,uint256 deadline)"
);
bytes32 constant ERC20_BURN_TYPE_HASH_V1 = keccak256(
    "ERC20Burn(address from,address token,uint256 amount,address feeRecipient,uint256 feeBPS,uint256 nonce,uint256 deadline)"
);
// ERC721
bytes32 constant ERC721_MINT_TYPE_HASH_V1 =
    keccak256("ERC721Mint(address recipient,address token,uint256 tokenID,uint256 nonce,uint256 deadline)");
bytes32 constant ERC721_TRANSFER_TYPE_HASH_V1 = keccak256(
    "ERC721Transfer(address recipient,address token,uint256 tokenID,uint256 nonce,uint256 deadline,bytes data)"
);
bytes32 constant ERC721_BURN_TYPE_HASH_V1 =
    keccak256("ERC721Burn(address from,address token,uint256 tokenID,uint256 nonce,uint256 deadline)");
// ERC1155
bytes32 constant ERC1155_MINT_TYPE_HASH_V1 = keccak256(
    "ERC1155Mint(address recipient,address token,uint256 tokenID,uint256 amount,uint256 nonce,uint256 deadline,bytes data)"
);
bytes32 constant ERC1155_TRANSFER_TYPE_HASH_V1 = keccak256(
    "ERC1155Transfer(address recipient,address token,uint256 tokenID,uint256 amount,uint256 nonce,uint256 deadline,bytes data)"
);
bytes32 constant ERC1155_BURN_TYPE_HASH_V1 =
    keccak256("ERC1155Burn(address from,address token,uint256 tokenID,uint256 amount,uint256 nonce,uint256 deadline)");
bytes32 constant ERC1155_MINT_BATCH_TYPE_HASH_V1 = keccak256(
    "ERC1155BatchMint(address recipient,address token,uint256[] tokenIDs,uint256[] amounts,uint256 nonce,uint256 deadline,bytes data)"
);
bytes32 constant ERC1155_TRANSFER_BATCH_TYPE_HASH_V1 = keccak256(
    "ERC1155BatchTransfer(address recipient,address token,uint256[] tokenIDs,uint256[] amounts,uint256 nonce,uint256 deadline,bytes data)"
);
bytes32 constant ERC1155_BURN_BATCH_TYPE_HASH_V1 = keccak256(
    "ERC1155BatchBurn(address from,address token,uint256[] tokenIDs,uint256[] amounts,uint256 nonce,uint256 deadline)"
);

// V2
// ERC20
bytes32 constant ERC20_MINT_TYPE_HASH_V2 = keccak256(
    "ERC20Mint(address token,uint256 amount,address feeRecipient,uint256 feeBPS,uint256 nonce,uint256 deadline)"
);
bytes32 constant ERC20_VALIDATOR_MINT_TYPE_HASH_V2 =
    keccak256("ValidatorERC20Mint(address recipient,bytes recipientSig)");

bytes32 constant ERC20_TRANSFER_TYPE_HASH_V2 = keccak256(
    "ERC20Transfer(address token,uint256 amount,address feeRecipient,uint256 feeBPS,uint256 nonce,uint256 deadline)"
);
bytes32 constant ERC20_VALIDATOR_TRANSFER_TYPE_HASH_V2 =
    keccak256("ValidatorERC20Transfer(address recipient,bytes recipientSig)");

bytes32 constant ERC20_TRANSFER_FROM_TYPE_HASH_V2 = keccak256(
    "ERC20TransferFrom(address token,uint256 amount,address feeRecipient,uint256 feeBPS,uint256 nonce,uint256 deadline)"
);
bytes32 constant ERC20_VALIDATOR_TRANSFER_FROM_TYPE_HASH_V2 =
    keccak256("ValidatorERC20TransferFrom(address from,bytes fromSig)");

bytes32 constant ERC20_BURN_TYPE_HASH_V2 = keccak256(
    "ERC20Burn(address token,uint256 amount,address feeRecipient,uint256 feeBPS,uint256 nonce,uint256 deadline)"
);
bytes32 constant ERC20_VALIDATOR_BURN_TYPE_HASH_V2 = keccak256("ValidatorERC20Burn(address from,bytes fromSig)");

// ERC721
bytes32 constant ERC721_MINT_TYPE_HASH_V2 =
    keccak256("ERC721Mint(address token,uint256 tokenID,uint256 nonce,uint256 deadline)");
bytes32 constant ERC721_VALIDATOR_MINT_TYPE_HASH_V2 =
    keccak256("ValidatorERC721Mint(address recipient,bytes recipientSig)");

bytes32 constant ERC721_TRANSFER_TYPE_HASH_V2 =
    keccak256("ERC721Transfer(address token,uint256 tokenID,uint256 nonce,uint256 deadline,bytes data)");
bytes32 constant ERC721_VALIDATOR_TRANSFER_TYPE_HASH_V2 =
    keccak256("ValidatorERC721Transfer(address recipient,bytes recipientSig)");

bytes32 constant ERC721_BURN_TYPE_HASH_V2 =
    keccak256("ERC721Burn(address token,uint256 tokenID,uint256 nonce,uint256 deadline)");
bytes32 constant ERC721_VALIDATOR_BURN_TYPE_HASH_V2 = keccak256("ValidatorERC721Burn(address from,bytes fromSig)");

// ERC1155
bytes32 constant ERC1155_MINT_TYPE_HASH_V2 =
    keccak256("ERC1155Mint(address token,uint256 tokenID,uint256 amount,uint256 nonce,uint256 deadline,bytes data)");
bytes32 constant ERC1155_VALIDATOR_MINT_TYPE_HASH_V2 =
    keccak256("ValidatorERC1155Mint(address recipient,bytes recipientSig)");

bytes32 constant ERC1155_TRANSFER_TYPE_HASH_V2 =
    keccak256("ERC1155Transfer(address token,uint256 tokenID,uint256 amount,uint256 nonce,uint256 deadline,bytes data)");
bytes32 constant ERC1155_VALIDATOR_TRANSFER_TYPE_HASH_V2 =
    keccak256("ValidatorERC1155Transfer(address recipient,bytes recipientSig)");

bytes32 constant ERC1155_BURN_TYPE_HASH_V2 =
    keccak256("ERC1155Burn(address token,uint256 tokenID,uint256 amount,uint256 nonce,uint256 deadline)");
bytes32 constant ERC1155_VALIDATOR_BURN_TYPE_HASH_V2 = keccak256("ValidatorERC1155Burn(address from,bytes fromSig)");

bytes32 constant ERC1155_MINT_BATCH_TYPE_HASH_V2 = keccak256(
    "ERC1155BatchMint(address token,uint256[] tokenIDs,uint256[] amounts,uint256 nonce,uint256 deadline,bytes data)"
);
bytes32 constant ERC1155_VALIDATOR_MINT_BATCH_TYPE_HASH_V2 =
    keccak256("ValidatorERC1155BatchMint(address recipient,bytes recipientSig)");
bytes32 constant ERC1155_TRANSFER_BATCH_TYPE_HASH_V2 = keccak256(
    "ERC1155BatchTransfer(address token,uint256[] tokenIDs,uint256[] amounts,uint256 nonce,uint256 deadline,bytes data)"
);
bytes32 constant ERC1155_VALIDATOR_TRANSFER_BATCH_TYPE_HASH_V2 =
    keccak256("ValidatorERC1155BatchTransfer(address recipient,bytes recipientSig)");
bytes32 constant ERC1155_BURN_BATCH_TYPE_HASH_V2 =
    keccak256("ERC1155BatchBurn(address token,uint256[] tokenIDs,uint256[] amounts,uint256 nonce,uint256 deadline)");
bytes32 constant ERC1155_VALIDATOR_BURN_BATCH_TYPE_HASH_V2 =
    keccak256("ValidatorERC1155BatchBurn(address from,bytes fromSig)");

// V3 == V1
// ERC20
bytes32 constant ERC20_MINT_TYPE_HASH_V3 = ERC20_MINT_TYPE_HASH_V1;
bytes32 constant ERC20_TRANSFER_TYPE_HASH_V3 = ERC20_TRANSFER_TYPE_HASH_V1;
bytes32 constant ERC20_TRANSFER_FROM_TYPE_HASH_V3 = ERC20_TRANSFER_FROM_TYPE_HASH_V1;
bytes32 constant ERC20_BURN_TYPE_HASH_V3 = ERC20_BURN_TYPE_HASH_V1;
// ERC721
bytes32 constant ERC721_MINT_TYPE_HASH_V3 = ERC721_MINT_TYPE_HASH_V1;
bytes32 constant ERC721_TRANSFER_TYPE_HASH_V3 = ERC721_TRANSFER_TYPE_HASH_V1;
bytes32 constant ERC721_BURN_TYPE_HASH_V3 = ERC721_BURN_TYPE_HASH_V1;
// ERC1155
bytes32 constant ERC1155_MINT_TYPE_HASH_V3 = ERC1155_MINT_TYPE_HASH_V1;
bytes32 constant ERC1155_TRANSFER_TYPE_HASH_V3 = ERC1155_TRANSFER_TYPE_HASH_V1;
bytes32 constant ERC1155_BURN_TYPE_HASH_V3 = ERC1155_BURN_TYPE_HASH_V1;

bytes32 constant ERC1155_MINT_BATCH_TYPE_HASH_V3 = ERC1155_MINT_BATCH_TYPE_HASH_V1;
bytes32 constant ERC1155_TRANSFER_BATCH_TYPE_HASH_V3 = ERC1155_TRANSFER_BATCH_TYPE_HASH_V1;
bytes32 constant ERC1155_BURN_BATCH_TYPE_HASH_V3 = ERC1155_BURN_BATCH_TYPE_HASH_V1;

// ERC20 Permit
bytes32 constant PERMIT_TYPE_HASH =
    keccak256("Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)");
