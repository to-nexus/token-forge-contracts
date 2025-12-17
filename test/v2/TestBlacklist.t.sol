// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {Test, console} from "forge-std-1.9.7/Test.sol";
import {Vm} from "forge-std-1.9.7/Vm.sol";

import {IDiamondCut} from "diamond-3-hardhat-1.0.0/interfaces/IDiamondCut.sol";
import {IDiamondLoupe} from "diamond-3-hardhat-1.0.0/interfaces/IDiamondLoupe.sol";
import {IERC173} from "diamond-3-hardhat-1.0.0/interfaces/IERC173.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts-5.3.0/proxy/ERC1967/ERC1967Proxy.sol";
import {MessageHashUtils} from "@openzeppelin-contracts-5.3.0/utils/cryptography/MessageHashUtils.sol";

import {ForgeProxyCode} from "../../src/ForgeProxy.sol";
import {Diamond3Facet} from "../../src/Diamond3Facet.sol";
import {ForgeFactory} from "../../src/v2/ForgeFactory.sol";
import "../../src/v2/BaseForge.sol";
import "../../src/v2/ForgeV1.sol";
import "../../src/v2/ForgeV2.sol";
import "../../src/v2/ForgeV3.sol";
import {IBaseForgeFacetV2} from "../../src/interfaces/IBaseForgeFacet.sol";

import {MockERC20} from "../mock/MockERC20.sol";
import {MockERC721} from "../mock/MockERC721.sol";
import {MockERC1155} from "../mock/MockERC1155.sol";

import "../mock/StructHash.sol";

contract TestBlacklist is Test {
    address public constant OWNER = address(bytes20("OWNER"));
    address public constant SERVICE_OWNER = address(bytes20("SERVICE_OWNER"));

    bytes32 public constant SERVICE_NAME = bytes32("TestService");
    Vm.Wallet public VALIDATOR;
    Vm.Wallet public ACCOUNT;
    Vm.Wallet public BLACKLIST_MANAGER;

    ForgeProxyCode public forgeProxyCode;
    Diamond3Facet public diamond3Facet;
    BaseForgeFacet public baseForgeFacet;
    ForgeV1 public forgeV1;
    ForgeV2 public forgeV2;
    ForgeV3 public forgeV3;
    ForgeFactory public forgeFactory;
    address public FORGE;
    bytes32 public DOMAIN_SEPARATOR;

    MockERC20 public mockERC20;
    MockERC721 public mockERC721;
    MockERC1155 public mockERC1155;

    function setUp() public {
        VALIDATOR = vm.createWallet("Validator");
        ACCOUNT = vm.createWallet("Account");
        BLACKLIST_MANAGER = vm.createWallet("BlacklistManager");

        vm.label(OWNER, "owner");
        vm.label(SERVICE_OWNER, "serviceOwner");
        vm.startPrank(OWNER);

        // deploy facets
        diamond3Facet = new Diamond3Facet();
        baseForgeFacet = new BaseForgeFacet();
        forgeV1 = new ForgeV1();
        forgeV2 = new ForgeV2();
        forgeV3 = new ForgeV3();

        // deploy proxy code
        forgeProxyCode = new ForgeProxyCode();

        // deploy factory
        ForgeFactory forgeFactoryImpl = new ForgeFactory();
        ERC1967Proxy forgeFactoryProxy = new ERC1967Proxy(
            address(forgeFactoryImpl),
            abi.encodeCall(
                ForgeFactory.initialize,
                (OWNER, address(forgeProxyCode), address(diamond3Facet), address(baseForgeFacet))
            )
        );

        // setup diamond cuts
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](3);
        {
            bytes4[] memory v1Selectors = new bytes4[](16);
            v1Selectors[0] = forgeV1.mintERC20.selector;
            v1Selectors[1] = forgeV1.mintERC721.selector;
            v1Selectors[2] = forgeV1.mintERC1155.selector;
            v1Selectors[3] = forgeV1.mintERC1155Batch.selector;
            v1Selectors[4] = forgeV1.transferERC20.selector;
            v1Selectors[5] = forgeV1.transferERC721.selector;
            v1Selectors[6] = forgeV1.transferERC1155.selector;
            v1Selectors[7] = forgeV1.transferERC1155Batch.selector;
            v1Selectors[8] = forgeV1.transferFromERC20.selector;
            v1Selectors[9] = forgeV1.burnERC20.selector;
            v1Selectors[10] = forgeV1.burnERC721.selector;
            v1Selectors[11] = forgeV1.burnERC1155.selector;
            v1Selectors[12] = forgeV1.burnERC1155Batch.selector;
            v1Selectors[13] = forgeV1.onERC721Received.selector;
            v1Selectors[14] = forgeV1.onERC1155Received.selector;
            v1Selectors[15] = forgeV1.onERC1155BatchReceived.selector;

            addCuts[0] = IDiamondCut.FacetCut({
                facetAddress: address(forgeV1), action: IDiamondCut.FacetCutAction.Add, functionSelectors: v1Selectors
            });
        }
        {
            bytes4[] memory v2Selectors = new bytes4[](12);
            v2Selectors[0] = forgeV2.mintERC20.selector;
            v2Selectors[1] = forgeV2.mintERC721.selector;
            v2Selectors[2] = forgeV2.mintERC1155.selector;
            v2Selectors[3] = forgeV2.mintERC1155Batch.selector;
            v2Selectors[4] = forgeV2.transferERC20.selector;
            v2Selectors[5] = forgeV2.transferERC721.selector;
            v2Selectors[6] = forgeV2.transferERC1155.selector;
            v2Selectors[7] = forgeV2.transferERC1155Batch.selector;
            v2Selectors[8] = forgeV2.burnERC20.selector;
            v2Selectors[9] = forgeV2.burnERC721.selector;
            v2Selectors[10] = forgeV2.burnERC1155.selector;
            v2Selectors[11] = forgeV2.burnERC1155Batch.selector;

            addCuts[1] = IDiamondCut.FacetCut({
                facetAddress: address(forgeV2), action: IDiamondCut.FacetCutAction.Add, functionSelectors: v2Selectors
            });
        }
        {
            bytes4[] memory v3Selectors = new bytes4[](8);
            v3Selectors[0] = forgeV3.mintERC20.selector;
            v3Selectors[1] = forgeV3.mintERC721.selector;
            v3Selectors[2] = forgeV3.mintERC1155.selector;
            v3Selectors[3] = forgeV3.mintERC1155Batch.selector;
            v3Selectors[4] = forgeV3.transferERC20.selector;
            v3Selectors[5] = forgeV3.transferERC721.selector;
            v3Selectors[6] = forgeV3.transferERC1155.selector;
            v3Selectors[7] = forgeV3.transferERC1155Batch.selector;

            addCuts[2] = IDiamondCut.FacetCut({
                facetAddress: address(forgeV3), action: IDiamondCut.FacetCutAction.Add, functionSelectors: v3Selectors
            });
        }

        forgeFactory = ForgeFactory(address(forgeFactoryProxy));
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock tokens
        mockERC20 = new MockERC20(FORGE);
        mockERC721 = new MockERC721(FORGE);
        mockERC1155 = new MockERC1155(FORGE);

        vm.stopPrank();
    }

    // ==================== Helper Functions ====================

    function _setupBlacklistManager(address manager) internal {
        address[] memory managers = new address[](1);
        managers[0] = manager;
        vm.prank(SERVICE_OWNER);
        IBaseForgeFacetV2(FORGE).setBlacklistManager(managers, true);
    }

    function _addToBlacklist(address account) internal {
        address[] memory accounts = new address[](1);
        accounts[0] = account;
        vm.prank(BLACKLIST_MANAGER.addr);
        IBaseForgeFacetV2(FORGE).updateBlacklist(accounts, true);
    }

    function _removeFromBlacklist(address account) internal {
        address[] memory accounts = new address[](1);
        accounts[0] = account;
        vm.prank(BLACKLIST_MANAGER.addr);
        IBaseForgeFacetV2(FORGE).updateBlacklist(accounts, false);
    }

    function _calcUUID(uint256 nonce) internal view returns (uint256) {
        return uint256(keccak256(abi.encode(FORGE, ACCOUNT.addr, nonce)));
    }

    function _sign_with_domain_separator(bytes32 structHash, Vm.Wallet memory signer) internal returns (bytes memory) {
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(signer, hash);
        return abi.encodePacked(r, s, v);
    }

    // ==================== 1. Blacklist Manager Tests ====================

    function test_setBlacklistManager_success() external {
        address[] memory managers = new address[](1);
        managers[0] = BLACKLIST_MANAGER.addr;

        assertFalse(BaseForge(FORGE).isBlacklistManager(BLACKLIST_MANAGER.addr));

        vm.expectEmit(true, true, false, false, FORGE);
        emit BaseForge.BlacklistManagerUpdated(BLACKLIST_MANAGER.addr, true);

        vm.prank(SERVICE_OWNER);
        IBaseForgeFacetV2(FORGE).setBlacklistManager(managers, true);

        assertTrue(BaseForge(FORGE).isBlacklistManager(BLACKLIST_MANAGER.addr));
    }

    function test_setBlacklistManager_remove() external {
        // setup manager first
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);
        assertTrue(BaseForge(FORGE).isBlacklistManager(BLACKLIST_MANAGER.addr));

        // remove manager
        address[] memory managers = new address[](1);
        managers[0] = BLACKLIST_MANAGER.addr;

        vm.expectEmit(true, true, false, false, FORGE);
        emit BaseForge.BlacklistManagerUpdated(BLACKLIST_MANAGER.addr, false);

        vm.prank(SERVICE_OWNER);
        IBaseForgeFacetV2(FORGE).setBlacklistManager(managers, false);

        assertFalse(BaseForge(FORGE).isBlacklistManager(BLACKLIST_MANAGER.addr));
    }

    function test_setBlacklistManager_multiple() external {
        address manager1 = address(0x1001);
        address manager2 = address(0x1002);
        address manager3 = address(0x1003);

        address[] memory managers = new address[](3);
        managers[0] = manager1;
        managers[1] = manager2;
        managers[2] = manager3;

        vm.prank(SERVICE_OWNER);
        IBaseForgeFacetV2(FORGE).setBlacklistManager(managers, true);

        assertTrue(BaseForge(FORGE).isBlacklistManager(manager1));
        assertTrue(BaseForge(FORGE).isBlacklistManager(manager2));
        assertTrue(BaseForge(FORGE).isBlacklistManager(manager3));
    }

    function test_setBlacklistManager_revert_notOwner() external {
        address[] memory managers = new address[](1);
        managers[0] = BLACKLIST_MANAGER.addr;

        vm.prank(ACCOUNT.addr);
        vm.expectRevert("LibDiamond: Must be contract owner");
        IBaseForgeFacetV2(FORGE).setBlacklistManager(managers, true);
    }

    function test_setBlacklistManager_revert_zeroAddress() external {
        address[] memory managers = new address[](1);
        managers[0] = address(0);

        vm.prank(SERVICE_OWNER);
        vm.expectRevert(abi.encodeWithSignature("BaseForge__ZeroAddress()"));
        IBaseForgeFacetV2(FORGE).setBlacklistManager(managers, true);
    }

    // ==================== 2. Blacklist Update Tests ====================

    function test_updateBlacklist_success() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);

        address[] memory accounts = new address[](1);
        accounts[0] = ACCOUNT.addr;

        assertFalse(BaseForge(FORGE).isBlacklisted(ACCOUNT.addr));

        vm.expectEmit(true, true, false, false, FORGE);
        emit BaseForge.BlacklistUpdated(ACCOUNT.addr, true);

        vm.prank(BLACKLIST_MANAGER.addr);
        IBaseForgeFacetV2(FORGE).updateBlacklist(accounts, true);

        assertTrue(BaseForge(FORGE).isBlacklisted(ACCOUNT.addr));
    }

    function test_updateBlacklist_remove() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);
        _addToBlacklist(ACCOUNT.addr);
        assertTrue(BaseForge(FORGE).isBlacklisted(ACCOUNT.addr));

        address[] memory accounts = new address[](1);
        accounts[0] = ACCOUNT.addr;

        vm.expectEmit(true, true, false, false, FORGE);
        emit BaseForge.BlacklistUpdated(ACCOUNT.addr, false);

        vm.prank(BLACKLIST_MANAGER.addr);
        IBaseForgeFacetV2(FORGE).updateBlacklist(accounts, false);

        assertFalse(BaseForge(FORGE).isBlacklisted(ACCOUNT.addr));
    }

    function test_updateBlacklist_multiple() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);

        address account1 = address(0x2001);
        address account2 = address(0x2002);
        address account3 = address(0x2003);

        address[] memory accounts = new address[](3);
        accounts[0] = account1;
        accounts[1] = account2;
        accounts[2] = account3;

        vm.prank(BLACKLIST_MANAGER.addr);
        IBaseForgeFacetV2(FORGE).updateBlacklist(accounts, true);

        assertTrue(BaseForge(FORGE).isBlacklisted(account1));
        assertTrue(BaseForge(FORGE).isBlacklisted(account2));
        assertTrue(BaseForge(FORGE).isBlacklisted(account3));
    }

    function test_updateBlacklist_revert_notManager() external {
        address[] memory accounts = new address[](1);
        accounts[0] = ACCOUNT.addr;

        vm.prank(ACCOUNT.addr);
        vm.expectRevert(abi.encodeWithSignature("BaseForge__NotBlacklistManager()"));
        IBaseForgeFacetV2(FORGE).updateBlacklist(accounts, true);
    }

    function test_updateBlacklist_revert_zeroAddress() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);

        address[] memory accounts = new address[](1);
        accounts[0] = address(0);

        vm.prank(BLACKLIST_MANAGER.addr);
        vm.expectRevert(abi.encodeWithSignature("BaseForge__ZeroAddress()"));
        IBaseForgeFacetV2(FORGE).updateBlacklist(accounts, true);
    }

    // ==================== 3. ERC20 Blacklist Tests ====================

    function test_blacklisted_mintERC20_v1() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);
        _addToBlacklist(ACCOUNT.addr);

        address token = address(mockERC20);
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);

        bytes32 structHash =
            keccak256(abi.encode(ERC20_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, amount, address(0), 0, nonce, deadline));
        bytes memory signature = _sign_with_domain_separator(structHash, VALIDATOR);

        vm.prank(ACCOUNT.addr);
        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV1(FORGE).mintERC20(token, amount, address(0), 0, deadline, signature);
    }

    function test_blacklisted_mintERC20_v2() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);
        _addToBlacklist(ACCOUNT.addr);

        address token = address(mockERC20);
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);

        bytes32 recipientStructHash =
            keccak256(abi.encode(ERC20_MINT_TYPE_HASH_V2, token, amount, address(0), 0, nonce, deadline));
        bytes memory recipientSig = _sign_with_domain_separator(recipientStructHash, ACCOUNT);

        bytes32 validatorStructHash =
            keccak256(abi.encode(ERC20_VALIDATOR_MINT_TYPE_HASH_V2, ACCOUNT.addr, keccak256(recipientSig)));
        bytes memory validatorSig = _sign_with_domain_separator(validatorStructHash, VALIDATOR);

        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV2(FORGE).mintERC20(ACCOUNT.addr, token, amount, address(0), 0, deadline, recipientSig, validatorSig);
    }

    function test_blacklisted_mintERC20_v3() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);
        _addToBlacklist(ACCOUNT.addr);

        address token = address(mockERC20);
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);

        bytes32 structHash =
            keccak256(abi.encode(ERC20_MINT_TYPE_HASH_V3, ACCOUNT.addr, token, amount, address(0), 0, nonce, deadline));
        bytes memory signature = _sign_with_domain_separator(structHash, VALIDATOR);

        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV3(FORGE).mintERC20(ACCOUNT.addr, token, amount, address(0), 0, deadline, signature);
    }

    function test_blacklisted_transferERC20_v1() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);
        _addToBlacklist(ACCOUNT.addr);

        // First mint tokens to FORGE
        mockERC20.forceMint(FORGE, 100 ether);

        address token = address(mockERC20);
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);

        bytes32 structHash = keccak256(
            abi.encode(ERC20_TRANSFER_TYPE_HASH_V1, ACCOUNT.addr, token, amount, address(0), 0, nonce, deadline)
        );
        bytes memory signature = _sign_with_domain_separator(structHash, VALIDATOR);

        vm.prank(ACCOUNT.addr);
        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV1(FORGE).transferERC20(token, amount, address(0), 0, deadline, signature);
    }

    function test_blacklisted_burnERC20_v1() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);

        // First mint tokens to ACCOUNT
        mockERC20.forceMint(ACCOUNT.addr, 100 ether);

        // Approve FORGE to burn
        vm.prank(ACCOUNT.addr);
        mockERC20.approve(FORGE, 100 ether);

        // Now blacklist the account
        _addToBlacklist(ACCOUNT.addr);

        address token = address(mockERC20);
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);

        bytes32 structHash =
            keccak256(abi.encode(ERC20_BURN_TYPE_HASH_V1, ACCOUNT.addr, token, amount, address(0), 0, nonce, deadline));
        bytes memory signature = _sign_with_domain_separator(structHash, VALIDATOR);

        vm.prank(ACCOUNT.addr);
        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV1(FORGE).burnERC20(token, amount, address(0), 0, deadline, signature);
    }

    // ==================== 4. ERC721 Blacklist Tests ====================

    function test_blacklisted_mintERC721_v1() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);
        _addToBlacklist(ACCOUNT.addr);

        address token = address(mockERC721);
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
        bytes memory data = "test";

        bytes32 structHash = keccak256(
            abi.encode(ERC721_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, keccak256(data), nonce, deadline)
        );
        bytes memory signature = _sign_with_domain_separator(structHash, VALIDATOR);

        vm.prank(ACCOUNT.addr);
        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV1(FORGE).mintERC721(token, tokenID, data, deadline, signature);
    }

    function test_blacklisted_mintERC721_v2() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);
        _addToBlacklist(ACCOUNT.addr);

        address token = address(mockERC721);
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
        bytes memory data = "test";

        bytes32 recipientStructHash =
            keccak256(abi.encode(ERC721_MINT_TYPE_HASH_V2, token, tokenID, keccak256(data), nonce, deadline));
        bytes memory recipientSig = _sign_with_domain_separator(recipientStructHash, ACCOUNT);

        bytes32 validatorStructHash =
            keccak256(abi.encode(ERC721_VALIDATOR_MINT_TYPE_HASH_V2, ACCOUNT.addr, keccak256(recipientSig)));
        bytes memory validatorSig = _sign_with_domain_separator(validatorStructHash, VALIDATOR);

        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV2(FORGE).mintERC721(ACCOUNT.addr, token, tokenID, data, deadline, recipientSig, validatorSig);
    }

    function test_blacklisted_mintERC721_v3() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);
        _addToBlacklist(ACCOUNT.addr);

        address token = address(mockERC721);
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
        bytes memory data = "test";

        bytes32 structHash = keccak256(
            abi.encode(ERC721_MINT_TYPE_HASH_V3, ACCOUNT.addr, token, tokenID, keccak256(data), nonce, deadline)
        );
        bytes memory signature = _sign_with_domain_separator(structHash, VALIDATOR);

        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV3(FORGE).mintERC721(ACCOUNT.addr, token, tokenID, data, deadline, signature);
    }

    function test_blacklisted_burnERC721_v1() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);

        // First mint NFT to ACCOUNT
        mockERC721.forceMint(ACCOUNT.addr, 1, "");

        // Approve FORGE to burn
        vm.prank(ACCOUNT.addr);
        mockERC721.approve(FORGE, 1);

        // Now blacklist the account
        _addToBlacklist(ACCOUNT.addr);

        address token = address(mockERC721);
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);

        bytes32 structHash =
            keccak256(abi.encode(ERC721_BURN_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, nonce, deadline));
        bytes memory signature = _sign_with_domain_separator(structHash, VALIDATOR);

        vm.prank(ACCOUNT.addr);
        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV1(FORGE).burnERC721(token, tokenID, deadline, signature);
    }

    // ==================== 5. ERC1155 Blacklist Tests ====================

    function test_blacklisted_mintERC1155_v1() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);
        _addToBlacklist(ACCOUNT.addr);

        address token = address(mockERC1155);
        uint256 tokenID = 1;
        uint256 amount = 10;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
        bytes memory data = "test";

        bytes32 structHash = keccak256(
            abi.encode(
                ERC1155_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, amount, keccak256(data), nonce, deadline
            )
        );
        bytes memory signature = _sign_with_domain_separator(structHash, VALIDATOR);

        vm.prank(ACCOUNT.addr);
        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV1(FORGE).mintERC1155(token, tokenID, amount, data, deadline, signature);
    }

    function test_blacklisted_mintERC1155_v3() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);
        _addToBlacklist(ACCOUNT.addr);

        address token = address(mockERC1155);
        uint256 tokenID = 1;
        uint256 amount = 10;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
        bytes memory data = "test";

        bytes32 structHash = keccak256(
            abi.encode(
                ERC1155_MINT_TYPE_HASH_V3, ACCOUNT.addr, token, tokenID, amount, keccak256(data), nonce, deadline
            )
        );
        bytes memory signature = _sign_with_domain_separator(structHash, VALIDATOR);

        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV3(FORGE).mintERC1155(ACCOUNT.addr, token, tokenID, amount, data, deadline, signature);
    }

    function test_blacklisted_mintERC1155Batch_v1() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);
        _addToBlacklist(ACCOUNT.addr);

        address token = address(mockERC1155);
        uint256[] memory tokenIDs = new uint256[](2);
        tokenIDs[0] = 1;
        tokenIDs[1] = 2;
        uint256[] memory amounts = new uint256[](2);
        amounts[0] = 10;
        amounts[1] = 20;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
        bytes memory data = "test";

        bytes32 structHash = keccak256(
            abi.encode(
                ERC1155_MINT_BATCH_TYPE_HASH_V1,
                ACCOUNT.addr,
                token,
                keccak256(abi.encodePacked(tokenIDs)),
                keccak256(abi.encodePacked(amounts)),
                keccak256(data),
                nonce,
                deadline
            )
        );
        bytes memory signature = _sign_with_domain_separator(structHash, VALIDATOR);

        vm.prank(ACCOUNT.addr);
        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV1(FORGE).mintERC1155Batch(token, tokenIDs, amounts, data, deadline, signature);
    }

    function test_blacklisted_burnERC1155_v1() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);

        // First mint tokens to ACCOUNT
        mockERC1155.forceMint(ACCOUNT.addr, 1, 100);

        // Approve FORGE
        vm.prank(ACCOUNT.addr);
        mockERC1155.setApprovalForAll(FORGE, true);

        // Now blacklist the account
        _addToBlacklist(ACCOUNT.addr);

        address token = address(mockERC1155);
        uint256 tokenID = 1;
        uint256 amount = 50;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);

        bytes32 structHash =
            keccak256(abi.encode(ERC1155_BURN_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, amount, nonce, deadline));
        bytes memory signature = _sign_with_domain_separator(structHash, VALIDATOR);

        vm.prank(ACCOUNT.addr);
        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV1(FORGE).burnERC1155(token, tokenID, amount, deadline, signature);
    }

    // ==================== 6. Unblacklist Tests ====================

    function test_unblacklist_then_mint_success() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);
        _addToBlacklist(ACCOUNT.addr);

        // Verify blacklisted
        assertTrue(BaseForge(FORGE).isBlacklisted(ACCOUNT.addr));

        // Remove from blacklist
        _removeFromBlacklist(ACCOUNT.addr);
        assertFalse(BaseForge(FORGE).isBlacklisted(ACCOUNT.addr));

        // Now mint should succeed
        address token = address(mockERC20);
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);

        bytes32 structHash =
            keccak256(abi.encode(ERC20_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, amount, address(0), 0, nonce, deadline));
        bytes memory signature = _sign_with_domain_separator(structHash, VALIDATOR);

        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).mintERC20(token, amount, address(0), 0, deadline, signature);

        assertEq(mockERC20.balanceOf(ACCOUNT.addr), amount);
    }

    // ==================== 7. View Function Tests ====================

    function test_isBlacklistManager_view() external {
        assertFalse(BaseForge(FORGE).isBlacklistManager(BLACKLIST_MANAGER.addr));

        _setupBlacklistManager(BLACKLIST_MANAGER.addr);

        assertTrue(BaseForge(FORGE).isBlacklistManager(BLACKLIST_MANAGER.addr));
        assertFalse(BaseForge(FORGE).isBlacklistManager(ACCOUNT.addr));
    }

    function test_isBlacklisted_view() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);

        assertFalse(BaseForge(FORGE).isBlacklisted(ACCOUNT.addr));

        _addToBlacklist(ACCOUNT.addr);

        assertTrue(BaseForge(FORGE).isBlacklisted(ACCOUNT.addr));
        assertFalse(BaseForge(FORGE).isBlacklisted(BLACKLIST_MANAGER.addr));
    }

    function test_getBlacklistManagers_empty() external view {
        address[] memory managers = BaseForge(FORGE).getBlacklistManagers();
        assertEq(managers.length, 0);
    }

    function test_getBlacklistManagers_single() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);

        address[] memory managers = BaseForge(FORGE).getBlacklistManagers();
        assertEq(managers.length, 1);
        assertEq(managers[0], BLACKLIST_MANAGER.addr);
    }

    function test_getBlacklistManagers_multiple() external {
        address manager1 = address(0x1001);
        address manager2 = address(0x1002);
        address manager3 = address(0x1003);

        address[] memory managersToAdd = new address[](3);
        managersToAdd[0] = manager1;
        managersToAdd[1] = manager2;
        managersToAdd[2] = manager3;

        vm.prank(SERVICE_OWNER);
        IBaseForgeFacetV2(FORGE).setBlacklistManager(managersToAdd, true);

        address[] memory managers = BaseForge(FORGE).getBlacklistManagers();
        assertEq(managers.length, 3);

        // Verify all managers are in the list (order may vary due to EnumerableSet)
        bool found1 = false;
        bool found2 = false;
        bool found3 = false;
        for (uint256 i = 0; i < managers.length; i++) {
            if (managers[i] == manager1) found1 = true;
            if (managers[i] == manager2) found2 = true;
            if (managers[i] == manager3) found3 = true;
        }
        assertTrue(found1 && found2 && found3);
    }

    function test_getBlacklistManagers_afterRemove() external {
        address manager1 = address(0x1001);
        address manager2 = address(0x1002);

        address[] memory managersToAdd = new address[](2);
        managersToAdd[0] = manager1;
        managersToAdd[1] = manager2;

        vm.prank(SERVICE_OWNER);
        IBaseForgeFacetV2(FORGE).setBlacklistManager(managersToAdd, true);

        assertEq(BaseForge(FORGE).getBlacklistManagers().length, 2);

        // Remove one manager
        address[] memory managerToRemove = new address[](1);
        managerToRemove[0] = manager1;
        vm.prank(SERVICE_OWNER);
        IBaseForgeFacetV2(FORGE).setBlacklistManager(managerToRemove, false);

        address[] memory managers = BaseForge(FORGE).getBlacklistManagers();
        assertEq(managers.length, 1);
        assertEq(managers[0], manager2);
    }

    function test_getBlacklistedAccounts_empty() external view {
        address[] memory blacklisted = BaseForge(FORGE).getBlacklistedAccounts();
        assertEq(blacklisted.length, 0);
    }

    function test_getBlacklistedAccounts_single() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);
        _addToBlacklist(ACCOUNT.addr);

        address[] memory blacklisted = BaseForge(FORGE).getBlacklistedAccounts();
        assertEq(blacklisted.length, 1);
        assertEq(blacklisted[0], ACCOUNT.addr);
    }

    function test_getBlacklistedAccounts_multiple() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);

        address account1 = address(0x2001);
        address account2 = address(0x2002);
        address account3 = address(0x2003);

        address[] memory accountsToAdd = new address[](3);
        accountsToAdd[0] = account1;
        accountsToAdd[1] = account2;
        accountsToAdd[2] = account3;

        vm.prank(BLACKLIST_MANAGER.addr);
        IBaseForgeFacetV2(FORGE).updateBlacklist(accountsToAdd, true);

        address[] memory blacklisted = BaseForge(FORGE).getBlacklistedAccounts();
        assertEq(blacklisted.length, 3);

        // Verify all accounts are in the list (order may vary due to EnumerableSet)
        bool found1 = false;
        bool found2 = false;
        bool found3 = false;
        for (uint256 i = 0; i < blacklisted.length; i++) {
            if (blacklisted[i] == account1) found1 = true;
            if (blacklisted[i] == account2) found2 = true;
            if (blacklisted[i] == account3) found3 = true;
        }
        assertTrue(found1 && found2 && found3);
    }

    function test_getBlacklistedAccounts_afterRemove() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);

        address account1 = address(0x2001);
        address account2 = address(0x2002);

        address[] memory accountsToAdd = new address[](2);
        accountsToAdd[0] = account1;
        accountsToAdd[1] = account2;

        vm.prank(BLACKLIST_MANAGER.addr);
        IBaseForgeFacetV2(FORGE).updateBlacklist(accountsToAdd, true);

        assertEq(BaseForge(FORGE).getBlacklistedAccounts().length, 2);

        // Remove one account
        address[] memory accountToRemove = new address[](1);
        accountToRemove[0] = account1;
        vm.prank(BLACKLIST_MANAGER.addr);
        IBaseForgeFacetV2(FORGE).updateBlacklist(accountToRemove, false);

        address[] memory blacklisted = BaseForge(FORGE).getBlacklistedAccounts();
        assertEq(blacklisted.length, 1);
        assertEq(blacklisted[0], account2);
    }

    // ==================== Additional ERC20 Tests ====================

    function test_blacklisted_transferERC20_v2() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);
        _addToBlacklist(ACCOUNT.addr);

        mockERC20.forceMint(FORGE, 100 ether);

        address token = address(mockERC20);
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);

        bytes32 recipientStructHash =
            keccak256(abi.encode(ERC20_TRANSFER_TYPE_HASH_V2, token, amount, address(0), 0, nonce, deadline));
        bytes memory recipientSig = _sign_with_domain_separator(recipientStructHash, ACCOUNT);

        bytes32 validatorStructHash =
            keccak256(abi.encode(ERC20_VALIDATOR_TRANSFER_TYPE_HASH_V2, ACCOUNT.addr, keccak256(recipientSig)));
        bytes memory validatorSig = _sign_with_domain_separator(validatorStructHash, VALIDATOR);

        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV2(FORGE).transferERC20(ACCOUNT.addr, token, amount, address(0), 0, deadline, recipientSig, validatorSig);
    }

    function test_blacklisted_transferERC20_v3() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);
        _addToBlacklist(ACCOUNT.addr);

        mockERC20.forceMint(FORGE, 100 ether);

        address token = address(mockERC20);
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);

        bytes32 structHash = keccak256(
            abi.encode(ERC20_TRANSFER_TYPE_HASH_V3, ACCOUNT.addr, token, amount, address(0), 0, nonce, deadline)
        );
        bytes memory signature = _sign_with_domain_separator(structHash, VALIDATOR);

        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV3(FORGE).transferERC20(ACCOUNT.addr, token, amount, address(0), 0, deadline, signature);
    }

    function test_blacklisted_transferFromERC20_v1() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);

        mockERC20.forceMint(ACCOUNT.addr, 100 ether);
        vm.prank(ACCOUNT.addr);
        mockERC20.approve(FORGE, 100 ether);

        _addToBlacklist(ACCOUNT.addr);

        address token = address(mockERC20);
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);

        bytes32 structHash = keccak256(
            abi.encode(ERC20_TRANSFER_FROM_TYPE_HASH_V1, ACCOUNT.addr, token, amount, address(0), 0, nonce, deadline)
        );
        bytes memory signature = _sign_with_domain_separator(structHash, VALIDATOR);

        vm.prank(ACCOUNT.addr);
        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV1(FORGE).transferFromERC20(token, amount, address(0), 0, deadline, signature);
    }

    function test_blacklisted_burnERC20_v2() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);

        mockERC20.forceMint(ACCOUNT.addr, 100 ether);
        vm.prank(ACCOUNT.addr);
        mockERC20.approve(FORGE, 100 ether);

        _addToBlacklist(ACCOUNT.addr);

        address token = address(mockERC20);
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);

        bytes32 fromStructHash =
            keccak256(abi.encode(ERC20_BURN_TYPE_HASH_V2, token, amount, address(0), 0, nonce, deadline));
        bytes memory fromSig = _sign_with_domain_separator(fromStructHash, ACCOUNT);

        bytes32 validatorStructHash =
            keccak256(abi.encode(ERC20_VALIDATOR_BURN_TYPE_HASH_V2, ACCOUNT.addr, keccak256(fromSig)));
        bytes memory validatorSig = _sign_with_domain_separator(validatorStructHash, VALIDATOR);

        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV2(FORGE).burnERC20(ACCOUNT.addr, token, amount, address(0), 0, deadline, fromSig, validatorSig);
    }

    // ==================== Additional ERC721 Tests ====================

    function test_blacklisted_transferERC721_v1() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);

        // First mint NFT to FORGE
        mockERC721.forceMint(FORGE, 1, "");

        _addToBlacklist(ACCOUNT.addr);

        address token = address(mockERC721);
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
        bytes memory data = "test";

        bytes32 structHash = keccak256(
            abi.encode(ERC721_TRANSFER_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, keccak256(data), nonce, deadline)
        );
        bytes memory signature = _sign_with_domain_separator(structHash, VALIDATOR);

        vm.prank(ACCOUNT.addr);
        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV1(FORGE).transferERC721(token, tokenID, data, deadline, signature);
    }

    function test_blacklisted_transferERC721_v3() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);

        mockERC721.forceMint(FORGE, 1, "");

        _addToBlacklist(ACCOUNT.addr);

        address token = address(mockERC721);
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
        bytes memory data = "test";

        bytes32 structHash = keccak256(
            abi.encode(ERC721_TRANSFER_TYPE_HASH_V3, ACCOUNT.addr, token, tokenID, keccak256(data), nonce, deadline)
        );
        bytes memory signature = _sign_with_domain_separator(structHash, VALIDATOR);

        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV3(FORGE).transferERC721(ACCOUNT.addr, token, tokenID, data, deadline, signature);
    }

    function test_blacklisted_burnERC721_v2() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);

        mockERC721.forceMint(ACCOUNT.addr, 1, "");
        vm.prank(ACCOUNT.addr);
        mockERC721.approve(FORGE, 1);

        _addToBlacklist(ACCOUNT.addr);

        address token = address(mockERC721);
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);

        bytes32 fromStructHash = keccak256(abi.encode(ERC721_BURN_TYPE_HASH_V2, token, tokenID, nonce, deadline));
        bytes memory fromSig = _sign_with_domain_separator(fromStructHash, ACCOUNT);

        bytes32 validatorStructHash =
            keccak256(abi.encode(ERC721_VALIDATOR_BURN_TYPE_HASH_V2, ACCOUNT.addr, keccak256(fromSig)));
        bytes memory validatorSig = _sign_with_domain_separator(validatorStructHash, VALIDATOR);

        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV2(FORGE).burnERC721(ACCOUNT.addr, token, tokenID, deadline, fromSig, validatorSig);
    }

    // ==================== Additional ERC1155 Tests ====================

    function test_blacklisted_mintERC1155_v2() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);
        _addToBlacklist(ACCOUNT.addr);

        address token = address(mockERC1155);
        uint256 tokenID = 1;
        uint256 amount = 10;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
        bytes memory data = "test";

        bytes32 recipientStructHash =
            keccak256(abi.encode(ERC1155_MINT_TYPE_HASH_V2, token, tokenID, amount, keccak256(data), nonce, deadline));
        bytes memory recipientSig = _sign_with_domain_separator(recipientStructHash, ACCOUNT);

        bytes32 validatorStructHash =
            keccak256(abi.encode(ERC1155_VALIDATOR_MINT_TYPE_HASH_V2, ACCOUNT.addr, keccak256(recipientSig)));
        bytes memory validatorSig = _sign_with_domain_separator(validatorStructHash, VALIDATOR);

        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV2(FORGE).mintERC1155(ACCOUNT.addr, token, tokenID, amount, data, deadline, recipientSig, validatorSig);
    }

    function test_blacklisted_transferERC1155_v1() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);

        mockERC1155.forceMint(FORGE, 1, 100);

        _addToBlacklist(ACCOUNT.addr);

        address token = address(mockERC1155);
        uint256 tokenID = 1;
        uint256 amount = 50;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
        bytes memory data = "test";

        bytes32 structHash = keccak256(
            abi.encode(
                ERC1155_TRANSFER_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, amount, keccak256(data), nonce, deadline
            )
        );
        bytes memory signature = _sign_with_domain_separator(structHash, VALIDATOR);

        vm.prank(ACCOUNT.addr);
        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV1(FORGE).transferERC1155(token, tokenID, amount, data, deadline, signature);
    }

    function test_blacklisted_transferERC1155_v3() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);

        mockERC1155.forceMint(FORGE, 1, 100);

        _addToBlacklist(ACCOUNT.addr);

        address token = address(mockERC1155);
        uint256 tokenID = 1;
        uint256 amount = 50;
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
        bytes memory data = "test";

        bytes32 structHash = keccak256(
            abi.encode(
                ERC1155_TRANSFER_TYPE_HASH_V3, ACCOUNT.addr, token, tokenID, amount, keccak256(data), nonce, deadline
            )
        );
        bytes memory signature = _sign_with_domain_separator(structHash, VALIDATOR);

        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV3(FORGE).transferERC1155(ACCOUNT.addr, token, tokenID, amount, data, deadline, signature);
    }

    function test_blacklisted_burnERC1155Batch_v1() external {
        _setupBlacklistManager(BLACKLIST_MANAGER.addr);

        uint256[] memory tokenIDs = new uint256[](2);
        tokenIDs[0] = 1;
        tokenIDs[1] = 2;
        uint256[] memory amounts = new uint256[](2);
        amounts[0] = 100;
        amounts[1] = 200;

        mockERC1155.forceMintBatch(ACCOUNT.addr, tokenIDs, amounts);
        vm.prank(ACCOUNT.addr);
        mockERC1155.setApprovalForAll(FORGE, true);

        _addToBlacklist(ACCOUNT.addr);

        address token = address(mockERC1155);
        uint256 deadline = block.timestamp + 1 days;
        uint256 nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);

        bytes32 structHash = keccak256(
            abi.encode(
                ERC1155_BURN_BATCH_TYPE_HASH_V1,
                ACCOUNT.addr,
                token,
                keccak256(abi.encodePacked(tokenIDs)),
                keccak256(abi.encodePacked(amounts)),
                nonce,
                deadline
            )
        );
        bytes memory signature = _sign_with_domain_separator(structHash, VALIDATOR);

        vm.prank(ACCOUNT.addr);
        vm.expectRevert(abi.encodeWithSignature("BaseForge__Blacklisted(address)", ACCOUNT.addr));
        ForgeV1(FORGE).burnERC1155Batch(token, tokenIDs, amounts, deadline, signature);
    }
}

