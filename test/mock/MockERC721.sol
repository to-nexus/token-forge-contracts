// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC721} from "@openzeppelin-contracts-5.3.0/token/ERC721/ERC721.sol";
import {Ownable} from "@openzeppelin-contracts-5.3.0/access/Ownable.sol";

contract MockERC721 is ERC721, Ownable {
    address public forge;

    constructor(address _forge) ERC721("MockERC721", "MCK") Ownable(_msgSender()) {
        forge = _forge;
    }

    function forceMint(address to, uint256 tokenID) external onlyOwner {
        _safeMint(to, tokenID);
    }

    function mint(address to, uint256 tokenID) external {
        require(_msgSender() == forge, "MockERC721: only forge can mint");
        _safeMint(to, tokenID);
    }

    function burnFrom(address from, uint256 tokenID) external {
        require(_msgSender() == forge, "MockERC721: only forge can burn");
        _checkAuthorized(from, _msgSender(), tokenID);
        _burn(tokenID);
    }
}
