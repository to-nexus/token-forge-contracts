// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {UUPSUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/proxy/utils/UUPSUpgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/access/OwnableUpgradeable.sol";
import {ERC165Upgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/utils/introspection/ERC165Upgradeable.sol";
import {ERC721Upgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC721/ERC721Upgradeable.sol";
import {ERC721HolderUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC721/utils/ERC721HolderUpgradeable.sol";
import {ERC721URIStorageUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol";
import {IERC721Forge} from "../../interfaces/IERC721Forge.sol";

abstract contract ERC721Base is
    IERC721Forge,
    ERC165Upgradeable,
    ERC721Upgradeable,
    ERC721HolderUpgradeable,
    ERC721URIStorageUpgradeable,
    OwnableUpgradeable,
    UUPSUpgradeable
{
    error ERC721Base__NullInput(bytes32 field);
    error ERC721Base__OnlyForge(address caller);

    /// @custom:storage-location erc7201:cross.storage.forge.erc721
    struct ERC721PresetStorage {
        address forge;
        string baseURI;
    }

    // keccak256(abi.encode(uint256(keccak256("cross.storage.forge.erc721")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant ERC721PresetStorageLocation =
        0xae399aba48b00e5fadd22cfd23d29924573cb21a1b3727798605d103f2a75a00;

    function _getERC721PresetStorage() private pure returns (ERC721PresetStorage storage $) {
        assembly {
            $.slot := ERC721PresetStorageLocation
        }
    }

    modifier onlyForge() {
        if (_msgSender() != _getERC721PresetStorage().forge) revert ERC721Base__OnlyForge(_msgSender());
        _;
    }

    function initialize(
        address owner,
        address forge,
        string calldata name,
        string calldata symbol,
        string calldata baseTokenURI
    ) external initializer {
        __ERC721Preset_init(forge, name, symbol, baseTokenURI);

        __ERC721_init(name, symbol);
        __ERC721Holder_init();
        __ERC721URIStorage_init();
        __Ownable_init(owner);
    }

    function __ERC721Preset_init(
        address forge,
        string calldata name,
        string calldata symbol,
        string calldata baseTokenURI
    ) internal virtual onlyInitializing {
        if (bytes(name).length == 0) revert ERC721Base__NullInput("name");
        if (bytes(symbol).length == 0) revert ERC721Base__NullInput("symbol");
        if (bytes(baseTokenURI).length == 0) revert ERC721Base__NullInput("symbol");
        ERC721PresetStorage storage $ = _getERC721PresetStorage();
        $.baseURI = baseTokenURI;
        $.forge = forge;
    }

    function mint(address to, uint256 tokenID) external onlyForge {
        _mint(to, tokenID);
    }

    function burnFrom(address from, uint256 tokenID) external {
        _checkAuthorized(from, _msgSender(), tokenID);
        _burn(tokenID);
    }

    function _baseURI() internal view override returns (string memory) {
        return _getERC721PresetStorage().baseURI;
    }

    function tokenURI(uint256 tokenId)
        public
        view
        override(ERC721Upgradeable, ERC721URIStorageUpgradeable)
        returns (string memory)
    {
        return ERC721URIStorageUpgradeable.tokenURI(tokenId);
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        virtual
        override(ERC165Upgradeable, ERC721Upgradeable, ERC721URIStorageUpgradeable)
        returns (bool)
    {
        return interfaceId == type(IERC721Forge).interfaceId || super.supportsInterface(interfaceId);
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}
}
