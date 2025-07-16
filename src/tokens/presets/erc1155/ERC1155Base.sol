// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {UUPSUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/proxy/utils/UUPSUpgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/access/OwnableUpgradeable.sol";
import {ERC165Upgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/utils/introspection/ERC165Upgradeable.sol";
import {ERC1155Upgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC1155/ERC1155Upgradeable.sol";
import {ERC1155HolderUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC1155/utils/ERC1155HolderUpgradeable.sol";
import {IERC1155Forge} from "../../../interfaces/IERC1155Forge.sol";

abstract contract ERC1155Base is
    IERC1155Forge,
    ERC165Upgradeable,
    ERC1155Upgradeable,
    ERC1155HolderUpgradeable,
    OwnableUpgradeable,
    UUPSUpgradeable
{
    error ERC1155Base__NullInput(bytes32 field);
    error ERC1155Base__OnlyForge(address caller);

    /// @custom:storage-location erc7201:cross.storage.forge.erc1155
    struct ERC1155BaseStorage {
        address forge;
    }

    // keccak256(abi.encode(uint256(keccak256("cross.storage.forge.erc1155")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant ERC1155BaseStorageLocation =
        0xa4673bb384131e0f99d53662bec7c02c60fe363c5b3f49f47bfea0107c7ded00;

    function _getERC1155BaseStorage() private pure returns (ERC1155BaseStorage storage $) {
        assembly {
            $.slot := ERC1155BaseStorageLocation
        }
    }

    modifier onlyForge() {
        if (_msgSender() != _getERC1155BaseStorage().forge) revert ERC1155Base__OnlyForge(_msgSender());
        _;
    }

    function initialize(address owner, address forge, string memory baseTokenURI, bytes memory)
        external
        virtual
        initializer
    {
        __ERC1155Base_init(owner, forge, baseTokenURI);
    }

    function __ERC1155Base_init(address owner, address forge, string memory baseTokenURI) internal onlyInitializing {
        __ERC1155Base_init_unchained(forge, baseTokenURI);

        __ERC1155_init(baseTokenURI);
        __ERC1155Holder_init();
        __Ownable_init(owner);
    }

    function __ERC1155Base_init_unchained(address forge, string memory baseTokenURI) private onlyInitializing {
        if (forge == address(0)) revert ERC1155Base__NullInput("forge");
        if (bytes(baseTokenURI).length == 0) revert ERC1155Base__NullInput("baseTokenURI");
        _getERC1155BaseStorage().forge = forge;
    }

    function mint(address to, uint256 tokenID, uint256 amount, bytes calldata data) external onlyForge {
        _mint(to, tokenID, amount, data);
    }

    function mintBatch(address to, uint256[] memory tokenIDs, uint256[] memory amounts, bytes calldata data)
        external
        onlyForge
    {
        _mintBatch(to, tokenIDs, amounts, data);
    }

    function burnFrom(address from, uint256 tokenID, uint256 amount) external {
        if (from != _msgSender() && !isApprovedForAll(from, _msgSender())) {
            revert ERC1155MissingApprovalForAll(_msgSender(), from);
        }

        _burn(from, tokenID, amount);
    }

    function burnFromBatch(address from, uint256[] memory tokenIDs, uint256[] memory amounts) external {
        if (from != _msgSender() && !isApprovedForAll(from, _msgSender())) {
            revert ERC1155MissingApprovalForAll(_msgSender(), from);
        }

        _burnBatch(from, tokenIDs, amounts);
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        virtual
        override(ERC165Upgradeable, ERC1155Upgradeable, ERC1155HolderUpgradeable)
        returns (bool)
    {
        return interfaceId == type(IERC1155Forge).interfaceId || super.supportsInterface(interfaceId);
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}
}
