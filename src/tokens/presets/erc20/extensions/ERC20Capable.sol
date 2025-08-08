// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {
    ERC20Upgradeable,
    ERC20CappedUpgradeable
} from "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC20/extensions/ERC20CappedUpgradeable.sol";
import {ERC20Base} from "../ERC20Base.sol";

abstract contract ERC20Capable is ERC20Upgradeable, ERC20Base, ERC20CappedUpgradeable {
    function __ERC20Capable_init(uint256 cap_) internal onlyInitializing {
        ERC20CappedUpgradeable.__ERC20Capped_init(cap_);
    }

    function remainingSupply() external view returns (uint256) {
        uint256 currentSupply = totalSupply();
        uint256 maxSupply = cap();
        unchecked {
            return maxSupply > currentSupply ? maxSupply - currentSupply : 0;
        }
    }

    function decimals() public view virtual override(ERC20Upgradeable, ERC20Base) returns (uint8) {
        return ERC20Base.decimals();
    }

    function _update(address from, address to, uint256 value)
        internal
        virtual
        override(ERC20Upgradeable, ERC20CappedUpgradeable)
    {
        ERC20CappedUpgradeable._update(from, to, value);
    }
}
