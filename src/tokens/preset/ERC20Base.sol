// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {UUPSUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/proxy/utils/UUPSUpgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/access/OwnableUpgradeable.sol";
import {ERC20Upgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC20/ERC20Upgradeable.sol";
import {ERC20PermitUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC20/extensions/ERC20PermitUpgradeable.sol";
import {ERC165Upgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/utils/introspection/ERC165Upgradeable.sol";
import {IERC20Forge} from "../../interfaces/IERC20Forge.sol";

abstract contract ERC20Base is
    IERC20Forge,
    ERC165Upgradeable,
    ERC20Upgradeable,
    ERC20PermitUpgradeable,
    OwnableUpgradeable,
    UUPSUpgradeable
{
    error ERC20Base__NullInput(bytes32 field);
    error ERC20Base__InvalidInitialSupply(uint256 initialSupply);
    error ERC20Base__OnlyForge(address caller);

    /// @custom:storage-location erc7201:cross.storage.forge.erc20
    struct ERC20BaseStorage {
        address forge;
        uint8 decimals;
    }

    // keccak256(abi.encode(uint256(keccak256("cross.storage.forge.erc20")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant ERC20BaseStorageLocation =
        0x5336721aabdf58b9f67dce2b9c89749de013c2e87912f00ce8440ff438c6c800;

    function _getERC20PresetStorage() private pure returns (ERC20BaseStorage storage $) {
        assembly {
            $.slot := ERC20BaseStorageLocation
        }
    }

    modifier onlyForge() {
        if (_msgSender() != _getERC20PresetStorage().forge) revert ERC20Base__OnlyForge(_msgSender());
        _;
    }

    function initialize(
        address _owner,
        address _forge,
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        uint256 _initialSupply
    ) external initializer {
        __ERC20Preset_init(_owner, _forge, _name, _symbol, _decimals, _initialSupply);

        __ERC20_init(_name, _symbol);
        __ERC20Permit_init(_name);
        __Ownable_init(_owner);
    }

    function __ERC20Preset_init(
        address _owner,
        address _forge,
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        uint256 _initialSupply
    ) internal virtual onlyInitializing {
        // owner 는 __Ownable_init() 에서 address(0) 을 확인함
        if (bytes(_name).length == 0) revert ERC20Base__NullInput("name");
        if (bytes(_symbol).length == 0) revert ERC20Base__NullInput("symbol");
        ERC20BaseStorage storage $ = _getERC20PresetStorage();
        $.decimals = _decimals;
        $.forge = _forge;
        if (_initialSupply != 0) {
            _mint(_owner, _initialSupply * (10 ** uint256(_decimals)));
        }
    }

    function mint(address to, uint256 amount) public virtual override onlyForge {
        _mint(to, amount);
    }

    function burnFrom(address from, uint256 amount) public virtual override {
        _spendAllowance(from, _msgSender(), amount);
        _burn(from, amount);
    }

    function decimals() public view override returns (uint8) {
        return _getERC20PresetStorage().decimals;
    }

    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
        return interfaceId == type(IERC20Forge).interfaceId || super.supportsInterface(interfaceId);
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}
}
