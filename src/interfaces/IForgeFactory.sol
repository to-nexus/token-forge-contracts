// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

enum TokenType {
    ERC20,
    ERC721,
    ERC1155
}

interface IForgeFactoryAlert {
    function alertMint(TokenType tokenType, uint256 uuid, address token, bytes calldata data) external;
    function alertTransfer(TokenType tokenType, uint256 uuid, address token, bytes calldata data) external;
    function alertTransferFrom(TokenType tokenType, uint256 uuid, address token, bytes calldata data) external;
    function alertBurn(TokenType tokenType, uint256 uuid, address token, bytes calldata data) external;
}

interface IForgeFactoryForgeByService {
    function forgeByService(string memory service) external view returns (address forge, bool running);
}
