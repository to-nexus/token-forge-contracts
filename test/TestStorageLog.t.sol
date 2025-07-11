// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std-1.9.7/Test.sol";

contract TestStorageLog is Test {
    function setUp() public {}

    function testLog() external pure {
        bytes32 location;
        // ForgeFactory.sol
        {
            location =
                keccak256(abi.encode(uint256(keccak256("cross.storage.ForgeFactory")) - 1)) & ~bytes32(uint256(0xff));
            console.log("\nTokenForgeFactory");
            console.logBytes32(location);
        }
    }
}
