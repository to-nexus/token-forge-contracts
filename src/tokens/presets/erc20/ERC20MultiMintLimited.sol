// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC20Base} from "./ERC20Base.sol";
import {ERC20Capable} from "./extensions/ERC20Capable.sol";
import {ERC20PeriodsMintLimit} from "./extensions/ERC20PeriodsMintLimit.sol";

contract ERC20MultiMintLimited is ERC20Base, ERC20Capable, ERC20PeriodsMintLimit {
    error ERC20MultiMintLimited__InvalidInitialData();
    error ERC20MultiMintLimited__CapTooLow(uint256 cap, uint256 initialSupply);

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
        if (data.length <= 32 * 3) revert ERC20MultiMintLimited__InvalidInitialData();
        (uint256 _cap, uint256[] memory durations, int256[] memory offsetSeconds, uint256[] memory limits) =
            abi.decode(data, (uint256, uint256[], int256[], uint256[]));

        // Validate cap
        if (_cap < initialSupply) revert ERC20MultiMintLimited__CapTooLow(_cap, initialSupply);

        // Initialize parent contracts
        __ERC20PeriodsMintLimit_init(durations, offsetSeconds, limits);
        __ERC20Capable_init(_cap);
    }

    function mint(address to, uint256 amount) public virtual override(ERC20Base, ERC20PeriodsMintLimit) {
        ERC20PeriodsMintLimit.mint(to, amount);
    }

    function _update(address from, address to, uint256 value) internal virtual override(ERC20Base, ERC20Capable) {
        ERC20Capable._update(from, to, value);
    }
}
