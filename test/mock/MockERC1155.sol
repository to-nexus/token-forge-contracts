// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1155Base} from "../../src/tokens/preset/ERC1155Base.sol";

contract MockERC1155 is ERC1155Base {
    function forceMint(address to, uint256 tokenID, uint256 amount) external onlyOwner {
        _mint(to, tokenID, amount, "");
    }

    function forceMintBatch(address to, uint256[] memory tokenIDs, uint256[] memory amounts) external onlyOwner {
        _mintBatch(to, tokenIDs, amounts, "");
    }
}
