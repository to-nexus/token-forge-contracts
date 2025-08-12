// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC20Base} from "./ERC20Base.sol";
import {ERC20Capable} from "./extensions/ERC20Capable.sol";
import {ERC20PeriodsMintLimit} from "./extensions/ERC20PeriodsMintLimit.sol";

contract ERC20MultiMintLimited is ERC20Base, ERC20Capable, ERC20PeriodsMintLimit {
    error ERC20MultiMintLimited__InvalidInitialData();
    error ERC20MultiMintLimited__CapTooLow(uint256 cap, uint256 initialSupply);

    function initialize(
        address _owner,
        address _manager,
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        uint256 _initialSupply,
        bytes memory _data
    ) external override initializer {
        // Decode cap from _data
        if (_data.length <= 32 * 3) revert ERC20MultiMintLimited__InvalidInitialData();
        (uint256 _cap, uint256[] memory periods, uint256[] memory limits) =
            abi.decode(_data, (uint256, uint256[], uint256[]));

        // Validate cap
        if (_cap < _initialSupply) revert ERC20MultiMintLimited__CapTooLow(_cap, _initialSupply);

        // Initialize parent contracts
        __ERC20PeriodsMintLimit_init(periods, limits);
        __ERC20Capable_init(_cap);
        __ERC20Base_init(_owner, _manager, _name, _symbol, _decimals, _initialSupply);
    }

    function mint(address to, uint256 amount) public virtual override(ERC20Base, ERC20PeriodsMintLimit) {
        ERC20PeriodsMintLimit.mint(to, amount);
    }

    function _update(address from, address to, uint256 value) internal virtual override(ERC20Base, ERC20Capable) {
        ERC20Capable._update(from, to, value);
    }
}
