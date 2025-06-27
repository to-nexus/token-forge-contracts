// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC20} from "@openzeppelin-contracts-5.3.0/token/ERC20/ERC20.sol";
import {Ownable} from "@openzeppelin-contracts-5.3.0/access/Ownable.sol";

contract MockERC20 is ERC20, Ownable {
    address public forge;

    constructor(address _forge) ERC20("MockERC20", "MCK") Ownable(_msgSender()) {
        forge = _forge;
    }

    function forceMint(address to, uint256 amount) external onlyOwner {
        _mint(to, amount);
    }

    function mint(address to, uint256 amount) external {
        require(_msgSender() == forge, "MockERC20: only forge can mint");
        _mint(to, amount);
    }

    function burnFrom(address from, uint256 amount) external {
        require(_msgSender() == forge, "MockERC20: only forge can burn");
        _spendAllowance(from, _msgSender(), amount);
        _burn(from, amount);
    }
}
