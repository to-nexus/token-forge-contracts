// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {IERC1155Forge} from "../../src/interfaces/IERC1155Forge.sol";
import {AccessControl} from "@openzeppelin-contracts-5.3.0/access/AccessControl.sol";
import {ERC1155} from "@openzeppelin-contracts-5.3.0/token/ERC1155/ERC1155.sol";

contract MockERC1155 is IERC1155Forge, ERC1155("https://token-cdn-domain/{id}.json"), AccessControl {
    constructor(address forge) {
        _grantRole(DEFAULT_ADMIN_ROLE, forge);
    }

    function forceMint(address to, uint256 tokenID, uint256 amount) external {
        _mint(to, tokenID, amount, "");
    }

    function forceMintBatch(address to, uint256[] memory tokenIDs, uint256[] memory amounts) external {
        _mintBatch(to, tokenIDs, amounts, "");
    }

    function mint(address to, uint256 tokenID, uint256 amount, bytes calldata data)
        external
        override
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        _mint(to, tokenID, amount, data);
    }

    function mintBatch(address to, uint256[] memory tokenIDs, uint256[] memory amounts, bytes calldata data)
        external
        override
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        _mintBatch(to, tokenIDs, amounts, data);
    }

    function burnFrom(address from, uint256 tokenID, uint256 amount) external override onlyRole(DEFAULT_ADMIN_ROLE) {
        require(from == msg.sender || isApprovedForAll(from, msg.sender), "NOT_AUTHORIZED");
        _burn(from, tokenID, amount);
    }

    function burnFromBatch(address from, uint256[] memory tokenIDs, uint256[] memory amounts)
        external
        override
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        require(from == msg.sender || isApprovedForAll(from, msg.sender), "NOT_AUTHORIZED");
        _burnBatch(from, tokenIDs, amounts);
    }

    function supportsInterface(bytes4 interfaceId) public view virtual override(ERC1155, AccessControl) returns (bool) {
        return super.supportsInterface(interfaceId);
    }
}
