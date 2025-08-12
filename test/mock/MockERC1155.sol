// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1155Base} from "../../src/tokens/presets/erc1155/ERC1155Base.sol";

contract MockERC1155 is ERC1155Base {
    function forceMint(address to, uint256 tokenID, uint256 amount) external onlyRole(DEFAULT_ADMIN_ROLE) {
        _mint(to, tokenID, amount, "");
    }

    function forceMintBatch(address to, uint256[] memory tokenIDs, uint256[] memory amounts)
        external
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        _mintBatch(to, tokenIDs, amounts, "");
    }
}
