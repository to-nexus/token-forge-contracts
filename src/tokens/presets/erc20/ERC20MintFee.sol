// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC20Base} from "./ERC20Base.sol";
import {ERC20MintingFee} from "./extensions/ERC20MintingFee.sol";

contract ERC20MintFee is ERC20Base, ERC20MintingFee {
    function initialize(
        address _owner,
        address _forge,
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        uint256 _initialSupply,
        bytes memory _data
    ) external override initializer {
        ERC20Base.__ERC20Base_init(_owner, _forge, _name, _symbol, _decimals, _initialSupply);
        (address feeRecipient, uint256 feeBPS) = abi.decode(_data, (address, uint256));
        ERC20MintingFee.__ERC20MintingFee_init(feeRecipient, feeBPS);
    }

    function mint(address to, uint256 amount) public override(ERC20Base, ERC20MintingFee) {
        super.mint(to, amount);
    }
}
