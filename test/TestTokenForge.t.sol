// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std-1.9.7/Test.sol";
import {Vm} from "forge-std-1.9.7/Vm.sol";

import {IDiamondCut} from "diamond-3-hardhat-1.0.0/interfaces/IDiamondCut.sol";
import {IDiamondLoupe} from "diamond-3-hardhat-1.0.0/interfaces/IDiamondLoupe.sol";
import {IERC173} from "diamond-3-hardhat-1.0.0/interfaces/IERC173.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts-5.3.0/proxy/ERC1967/ERC1967Proxy.sol";
import {MessageHashUtils} from "@openzeppelin-contracts-5.3.0/utils/cryptography/MessageHashUtils.sol";

import {ForgeProxyCode} from "../src/ForgeProxy.sol";
import {Diamond3Facet} from "../src/Diamond3Facet.sol";
import {TokenForgeFactory} from "../src/TokenForgeFactory.sol";
import "../src/BaseForge.sol";
import "../src/ForgeV1.sol";
import "../src/ForgeV2.sol";
import "../src/ForgeV3.sol";

import {MockERC20} from "./mock/MockERC20.sol";
import {MockERC721} from "./mock/MockERC721.sol";
import {MockERC1155} from "./mock/MockERC1155.sol";

import "./mock/StructHash.sol";

contract TestTokenForgeFactory is Test {
    address public constant OWNER = address(bytes20("OWNER"));
    address public constant SERVICE_OWNER = address(bytes20("SERVICE_OWNER"));

    string public constant SERVICE_NAME = "TestService";
    Vm.Wallet public VALIDATOR = vm.createWallet("Validator");
    Vm.Wallet public ACCOUNT = vm.createWallet("Account");

    ForgeProxyCode public forgeProxyCode;
    Diamond3Facet public diamond3Facet;
    BaseForgeFacet public baseForgeFacet;
    ForgeV1 public forgeV1;
    ForgeV2 public forgeV2;
    ForgeV3 public forgeV3;
    TokenForgeFactory public tokenForgeFactory;
    address public FORGE;

    function setUp() public {
        vm.label(OWNER, "owner");
        vm.startPrank(OWNER);
        // deploy contracts
        // deploy faucets
        diamond3Facet = new Diamond3Facet();
        baseForgeFacet = new BaseForgeFacet();
        forgeV1 = new ForgeV1();
        forgeV2 = new ForgeV2();
        forgeV3 = new ForgeV3();
        // deploy proxy code
        forgeProxyCode = new ForgeProxyCode();
        // deploy factory
        TokenForgeFactory tokenForgeFactoryImpl = new TokenForgeFactory();
        ERC1967Proxy tokenForgeFactoryProxy = new ERC1967Proxy(
            address(tokenForgeFactoryImpl),
            abi.encodeCall(
                TokenForgeFactory.initialize,
                (OWNER, address(forgeProxyCode), address(diamond3Facet), address(baseForgeFacet))
            )
        );
        tokenForgeFactory = TokenForgeFactory(address(tokenForgeFactoryProxy));

        vm.stopPrank();
    }

    function test_mint_erc20_v1() external {
        // set erc20Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV1.mintERC20.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = address(new MockERC20(FORGE));
        vm.label(token, "MockERC20");

        // validator signature
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);

        bytes32 structHash =
            keccak256(abi.encode(ERC20_MINT_TYPE_HASH_V1, uuid, ACCOUNT.addr, token, amount, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).mintERC20(token, amount, deadline, abi.encodePacked(r, s, v));
        assertEq(amount, MockERC20(token).balanceOf(ACCOUNT.addr), "Minted amount mismatch");
    }

    function test_mint_erc721_v1() external {
        // set erc721Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV1.mintERC721.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc721
        address token = address(new MockERC721(FORGE));
        vm.label(token, "MockERC721");

        // validator signature
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);

        bytes32 structHash =
            keccak256(abi.encode(ERC721_MINT_TYPE_HASH_V1, uuid, ACCOUNT.addr, token, tokenID, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).mintERC721(token, tokenID, deadline, abi.encodePacked(r, s, v));
        assertEq(ACCOUNT.addr, MockERC721(token).ownerOf(tokenID), "Minted tokenID mismatch");
    }

    function test_mint_erc1155_v1() external {
        // set erc1155Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV1.mintERC1155.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = address(new MockERC1155(FORGE));
        vm.label(token, "MockERC1155");

        // validator signature
        uint256 tokenID = 1;
        uint256 amount = 10 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);
        bytes memory data = "Minting ERC1155 token";

        bytes32 structHash = keccak256(
            abi.encode(ERC1155_MINT_TYPE_HASH_V1, uuid, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).mintERC1155(token, tokenID, amount, deadline, abi.encodePacked(r, s, v), data);
        assertEq(amount, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenID), "Minted tokenID mismatch");
    }

    function test_mint_batch_erc1155_v1() external {
        // set erc1155Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV1.mintERC1155Batch.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = address(new MockERC1155(FORGE));
        vm.label(token, "MockERC1155");

        // validator signature
        uint256[] memory tokenIDs = new uint256[](2);
        tokenIDs[0] = 1;
        tokenIDs[1] = 2;
        uint256[] memory amounts = new uint256[](2);
        amounts[0] = 10 ether;
        amounts[1] = 20 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);
        bytes memory data = "Minting ERC1155 token";

        bytes32 structHash = keccak256(
            abi.encode(ERC1155_MINT_BATCH_TYPE_HASH_V1, uuid, ACCOUNT.addr, token, tokenIDs, amounts, nonce, deadline)
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).mintERC1155Batch(token, tokenIDs, amounts, deadline, abi.encodePacked(r, s, v), data);

        assertEq(amounts[0], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[0]), "Minted tokenID mismatch");
        assertEq(amounts[1], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[1]), "Minted tokenID mismatch");
    }

    function test_transfer_erc20_v1() external {
        // set erc20Transfer
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV1.transferERC20.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = address(new MockERC20(FORGE));
        vm.label(token, "MockERC20");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        // charge token to forge
        MockERC20(token).forceMint(FORGE, amount);

        bytes32 structHash =
            keccak256(abi.encode(ERC20_TRANSFER_TYPE_HASH_V1, uuid, ACCOUNT.addr, token, amount, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).transferERC20(token, amount, deadline, abi.encodePacked(r, s, v));
        assertEq(amount, MockERC20(token).balanceOf(ACCOUNT.addr), "Transfered amount mismatch");
    }

    function test_transfer_erc721_v1() external {
        // set erc721Transfer
        bytes4[] memory functionSelectors = new bytes4[](2);
        functionSelectors[0] = forgeV1.transferERC721.selector;
        functionSelectors[1] = forgeV1.onERC721Received.selector;

        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc721
        address token = address(new MockERC721(FORGE));
        vm.label(token, "MockERC721");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        bytes memory data = "Transfer ERC721 token";

        // charge token to forge
        MockERC721(token).forceMint(FORGE, tokenID);

        bytes32 structHash =
            keccak256(abi.encode(ERC721_TRANSFER_TYPE_HASH_V1, uuid, ACCOUNT.addr, token, tokenID, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).transferERC721(token, tokenID, deadline, abi.encodePacked(r, s, v), data);
        assertEq(ACCOUNT.addr, MockERC721(token).ownerOf(tokenID), "Transfered tokenID mismatch");
    }

    function test_transfer_erc1155_v1() external {
        // set erc1155Transfer
        bytes4[] memory functionSelectors = new bytes4[](3);
        functionSelectors[0] = forgeV1.transferERC1155.selector;
        functionSelectors[1] = forgeV1.onERC1155Received.selector;
        functionSelectors[2] = forgeV1.onERC1155BatchReceived.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = address(new MockERC1155(FORGE));
        vm.label(token, "MockERC1155");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256 tokenID = 1;
        uint256 amount = 10 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        bytes memory data = "Transfer ERC1155 token";

        // charge token to forge
        MockERC1155(token).forceMint(FORGE, tokenID, amount);

        bytes32 structHash = keccak256(
            abi.encode(ERC1155_TRANSFER_TYPE_HASH_V1, uuid, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).transferERC1155(token, tokenID, amount, deadline, abi.encodePacked(r, s, v), data);
        assertEq(amount, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenID), "Transfered tokenID mismatch");
    }

    function test_transfer_batch_erc1155_v1() external {
        // set erc1155Transfer
        bytes4[] memory functionSelectors = new bytes4[](3);
        functionSelectors[0] = forgeV1.transferERC1155Batch.selector;
        functionSelectors[1] = forgeV1.onERC1155Received.selector;
        functionSelectors[2] = forgeV1.onERC1155BatchReceived.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = address(new MockERC1155(FORGE));
        vm.label(token, "MockERC1155");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256[] memory tokenIDs = new uint256[](2);
        tokenIDs[0] = 1;
        tokenIDs[1] = 2;
        uint256[] memory amounts = new uint256[](2);
        amounts[0] = 10 ether;
        amounts[1] = 20 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        bytes memory data = "Transfer ERC1155 token";

        // charge token to forge
        MockERC1155(token).forceMintBatch(FORGE, tokenIDs, amounts);

        bytes32 structHash = keccak256(
            abi.encode(
                ERC1155_TRANSFER_BATCH_TYPE_HASH_V1, uuid, ACCOUNT.addr, token, tokenIDs, amounts, nonce, deadline
            )
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).transferERC1155Batch(token, tokenIDs, amounts, deadline, abi.encodePacked(r, s, v), data);

        assertEq(amounts[0], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[0]), "Transfered tokenID mismatch");
        assertEq(amounts[1], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[1]), "Transfered tokenID mismatch");
    }

    function test_burn_erc20_v1() external {
        // set erc20Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV1.burnERC20.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = address(new MockERC20(FORGE));
        vm.label(token, "MockERC20");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        // charge token to account
        MockERC20(token).forceMint(ACCOUNT.addr, amount);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC20(token).approve(FORGE, amount);

        bytes32 structHash =
            keccak256(abi.encode(ERC20_BURN_TYPE_HASH_V1, uuid, ACCOUNT.addr, token, amount, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).burnERC20(token, amount, deadline, abi.encodePacked(r, s, v));
        assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "Burned amount mismatch");
    }

    function test_burn_erc721_v1() external {
        // set erc721Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV1.burnERC721.selector;

        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc721
        address token = address(new MockERC721(FORGE));
        vm.label(token, "MockERC721");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        // charge token to account
        MockERC721(token).forceMint(ACCOUNT.addr, tokenID);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC721(token).approve(FORGE, tokenID);

        bytes32 structHash =
            keccak256(abi.encode(ERC721_BURN_TYPE_HASH_V1, uuid, ACCOUNT.addr, token, tokenID, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).burnERC721(token, tokenID, deadline, abi.encodePacked(r, s, v));
        vm.expectRevert(abi.encodeWithSignature("ERC721NonexistentToken(uint256)", tokenID));
        MockERC721(token).ownerOf(tokenID);
    }

    function test_burn_erc1155_v1() external {
        // set erc1155Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV1.burnERC1155.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = address(new MockERC1155(FORGE));
        vm.label(token, "MockERC1155");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256 tokenID = 1;
        uint256 amount = 10 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        // charge token to acount
        MockERC1155(token).forceMint(ACCOUNT.addr, tokenID, amount);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC1155(token).setApprovalForAll(FORGE, true);

        bytes32 structHash = keccak256(
            abi.encode(ERC1155_BURN_TYPE_HASH_V1, uuid, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).burnERC1155(token, tokenID, amount, deadline, abi.encodePacked(r, s, v));
        assertEq(0, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenID), "Burned tokenID mismatch");
    }

    function test_burn_batch_erc1155_v1() external {
        // set erc1155Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV1.burnERC1155Batch.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = address(new MockERC1155(FORGE));
        vm.label(token, "MockERC1155");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256[] memory tokenIDs = new uint256[](2);
        tokenIDs[0] = 1;
        tokenIDs[1] = 2;
        uint256[] memory amounts = new uint256[](2);
        amounts[0] = 10 ether;
        amounts[1] = 20 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        // charge token to account
        MockERC1155(token).forceMintBatch(ACCOUNT.addr, tokenIDs, amounts);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC1155(token).setApprovalForAll(FORGE, true);

        bytes32 structHash = keccak256(
            abi.encode(ERC1155_BURN_BATCH_TYPE_HASH_V1, uuid, ACCOUNT.addr, token, tokenIDs, amounts, nonce, deadline)
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).burnERC1155Batch(token, tokenIDs, amounts, deadline, abi.encodePacked(r, s, v));

        assertEq(0, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[0]), "Burned tokenID mismatch");
        assertEq(0, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[1]), "Burned tokenID mismatch");
    }

    function test_mint_erc20_v2() external {
        // set erc20Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV2.mintERC20.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV2),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = address(new MockERC20(FORGE));
        vm.label(token, "MockERC20");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        bytes memory recipientSig;

        {
            bytes32 recipientStructHash = keccak256(abi.encode(ERC20_MINT_TYPE_HASH_V2, token, amount, nonce, deadline));
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC20_VALIDATOR_MINT_TYPE_HASH_V2, uuid, ACCOUNT.addr, recipientSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE).mintERC20(ACCOUNT.addr, token, amount, deadline, recipientSig, validatorSig);
        assertEq(amount, MockERC20(token).balanceOf(ACCOUNT.addr), "Minted amount mismatch");
    }

    function test_mint_erc721_v2() external {
        // set erc721Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV2.mintERC721.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV2),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc721
        address token = address(new MockERC721(FORGE));
        vm.label(token, "MockERC721");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        bytes memory recipientSig;

        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC721_MINT_TYPE_HASH_V2, token, tokenID, nonce, deadline));
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC721_VALIDATOR_MINT_TYPE_HASH_V2, uuid, ACCOUNT.addr, recipientSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE).mintERC721(ACCOUNT.addr, token, tokenID, deadline, recipientSig, validatorSig);
        assertEq(ACCOUNT.addr, MockERC721(token).ownerOf(tokenID), "Minted tokenID mismatch");
    }

    function test_mint_erc1155_v2() external {
        // set erc1155Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV2.mintERC1155.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV2),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = address(new MockERC1155(FORGE));
        vm.label(token, "MockERC1155");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256 tokenID = 1;
        uint256 amount = 10 ether;
        bytes memory data = "Minting ERC1155 token";
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        bytes memory recipientSig;

        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC1155_MINT_TYPE_HASH_V2, token, tokenID, amount, nonce, deadline));
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_MINT_TYPE_HASH_V2, uuid, ACCOUNT.addr, recipientSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE).mintERC1155(ACCOUNT.addr, token, tokenID, amount, deadline, recipientSig, validatorSig, data);
        assertEq(amount, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenID), "Minted tokenID mismatch");
    }

    function test_mint_batch_erc1155_v2() external {
        // set erc1155Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV2.mintERC1155Batch.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV2),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = address(new MockERC1155(FORGE));
        vm.label(token, "MockERC1155");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256[] memory tokenIDs = new uint256[](2);
        tokenIDs[0] = 1;
        tokenIDs[1] = 2;
        uint256[] memory amounts = new uint256[](2);
        amounts[0] = 10 ether;
        amounts[1] = 20 ether;

        bytes memory data = "Minting ERC1155 token";
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        bytes memory recipientSig;

        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC1155_MINT_BATCH_TYPE_HASH_V2, token, tokenIDs, amounts, nonce, deadline));
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_MINT_BATCH_TYPE_HASH_V2, uuid, ACCOUNT.addr, recipientSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE).mintERC1155Batch(
            ACCOUNT.addr, token, tokenIDs, amounts, deadline, recipientSig, validatorSig, data
        );
        assertEq(amounts[0], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[0]), "Minted tokenID mismatch");
        assertEq(amounts[1], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[1]), "Minted tokenID mismatch");
    }

    function test_transfer_erc20_v2() external {
        // set erc20Transfer
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV2.transferERC20.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV2),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = address(new MockERC20(FORGE));
        vm.label(token, "MockERC20");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        MockERC20(token).forceMint(FORGE, amount);

        bytes memory recipientSig;

        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC20_TRANSFER_TYPE_HASH_V2, token, amount, nonce, deadline));
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC20_VALIDATOR_TRANSFER_TYPE_HASH_V2, uuid, ACCOUNT.addr, recipientSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE).transferERC20(ACCOUNT.addr, token, amount, deadline, recipientSig, validatorSig);
        assertEq(amount, MockERC20(token).balanceOf(ACCOUNT.addr), "Transfered amount mismatch");
    }

    function test_transfer_erc721_v2() external {
        // set erc721Transfer
        bytes4[] memory functionSelectors = new bytes4[](2);
        functionSelectors[0] = forgeV2.transferERC721.selector;
        functionSelectors[1] = forgeV2.onERC721Received.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV2),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc721
        address token = address(new MockERC721(FORGE));
        vm.label(token, "MockERC721");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        bytes memory data = "Transfer ERC721 token";

        // charge token to forge
        MockERC721(token).forceMint(FORGE, tokenID);

        bytes memory recipientSig;

        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC721_TRANSFER_TYPE_HASH_V2, token, tokenID, nonce, deadline));
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC721_VALIDATOR_TRANSFER_TYPE_HASH_V2, uuid, ACCOUNT.addr, recipientSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE).transferERC721(ACCOUNT.addr, token, tokenID, deadline, recipientSig, validatorSig, data);
        assertEq(ACCOUNT.addr, MockERC721(token).ownerOf(tokenID), "Transferred tokenID mismatch");
    }

    function test_transfer_erc1155_v2() external {
        // set erc1155Transfer
        bytes4[] memory functionSelectors = new bytes4[](3);
        functionSelectors[0] = forgeV2.transferERC1155.selector;
        functionSelectors[1] = forgeV2.onERC1155Received.selector;
        functionSelectors[2] = forgeV2.onERC1155BatchReceived.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV2),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = address(new MockERC1155(FORGE));
        vm.label(token, "MockERC1155");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256 tokenID = 1;
        uint256 amount = 10 ether;
        bytes memory data = "Transfer ERC1155 token";
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        // charge token to forge
        MockERC1155(token).forceMint(FORGE, tokenID, amount);

        bytes memory recipientSig;

        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC1155_TRANSFER_TYPE_HASH_V2, token, tokenID, amount, nonce, deadline));
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_TRANSFER_TYPE_HASH_V2, uuid, ACCOUNT.addr, recipientSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE).transferERC1155(ACCOUNT.addr, token, tokenID, amount, deadline, recipientSig, validatorSig, data);
        assertEq(amount, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenID), "Transfered tokenID mismatch");
    }

    function test_transfer_batch_erc1155_v2() external {
        // set erc1155Transfer
        bytes4[] memory functionSelectors = new bytes4[](3);
        functionSelectors[0] = forgeV2.transferERC1155Batch.selector;
        functionSelectors[1] = forgeV2.onERC1155Received.selector;
        functionSelectors[2] = forgeV2.onERC1155BatchReceived.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV2),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = address(new MockERC1155(FORGE));
        vm.label(token, "MockERC1155");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256[] memory tokenIDs = new uint256[](2);
        tokenIDs[0] = 1;
        tokenIDs[1] = 2;
        uint256[] memory amounts = new uint256[](2);
        amounts[0] = 10 ether;
        amounts[1] = 20 ether;

        bytes memory data = "Transfer ERC1155 token";
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        // charge token to forge
        MockERC1155(token).forceMintBatch(FORGE, tokenIDs, amounts);

        bytes memory recipientSig;

        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC1155_TRANSFER_BATCH_TYPE_HASH_V2, token, tokenIDs, amounts, nonce, deadline));
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_TRANSFER_BATCH_TYPE_HASH_V2, uuid, ACCOUNT.addr, recipientSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE).transferERC1155Batch(
            ACCOUNT.addr, token, tokenIDs, amounts, deadline, recipientSig, validatorSig, data
        );
        assertEq(amounts[0], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[0]), "Transfered tokenID mismatch");
        assertEq(amounts[1], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[1]), "Transfered tokenID mismatch");
    }

    function test_burn_erc20_v2() external {
        // set erc20Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV2.burnERC20.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV2),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = address(new MockERC20(FORGE));
        vm.label(token, "MockERC20");

        // validator signature
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);

        // charge token to account
        MockERC20(token).forceMint(ACCOUNT.addr, amount);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC20(token).approve(FORGE, amount);

        bytes memory fromSig;
        {
            bytes32 fromStructHash = keccak256(abi.encode(ERC20_BURN_TYPE_HASH_V2, token, amount, nonce, deadline));
            bytes32 fromHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, fromStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, fromHash);
            fromSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC20_VALIDATOR_BURN_TYPE_HASH_V2, uuid, ACCOUNT.addr, fromSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE).burnERC20(ACCOUNT.addr, token, amount, deadline, fromSig, validatorSig);
        assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "Burned amount mismatch");
    }

    function test_burn_erc721_v2() external {
        // set erc721Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV2.burnERC721.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV2),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc721
        address token = address(new MockERC721(FORGE));
        vm.label(token, "MockERC721");

        // validator signature
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);

        // charge token to account
        MockERC721(token).forceMint(ACCOUNT.addr, tokenID);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC721(token).approve(FORGE, tokenID);

        bytes memory fromSig;

        {
            bytes32 fromStructHash = keccak256(abi.encode(ERC721_BURN_TYPE_HASH_V2, token, tokenID, nonce, deadline));
            bytes32 fromHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, fromStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, fromHash);
            fromSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC721_VALIDATOR_BURN_TYPE_HASH_V2, uuid, ACCOUNT.addr, fromSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE).burnERC721(ACCOUNT.addr, token, tokenID, deadline, fromSig, validatorSig);
        vm.expectRevert(abi.encodeWithSignature("ERC721NonexistentToken(uint256)", tokenID));
        MockERC721(token).ownerOf(tokenID);
    }

    function test_burn_erc1155_v2() external {
        // set erc1155Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV2.burnERC1155.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV2),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = address(new MockERC1155(FORGE));
        vm.label(token, "MockERC1155");

        // validator signature
        uint256 tokenID = 1;
        uint256 amount = 10 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);

        // charge token to account
        MockERC1155(token).forceMint(ACCOUNT.addr, tokenID, amount);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC1155(token).setApprovalForAll(FORGE, true);

        bytes memory fromSig;
        {
            bytes32 fromStructHash =
                keccak256(abi.encode(ERC1155_BURN_TYPE_HASH_V2, token, tokenID, amount, nonce, deadline));
            bytes32 fromHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, fromStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, fromHash);
            fromSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_BURN_TYPE_HASH_V2, uuid, ACCOUNT.addr, fromSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE).burnERC1155(ACCOUNT.addr, token, tokenID, amount, deadline, fromSig, validatorSig);
        assertEq(0, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenID), "Burned tokenID mismatch");
    }

    function test_burn_batch_erc1155_v2() external {
        // set erc1155Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV2.burnERC1155Batch.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV2),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = address(new MockERC1155(FORGE));
        vm.label(token, "MockERC1155");

        // validator signature
        uint256[] memory tokenIDs = new uint256[](2);
        tokenIDs[0] = 1;
        tokenIDs[1] = 2;
        uint256[] memory amounts = new uint256[](2);
        amounts[0] = 10 ether;
        amounts[1] = 20 ether;

        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);
        // charge token to account
        MockERC1155(token).forceMintBatch(ACCOUNT.addr, tokenIDs, amounts);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC1155(token).setApprovalForAll(FORGE, true);

        bytes memory fromSig;

        {
            bytes32 fromStructHash =
                keccak256(abi.encode(ERC1155_BURN_BATCH_TYPE_HASH_V2, token, tokenIDs, amounts, nonce, deadline));
            bytes32 fromHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, fromStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, fromHash);
            fromSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_BURN_BATCH_TYPE_HASH_V2, uuid, ACCOUNT.addr, fromSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE).burnERC1155Batch(ACCOUNT.addr, token, tokenIDs, amounts, deadline, fromSig, validatorSig);
        assertEq(0, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[0]), "Burned tokenID mismatch");
        assertEq(0, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[1]), "Burned tokenID mismatch");
    }

    function test_mint_erc20_v3() external {
        // set erc20Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV3.mintERC20.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV3),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = address(new MockERC20(FORGE));
        vm.label(token, "MockERC20");

        // validator signature

        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);
        bytes32 structHash =
            keccak256(abi.encode(ERC20_MINT_TYPE_HASH_V1, uuid, ACCOUNT.addr, token, amount, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        ForgeV3(FORGE).mintERC20(ACCOUNT.addr, token, amount, deadline, abi.encodePacked(r, s, v));
        assertEq(amount, MockERC20(token).balanceOf(ACCOUNT.addr), "Minted amount mismatch");
    }

    function test_mint_erc721_v3() external {
        // set erc721Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV3.mintERC721.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV3),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc721
        address token = address(new MockERC721(FORGE));
        vm.label(token, "MockERC721");

        // validator signature

        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);
        bytes32 structHash =
            keccak256(abi.encode(ERC721_MINT_TYPE_HASH_V1, uuid, ACCOUNT.addr, token, tokenID, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        ForgeV3(FORGE).mintERC721(ACCOUNT.addr, token, tokenID, deadline, abi.encodePacked(r, s, v));
        assertEq(ACCOUNT.addr, MockERC721(token).ownerOf(tokenID), "Minted tokenID mismatch");
    }

    function test_mint_erc1155_v3() external {
        // set erc1155Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV3.mintERC1155.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV3),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = address(new MockERC1155(FORGE));
        vm.label(token, "MockERC1155");

        // validator signature

        uint256 tokenID = 1;
        uint256 amount = 10 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);
        bytes memory data = "Minting ERC1155 token";

        bytes32 structHash = keccak256(
            abi.encode(ERC1155_MINT_TYPE_HASH_V1, uuid, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        ForgeV3(FORGE).mintERC1155(ACCOUNT.addr, token, tokenID, amount, deadline, abi.encodePacked(r, s, v), data);
        assertEq(amount, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenID), "Minted tokenID mismatch");
    }

    function test_mint_batch_erc1155_v3() external {
        // set erc1155Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV3.mintERC1155Batch.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV3),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = address(new MockERC1155(FORGE));
        vm.label(token, "MockERC1155");

        // validator signature

        uint256[] memory tokenIDs = new uint256[](2);
        tokenIDs[0] = 1;
        tokenIDs[1] = 2;
        uint256[] memory amounts = new uint256[](2);
        amounts[0] = 10 ether;
        amounts[1] = 20 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);
        bytes memory data = "Minting ERC1155 token";

        bytes32 structHash = keccak256(
            abi.encode(ERC1155_MINT_BATCH_TYPE_HASH_V1, uuid, ACCOUNT.addr, token, tokenIDs, amounts, nonce, deadline)
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        ForgeV3(FORGE).mintERC1155Batch(
            ACCOUNT.addr, token, tokenIDs, amounts, deadline, abi.encodePacked(r, s, v), data
        );

        assertEq(amounts[0], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[0]), "Minted tokenID mismatch");
        assertEq(amounts[1], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[1]), "Minted tokenID mismatch");
    }

    function test_transfer_erc20_v3() external {
        // set erc20Transfer
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV3.transferERC20.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV3),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = address(new MockERC20(FORGE));
        vm.label(token, "MockERC20");

        // validator signature

        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);
        // charge token to forge
        MockERC20(token).forceMint(FORGE, amount);

        bytes32 structHash =
            keccak256(abi.encode(ERC20_TRANSFER_TYPE_HASH_V1, uuid, ACCOUNT.addr, token, amount, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        ForgeV3(FORGE).transferERC20(ACCOUNT.addr, token, amount, deadline, abi.encodePacked(r, s, v));
        assertEq(amount, MockERC20(token).balanceOf(ACCOUNT.addr), "Transfered amount mismatch");
    }

    function test_transfer_erc721_v3() external {
        // set erc721Transfer
        bytes4[] memory functionSelectors = new bytes4[](2);
        functionSelectors[0] = forgeV3.transferERC721.selector;
        functionSelectors[1] = forgeV3.onERC721Received.selector;

        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV3),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc721
        address token = address(new MockERC721(FORGE));
        vm.label(token, "MockERC721");

        // validator signature
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);
        bytes memory data = "Transfer ERC721 token";

        // charge token to forge
        MockERC721(token).forceMint(FORGE, tokenID);

        bytes32 structHash =
            keccak256(abi.encode(ERC721_TRANSFER_TYPE_HASH_V1, uuid, ACCOUNT.addr, token, tokenID, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        ForgeV3(FORGE).transferERC721(ACCOUNT.addr, token, tokenID, deadline, abi.encodePacked(r, s, v), data);
        assertEq(ACCOUNT.addr, MockERC721(token).ownerOf(tokenID), "Transfered tokenID mismatch");
    }

    function test_transfer_erc1155_v3() external {
        // set erc1155Transfer
        bytes4[] memory functionSelectors = new bytes4[](3);
        functionSelectors[0] = forgeV3.transferERC1155.selector;
        functionSelectors[1] = forgeV3.onERC1155Received.selector;
        functionSelectors[2] = forgeV3.onERC1155BatchReceived.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV3),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = address(new MockERC1155(FORGE));
        vm.label(token, "MockERC1155");

        // validator signature
        uint256 tokenID = 1;
        uint256 amount = 10 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);
        bytes memory data = "Transfer ERC1155 token";

        // charge token to forge
        MockERC1155(token).forceMint(FORGE, tokenID, amount);

        bytes32 structHash = keccak256(
            abi.encode(ERC1155_TRANSFER_TYPE_HASH_V1, uuid, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        ForgeV3(FORGE).transferERC1155(ACCOUNT.addr, token, tokenID, amount, deadline, abi.encodePacked(r, s, v), data);
        assertEq(amount, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenID), "Transfered tokenID mismatch");
    }

    function test_transfer_batch_erc1155_v3() external {
        // set erc1155Transfer
        bytes4[] memory functionSelectors = new bytes4[](3);
        functionSelectors[0] = forgeV3.transferERC1155Batch.selector;
        functionSelectors[1] = forgeV3.onERC1155Received.selector;
        functionSelectors[2] = forgeV3.onERC1155BatchReceived.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV3),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = address(new MockERC1155(FORGE));
        vm.label(token, "MockERC1155");

        // validator signature
        uint256[] memory tokenIDs = new uint256[](2);
        tokenIDs[0] = 1;
        tokenIDs[1] = 2;
        uint256[] memory amounts = new uint256[](2);
        amounts[0] = 10 ether;
        amounts[1] = 20 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);
        bytes memory data = "Transfer ERC1155 token";

        // charge token to forge
        MockERC1155(token).forceMintBatch(FORGE, tokenIDs, amounts);

        bytes32 structHash = keccak256(
            abi.encode(
                ERC1155_TRANSFER_BATCH_TYPE_HASH_V1, uuid, ACCOUNT.addr, token, tokenIDs, amounts, nonce, deadline
            )
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        ForgeV3(FORGE).transferERC1155Batch(
            ACCOUNT.addr, token, tokenIDs, amounts, deadline, abi.encodePacked(r, s, v), data
        );

        assertEq(amounts[0], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[0]), "Transfered tokenID mismatch");
        assertEq(amounts[1], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[1]), "Transfered tokenID mismatch");
    }

    function test_burn_erc20_v3() external {
        // set erc20Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV3.burnERC20.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV3),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = address(new MockERC20(FORGE));
        vm.label(token, "MockERC20");

        // validator signature
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);

        // charge token to account
        MockERC20(token).forceMint(ACCOUNT.addr, amount);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC20(token).approve(FORGE, amount);

        bytes32 structHash =
            keccak256(abi.encode(ERC20_BURN_TYPE_HASH_V1, uuid, ACCOUNT.addr, token, amount, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        ForgeV3(FORGE).burnERC20(ACCOUNT.addr, token, amount, deadline, abi.encodePacked(r, s, v));
        assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "Burned amount mismatch");
    }

    function test_burn_erc721_v3() external {
        // set erc721Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV3.burnERC721.selector;

        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV3),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc721
        address token = address(new MockERC721(FORGE));
        vm.label(token, "MockERC721");

        // validator signature
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);

        // charge token to account
        MockERC721(token).forceMint(ACCOUNT.addr, tokenID);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC721(token).approve(FORGE, tokenID);

        bytes32 structHash =
            keccak256(abi.encode(ERC721_BURN_TYPE_HASH_V1, uuid, ACCOUNT.addr, token, tokenID, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        ForgeV3(FORGE).burnERC721(ACCOUNT.addr, token, tokenID, deadline, abi.encodePacked(r, s, v));
        vm.expectRevert(abi.encodeWithSignature("ERC721NonexistentToken(uint256)", tokenID));
        MockERC721(token).ownerOf(tokenID);
    }

    function test_burn_erc1155_v3() external {
        // set erc1155Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV3.burnERC1155.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV3),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = address(new MockERC1155(FORGE));
        vm.label(token, "MockERC1155");

        // validator signature
        uint256 tokenID = 1;
        uint256 amount = 10 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);

        // charge token to acount
        MockERC1155(token).forceMint(ACCOUNT.addr, tokenID, amount);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC1155(token).setApprovalForAll(FORGE, true);

        bytes32 structHash = keccak256(
            abi.encode(ERC1155_BURN_TYPE_HASH_V1, uuid, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        ForgeV3(FORGE).burnERC1155(ACCOUNT.addr, token, tokenID, amount, deadline, abi.encodePacked(r, s, v));
        assertEq(0, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenID), "Burned tokenID mismatch");
    }

    function test_burn_batch_erc1155_v3() external {
        // set erc1155Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV3.burnERC1155Batch.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV3),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = address(new MockERC1155(FORGE));
        vm.label(token, "MockERC1155");

        // validator signature
        uint256[] memory tokenIDs = new uint256[](2);
        tokenIDs[0] = 1;
        tokenIDs[1] = 2;
        uint256[] memory amounts = new uint256[](2);
        amounts[0] = 10 ether;
        amounts[1] = 20 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);

        // charge token to account
        MockERC1155(token).forceMintBatch(ACCOUNT.addr, tokenIDs, amounts);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC1155(token).setApprovalForAll(FORGE, true);

        bytes32 structHash = keccak256(
            abi.encode(ERC1155_BURN_BATCH_TYPE_HASH_V3, uuid, ACCOUNT.addr, token, tokenIDs, amounts, nonce, deadline)
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        ForgeV3(FORGE).burnERC1155Batch(ACCOUNT.addr, token, tokenIDs, amounts, deadline, abi.encodePacked(r, s, v));

        assertEq(0, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[0]), "Burned tokenID mismatch");
        assertEq(0, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[1]), "Burned tokenID mismatch");
    }

    function _calcUUID(uint256 nonce) internal view returns (uint256) {
        return uint256(keccak256(abi.encode(FORGE, ACCOUNT.addr, nonce)));
    }
}
