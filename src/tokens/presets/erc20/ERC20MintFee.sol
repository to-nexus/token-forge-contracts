// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC20Base} from "./ERC20Base.sol";
import {ERC20MintingFee} from "./extensions/ERC20MintingFee.sol";

contract ERC20MintFee is ERC20Base, ERC20MintingFee {
    function mint(address to, uint256 amount) public override(ERC20Base, ERC20MintingFee) {
        super.mint(to, amount);
    }
}
