// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

interface IERC1155Forge {
    function initialize(address owner, string calldata uri, bytes calldata data) external;

    function mint(address to, uint256 tokenID, uint256 amount, bytes calldata data) external;
    function mintBatch(address to, uint256[] memory tokenIDs, uint256[] memory amounts, bytes calldata data) external;

    function burnFrom(address from, uint256 tokenID, uint256 amount) external;
    function burnFromBatch(address from, uint256[] memory tokenIDs, uint256[] memory amounts) external;
}
