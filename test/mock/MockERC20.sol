// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {IERC20Forge} from "../../src/interfaces/IERC20Forge.sol";
import {AccessControl} from "@openzeppelin-contracts-5.3.0/access/AccessControl.sol";
import {ERC20, ERC20Permit} from "@openzeppelin-contracts-5.3.0/token/ERC20/extensions/ERC20Permit.sol";

contract MockERC20 is IERC20Forge, ERC20Permit, AccessControl {
    constructor(address forge) ERC20("mock20", "m20") ERC20Permit("mock20") {
        _grantRole(DEFAULT_ADMIN_ROLE, forge);
    }

    function forceMint(address to, uint256 amount) external {
        _mint(to, amount);
    }

    function mint(address to, uint256 amount) external override onlyRole(DEFAULT_ADMIN_ROLE) {
        _mint(to, amount);
    }

    function burnFrom(address from, uint256 amount) external override onlyRole(DEFAULT_ADMIN_ROLE) {
        if (from != msg.sender) _spendAllowance(from, msg.sender, amount);
        _burn(from, amount);
    }
}
