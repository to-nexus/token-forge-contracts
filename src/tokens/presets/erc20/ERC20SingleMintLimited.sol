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
        address owner,
        address manager,
        string memory name,
        string memory symbol,
        uint8 decimals,
        uint256 initialSupply,
        address initialRecipient,
        bytes memory data
    ) external override initializer {
        __ERC20Base_init(owner, manager, name, symbol, decimals, initialSupply, initialRecipient);

        // Decode cap from _data
        if (data.length != 32 * 4) revert ERC20SingleMintLimited__InvalidInitialData();
        (uint256 _cap, uint256 duration, int256 offsetSeconds, uint256 limit) =
            abi.decode(data, (uint256, uint256, int256, uint256));

        // Validate cap
        if (_cap < initialSupply) revert ERC20SingleMintLimited__CapTooLow(_cap, initialSupply);

        // Initialize parent contracts
        __ERC20PeriodMintLimit_init(duration, offsetSeconds, limit);
        __ERC20Capable_init(_cap);
    }

    function mint(address to, uint256 amount) public virtual override(ERC20Base, ERC20PeriodMintLimit) {
        ERC20PeriodMintLimit.mint(to, amount);
    }

    function _update(address from, address to, uint256 value) internal virtual override(ERC20Base, ERC20Capable) {
        ERC20Capable._update(from, to, value);
    }
}
