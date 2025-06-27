// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1155} from "@openzeppelin-contracts-5.3.0/token/ERC1155/ERC1155.sol";
import {Ownable} from "@openzeppelin-contracts-5.3.0/access/Ownable.sol";

contract MockERC1155 is ERC1155, Ownable {
    address public forge;

    constructor(address _forge) ERC1155("") Ownable(_msgSender()) {
        forge = _forge;
    }

    function forceMint(address to, uint256 tokenID, uint256 amount) external onlyOwner {
        _mint(to, tokenID, amount, hex"");
    }

    function forceMintBatch(address to, uint256[] memory tokenIDs, uint256[] memory amounts) external onlyOwner {
        _mintBatch(to, tokenIDs, amounts, hex"");
    }

    function mint(address to, uint256 tokenID, uint256 amount, bytes memory data) external {
        require(_msgSender() == forge, "MockERC1155: only forge can mint");
        _mint(to, tokenID, amount, data);
    }

    function mintBatch(address to, uint256[] memory tokenIDs, uint256[] memory amounts, bytes memory data) external {
        require(_msgSender() == forge, "MockERC1155: only forge can mint");
        _mintBatch(to, tokenIDs, amounts, data);
    }

    function burnFrom(address from, uint256 tokenID, uint256 amount) external {
        require(_msgSender() == forge, "MockERC1155: only forge can burn");
        if (!isApprovedForAll(from, _msgSender())) revert ERC1155MissingApprovalForAll(_msgSender(), from);
        _burn(from, tokenID, amount);
    }

    function burnFromBatch(address from, uint256[] memory tokenIDs, uint256[] memory amounts) external {
        require(_msgSender() == forge, "MockERC1155: only forge can burn");
        if (!isApprovedForAll(from, _msgSender())) revert ERC1155MissingApprovalForAll(_msgSender(), from);
        _burnBatch(from, tokenIDs, amounts);
    }
}
