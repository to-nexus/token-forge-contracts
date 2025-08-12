// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC20Base} from "./ERC20Base.sol";
import {ERC20Capable} from "./extensions/ERC20Capable.sol";
import {ERC20PeriodMintLimit} from "./extensions/ERC20PeriodMintLimit.sol";
import {ERC20PeriodsMintLimit} from "./extensions/ERC20PeriodsMintLimit.sol";

contract ERC20SingleMintLimited is ERC20Base, ERC20Capable, ERC20PeriodMintLimit {
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
        (uint256 _cap, uint256 period, uint256 limit) = abi.decode(_data, (uint256, uint256, uint256));

        // Validate cap
        if (_cap < _initialSupply) revert ERC20SingleMintLimited__CapTooLow(_cap, _initialSupply);

        // Initialize parent contracts
        __ERC20PeriodMintLimit_init(period, limit);
        __ERC20Capable_init(_cap);
        __ERC20Base_init(_owner, _name, _symbol, _decimals, _initialSupply);
    }

    function mint(address to, uint256 amount) public virtual override(ERC20Base, ERC20PeriodMintLimit) {
        ERC20PeriodMintLimit.mint(to, amount);
    }

    function _update(address from, address to, uint256 value) internal virtual override(ERC20Base, ERC20Capable) {
        ERC20Capable._update(from, to, value);
    }
}
