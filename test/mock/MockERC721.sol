// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {IERC721Forge} from "../../src/interfaces/IERC721Forge.sol";
import {AccessControl} from "@openzeppelin-contracts-5.3.0/access/AccessControl.sol";
import {ERC721} from "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract MockERC721 is IERC721Forge, ERC721("mock721", "m721"), AccessControl {
    constructor(address forge) {
        _grantRole(DEFAULT_ADMIN_ROLE, forge);
    }

    function forceMint(address to, uint256 tokenID, bytes memory) external {
        _safeMint(to, tokenID);
    }

    function mint(address to, uint256 tokenID, bytes memory data)
        external
        override
        onlyRole(DEFAULT_ADMIN_ROLE)
        returns (uint256)
    {
        _safeMint(to, tokenID, data);
        return tokenID;
    }

    function burnFrom(address from, uint256 tokenID) external override onlyRole(DEFAULT_ADMIN_ROLE) {
        require(_isAuthorized(from, msg.sender, tokenID), "NOT_AUTHORIZED");
        _burn(tokenID);
    }

    function supportsInterface(bytes4 interfaceId) public view virtual override(ERC721, AccessControl) returns (bool) {
        return super.supportsInterface(interfaceId);
    }
}
