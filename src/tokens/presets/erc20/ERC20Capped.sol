// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC20Base, ERC20Capable} from "./extensions/ERC20Capable.sol";

contract ERC20Capped is ERC20Capable {
    error ERC20Capped__InvalidCapData();
    error ERC20Capped__CapTooLow(uint256 cap, uint256 initialSupply);

    function initialize(
        address _owner,
        address manager,
        string memory name,
        string memory symbol,
        uint8 decimals,
        uint256 initialSupply,
        address initialRecipient,
        bytes memory data
    ) external override initializer {
        __ERC20Base_init(_owner, manager, name, symbol, decimals, initialSupply, initialRecipient);

        // Decode cap from _data
        if (data.length != 32) revert ERC20Capped__InvalidCapData();
        uint256 cap_ = abi.decode(data, (uint256));

        // Validate cap
        if (cap_ == 0) revert TokenBase__NullInput("cap");
        if (cap_ < initialSupply) revert ERC20Capped__CapTooLow(cap_, initialSupply);

        // Initialize parent contracts
        __ERC20Capable_init(cap_);
    }
}
