// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {Test, console} from "forge-std-1.9.7/Test.sol";
import {Vm} from "forge-std-1.9.7/Vm.sol";

import {IDiamondCut} from "diamond-3-hardhat-1.0.0/interfaces/IDiamondCut.sol";
import {IDiamondLoupe} from "diamond-3-hardhat-1.0.0/interfaces/IDiamondLoupe.sol";
import {IERC173} from "diamond-3-hardhat-1.0.0/interfaces/IERC173.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts-5.3.0/proxy/ERC1967/ERC1967Proxy.sol";
import {UUPSUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/proxy/utils/UUPSUpgradeable.sol";
import {IAccessControl} from "@openzeppelin-contracts-5.3.0/access/IAccessControl.sol";

import {ForgeProxyCode} from "../../src/ForgeProxy.sol";
import {Diamond3Facet} from "../../src/Diamond3Facet.sol";

// v1 imports
import {ForgeFactory as ForgeFactoryV1} from "../../src/v1/ForgeFactory.sol";
import {BaseForge as BaseForgeV1, BaseForgeFacet as BaseForgeFacetV1} from "../../src/v1/BaseForge.sol";
import {ForgeV1 as ForgeV1FacetV1} from "../../src/v1/ForgeV1.sol";

// v2 imports
import {ForgeFactory as ForgeFactoryV2} from "../../src/v2/ForgeFactory.sol";
import {BaseForge as BaseForgeV2, BaseForgeFacet as BaseForgeFacetV2} from "../../src/v2/BaseForge.sol";
import {ForgeV1 as ForgeV1FacetV2} from "../../src/v2/ForgeV1.sol";
import {IBaseForgeFacetV2} from "../../src/interfaces/IBaseForgeFacet.sol";

// Mock tokens for testing
import {MockERC20} from "../mock/MockERC20.sol";
import {MessageHashUtils} from "@openzeppelin-contracts-5.3.0/utils/cryptography/MessageHashUtils.sol";

contract TestUpgradeV1ToV2 is Test {
    address public constant OWNER = address(bytes20("OWNER"));
    address public constant SERVICE_OWNER = address(bytes20("SERVICE_OWNER"));
    bytes32 public constant SERVICE_NAME = bytes32("TestService");

    Vm.Wallet public VALIDATOR;

    // Common contracts
    ForgeProxyCode public forgeProxyCode;
    Diamond3Facet public diamond3Facet;

    // v1 contracts
    BaseForgeFacetV1 public baseForgeFacetV1;
    ForgeV1FacetV1 public forgeV1FacetV1;
    ForgeFactoryV1 public forgeFactoryV1Impl;

    // v2 contracts
    BaseForgeFacetV2 public baseForgeFacetV2;
    ForgeV1FacetV2 public forgeV1FacetV2;
    ForgeFactoryV2 public forgeFactoryV2Impl;

    // Proxy (used for both v1 and v2)
    address public forgeFactoryProxy;

    // Mock tokens
    MockERC20 public mockERC20;

    function setUp() public {
        VALIDATOR = vm.createWallet("Validator");

        vm.label(OWNER, "owner");
        vm.label(SERVICE_OWNER, "serviceOwner");
        vm.startPrank(OWNER);

        // Deploy common contracts
        forgeProxyCode = new ForgeProxyCode();
        diamond3Facet = new Diamond3Facet();

        // Deploy v1 contracts
        baseForgeFacetV1 = new BaseForgeFacetV1();
        forgeV1FacetV1 = new ForgeV1FacetV1();
        forgeFactoryV1Impl = new ForgeFactoryV1();

        // Deploy v2 contracts
        baseForgeFacetV2 = new BaseForgeFacetV2();
        forgeV1FacetV2 = new ForgeV1FacetV2();
        forgeFactoryV2Impl = new ForgeFactoryV2();

        // Deploy ForgeFactory proxy with v1 implementation
        ERC1967Proxy proxy = new ERC1967Proxy(
            address(forgeFactoryV1Impl),
            abi.encodeCall(
                ForgeFactoryV1.initialize,
                (OWNER, address(forgeProxyCode), address(diamond3Facet), address(baseForgeFacetV1))
            )
        );
        forgeFactoryProxy = address(proxy);

        vm.stopPrank();
    }

    // ==================== Test: Basic Upgrade Flow ====================

    function test_upgrade_forgeFactory_v1_to_v2() external {
        // Step 1: Verify v1 is working
        ForgeFactoryV1 factoryV1 = ForgeFactoryV1(forgeFactoryProxy);
        assertEq(factoryV1.lengthAllForges(), 0);

        // Step 2: Add a service with v1
        IDiamondCut.FacetCut[] memory addCuts;
        vm.prank(OWNER);
        address forgeV1 = factoryV1.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        assertNotEq(forgeV1, address(0));

        // Verify v1 forge uses EIP712 version "1"
        bytes32 domainSeparatorV1 = BaseForgeV1(forgeV1).DOMAIN_SEPARATOR();
        assertNotEq(domainSeparatorV1, bytes32(0));

        // Step 3: Upgrade to v2 implementation
        vm.prank(OWNER);
        UUPSUpgradeable(forgeFactoryProxy).upgradeToAndCall(address(forgeFactoryV2Impl), "");

        // Step 4: Call reinitialize with v2 BaseForgeFacet
        vm.prank(OWNER);
        ForgeFactoryV2(forgeFactoryProxy).reinitialize(address(baseForgeFacetV2));

        // Step 5: Verify state is preserved
        ForgeFactoryV2 factoryV2 = ForgeFactoryV2(forgeFactoryProxy);
        assertEq(factoryV2.lengthAllForges(), 1);

        (address forge, bool running) = factoryV2.forgeByService(SERVICE_NAME);
        assertEq(forge, forgeV1);
        assertTrue(running);
    }

    // ==================== Test: New Service with Blacklist ====================

    function test_upgrade_then_add_new_service_with_blacklist() external {
        // Upgrade to v2
        vm.startPrank(OWNER);
        UUPSUpgradeable(forgeFactoryProxy).upgradeToAndCall(address(forgeFactoryV2Impl), "");
        ForgeFactoryV2(forgeFactoryProxy).reinitialize(address(baseForgeFacetV2));

        // Add new service after upgrade
        bytes32 newServiceName = bytes32("NewServiceV2");
        IDiamondCut.FacetCut[] memory addCuts;
        address newForge =
            ForgeFactoryV2(forgeFactoryProxy).addService(SERVICE_OWNER, VALIDATOR.addr, newServiceName, addCuts);
        vm.stopPrank();

        // Verify new forge has blacklist functionality
        assertFalse(BaseForgeV2(newForge).isBlacklistManager(SERVICE_OWNER));
        assertFalse(BaseForgeV2(newForge).isBlacklisted(address(0x1234)));

        // Set blacklist manager
        address[] memory managers = new address[](1);
        managers[0] = SERVICE_OWNER;
        vm.prank(SERVICE_OWNER);
        IBaseForgeFacetV2(newForge).setBlacklistManager(managers, true);

        assertTrue(BaseForgeV2(newForge).isBlacklistManager(SERVICE_OWNER));

        // Add to blacklist
        address[] memory accounts = new address[](1);
        accounts[0] = address(0x1234);
        vm.prank(SERVICE_OWNER);
        IBaseForgeFacetV2(newForge).updateBlacklist(accounts, true);

        assertTrue(BaseForgeV2(newForge).isBlacklisted(address(0x1234)));

        // Verify EIP712 version is "2" for new forge
        bytes32 TYPE_HASH =
            keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)");
        bytes32 nameHash = keccak256(abi.encodePacked(newServiceName));
        bytes32 versionHash = keccak256(bytes("2")); // v2 uses version "2"
        bytes32 expectedDomainSeparator =
            keccak256(abi.encode(TYPE_HASH, nameHash, versionHash, block.chainid, newForge));
        assertEq(BaseForgeV2(newForge).DOMAIN_SEPARATOR(), expectedDomainSeparator);
    }

    // ==================== Test: Existing Forge Facet Upgrade ====================

    function test_upgrade_existing_forge_facets() external {
        // Add service with v1
        IDiamondCut.FacetCut[] memory addCuts;
        vm.prank(OWNER);
        address forgeV1 =
            ForgeFactoryV1(forgeFactoryProxy).addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);

        // Verify v1 forge doesn't have blacklist functions (will revert)
        vm.expectRevert("Diamond: Function does not exist");
        BaseForgeV2(forgeV1).isBlacklistManager(SERVICE_OWNER);

        // Upgrade ForgeFactory to v2
        vm.startPrank(OWNER);
        UUPSUpgradeable(forgeFactoryProxy).upgradeToAndCall(address(forgeFactoryV2Impl), "");
        ForgeFactoryV2(forgeFactoryProxy).reinitialize(address(baseForgeFacetV2));
        vm.stopPrank();

        // Existing forge still doesn't have v2 functions
        vm.expectRevert("Diamond: Function does not exist");
        BaseForgeV2(forgeV1).isBlacklistManager(SERVICE_OWNER);

        // Upgrade existing forge's facets via diamondCut
        // v1 BaseForgeFacet selectors (from BASEFORGE_FACET_FUNCTIONS):
        // eip712Domain, nonces, DOMAIN_SEPARATOR, validator, setValidator
        bytes4[] memory oldSelectors = new bytes4[](5);
        oldSelectors[0] = bytes4(keccak256("eip712Domain()")); // EIP712Upgradeable.eip712Domain.selector
        oldSelectors[1] = bytes4(keccak256("nonces(address)")); // NoncesUpgradeable.nonces.selector
        oldSelectors[2] = BaseForgeV1.DOMAIN_SEPARATOR.selector;
        oldSelectors[3] = BaseForgeV1.validator.selector;
        oldSelectors[4] = BaseForgeFacetV1.setValidator.selector;

        // New v2 selectors to add
        bytes4[] memory newSelectors = new bytes4[](4);
        newSelectors[0] = BaseForgeV2.isBlacklistManager.selector;
        newSelectors[1] = BaseForgeV2.isBlacklisted.selector;
        newSelectors[2] = IBaseForgeFacetV2.setBlacklistManager.selector;
        newSelectors[3] = IBaseForgeFacetV2.updateBlacklist.selector;

        // Prepare diamond cut
        IDiamondCut.FacetCut[] memory cuts = new IDiamondCut.FacetCut[](2);
        cuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(baseForgeFacetV2),
            action: IDiamondCut.FacetCutAction.Replace,
            functionSelectors: oldSelectors
        });
        cuts[1] = IDiamondCut.FacetCut({
            facetAddress: address(baseForgeFacetV2),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: newSelectors
        });

        // Service owner performs the diamond cut
        vm.prank(SERVICE_OWNER);
        IDiamondCut(forgeV1).diamondCut(cuts, address(0), "");

        // Now v1 forge has blacklist functions
        assertFalse(BaseForgeV2(forgeV1).isBlacklistManager(SERVICE_OWNER));
        assertFalse(BaseForgeV2(forgeV1).isBlacklisted(address(0x1234)));

        // Note: EIP712 version remains "1" because it was initialized with v1
        bytes32 TYPE_HASH =
            keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)");
        bytes32 nameHash = keccak256(abi.encodePacked(SERVICE_NAME));
        bytes32 versionHashV1 = keccak256(bytes("1")); // Still version "1"
        bytes32 expectedDomainSeparator =
            keccak256(abi.encode(TYPE_HASH, nameHash, versionHashV1, block.chainid, forgeV1));
        assertEq(BaseForgeV1(forgeV1).DOMAIN_SEPARATOR(), expectedDomainSeparator);
    }

    // ==================== Test: Full Forge V1 to V2 Upgrade ====================

    function test_upgrade_forge_v1_to_v2_with_full_functionality() external {
        // Step 1: Deploy Forge with v1 facets (BaseForgeFacet + ForgeV1)
        Vm.Wallet memory USER = vm.createWallet("User");

        // Prepare v1 ForgeV1 facet cuts
        bytes4[] memory v1Selectors = new bytes4[](4);
        v1Selectors[0] = bytes4(keccak256("mintERC20(address,uint256,address,uint256,uint256,bytes)"));
        v1Selectors[1] = bytes4(keccak256("transferERC20(address,uint256,address,uint256,uint256,bytes)"));
        v1Selectors[2] = bytes4(keccak256("transferFromERC20(address,uint256,address,uint256,uint256,bytes)"));
        v1Selectors[3] = bytes4(keccak256("burnERC20(address,uint256,address,uint256,uint256,bytes)"));

        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1FacetV1),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: v1Selectors
        });

        vm.prank(OWNER);
        address forge =
            ForgeFactoryV1(forgeFactoryProxy).addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);

        // Deploy mock ERC20 with forge as minter
        vm.prank(OWNER);
        mockERC20 = new MockERC20(forge);

        // Step 2: Verify v1 mint works
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForgeV1(forge).nonces(USER.addr);

        bytes32 ERC20_MINT_TYPE_HASH = keccak256(
            "ERC20Mint(address recipient,address token,uint256 amount,address feeRecipient,uint256 feeBPS,uint256 nonce,uint256 deadline)"
        );
        bytes32 structHash = keccak256(
            abi.encode(ERC20_MINT_TYPE_HASH, USER.addr, address(mockERC20), amount, address(0), 0, nonce, deadline)
        );
        bytes32 domainSeparator = BaseForgeV1(forge).DOMAIN_SEPARATOR();
        bytes32 hash = MessageHashUtils.toTypedDataHash(domainSeparator, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);
        bytes memory validatorSig = abi.encodePacked(r, s, v);

        vm.prank(USER.addr);
        ForgeV1FacetV1(forge).mintERC20(address(mockERC20), amount, address(0), 0, deadline, validatorSig);
        assertEq(mockERC20.balanceOf(USER.addr), amount);

        // Step 3: Verify v1 forge doesn't have blacklist (will revert)
        vm.expectRevert("Diamond: Function does not exist");
        BaseForgeV2(forge).isBlacklistManager(SERVICE_OWNER);

        // Step 4: Upgrade Forge facets to v2
        // Replace BaseForgeFacet selectors
        bytes4[] memory baseOldSelectors = new bytes4[](5);
        baseOldSelectors[0] = bytes4(keccak256("eip712Domain()"));
        baseOldSelectors[1] = bytes4(keccak256("nonces(address)"));
        baseOldSelectors[2] = BaseForgeV1.DOMAIN_SEPARATOR.selector;
        baseOldSelectors[3] = BaseForgeV1.validator.selector;
        baseOldSelectors[4] = BaseForgeFacetV1.setValidator.selector;

        // New blacklist selectors to add
        bytes4[] memory baseNewSelectors = new bytes4[](4);
        baseNewSelectors[0] = BaseForgeV2.isBlacklistManager.selector;
        baseNewSelectors[1] = BaseForgeV2.isBlacklisted.selector;
        baseNewSelectors[2] = IBaseForgeFacetV2.setBlacklistManager.selector;
        baseNewSelectors[3] = IBaseForgeFacetV2.updateBlacklist.selector;

        // Replace ForgeV1 selectors (v2 has notBlacklisted modifier)
        bytes4[] memory forgeV1Selectors = new bytes4[](4);
        forgeV1Selectors[0] = bytes4(keccak256("mintERC20(address,uint256,address,uint256,uint256,bytes)"));
        forgeV1Selectors[1] = bytes4(keccak256("transferERC20(address,uint256,address,uint256,uint256,bytes)"));
        forgeV1Selectors[2] = bytes4(keccak256("transferFromERC20(address,uint256,address,uint256,uint256,bytes)"));
        forgeV1Selectors[3] = bytes4(keccak256("burnERC20(address,uint256,address,uint256,uint256,bytes)"));

        IDiamondCut.FacetCut[] memory upgradeCuts = new IDiamondCut.FacetCut[](3);
        upgradeCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(baseForgeFacetV2),
            action: IDiamondCut.FacetCutAction.Replace,
            functionSelectors: baseOldSelectors
        });
        upgradeCuts[1] = IDiamondCut.FacetCut({
            facetAddress: address(baseForgeFacetV2),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: baseNewSelectors
        });
        upgradeCuts[2] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1FacetV2),
            action: IDiamondCut.FacetCutAction.Replace,
            functionSelectors: forgeV1Selectors
        });

        vm.prank(SERVICE_OWNER);
        IDiamondCut(forge).diamondCut(upgradeCuts, address(0), "");

        // Step 5: Verify v2 blacklist functions now work
        assertFalse(BaseForgeV2(forge).isBlacklistManager(SERVICE_OWNER));
        assertFalse(BaseForgeV2(forge).isBlacklisted(USER.addr));

        // Step 6: Setup blacklist manager and blacklist user
        address[] memory managers = new address[](1);
        managers[0] = SERVICE_OWNER;
        vm.prank(SERVICE_OWNER);
        IBaseForgeFacetV2(forge).setBlacklistManager(managers, true);

        address[] memory accounts = new address[](1);
        accounts[0] = USER.addr;
        vm.prank(SERVICE_OWNER);
        IBaseForgeFacetV2(forge).updateBlacklist(accounts, true);

        assertTrue(BaseForgeV2(forge).isBlacklisted(USER.addr));

        // Step 7: Verify blacklisted user cannot mint (v2 ForgeV1 has notBlacklisted modifier)
        nonce = BaseForgeV2(forge).nonces(USER.addr);
        structHash = keccak256(
            abi.encode(ERC20_MINT_TYPE_HASH, USER.addr, address(mockERC20), amount, address(0), 0, nonce, deadline)
        );
        hash = MessageHashUtils.toTypedDataHash(domainSeparator, structHash);
        (v, r, s) = vm.sign(VALIDATOR, hash);
        validatorSig = abi.encodePacked(r, s, v);

        vm.prank(USER.addr);
        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", USER.addr));
        ForgeV1FacetV2(forge).mintERC20(address(mockERC20), amount, address(0), 0, deadline, validatorSig);

        // Step 8: Remove from blacklist and verify mint works again
        vm.prank(SERVICE_OWNER);
        IBaseForgeFacetV2(forge).updateBlacklist(accounts, false);

        assertFalse(BaseForgeV2(forge).isBlacklisted(USER.addr));

        // Mint should now succeed
        vm.prank(USER.addr);
        ForgeV1FacetV2(forge).mintERC20(address(mockERC20), amount, address(0), 0, deadline, validatorSig);
        assertEq(mockERC20.balanceOf(USER.addr), amount * 2); // 200 ether total
    }

    function test_upgrade_forge_preserves_state_after_facet_upgrade() external {
        // Deploy Forge with v1 facets
        Vm.Wallet memory USER = vm.createWallet("User");

        bytes4[] memory v1Selectors = new bytes4[](1);
        v1Selectors[0] = bytes4(keccak256("mintERC20(address,uint256,address,uint256,uint256,bytes)"));

        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1FacetV1),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: v1Selectors
        });

        vm.prank(OWNER);
        address forge =
            ForgeFactoryV1(forgeFactoryProxy).addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);

        vm.prank(OWNER);
        mockERC20 = new MockERC20(forge);

        // Mint some tokens to increase nonce
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 1 days;

        bytes32 ERC20_MINT_TYPE_HASH = keccak256(
            "ERC20Mint(address recipient,address token,uint256 amount,address feeRecipient,uint256 feeBPS,uint256 nonce,uint256 deadline)"
        );

        for (uint256 i = 0; i < 3; i++) {
            uint256 nonce = BaseForgeV1(forge).nonces(USER.addr);
            bytes32 structHash = keccak256(
                abi.encode(ERC20_MINT_TYPE_HASH, USER.addr, address(mockERC20), amount, address(0), 0, nonce, deadline)
            );
            bytes32 hash = MessageHashUtils.toTypedDataHash(BaseForgeV1(forge).DOMAIN_SEPARATOR(), structHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);
            bytes memory validatorSig = abi.encodePacked(r, s, v);

            vm.prank(USER.addr);
            ForgeV1FacetV1(forge).mintERC20(address(mockERC20), amount, address(0), 0, deadline, validatorSig);
        }

        // Record state before upgrade
        uint256 nonceBeforeUpgrade = BaseForgeV1(forge).nonces(USER.addr);
        address validatorBeforeUpgrade = BaseForgeV1(forge).validator();
        bytes32 domainSeparatorBeforeUpgrade = BaseForgeV1(forge).DOMAIN_SEPARATOR();
        uint256 balanceBeforeUpgrade = mockERC20.balanceOf(USER.addr);

        assertEq(nonceBeforeUpgrade, 3);
        assertEq(validatorBeforeUpgrade, VALIDATOR.addr);
        assertEq(balanceBeforeUpgrade, amount * 3);

        // Upgrade to v2 facets
        bytes4[] memory baseOldSelectors = new bytes4[](5);
        baseOldSelectors[0] = bytes4(keccak256("eip712Domain()"));
        baseOldSelectors[1] = bytes4(keccak256("nonces(address)"));
        baseOldSelectors[2] = BaseForgeV1.DOMAIN_SEPARATOR.selector;
        baseOldSelectors[3] = BaseForgeV1.validator.selector;
        baseOldSelectors[4] = BaseForgeFacetV1.setValidator.selector;

        bytes4[] memory baseNewSelectors = new bytes4[](4);
        baseNewSelectors[0] = BaseForgeV2.isBlacklistManager.selector;
        baseNewSelectors[1] = BaseForgeV2.isBlacklisted.selector;
        baseNewSelectors[2] = IBaseForgeFacetV2.setBlacklistManager.selector;
        baseNewSelectors[3] = IBaseForgeFacetV2.updateBlacklist.selector;

        bytes4[] memory forgeV1Selectors = new bytes4[](1);
        forgeV1Selectors[0] = bytes4(keccak256("mintERC20(address,uint256,address,uint256,uint256,bytes)"));

        IDiamondCut.FacetCut[] memory upgradeCuts = new IDiamondCut.FacetCut[](3);
        upgradeCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(baseForgeFacetV2),
            action: IDiamondCut.FacetCutAction.Replace,
            functionSelectors: baseOldSelectors
        });
        upgradeCuts[1] = IDiamondCut.FacetCut({
            facetAddress: address(baseForgeFacetV2),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: baseNewSelectors
        });
        upgradeCuts[2] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1FacetV2),
            action: IDiamondCut.FacetCutAction.Replace,
            functionSelectors: forgeV1Selectors
        });

        vm.prank(SERVICE_OWNER);
        IDiamondCut(forge).diamondCut(upgradeCuts, address(0), "");

        // Verify state is preserved after upgrade
        uint256 nonceAfterUpgrade = BaseForgeV2(forge).nonces(USER.addr);
        address validatorAfterUpgrade = BaseForgeV2(forge).validator();
        bytes32 domainSeparatorAfterUpgrade = BaseForgeV2(forge).DOMAIN_SEPARATOR();
        uint256 balanceAfterUpgrade = mockERC20.balanceOf(USER.addr);

        assertEq(nonceAfterUpgrade, nonceBeforeUpgrade, "Nonce should be preserved");
        assertEq(validatorAfterUpgrade, validatorBeforeUpgrade, "Validator should be preserved");
        assertEq(domainSeparatorAfterUpgrade, domainSeparatorBeforeUpgrade, "Domain separator should be preserved");
        assertEq(balanceAfterUpgrade, balanceBeforeUpgrade, "Balance should be preserved");

        // Verify new blacklist storage is initialized to empty (no data collision)
        assertFalse(BaseForgeV2(forge).isBlacklistManager(SERVICE_OWNER));
        assertFalse(BaseForgeV2(forge).isBlacklisted(USER.addr));
    }

    // ==================== Test: Access Control ====================

    function test_reinitialize_only_admin() external {
        // Upgrade to v2
        vm.prank(OWNER);
        UUPSUpgradeable(forgeFactoryProxy).upgradeToAndCall(address(forgeFactoryV2Impl), "");

        // Non-admin cannot call reinitialize
        vm.prank(SERVICE_OWNER);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector,
                SERVICE_OWNER,
                bytes32(0) // DEFAULT_ADMIN_ROLE
            )
        );
        ForgeFactoryV2(forgeFactoryProxy).reinitialize(address(baseForgeFacetV2));

        // Admin can call reinitialize
        vm.prank(OWNER);
        ForgeFactoryV2(forgeFactoryProxy).reinitialize(address(baseForgeFacetV2));
    }

    function test_reinitialize_only_once() external {
        // Upgrade to v2
        vm.prank(OWNER);
        UUPSUpgradeable(forgeFactoryProxy).upgradeToAndCall(address(forgeFactoryV2Impl), "");

        // First reinitialize succeeds
        vm.prank(OWNER);
        ForgeFactoryV2(forgeFactoryProxy).reinitialize(address(baseForgeFacetV2));

        // Second reinitialize fails
        vm.prank(OWNER);
        vm.expectRevert();
        ForgeFactoryV2(forgeFactoryProxy).reinitialize(address(baseForgeFacetV2));
    }

    // ==================== Test: Invalid Input ====================

    function test_reinitialize_revert_zero_address() external {
        vm.prank(OWNER);
        UUPSUpgradeable(forgeFactoryProxy).upgradeToAndCall(address(forgeFactoryV2Impl), "");

        vm.prank(OWNER);
        vm.expectRevert(abi.encodeWithSignature("TokenForgeFactory__InvalidData(bytes32)", bytes32("baseImpl")));
        ForgeFactoryV2(forgeFactoryProxy).reinitialize(address(0));
    }

    function test_reinitialize_revert_invalid_facet() external {
        vm.prank(OWNER);
        UUPSUpgradeable(forgeFactoryProxy).upgradeToAndCall(address(forgeFactoryV2Impl), "");

        // Pass an address that doesn't implement IDefaultDiamondCut correctly
        // Use an EOA address which will revert when calling defaultDiamondFacetCut
        vm.prank(OWNER);
        vm.expectRevert(); // Will revert because EOA can't respond to defaultDiamondFacetCut call
        ForgeFactoryV2(forgeFactoryProxy).reinitialize(address(0x1234));
    }
}

