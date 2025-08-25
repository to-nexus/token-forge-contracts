// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC20Base} from "./ERC20Base.sol";

contract ERC20Mintable is ERC20Base {
    error ERC20Mintable__MintingNotAllowed();

    function initialize(
        address _owner,
        address _manager,
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        uint256 _initialSupply,
        address _initialRecipient,
        bytes memory
    ) external override initializer {
        ERC20Base.__ERC20Base_init(_owner, _manager, _name, _symbol, _decimals, _initialSupply, _initialRecipient);
    }
}
