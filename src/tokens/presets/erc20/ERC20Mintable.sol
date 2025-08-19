// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC20Base} from "./ERC20Base.sol";

contract ERC20Mintable is ERC20Base {
    function initialize(
        address owner,
        address manager,
        string memory name,
        string memory symbol,
        uint8 decimals,
        uint256 initialSupply,
        address initialRecipient,
        bytes memory
    ) external override initializer {
        __ERC20Base_init(owner, manager, name, symbol, decimals, initialSupply, initialRecipient);
    }
}
