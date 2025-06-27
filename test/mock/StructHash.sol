// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

// V1
// ERC20
bytes32 constant ERC20_MINT_TYPE_HASH_V1 = keccak256(
    "ERC20Mint(uint256 uuid, address recipient, address token, uint256 amount, uint256 nonce, uint256 deadline)"
);
bytes32 constant ERC20_TRANSFER_TYPE_HASH_V1 = keccak256(
    "ERC20Transfer(uint256 uuid, address recipient, address token, uint256 amount, uint256 nonce, uint256 deadline)"
);
bytes32 constant ERC20_BURN_TYPE_HASH_V1 =
    keccak256("ERC20Burn(uint256 uuid, address from, address token, uint256 amount, uint256 nonce, uint256 deadline)");
// ERC721
bytes32 constant ERC721_MINT_TYPE_HASH_V1 = keccak256(
    "ERC721Mint(uint256 uuid, address recipient, address token, uint256 tokenID, uint256 nonce, uint256 deadline)"
);
bytes32 constant ERC721_TRANSFER_TYPE_HASH_V1 = keccak256(
    "ERC721Transfer(uint256 uuid, address recipient, address token, uint256 tokenID, uint256 nonce, uint256 deadline)"
);
bytes32 constant ERC721_BURN_TYPE_HASH_V1 =
    keccak256("ERC721Burn(uint256 uuid, address from, address token, uint256 tokenID, uint256 nonce, uint256 deadline)");
// ERC1155
bytes32 constant ERC1155_MINT_TYPE_HASH_V1 = keccak256(
    "ERC1155Mint(uint256 uuid, address recipient, address token, uint256 tokenID, uint256 amount, uint256 nonce, uint256 deadline)"
);
bytes32 constant ERC1155_TRANSFER_TYPE_HASH_V1 = keccak256(
    "ERC1155Transfer(uint256 uuid, address recipient, address token, uint256 tokenID, uint256 amount, uint256 nonce, uint256 deadline)"
);
bytes32 constant ERC1155_BURN_TYPE_HASH_V1 = keccak256(
    "ERC1155Burn(uint256 uuid, address from, address token, uint256 tokenID, uint256 amount, uint256 nonce, uint256 deadline)"
);
bytes32 constant ERC1155_MINT_BATCH_TYPE_HASH_V1 = keccak256(
    "ERC1155BatchMint(uint256 uuid, address recipient, address token, uint256[] tokenIDs, uint256[] amounts, uint256 nonce, uint256 deadline)"
);
bytes32 constant ERC1155_TRANSFER_BATCH_TYPE_HASH_V1 = keccak256(
    "ERC1155BatchTransfer(uint256 uuid, address recipient, address token, uint256[] tokenIDs, uint256[] amounts, uint256 nonce, uint256 deadline)"
);
bytes32 constant ERC1155_BURN_BATCH_TYPE_HASH_V1 = keccak256(
    "ERC1155BatchBurn(uint256 uuid, address from, address token, uint256[] tokenIDs, uint256[] amounts, uint256 nonce, uint256 deadline)"
);

// V2
// ERC20
bytes32 constant ERC20_MINT_TYPE_HASH_V2 =
    keccak256("ERC20Mint(address token, uint256 amount, uint256 nonce, uint256 deadline)");
bytes32 constant ERC20_VALIDATOR_MINT_TYPE_HASH_V2 =
    keccak256("ValidatorERC20Mint(uint256 uuid, address recipient, bytes recipientSig)");

bytes32 constant ERC20_TRANSFER_TYPE_HASH_V2 =
    keccak256("ERC20Transfer(address token, uint256 amount, uint256 nonce, uint256 deadline)");
bytes32 constant ERC20_VALIDATOR_TRANSFER_TYPE_HASH_V2 =
    keccak256("ValidatorERC20Transfer(uint256 uuid, address recipient, bytes recipientSig)");

bytes32 constant ERC20_BURN_TYPE_HASH_V2 =
    keccak256("ERC20Burn(address token, uint256 amount, uint256 nonce, uint256 deadline)");
bytes32 constant ERC20_VALIDATOR_BURN_TYPE_HASH_V2 =
    keccak256("ValidatorERC20Burn(uint256 uuid, address from, bytes fromSig)");

// ERC721
bytes32 constant ERC721_MINT_TYPE_HASH_V2 =
    keccak256("ERC721Mint(address token, uint256 tokenID, uint256 nonce, uint256 deadline)");
bytes32 constant ERC721_VALIDATOR_MINT_TYPE_HASH_V2 =
    keccak256("ValidatorERC721Mint(uint256 uuid, address recipient, bytes recipientSig)");

bytes32 constant ERC721_TRANSFER_TYPE_HASH_V2 =
    keccak256("ERC721Transfer(address token, uint256 tokenID, uint256 nonce, uint256 deadline)");
bytes32 constant ERC721_VALIDATOR_TRANSFER_TYPE_HASH_V2 =
    keccak256("ValidatorERC721Transfer(uint256 uuid, address recipient, bytes recipientSig)");

bytes32 constant ERC721_BURN_TYPE_HASH_V2 =
    keccak256("ERC721Burn(address token, uint256 tokenID, uint256 nonce, uint256 deadline)");
bytes32 constant ERC721_VALIDATOR_BURN_TYPE_HASH_V2 =
    keccak256("ValidatorERC721Burn(uint256 uuid, address from, bytes fromSig)");

// ERC1155
bytes32 constant ERC1155_MINT_TYPE_HASH_V2 =
    keccak256("ERC1155Mint(address token, uint256 tokenID, uint256 amount, uint256 nonce, uint256 deadline)");
bytes32 constant ERC1155_VALIDATOR_MINT_TYPE_HASH_V2 =
    keccak256("ValidatorERC1155Mint(uint256 uuid, address recipient, bytes recipientSig)");

bytes32 constant ERC1155_TRANSFER_TYPE_HASH_V2 =
    keccak256("ERC1155Transfer(address token, uint256 tokenID, uint256 amount, uint256 nonce, uint256 deadline)");
bytes32 constant ERC1155_VALIDATOR_TRANSFER_TYPE_HASH_V2 =
    keccak256("ValidatorERC1155Transfer(uint256 uuid, address recipient, bytes recipientSig)");

bytes32 constant ERC1155_BURN_TYPE_HASH_V2 =
    keccak256("ERC1155Burn(address token, uint256 tokenID, uint256 amount, uint256 nonce, uint256 deadline)");
bytes32 constant ERC1155_VALIDATOR_BURN_TYPE_HASH_V2 =
    keccak256("ValidatorERC1155Burn(uint256 uuid, address from, bytes fromSig)");

bytes32 constant ERC1155_MINT_BATCH_TYPE_HASH_V2 =
    keccak256("ERC1155BatchMint(address token, uint256[] tokenIDs, uint256[] amounts, uint256 nonce, uint256 deadline)");
bytes32 constant ERC1155_VALIDATOR_MINT_BATCH_TYPE_HASH_V2 =
    keccak256("ValidatorERC1155BatchMint(uint256 uuid, address recipient, bytes recipientSig)");
bytes32 constant ERC1155_TRANSFER_BATCH_TYPE_HASH_V2 = keccak256(
    "ERC1155BatchTransfer(address token, uint256[] tokenIDs, uint256[] amounts, uint256 nonce, uint256 deadline)"
);
bytes32 constant ERC1155_VALIDATOR_TRANSFER_BATCH_TYPE_HASH_V2 =
    keccak256("ValidatorERC1155BatchTransfer(uint256 uuid, address recipient, bytes recipientSig)");
bytes32 constant ERC1155_BURN_BATCH_TYPE_HASH_V2 =
    keccak256("ERC1155BatchBurn(address token, uint256[] tokenIDs, uint256[] amounts, uint256 nonce, uint256 deadline)");
bytes32 constant ERC1155_VALIDATOR_BURN_BATCH_TYPE_HASH_V2 =
    keccak256("ValidatorERC1155BatchBurn(uint256 uuid, address from, bytes fromSig)");

// V3 == V1
// ERC20
bytes32 constant ERC20_MINT_TYPE_HASH_V3 = ERC20_MINT_TYPE_HASH_V1;
bytes32 constant ERC20_TRANSFER_TYPE_HASH_V3 = ERC20_TRANSFER_TYPE_HASH_V1;
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
