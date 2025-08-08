// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC20Base, ERC20Capable} from "./extensions/ERC20Capable.sol";

contract ERC20Capped is ERC20Capable {
    error ERC20Capped__InvalidInitialData();
    error ERC20Capped__CapTooLow(uint256 cap, uint256 initialSupply);

    function initialize(
        address _owner,
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        uint256 _initialSupply,
        bytes memory _data
    ) external override initializer {
        // Decode cap from _data
        if (_data.length != 32) revert ERC20Capped__InvalidInitialData();
        uint256 cap_ = abi.decode(_data, (uint256));

        // Validate cap
        if (cap_ == 0) revert TokenBase__NullInput("cap");
        if (cap_ < _initialSupply) revert ERC20Capped__CapTooLow(cap_, _initialSupply);

        // Initialize parent contracts
        __ERC20Capable_init(cap_);
        __ERC20Base_init(_owner, _name, _symbol, _decimals, _initialSupply);
    }
}
