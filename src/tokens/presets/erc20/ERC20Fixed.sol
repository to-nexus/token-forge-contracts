// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC20Base} from "./ERC20Base.sol";

contract ERC20Fixed is ERC20Base {
    error ERC20Fixed__MintingNotAllowed();
    error ERC20Fixed__BurningNotAllowed();

    function initialize(
        address owner,
        address manager,
        string memory name,
        string memory symbol,
        uint8 decimals,
        uint256 initialSupply,
        address initialRecipient,
        bytes memory
    ) external override initializer {
        if (initialSupply == 0) revert TokenBase__NullInput("initialSupply");
        ERC20Base.__ERC20Base_init(owner, manager, name, symbol, decimals, initialSupply, initialRecipient);
    }

    function mint(address, uint256) public pure override {
        revert ERC20Fixed__MintingNotAllowed();
    }

    function burnFrom(address, uint256) public pure override {
        revert ERC20Fixed__BurningNotAllowed();
    }
}
