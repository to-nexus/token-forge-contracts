// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC20Base, ERC20Capable, ERC20PeriodMintLimit} from "./extensions/ERC20PeriodMintLimit.sol";

contract ERC20SingleMintLimited is ERC20PeriodMintLimit {
    error ERC20SingleMintLimited__InvalidInitialData();
    error ERC20SingleMintLimited__CapTooLow(uint256 cap, uint256 initialSupply);

    function initialize(
        address _owner,
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        uint256 _initialSupply,
        bytes memory _data
    ) external override initializer {
        // Decode cap from _data
        if (_data.length != 32 * 3) revert ERC20SingleMintLimited__InvalidInitialData();
        (uint256 cap_, uint256 period, uint256 limit) = abi.decode(_data, (uint256, uint256, uint256));

        // Validate cap
        if (cap_ < _initialSupply) revert ERC20SingleMintLimited__CapTooLow(cap_, _initialSupply);

        // Initialize parent contracts
        __ERC20Capable_init(cap_);
        __ERC20PeriodMintLimit_init(period, limit);
        __ERC20Base_init(_owner, _name, _symbol, _decimals, _initialSupply);
    }
}
