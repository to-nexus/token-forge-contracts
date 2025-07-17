// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC20Base} from "./ERC20Base.sol";
import {ERC20Fee} from "./extensions/ERC20Fee.sol";

contract ERC20MintingFee is ERC20Base, ERC20Fee {
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
        ERC20Fee.__ERC20Fee_init(feeRecipient, feeBPS);
    }

    function mint(address to, uint256 amount) public override {
        uint256 feeBPS = feeBPS();
        if (feeBPS == 0) {
            ERC20Base.mint(to, amount);
        } else {
            address feeRecipient = feeRecipient();
            uint256 fee = (amount * feeBPS) / 10000;
            if (fee != 0) {
                amount -= fee;
                ERC20Base.mint(feeRecipient, fee);
                emit MintingFeeCollected(feeRecipient, fee);
            }
            ERC20Base.mint(to, amount);
        }
    }
}
