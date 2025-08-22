// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC721Upgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC721/ERC721Upgradeable.sol";
import {ERC721HolderUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC721/utils/ERC721HolderUpgradeable.sol";
import {ERC721URIStorageUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol";
import {IERC721Forge} from "../../../interfaces/IERC721Forge.sol";
import {TokenBase} from "../../TokenBase.sol";

abstract contract ERC721Base is
    TokenBase,
    IERC721Forge,
    ERC721Upgradeable,
    ERC721HolderUpgradeable,
    ERC721URIStorageUpgradeable
{
    // keccak256(abi.encode(uint256(keccak256("cross.storage.forge.erc721.baseURI")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant ERC721BaseURIStorageLocation =
        0x614112defb2808fb2f424451bd65cd4db97803d41a04f6f5d04c43961d7c9300;

    function initialize(
        address owner,
        string memory name,
        string memory symbol,
        string memory baseTokenURI,
        bytes memory
    ) external virtual override;

    function __ERC721Base_init(address owner, string memory name, string memory symbol, string memory baseTokenURI)
        internal
        onlyInitializing
    {
        __TokenBase_init(owner);
        __ERC721Base_init_unchained(name, symbol, baseTokenURI);

        __ERC721_init(name, symbol);
        __ERC721Holder_init();
        __ERC721URIStorage_init();
    }

    function __ERC721Base_init_unchained(string memory name, string memory symbol, string memory baseTokenURI)
        private
        onlyInitializing
    {
        if (bytes(name).length == 0) revert TokenBase__NullInput("name");
        if (bytes(symbol).length == 0) revert TokenBase__NullInput("symbol");
        if (bytes(baseTokenURI).length == 0) revert TokenBase__NullInput("baseTokenURI");
        assembly {
            sstore(ERC721BaseURIStorageLocation, baseTokenURI)
        }
    }

    function mint(address to, uint256 tokenID) public virtual override onlyForge returns (uint256) {
        _mint(to, tokenID);
        return tokenID;
    }

    function burnFrom(address from, uint256 tokenID) public virtual override {
        _checkAuthorized(from, _msgSender(), tokenID);
        _burn(tokenID);
    }

    function _baseURI() internal view override returns (string memory) {
        string memory uri;
        assembly {
            uri := sload(ERC721BaseURIStorageLocation)
        }
        return uri;
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
        override(TokenBase, ERC721Upgradeable, ERC721URIStorageUpgradeable)
        returns (bool)
    {
        return interfaceId == type(IERC721Forge).interfaceId || super.supportsInterface(interfaceId);
    }
}
