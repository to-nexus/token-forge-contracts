// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {
    ERC20Upgradeable,
    ERC20CappedUpgradeable
} from "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC20/extensions/ERC20CappedUpgradeable.sol";
import {ERC20Base} from "./ERC20Base.sol";

contract ERC20Capped is ERC20Upgradeable, ERC20Base, ERC20CappedUpgradeable {
    error ERC20Capped__InvalidCapData();
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
        if (_data.length != 32) revert ERC20Capped__InvalidCapData();
        uint256 cap_ = abi.decode(_data, (uint256));

        // Validate cap
        if (cap_ == 0) revert ERC20InvalidCap(0);
        if (cap_ < _initialSupply) revert ERC20Capped__CapTooLow(cap_, _initialSupply);

        // Initialize parent contracts
        __ERC20Base_init(_owner, _name, _symbol, _decimals, _initialSupply);
        __ERC20Capped_init(cap_);
    }

    function remainingSupply() external view returns (uint256) {
        uint256 currentSupply = totalSupply();
        uint256 maxSupply = cap();
        unchecked {
            return maxSupply > currentSupply ? maxSupply - currentSupply : 0;
        }
    }

    function decimals() public view override(ERC20Upgradeable, ERC20Base) returns (uint8) {
        return ERC20Base.decimals();
    }

    function isCapReached() external view returns (bool) {
        return totalSupply() >= cap();
    }

    function _update(address from, address to, uint256 value)
        internal
        override(ERC20Upgradeable, ERC20CappedUpgradeable)
    {
        ERC20CappedUpgradeable._update(from, to, value);
    }
}
