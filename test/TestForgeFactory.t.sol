// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std-1.9.7/Test.sol";
import {Vm} from "forge-std-1.9.7/Vm.sol";

import {IDiamondCut} from "diamond-3-hardhat-1.0.0/interfaces/IDiamondCut.sol";
import {IDiamondLoupe} from "diamond-3-hardhat-1.0.0/interfaces/IDiamondLoupe.sol";
import {IERC173} from "diamond-3-hardhat-1.0.0/interfaces/IERC173.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts-5.3.0/proxy/ERC1967/ERC1967Proxy.sol";

import {ForgeProxyCode} from "../src/ForgeProxy.sol";
import {Diamond3Facet} from "../src/Diamond3Facet.sol";
import {BaseForge, BaseForgeFacet} from "../src/BaseForge.sol";
import {ForgeFactory} from "../src/ForgeFactory.sol";

contract TestForgeFactory is Test {
    address public constant OWNER = address(bytes20("OWNER"));

    ForgeProxyCode public forgeProxyCode;
    Diamond3Facet public diamond3Facet;
    BaseForgeFacet public baseForgeFacet;
    ForgeFactory public forgeFactory;

    function setUp() public {
        vm.label(OWNER, "owner");
        vm.startPrank(OWNER);
        // deploy contracts
        // deploy faucets
        diamond3Facet = new Diamond3Facet();
        baseForgeFacet = new BaseForgeFacet();
        // deploy proxy code
        forgeProxyCode = new ForgeProxyCode();
        // deploy factory
        ForgeFactory tokenForgeFactoryImpl = new ForgeFactory();
        ERC1967Proxy tokenForgeFactoryProxy = new ERC1967Proxy(
            address(tokenForgeFactoryImpl),
            abi.encodeCall(
                ForgeFactory.initialize,
                (OWNER, address(forgeProxyCode), address(diamond3Facet), address(baseForgeFacet))
            )
        );
        forgeFactory = ForgeFactory(address(tokenForgeFactoryProxy));
        vm.stopPrank();
    }

    function test_check_deployed_faucets() external view {
        IDiamondCut.FacetCut memory cut;
        cut = diamond3Facet.defaultDiamondFacetCut();
        assertEq(address(diamond3Facet), cut.facetAddress, "Diamond3Facet address mismatch");

        cut = baseForgeFacet.defaultDiamondFacetCut();
        assertEq(address(baseForgeFacet), cut.facetAddress, "BaseForgeFacet address mismatch");
    }

    function test_factory_add_pause_remove_service() external {
        bytes32 service = bytes32("TestService");
        address serviceOwner = address(bytes20("TestServiceOwner"));

        Vm.Wallet memory validator = vm.createWallet(string(abi.encodePacked(service)));
        IDiamondCut.FacetCut[] memory addCuts;
        vm.prank(OWNER);
        address forgeProxy = forgeFactory.addService(serviceOwner, validator.addr, service, addCuts);
        assertNotEq(forgeProxy, address(0), "ForgeProxy address should not be zero");
        // check domain separator
        {
            bytes32 domainSeparator = BaseForge(forgeProxy).DOMAIN_SEPARATOR();
            assertNotEq(domainSeparator, bytes32(0), "Domain separator should not be zero");
            bytes32 TYPE_HASH =
                keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)");
            bytes32 nameHash = keccak256(abi.encodePacked(service));
            bytes32 versionHash = keccak256(bytes("1"));
            uint256 chainID = block.chainid;
            address thisAddress = forgeProxy;
            bytes32 expectedDomainSeparator =
                keccak256(abi.encode(TYPE_HASH, nameHash, versionHash, chainID, thisAddress));
            assertEq(domainSeparator, expectedDomainSeparator);
        }

        // check facets in the forge proxy
        {
            IDiamondLoupe.Facet[] memory faucets = IDiamondLoupe(forgeProxy).facets();
            assertEq(faucets.length, 2, "Expected 2 facets in the forge proxy");
            (IDiamondCut.FacetCut memory cut, IDiamondLoupe.Facet memory facet) =
                (diamond3Facet.defaultDiamondFacetCut(), faucets[0]);
            assertEq(cut.facetAddress, facet.facetAddress, "Facet address mismatch for Diamond3Facet");
            uint256 length = cut.functionSelectors.length;
            assertEq(length, 8, "Expected 8 function selectors for Diamond3Facet");
            assertEq(length, facet.functionSelectors.length, "Function selectors length mismatch for Diamond3Facet");
            for (uint256 i = 0; i < length; i++) {
                assertEq(
                    cut.functionSelectors[i], facet.functionSelectors[i], "Function selector mismatch for Diamond3Facet"
                );
            }
            (cut, facet) = (baseForgeFacet.defaultDiamondFacetCut(), faucets[1]);
            assertEq(cut.facetAddress, facet.facetAddress, "Facet address mismatch for BaseForgeFacet");
            length = cut.functionSelectors.length;
            assertEq(length, 5, "Expected 5 function selectors for BaseForgeFacet");
            assertEq(length, facet.functionSelectors.length, "Function selectors length mismatch for BaseForgeFacet");
            for (uint256 i = 0; i < length; i++) {
                assertEq(
                    cut.functionSelectors[i],
                    facet.functionSelectors[i],
                    "Function selector mismatch for BaseForgeFacet"
                );
            }
        }
        // check initial data
        {
            assertEq(BaseForge(forgeProxy).validator(), validator.addr, "Validator address mismatch");
            assertEq(IERC173(forgeProxy).owner(), serviceOwner, "Service Owner address mismatch");
        }
        // check factory data
        {
            // allForges
            (bytes32[] memory services, address[] memory forges) = forgeFactory.allForges();
            assertEq(services.length, 1, "Expected 1 service in the factory");
            assertEq(services.length, forges.length, "Services and forges length mismatch");
            assertEq(services[0], service, "Service name mismatch");
            assertEq(forges[0], forgeProxy, "Forge address mismatch");

            // lengthAllForges
            uint256 lengthAllForges = forgeFactory.lengthAllForges();
            assertEq(lengthAllForges, 1, "Expected 1 forge in the factory");

            // forgeByIndex
            (bytes32 serviceName, address forge, bool running) = forgeFactory.forgeByIndex(0);
            assertEq(service, serviceName, "Service name mismatch by index");
            assertEq(forgeProxy, forge, "Service name mismatch by index");
            assertTrue(running, "Service should be running");

            // forgeByService
            (forge, running) = forgeFactory.forgeByService(service);
            assertEq(forgeProxy, forge, "Forge address mismatch by service");
            assertTrue(running, "Service should be running");
        }

        // transfer ownership
        {
            vm.prank(serviceOwner);
            IERC173(forgeProxy).transferOwnership(OWNER);
            assertEq(IERC173(forgeProxy).owner(), OWNER, "Owner address mismatch after transfer");
        }
        // change validator
        {
            vm.prank(OWNER);
            BaseForgeFacet(forgeProxy).setValidator(OWNER);
            assertEq(BaseForge(forgeProxy).validator(), OWNER, "Validator address mismatch after change");
        }
        // pause service
        {
            vm.prank(OWNER);
            forgeFactory.pauseService(service, true);
            (address forge, bool running) = forgeFactory.forgeByService(service);
            assertEq(forge, forgeProxy, "Forge address should not change after pause");
            assertFalse(running, "Service should not be running after pause");

            // check facets in the forge proxy
            IDiamondLoupe.Facet[] memory faucets = IDiamondLoupe(forgeProxy).facets();
            assertEq(faucets.length, 2, "Expected 2 facets in the forge proxy after pause");
        }
        // resume service
        {
            vm.prank(OWNER);
            forgeFactory.pauseService(service, false);
            (address forge, bool running) = forgeFactory.forgeByService(service);
            assertEq(forge, forgeProxy, "Forge address should not change after resume");
            assertTrue(running, "Service should be running after resume");

            // check facets in the forge proxy
            IDiamondLoupe.Facet[] memory faucets = IDiamondLoupe(forgeProxy).facets();
            assertEq(faucets.length, 2, "Expected 2 facets in the forge proxy after resume");
        }
        // check facets in the forge proxy
        {
            vm.prank(OWNER);
            forgeFactory.removeService(service);
            // allForges
            (bytes32[] memory services, address[] memory forges) = forgeFactory.allForges();
            assertEq(services.length, 0, "Expected 0 services in the factory after removal");
            assertEq(services.length, forges.length, "Services and forges length mismatch after removal");
            // lengthAllForges
            uint256 lengthAllForges = forgeFactory.lengthAllForges();
            assertEq(lengthAllForges, 0, "Expected 0 forges in the factory after removal");
            // forgeByIndex out of bounds
            // (string memory serviceName, address forge, bool running) = tokenForgeFactory.forgeByIndex(0);
            // assertEq(serviceName, "", "Service name should be empty after removal");
            // assertEq(forge, address(0), "Forge address should be zero after removal");
            // assertFalse(running, "Service should not be running after removal");
            // forgeByService
            (address forge, bool running) = forgeFactory.forgeByService(service);
            assertEq(forge, address(0), "Forge address should be zero after removal");
            assertFalse(running, "Service should not be running after removal");
        }
    }

    function test_factory_2forge() external {
        (bytes32 service1, address serviceOwner1, Vm.Wallet memory validator1) =
            ("TestService1", address(bytes20("TestServiceOwner1")), vm.createWallet("TestServiceOwner1"));
        (bytes32 service2, address serviceOwner2, Vm.Wallet memory validator2) =
            ("TestService2", address(bytes20("TestServiceOwner2")), vm.createWallet("TestServiceOwner2"));
        assertNotEq(service1, service2, "Service names should be different");
        assertNotEq(serviceOwner1, serviceOwner2, "Service owners should be different");
        assertNotEq(validator1.addr, validator2.addr, "Validators should be different");

        IDiamondCut.FacetCut[] memory addCuts;
        vm.prank(OWNER);
        address forgeProxy1 = forgeFactory.addService(serviceOwner1, validator1.addr, service1, addCuts);
        vm.prank(OWNER);
        address forgeProxy2 = forgeFactory.addService(serviceOwner2, validator2.addr, service2, addCuts);
        assertNotEq(forgeProxy1, forgeProxy2, "Forge proxies should be different");
        assertNotEq(
            BaseForge(forgeProxy1).DOMAIN_SEPARATOR(),
            BaseForge(forgeProxy2).DOMAIN_SEPARATOR(),
            "Domain separators should be different"
        );
    }
}
