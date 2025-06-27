// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import {IERC1155} from "@openzeppelin-contracts-5.3.0/token/ERC1155/IERC1155.sol";

interface IERC1155Forge is IERC1155 {
    function mint(address to, uint256 tokenID, uint256 amount, bytes calldata data) external;
    function mintBatch(address to, uint256[] memory tokenIDs, uint256[] memory amounts, bytes calldata data) external;

    function burnFrom(address from, uint256 tokenID, uint256 amount) external;
    function burnFromBatch(address from, uint256[] memory tokenIDs, uint256[] memory amounts) external;
}
