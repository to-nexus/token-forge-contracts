// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC20Upgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC20/ERC20Upgradeable.sol";
import {ERC20PermitUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC20/extensions/ERC20PermitUpgradeable.sol";
import {IERC20Forge} from "../../../interfaces/IERC20Forge.sol";
import {TokenBase} from "../../TokenBase.sol";

abstract contract ERC20Base is TokenBase, IERC20Forge, ERC20Upgradeable, ERC20PermitUpgradeable {
    // keccak256(abi.encode(uint256(keccak256("cross.storage.forge.erc20.decimals")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant ERC20DecimalsStorageLocation =
        0x7c3cec2c3c3d573fa006ae6cf08c04c65698a48d9586626c622c1a81a45caf00;

    function initialize(
        address owner,
        address manager,
        string memory name,
        string memory symbol,
        uint8 _decimals,
        uint256 initialSupply,
        address initialRecipient,
        bytes memory
    ) external virtual override initializer {
        __ERC20Base_init(owner, manager, name, symbol, _decimals, initialSupply, initialRecipient);
    }

    function __ERC20Base_init(
        address owner,
        address manager,
        string memory name,
        string memory symbol,
        uint8 _decimals,
        uint256 initialSupply,
        address initialRecipient
    ) internal onlyInitializing {
        if (bytes(name).length == 0) revert TokenBase__NullInput("name");
        if (bytes(symbol).length == 0) revert TokenBase__NullInput("symbol");

        __TokenBase_init(owner, manager);

        __ERC20_init(name, symbol);
        __ERC20Permit_init(name);
        __ERC20Base_init_unchained(_decimals, initialSupply, initialRecipient);
    }

    function __ERC20Base_init_unchained(uint8 _decimals, uint256 initialSupply, address initialRecipient)
        private
        onlyInitializing
    {
        assembly {
            sstore(ERC20DecimalsStorageLocation, _decimals)
        }
        if (initialSupply != 0) {
            if (initialRecipient == address(0)) revert TokenBase__NullInput("initialRecipient");
            ERC20Upgradeable._update(address(0), initialRecipient, initialSupply);
        }
    }

    function mint(address to, uint256 amount) public virtual override onlyForge {
        _mint(to, amount);
    }

    function burnFrom(address from, uint256 amount) public virtual override {
        _spendAllowance(from, _msgSender(), amount);
        _burn(from, amount);
    }

    function decimals() public view virtual override returns (uint8) {
        uint8 _decimals;
        assembly {
            _decimals := sload(ERC20DecimalsStorageLocation)
        }
        return _decimals;
    }

    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
        return interfaceId == type(IERC20Forge).interfaceId || super.supportsInterface(interfaceId);
    }

    function _update(address from, address to, uint256 value) internal virtual override {
        ERC20Upgradeable._update(from, to, value);
    }

    function transfer(address to, uint256 value) public virtual override returns (bool) {
        return ERC20Upgradeable.transfer(to, value);
    }

    function transferFrom(address from, address to, uint256 value) public virtual override returns (bool) {
        return ERC20Upgradeable.transferFrom(from, to, value);
    }
}
