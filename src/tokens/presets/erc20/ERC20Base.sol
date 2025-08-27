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
        address _owner,
        address _manager,
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        uint256 _initialSupply,
        address _initialRecipient,
        bytes memory
    ) external virtual override;

    function __ERC20Base_init(
        address _owner,
        address _manager,
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        uint256 _initialSupply,
        address _initialRecipient
    ) internal onlyInitializing {
        __TokenBase_init(_owner, _manager);
        __ERC20Base_init_unchained(_owner, _name, _symbol, _decimals, _initialSupply, _initialRecipient);

        __ERC20_init(_name, _symbol);
        __ERC20Permit_init(_name);
    }

    function __ERC20Base_init_unchained(
        address _owner,
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        uint256 _initialSupply,
        address _initialRecipient
    ) internal onlyInitializing {
        if (bytes(_name).length == 0) revert TokenBase__NullInput("name");
        if (bytes(_symbol).length == 0) revert TokenBase__NullInput("symbol");
        assembly {
            sstore(ERC20DecimalsStorageLocation, _decimals)
        }
        if (_initialSupply != 0) {
            if (_initialRecipient == address(0)) revert TokenBase__NullInput("initialRecipient");
            _mint(_initialRecipient, _initialSupply);
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
}
