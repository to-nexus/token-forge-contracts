// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC20Base} from "../../src/tokens/preset/ERC20Base.sol";

contract MockERC20 is ERC20Base {
    function forceMint(address to, uint256 amount) external onlyOwner {
        _mint(to, amount);
    }
}
