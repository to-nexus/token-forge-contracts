// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

interface IERC721Forge {
    function mint(address to, uint256 tokenID, bytes memory data) external returns (uint256);
    function burnFrom(address from, uint256 tokenID) external;
}
