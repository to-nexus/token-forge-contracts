// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC20Base} from "./ERC20Base.sol";

contract ERC20Fixed is ERC20Base {
    error ERC20Fixed__MintingNotAllowed();
    error ERC20Fixed__BurningNotAllowed();

    function __ERC20Base_init(
        address _owner,
        address _forge,
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        uint256 _initialSupply,
        bytes memory _data
    ) internal override {
        if (_initialSupply == 0) revert ERC20Base__NullInput("initialSupply");
        ERC20Base.__ERC20Base_init(_owner, _forge, _name, _symbol, _decimals, _initialSupply, _data);
    }

    function mint(address, uint256) public pure override {
        revert ERC20Fixed__MintingNotAllowed();
    }

    function burnFrom(address, uint256) public pure override {
        revert ERC20Fixed__BurningNotAllowed();
    }
}
