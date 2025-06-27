// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

bytes32 constant SPEND_NONCE_TYPEHASH =
    keccak256("SpendNonce(uint256 uuid, address spender, uint256 nonce, uint256 deadline)");

// ERC20
bytes32 constant ERC20_MINT_TYPEHASH = keccak256(
    "ERC20Mint(uint256 uuid, address recipient, address token, uint256 amount, uint256 nonce, uint256 deadline)"
);
bytes32 constant ERC20_MINTTO_TYPEHASH =
    keccak256("ERC20MintTo(address token, uint256 amount, uint256 nonce, uint256 deadline)");
bytes32 constant ERC20_VALIDATOR_MINTTO_TYPEHASH =
    keccak256("ValidatorERC20MintTo(uint256 uuid, address recipient, bytes recipientSig)");

bytes32 constant ERC20_TRANSFER_TYPEHASH = keccak256(
    "ERC20Transfer(uint256 uuid, address recipient, address token, uint256 amount, uint256 nonce, uint256 deadline)"
);
bytes32 constant ERC20_TRANSFERTO_TYPEHASH =
    keccak256("ERC20TransferTo(address token, uint256 amount, uint256 nonce, uint256 deadline)");
bytes32 constant ERC20_VALIDATOR_TRANSFERTO_TYPEHASH =
    keccak256("ValidatorERC20TransferTo(uint256 uuid, address recipient, bytes recipientSig)");

bytes32 constant ERC20_BURN_TYPEHASH =
    keccak256("ERC20Burn(uint256 uuid, address from, address token, uint256 amount, uint256 nonce, uint256 deadline)");
bytes32 constant ERC20_BURNFROM_TYPEHASH =
    keccak256("ERC20BurnFrom(address token, uint256 amount, uint256 nonce, uint256 deadline)");
bytes32 constant ERC20_VALIDATOR_BURNFROM_TYPEHASH =
    keccak256("ValidatorERC20BurnFrom(uint256 uuid, address from, bytes fromSig)");

// ERC721
bytes32 constant ERC721_MINT_TYPEHASH = keccak256(
    "ERC721Mint(uint256 uuid, address recipient, address token, uint256 tokenID, uint256 nonce, uint256 deadline)"
);
bytes32 constant ERC721_MINTTO_TYPEHASH =
    keccak256("ERC721MintTo(address token, uint256 tokenID, uint256 nonce, uint256 deadline)");
bytes32 constant ERC721_VALIDATOR_MINTTO_TYPEHASH =
    keccak256("ValidatorERC721MintTo(uint256 uuid, address recipient, bytes recipientSig)");

bytes32 constant ERC721_TRANSFER_TYPEHASH = keccak256(
    "ERC721Transfer(uint256 uuid, address recipient, address token, uint256 tokenID, uint256 nonce, uint256 deadline)"
);
bytes32 constant ERC721_TRANSFERTO_TYPEHASH =
    keccak256("ERC721TransferTo(address token, uint256 tokenID, uint256 nonce, uint256 deadline)");
bytes32 constant ERC721_VALIDATOR_TRANSFERTO_TYPEHASH =
    keccak256("ValidatorERC721TransferTo(uint256 uuid, address recipient, bytes recipientSig)");

bytes32 constant ERC721_BURN_TYPEHASH =
    keccak256("ERC721Burn(uint256 uuid, address from, address token, uint256 tokenID, uint256 nonce, uint256 deadline)");
bytes32 constant ERC721_BURNFROM_TYPEHASH =
    keccak256("ERC721BurnFrom(address token, uint256 tokenID, uint256 nonce, uint256 deadline)");
bytes32 constant ERC721_VALIDATOR_BURNFROM_TYPEHASH =
    keccak256("ValidatorERC721BurnFrom(uint256 uuid, address from, bytes fromSig)");

// ERC1155
bytes32 constant ERC1155_MINT_TYPEHASH = keccak256(
    "ERC1155Mint(uint256 uuid, address recipient, address token, uint256 tokenID, uint256 amount, uint256 nonce, uint256 deadline)"
);
bytes32 constant ERC1155_MINT_BATCH_TYPEHASH = keccak256(
    "ERC1155MintBatch(uint256 uuid, address recipient, address token, uint256[] tokenIDs, uint256[] amounts, uint256 nonce, uint256 deadline)"
);
bytes32 constant ERC1155_MINTTO_TYPEHASH =
    keccak256("ERC1155MintTo(address token, uint256 tokenID, uint256 amount, uint256 nonce, uint256 deadline)");
bytes32 constant ERC1155_MINTTO_BATCH_TYPEHASH = keccak256(
    "ERC1155MintToBatch(address token, uint256[] tokenIDs, uint256[] amounts, uint256 nonce, uint256 deadline)"
);
bytes32 constant ERC1155_VALIDATOR_MINTTO_TYPEHASH =
    keccak256("ValidatorERC1155MintTo(uint256 uuid, address recipient, bytes recipientSig)");
bytes32 constant ERC1155_VALIDATOR_MINTTO_BATCH_TYPEHASH =
    keccak256("ValidatorERC1155MintToBatch(uint256 uuid, address recipient, bytes recipientSig)");

bytes32 constant ERC1155_TRANSFER_TYPEHASH = keccak256(
    "ERC1155Transfer(uint256 uuid, address recipient, address token, uint256 tokenID, uint256 amount, uint256 nonce, uint256 deadline)"
);
bytes32 constant ERC1155_TRANSFER_BATCH_TYPEHASH = keccak256(
    "ERC1155TransferBatch(uint256 uuid, address recipient, address token, uint256[] tokenIDs, uint256[] amounts, uint256 nonce, uint256 deadline)"
);

bytes32 constant ERC1155_TRANSFERTO_TYPEHASH =
    keccak256("ERC1155TransferTo(address token, uint256 tokenID, uint256 amount, uint256 nonce, uint256 deadline)");
bytes32 constant ERC1155_TRANSFERTO_BATCH_TYPEHASH = keccak256(
    "ERC1155TransferToBatch(address token, uint256[] tokenIDs, uint256[] amounts, uint256 nonce, uint256 deadline)"
);
bytes32 constant ERC1155_VALIDATOR_TRANSFERTO_TYPEHASH =
    keccak256("ValidatorERC1155TransferTo(uint256 uuid, address recipient, bytes recipientSig)");
bytes32 constant ERC1155_VALIDATOR_TRANSFERTO_BATCH_TYPEHASH =
    keccak256("ValidatorERC1155TransferToBatch(uint256 uuid, address recipient, bytes recipientSig)");

bytes32 constant ERC1155_BURN_TYPEHASH = keccak256(
    "ERC1155Burn(uint256 uuid, address from, address token, uint256 tokenID, uint256 amount, uint256 nonce, uint256 deadline)"
);
bytes32 constant ERC1155_BURN_BATCH_TYPEHASH = keccak256(
    "ERC1155BurnBatch(uint256 uuid, address from, address token, uint256[] tokenIDs, uint256[] amounts, uint256 nonce, uint256 deadline)"
);

bytes32 constant ERC1155_BURNFROM_TYPEHASH =
    keccak256("ERC1155BurnFrom(address token, uint256 tokenID, uint256 amount, uint256 nonce, uint256 deadline)");
bytes32 constant ERC1155_BURNFROM_BATCH_TYPEHASH = keccak256(
    "ERC1155BurnFromBatch(address token, uint256[] tokenIDs, uint256[] amounts, uint256 nonce, uint256 deadline)"
);
bytes32 constant ERC1155_VALIDATOR_BURNFROM_TYPEHASH =
    keccak256("ValidatorERC1155BurnFrom(uint256 uuid, address from, bytes fromSig)");
bytes32 constant ERC1155_VALIDATOR_BURNFROM_BATCH_TYPEHASH =
    keccak256("ValidatorERC1155BurnFromBatch(uint256 uuid, address from, bytes fromSig)");
