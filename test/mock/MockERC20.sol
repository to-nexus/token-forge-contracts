// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC20Base} from "../../src/tokens/presets/erc20/ERC20Base.sol";
import {ERC20Permit} from "@openzeppelin-contracts-5.3.0/token/ERC20/extensions/ERC20Permit.sol";

contract MockERC20 is ERC20Base {
    function initialize(
        address _owner,
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        uint256 _initialSupply,
        bytes memory
    ) external virtual override initializer {
        __ERC20Base_init(_owner, _name, _symbol, _decimals, _initialSupply);
    }

    function forceMint(address to, uint256 amount) external onlyOwner {
        _mint(to, amount);
    }
}
