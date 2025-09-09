// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.3.0/proxy/ERC1967/ERC1967Proxy.sol";
import {IDiamondCut} from "diamond-3-hardhat-1.0.0/interfaces/IDiamondCut.sol";
import {Script, console} from "forge-std-1.9.7/Script.sol";

import "../src/ForgeFactory.sol";
import "../src/Diamond3Facet.sol";
import "../src/BaseForge.sol";
import "../src/ForgeProxy.sol";
import "../src/ForgeV1.sol";
import "../src/ForgeV2.sol";
import "../src/ForgeV3.sol";
import "../test/mock/MockERC20.sol";
import "../test/mock/MockERC721.sol";
import "../test/mock/MockERC1155.sol";

contract DeployScript is Script {
    function setUp() public {}

    // deployForgeImpls
    function deployForgeImpls() external {
        vm.startBroadcast();
        address forgeFactoryImpl = address(new ForgeFactory());
        address baseForgeFacetImpl = address(new BaseForgeFacet());
        address diamond3FacetImpl = address(new Diamond3Facet());
        address forgeProxyCode = address(new ForgeProxyCode());
        address forgeV1Impl = address(new ForgeV1());
        address forgeV2Impl = address(new ForgeV2());
        address forgeV3Impl = address(new ForgeV3());
        vm.stopBroadcast();
        console.log("address private forgeFactoryImpl= ", forgeFactoryImpl, ";");
        console.log("address private forgeProxyCode= ", forgeProxyCode, ";");
        console.log("address private diamond3FacetImpl= ", diamond3FacetImpl, ";");
        console.log("address private baseForgeFacetImpl= ", baseForgeFacetImpl, ";");
        console.log("address private forgeV1Impl= ", forgeV1Impl, ";");
        console.log("address private forgeV2Impl= ", forgeV2Impl, ";");
        console.log("address private forgeV3Impl= ", forgeV3Impl, ";");
    }

    // initializeForgeImpls
    function initializeForgeImpls(
        address owner,
        address forgeFactoryImpl,
        address forgeProxyCode,
        address diamond3FacetImpl,
        address baseForgeFacetImpl
    ) external {
        vm.startBroadcast();
        address forgeFactory = address(
            new ERC1967Proxy(
                forgeFactoryImpl,
                abi.encodeCall(ForgeFactory.initialize, (owner, forgeProxyCode, diamond3FacetImpl, baseForgeFacetImpl))
            )
        );

        vm.stopBroadcast();
        console.log("address private forgeFactory= ", forgeFactory, ";");
    }
}
