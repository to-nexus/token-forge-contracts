// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1155Upgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC1155/ERC1155Upgradeable.sol";
import {ERC1155HolderUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC1155/utils/ERC1155HolderUpgradeable.sol";
import {IERC1155Forge} from "../../../interfaces/IERC1155Forge.sol";
import {TokenBase} from "../../TokenBase.sol";

abstract contract ERC1155Base is TokenBase, IERC1155Forge, ERC1155Upgradeable, ERC1155HolderUpgradeable {
    function initialize(address owner, string memory baseTokenURI, bytes memory) external virtual override;

    function __ERC1155Base_init(address owner, string memory baseTokenURI) internal onlyInitializing {
        __TokenBase_init(owner);
        __ERC1155Base_init_unchained(baseTokenURI);

        __ERC1155_init(baseTokenURI);
        __ERC1155Holder_init();
    }

    function __ERC1155Base_init_unchained(string memory baseTokenURI) internal onlyInitializing {
        if (bytes(baseTokenURI).length == 0) revert TokenBase__NullInput("baseTokenURI");
    }

    function mint(address to, uint256 tokenID, uint256 amount, bytes calldata data) public virtual override onlyForge {
        _mint(to, tokenID, amount, data);
    }

    function mintBatch(address to, uint256[] memory tokenIDs, uint256[] memory amounts, bytes calldata data)
        public
        virtual
        override
        onlyForge
    {
        _mintBatch(to, tokenIDs, amounts, data);
    }

    function burnFrom(address from, uint256 tokenID, uint256 amount) public virtual override {
        if (from != _msgSender() && !isApprovedForAll(from, _msgSender())) {
            revert ERC1155MissingApprovalForAll(_msgSender(), from);
        }

        _burn(from, tokenID, amount);
    }

    function burnFromBatch(address from, uint256[] memory tokenIDs, uint256[] memory amounts) public virtual override {
        if (from != _msgSender() && !isApprovedForAll(from, _msgSender())) {
            revert ERC1155MissingApprovalForAll(_msgSender(), from);
        }

        _burnBatch(from, tokenIDs, amounts);
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        virtual
        override(TokenBase, ERC1155Upgradeable, ERC1155HolderUpgradeable)
        returns (bool)
    {
        return interfaceId == type(IERC1155Forge).interfaceId || super.supportsInterface(interfaceId);
    }
}
