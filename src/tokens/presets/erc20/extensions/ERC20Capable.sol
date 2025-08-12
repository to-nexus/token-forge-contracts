// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC20Base} from "../ERC20Base.sol";

abstract contract ERC20Capable is ERC20Base {
    /**
     * @dev Total supply cap has been exceeded.
     */
    error ERC20Capable__ERC20ExceededCap(uint256 increasedSupply, uint256 cap);

    // keccak256(abi.encode(uint256(keccak256("cross.storage.forge.erc20.ERC20Capable")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant ERC20CapableStorageLocation =
        0x8007109065499a57d273132fba2e55750074b17ed6f6893f5ca2c5852612b300;

    function __ERC20Capable_init(uint256 cap_) internal onlyInitializing {
        if (cap_ == 0) {
            revert TokenBase__NullInput("cap");
        }
        assembly {
            sstore(ERC20CapableStorageLocation, cap_)
        }
    }

    function remainingSupply() external view returns (uint256) {
        uint256 currentSupply = totalSupply();
        uint256 maxSupply = cap();
        unchecked {
            return maxSupply > currentSupply ? maxSupply - currentSupply : 0;
        }
    }

    function cap() public view returns (uint256) {
        uint256 _cap;
        assembly {
            _cap := sload(ERC20CapableStorageLocation)
        }
        return _cap;
    }

    function _update(address from, address to, uint256 value) internal virtual override {
        super._update(from, to, value);

        if (from == address(0)) {
            uint256 maxSupply = cap();
            uint256 supply = totalSupply();
            if (supply > maxSupply) {
                revert ERC20Capable__ERC20ExceededCap(supply, maxSupply);
            }
        }
    }
}
