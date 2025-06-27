// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import {IERC721} from "@openzeppelin-contracts-5.3.0/token/ERC721/IERC721.sol";

interface IERC721Forge is IERC721 {
    function mint(address to, uint256 tokenID) external;
    function burnFrom(address from, uint256 tokenID) external;
}
