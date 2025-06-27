// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

import {IERC20} from "@openzeppelin-contracts-5.3.0/token/ERC20/IERC20.sol";
import {IERC20Metadata} from "@openzeppelin-contracts-5.3.0/token/ERC20/extensions/IERC20Metadata.sol";

interface IERC20Forge is IERC20, IERC20Metadata {
    function mint(address to, uint256 amount) external;
    function burnFrom(address from, uint256 amount) external;
}
