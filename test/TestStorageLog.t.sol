// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std-1.9.7/Test.sol";

contract TestStorageLog is Test {
    function setUp() public {}

    function test_storage_forge_factory() external pure {
        bytes32 location;
        // ForgeFactory.sol
        {
            location = keccak256(abi.encode(uint256(keccak256("cross.storage.forge.ForgeFactory")) - 1))
                & ~bytes32(uint256(0xff));
            console.log("\nForgeFactory");
            console.logBytes32(location);
        }
    }

    function test_storage_base_forge() external pure {
        bytes32 location;
        {
            location =
                keccak256(abi.encode(uint256(keccak256("cross.storage.forge.BaseForge")) - 1)) & ~bytes32(uint256(0xff));
            console.log("\nBaseForge");
            console.logBytes32(location);
        }
    }

    function test_storage_token_factory() external pure {
        bytes32 location;
        {
            location = keccak256(abi.encode(uint256(keccak256("cross.storage.forge.TokenFactory")) - 1))
                & ~bytes32(uint256(0xff));
            console.log("\nTokenFactory");
            console.logBytes32(location);
        }
    }

    function test_storage_token_erc20_decimals() external pure {
        bytes32 location;
        {
            location = keccak256(abi.encode(uint256(keccak256("cross.storage.forge.erc20.decimals")) - 1))
                & ~bytes32(uint256(0xff));
            console.log("\nERC20 Decimals");
            console.logBytes32(location);
        }
    }

    function test_storage_token_erc721() external pure {
        bytes32 location;
        {
            location = keccak256(abi.encode(uint256(keccak256("cross.storage.forge.erc721.baseURI")) - 1))
                & ~bytes32(uint256(0xff));
            console.log("\nERC721 BaseURI");
            console.logBytes32(location);
        }
    }

    function test_storage_token_erc20_mintingfee() external pure {
        bytes32 location;
        {
            location =
                keccak256(abi.encode(uint256(keccak256("cross.storage.forge.erc20.fee")) - 1)) & ~bytes32(uint256(0xff));
            console.log("\nERC20Fee");
            console.logBytes32(location);
        }
    }

    function test_storage_token_forges() external pure {
        bytes32 location;
        {
            location =
                keccak256(abi.encode(uint256(keccak256("cross.storage.forge.token")) - 1)) & ~bytes32(uint256(0xff));
            console.log("\nToken Forges");
            console.logBytes32(location);
        }
    }
}
