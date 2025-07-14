// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {UUPSUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/proxy/utils/UUPSUpgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/access/OwnableUpgradeable.sol";
import {ERC165Upgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/utils/introspection/ERC165Upgradeable.sol";
import {ERC1155Upgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC1155/ERC1155Upgradeable.sol";
import {ERC1155HolderUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC1155/utils/ERC1155HolderUpgradeable.sol";
import {IERC1155Forge} from "../../interfaces/IERC1155Forge.sol";

abstract contract ERC1155Base is
    IERC1155Forge,
    ERC165Upgradeable,
    ERC1155Upgradeable,
    ERC1155HolderUpgradeable,
    OwnableUpgradeable,
    UUPSUpgradeable
{
    error ERC1155Preset__NullInput(bytes32 field);

    /// @custom:storage-location erc7201:cross.storage.forge.erc1155
    struct ERC1155BaseStorage {
        address forge;
    }

    // keccak256(abi.encode(uint256(keccak256("cross.storage.forge.erc1155")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant ERC1155BaseStorageLocation =
        0xa4673bb384131e0f99d53662bec7c02c60fe363c5b3f49f47bfea0107c7ded00;

    function _getERC1155PresetStorage() private pure returns (ERC1155BaseStorage storage $) {
        assembly {
            $.slot := ERC1155BaseStorageLocation
        }
    }

    function initialize(address owner, address forge, string calldata baseTokenURI) external initializer {
        __ERC1155Preset_init(forge, baseTokenURI);

        __ERC1155_init(baseTokenURI);
        __ERC1155Holder_init();
        __Ownable_init(owner);
    }

    function __ERC1155Preset_init(address forge, string calldata baseTokenURI) internal virtual onlyInitializing {
        if (forge == address(0)) revert ERC1155Preset__NullInput("forge");
        if (bytes(baseTokenURI).length == 0) revert ERC1155Preset__NullInput("baseTokenURI");
        _getERC1155PresetStorage().forge = forge;
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
