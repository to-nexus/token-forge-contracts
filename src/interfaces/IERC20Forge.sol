// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

interface IERC20Forge {
    function initialize(
        address owner,
        address manager,
        string calldata name,
        string calldata symbol,
        uint8 decimals,
        uint256 initialSupply,
        bytes calldata data
    ) external;

    function mint(address to, uint256 amount) external;
    function burnFrom(address from, uint256 amount) external;
}
