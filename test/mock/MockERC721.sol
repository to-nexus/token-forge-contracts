// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC721Base} from "../../src/tokens/presets/erc721/ERC721Base.sol";

contract MockERC721 is ERC721Base {
    function initialize(
        address owner,
        address manager,
        string memory name,
        string memory symbol,
        string memory baseTokenURI,
        bytes memory
    ) external virtual override initializer {
        ERC721Base.__ERC721Base_init(owner, manager, name, symbol, baseTokenURI);
    }

    function forceMint(address to, uint256 tokenID, bytes memory) external onlyRole(MANAGER_ROLE) {
        _safeMint(to, tokenID);
    }
}
