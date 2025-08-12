// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC721Base} from "../../src/tokens/presets/erc721/ERC721Base.sol";

contract MockERC721 is ERC721Base {
    function initialize(
        address owner,
        string memory name,
        string memory symbol,
        string memory baseTokenURI,
        bytes memory
    ) external virtual override initializer {
        __ERC721Base_init(owner, name, symbol, baseTokenURI);
    }

    function forceMint(address to, uint256 tokenID) external onlyOwner {
        _safeMint(to, tokenID);
    }
}
