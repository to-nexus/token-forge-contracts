// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

interface IERC20Forge {
    function initialize(
        address owner,
        address manager,
        string memory name,
        string memory symbol,
        uint8 decimals,
        uint256 initialSupply,
        address initialRecipient,
        bytes memory data
    ) external;

    function mint(address to, uint256 amount) external;
    function burnFrom(address from, uint256 amount) external;
}
