// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std-1.9.7/Test.sol";
import {Vm} from "forge-std-1.9.7/Vm.sol";

import {IDiamondCut} from "diamond-3-hardhat-1.0.0/interfaces/IDiamondCut.sol";
import {IDiamondLoupe} from "diamond-3-hardhat-1.0.0/interfaces/IDiamondLoupe.sol";
import {IERC173} from "diamond-3-hardhat-1.0.0/interfaces/IERC173.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts-5.3.0/proxy/ERC1967/ERC1967Proxy.sol";
import {MessageHashUtils} from "@openzeppelin-contracts-5.3.0/utils/cryptography/MessageHashUtils.sol";

import {ForgeProxyCode} from "../../src/ForgeProxy.sol";
import {Diamond3Facet} from "../../src/Diamond3Facet.sol";
import {ForgeFactory} from "../../src/v1/ForgeFactory.sol";
import "../../src/v1/BaseForge.sol";
import "../../src/v1/ForgeV1.sol";
import "../../src/v1/ForgeV2.sol";
import "../../src/v1/ForgeV3.sol";

import {IERC20} from "@openzeppelin-contracts-5.3.0/token/ERC20/IERC20.sol";
import {IERC20Permit} from "@openzeppelin-contracts-5.3.0/token/ERC20/extensions/IERC20Permit.sol";
import {IERC721} from "@openzeppelin-contracts-5.3.0/token/ERC721/IERC721.sol";
import {IERC1155} from "@openzeppelin-contracts-5.3.0/token/ERC1155/IERC1155.sol";

import {MockERC20} from "../mock/MockERC20.sol";
import {MockERC721} from "../mock/MockERC721.sol";
import {MockERC1155} from "../mock/MockERC1155.sol";

import "../mock/StructHash.sol";

contract TestTokenForgeFactory is Test {
    address public constant OWNER = address(bytes20("OWNER"));
    address public constant SERVICE_OWNER = address(bytes20("SERVICE_OWNER"));

    bytes32 public constant SERVICE_NAME = bytes32("TestService");
    Vm.Wallet public VALIDATOR = vm.createWallet("Validator");
    Vm.Wallet public ACCOUNT = vm.createWallet("Account");

    ForgeProxyCode public forgeProxyCode;
    Diamond3Facet public diamond3Facet;
    BaseForgeFacet public baseForgeFacet;
    ForgeV1 public forgeV1;
    ForgeV2 public forgeV2;
    ForgeV3 public forgeV3;
    ForgeFactory public forgeFactory;
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
        ForgeFactory forgeFactoryImpl = new ForgeFactory();
        ERC1967Proxy forgeFactoryProxy = new ERC1967Proxy(
            address(forgeFactoryImpl),
            abi.encodeCall(
                ForgeFactory.initialize,
                (OWNER, address(forgeProxyCode), address(diamond3Facet), address(baseForgeFacet))
            )
        );
        forgeFactory = ForgeFactory(address(forgeFactoryProxy));

        vm.stopPrank();
    }

    function _deployERC20() internal returns (address) {
        return address(new MockERC20(FORGE));
    }

    function _deployERC721() internal returns (address) {
        return address(new MockERC721(FORGE));
    }

    function _deployERC1155() internal returns (address) {
        return address(new MockERC1155(FORGE));
    }

    function test_mint_erc20_v1() external {
        // set erc20Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV1.mintERC20.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = _deployERC20();
        vm.label(token, "MockERC20");

        // validator signature
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);

        bytes32 structHash =
            keccak256(abi.encode(ERC20_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, amount, address(0), 0, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // expect emit
        vm.expectEmit();
        emit IERC20.Transfer(address(0), ACCOUNT.addr, amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC20Minted(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).mintERC20(token, amount, address(0), 0, deadline, abi.encodePacked(r, s, v));
        assertEq(amount, MockERC20(token).balanceOf(ACCOUNT.addr), "Minted amount mismatch");
    }

    function test_mint_erc721_v1() external {
        // set erc721Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV1.mintERC721.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc721
        address token = _deployERC721();
        vm.label(token, "MockERC721");

        // validator signature
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);
        bytes memory data = "test_mint_erc721_v1";

        bytes32 structHash = keccak256(
            abi.encode(ERC721_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, keccak256(data), nonce, deadline)
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // expect emit
        vm.expectEmit();
        emit IERC721.Transfer(address(0), ACCOUNT.addr, tokenID);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC721Minted(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenID);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).mintERC721(token, tokenID, data, deadline, abi.encodePacked(r, s, v));
        assertEq(ACCOUNT.addr, MockERC721(token).ownerOf(tokenID), "Minted tokenID mismatch");
    }

    function test_mint_erc1155_v1() external {
        // set erc1155Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV1.mintERC1155.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = _deployERC1155();
        vm.label(token, "MockERC1155");

        // validator signature
        uint256 tokenID = 1;
        uint256 amount = 10 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);
        bytes memory data = "Minting ERC1155 token";

        bytes32 structHash = keccak256(
            abi.encode(
                ERC1155_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, amount, keccak256(data), nonce, deadline
            )
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // expect emit
        vm.expectEmit();
        emit IERC1155.TransferSingle(address(FORGE), address(0), ACCOUNT.addr, tokenID, amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Minted(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenID, amount);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).mintERC1155(token, tokenID, amount, data, deadline, abi.encodePacked(r, s, v));
        assertEq(amount, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenID), "Minted tokenID mismatch");
    }

    function test_mint_batch_erc1155_v1() external {
        // set erc1155Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV1.mintERC1155Batch.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = _deployERC1155();
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
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // expect emit
        vm.expectEmit();
        emit IERC1155.TransferBatch(address(FORGE), address(0), ACCOUNT.addr, tokenIDs, amounts);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Minted(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenIDs[0], amounts[0]);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Minted(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenIDs[1], amounts[1]);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).mintERC1155Batch(token, tokenIDs, amounts, data, deadline, abi.encodePacked(r, s, v));
        assertEq(amounts[0], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[0]), "Minted tokenID mismatch");
        assertEq(amounts[1], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[1]), "Minted tokenID mismatch");
    }

    function test_transfer_erc20_v1() external {
        // set erc20Transfer
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV1.transferERC20.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = _deployERC20();
        vm.label(token, "MockERC20");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        // charge token to forge
        vm.prank(OWNER);
        MockERC20(token).forceMint(FORGE, amount);

        bytes32 structHash = keccak256(
            abi.encode(ERC20_TRANSFER_TYPE_HASH_V1, ACCOUNT.addr, token, amount, address(0), 0, nonce, deadline)
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // expect emit
        vm.expectEmit();
        emit IERC20.Transfer(address(FORGE), ACCOUNT.addr, amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC20Transferred(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).transferERC20(token, amount, address(0), 0, deadline, abi.encodePacked(r, s, v));
        assertEq(amount, MockERC20(token).balanceOf(ACCOUNT.addr), "Transfered amount mismatch");
    }

    function test_transfer_erc721_v1() external {
        // set erc721Transfer
        bytes4[] memory functionSelectors = new bytes4[](2);
        functionSelectors[0] = forgeV1.transferERC721.selector;
        functionSelectors[1] = forgeV1.onERC721Received.selector;

        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc721
        address token = _deployERC721();
        vm.label(token, "MockERC721");

        // validator signature
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);
        bytes memory data = "Transfer ERC721 token";

        // charge token to forge
        vm.prank(OWNER);
        MockERC721(token).forceMint(FORGE, tokenID, "");

        bytes32 structHash = keccak256(
            abi.encode(ERC721_TRANSFER_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, keccak256(data), nonce, deadline)
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // expect emit
        vm.expectEmit();
        emit IERC721.Transfer(address(FORGE), ACCOUNT.addr, tokenID);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC721Transferred(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenID);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).transferERC721(token, tokenID, data, deadline, abi.encodePacked(r, s, v));
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
            facetAddress: address(forgeV1), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = _deployERC1155();
        vm.label(token, "MockERC1155");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256 tokenID = 1;
        uint256 amount = 10 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        bytes memory data = "Transfer ERC1155 token";

        // charge token to forge
        vm.prank(OWNER);
        MockERC1155(token).forceMint(FORGE, tokenID, amount);

        bytes32 structHash = keccak256(
            abi.encode(
                ERC1155_TRANSFER_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, amount, keccak256(data), nonce, deadline
            )
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // expect emit
        vm.expectEmit();
        emit IERC1155.TransferSingle(address(FORGE), address(FORGE), ACCOUNT.addr, tokenID, amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Transferred(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenID, amount);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).transferERC1155(token, tokenID, amount, data, deadline, abi.encodePacked(r, s, v));
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
            facetAddress: address(forgeV1), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = _deployERC1155();
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
        vm.prank(OWNER);
        MockERC1155(token).forceMintBatch(FORGE, tokenIDs, amounts);

        bytes32 structHash = keccak256(
            abi.encode(
                ERC1155_TRANSFER_BATCH_TYPE_HASH_V1,
                ACCOUNT.addr,
                token,
                keccak256(abi.encodePacked(tokenIDs)),
                keccak256(abi.encodePacked(amounts)),
                keccak256(data),
                nonce,
                deadline
            )
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // expect emit
        vm.expectEmit();
        emit IERC1155.TransferBatch(address(FORGE), address(FORGE), ACCOUNT.addr, tokenIDs, amounts);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Transferred(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenIDs[0], amounts[0]);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Transferred(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenIDs[1], amounts[1]);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).transferERC1155Batch(token, tokenIDs, amounts, data, deadline, abi.encodePacked(r, s, v));

        assertEq(amounts[0], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[0]), "Transfered tokenID mismatch");
        assertEq(amounts[1], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[1]), "Transfered tokenID mismatch");
    }

    function test_burn_erc20_v1() external {
        // set erc20Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV1.burnERC20.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = _deployERC20();
        vm.label(token, "MockERC20");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        // charge token to account
        vm.prank(OWNER);
        MockERC20(token).forceMint(ACCOUNT.addr, amount);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC20(token).approve(FORGE, amount);

        bytes32 structHash =
            keccak256(abi.encode(ERC20_BURN_TYPE_HASH_V1, ACCOUNT.addr, token, amount, address(0), 0, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // expect emit
        vm.expectEmit();
        emit IERC20.Transfer(ACCOUNT.addr, address(0), amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC20Burned(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).burnERC20(token, amount, address(0), 0, deadline, abi.encodePacked(r, s, v));
        assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "Burned amount mismatch");
    }

    function test_burn_erc721_v1() external {
        // set erc721Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV1.burnERC721.selector;

        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc721
        address token = _deployERC721();
        vm.label(token, "MockERC721");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        // charge token to account
        vm.prank(OWNER);
        MockERC721(token).forceMint(ACCOUNT.addr, tokenID, "");
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC721(token).approve(FORGE, tokenID);

        bytes32 structHash =
            keccak256(abi.encode(ERC721_BURN_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // expect emit
        vm.expectEmit();
        emit IERC721.Transfer(ACCOUNT.addr, address(0), tokenID);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC721Burned(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenID);

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
            facetAddress: address(forgeV1), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = _deployERC1155();
        vm.label(token, "MockERC1155");

        // validator signature
        uint256 tokenID = 1;
        uint256 amount = 10 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);

        // charge token to account
        vm.prank(OWNER);
        MockERC1155(token).forceMint(ACCOUNT.addr, tokenID, amount);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC1155(token).setApprovalForAll(FORGE, true);

        bytes32 structHash =
            keccak256(abi.encode(ERC1155_BURN_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, amount, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // expect emit
        vm.expectEmit();
        emit IERC1155.TransferSingle(address(FORGE), ACCOUNT.addr, address(0), tokenID, amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Burned(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenID, amount);

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
            facetAddress: address(forgeV1), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = _deployERC1155();
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
        vm.prank(OWNER);
        MockERC1155(token).forceMintBatch(ACCOUNT.addr, tokenIDs, amounts);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC1155(token).setApprovalForAll(FORGE, true);

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
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // expect emit
        vm.expectEmit();
        emit IERC1155.TransferBatch(address(FORGE), ACCOUNT.addr, address(0), tokenIDs, amounts);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Burned(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenIDs[0], amounts[0]);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Burned(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenIDs[1], amounts[1]);

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
            facetAddress: address(forgeV2), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = _deployERC20();
        vm.label(token, "MockERC20");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        bytes memory recipientSig;
        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC20_MINT_TYPE_HASH_V2, token, amount, address(0), 0, nonce, deadline));
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC20_VALIDATOR_MINT_TYPE_HASH_V2, ACCOUNT.addr, keccak256(recipientSig)));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // expect emit
        vm.expectEmit();
        emit IERC20.Transfer(address(0), ACCOUNT.addr, amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC20Minted(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE).mintERC20(ACCOUNT.addr, token, amount, address(0), 0, deadline, recipientSig, validatorSig);
        assertEq(amount, MockERC20(token).balanceOf(ACCOUNT.addr), "Minted amount mismatch");
    }

    function test_mint_erc721_v2() external {
        // set erc721Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV2.mintERC721.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV2), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc721
        address token = _deployERC721();
        vm.label(token, "MockERC721");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        bytes memory data = "test_mint_erc721_v2";

        bytes memory recipientSig;

        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC721_MINT_TYPE_HASH_V2, token, tokenID, keccak256(data), nonce, deadline));
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC721_VALIDATOR_MINT_TYPE_HASH_V2, ACCOUNT.addr, keccak256(recipientSig)));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // expect emit
        vm.expectEmit();
        emit IERC721.Transfer(address(0), ACCOUNT.addr, tokenID);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC721Minted(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenID);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE).mintERC721(ACCOUNT.addr, token, tokenID, data, deadline, recipientSig, validatorSig);
        assertEq(ACCOUNT.addr, MockERC721(token).ownerOf(tokenID), "Minted tokenID mismatch");
    }

    function test_mint_erc1155_v2() external {
        // set erc1155Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV2.mintERC1155.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV2), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = _deployERC1155();
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
            bytes32 recipientStructHash = keccak256(
                abi.encode(ERC1155_MINT_TYPE_HASH_V2, token, tokenID, amount, keccak256(data), nonce, deadline)
            );
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_MINT_TYPE_HASH_V2, ACCOUNT.addr, keccak256(recipientSig)));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // expect emit
        vm.expectEmit();
        emit IERC1155.TransferSingle(address(FORGE), address(0), ACCOUNT.addr, tokenID, amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Minted(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenID, amount);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE).mintERC1155(ACCOUNT.addr, token, tokenID, amount, data, deadline, recipientSig, validatorSig);
        assertEq(amount, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenID), "Minted tokenID mismatch");
    }

    function test_mint_batch_erc1155_v2() external {
        // set erc1155Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV2.mintERC1155Batch.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV2), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = _deployERC1155();
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
            bytes32 recipientStructHash = keccak256(
                abi.encode(
                    ERC1155_MINT_BATCH_TYPE_HASH_V2,
                    token,
                    keccak256(abi.encodePacked(tokenIDs)),
                    keccak256(abi.encodePacked(amounts)),
                    keccak256(data),
                    nonce,
                    deadline
                )
            );
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_MINT_BATCH_TYPE_HASH_V2, ACCOUNT.addr, keccak256(recipientSig)));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // expect emit
        vm.expectEmit();
        emit IERC1155.TransferBatch(address(FORGE), address(0), ACCOUNT.addr, tokenIDs, amounts);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Minted(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenIDs[0], amounts[0]);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Minted(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenIDs[1], amounts[1]);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE)
            .mintERC1155Batch(ACCOUNT.addr, token, tokenIDs, amounts, data, deadline, recipientSig, validatorSig);
        assertEq(amounts[0], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[0]), "Minted tokenID mismatch");
        assertEq(amounts[1], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[1]), "Minted tokenID mismatch");
    }

    function test_transfer_erc20_v2() external {
        // set erc20Transfer
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV2.transferERC20.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV2), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = _deployERC20();
        vm.label(token, "MockERC20");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        vm.prank(OWNER);
        MockERC20(token).forceMint(FORGE, amount);

        bytes memory recipientSig;
        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC20_TRANSFER_TYPE_HASH_V2, token, amount, address(0), 0, nonce, deadline));
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC20_VALIDATOR_TRANSFER_TYPE_HASH_V2, ACCOUNT.addr, keccak256(recipientSig)));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // expect emit
        vm.expectEmit();
        emit IERC20.Transfer(address(FORGE), ACCOUNT.addr, amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC20Transferred(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE).transferERC20(ACCOUNT.addr, token, amount, address(0), 0, deadline, recipientSig, validatorSig);
        assertEq(amount, MockERC20(token).balanceOf(ACCOUNT.addr), "Transfered amount mismatch");
    }

    function test_transfer_erc721_v2() external {
        // set erc721Transfer
        bytes4[] memory functionSelectors = new bytes4[](2);
        functionSelectors[0] = forgeV2.transferERC721.selector;
        functionSelectors[1] = forgeV2.onERC721Received.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV2), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc721
        address token = _deployERC721();
        vm.label(token, "MockERC721");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        bytes memory data = "Transfer ERC721 token";

        // charge token to forge
        vm.prank(OWNER);
        MockERC721(token).forceMint(FORGE, tokenID, "");

        bytes memory recipientSig;

        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC721_TRANSFER_TYPE_HASH_V2, token, tokenID, keccak256(data), nonce, deadline));
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC721_VALIDATOR_TRANSFER_TYPE_HASH_V2, ACCOUNT.addr, keccak256(recipientSig)));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // expect emit
        vm.expectEmit();
        emit IERC721.Transfer(address(FORGE), ACCOUNT.addr, tokenID);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC721Transferred(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenID);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE).transferERC721(ACCOUNT.addr, token, tokenID, data, deadline, recipientSig, validatorSig);
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
            facetAddress: address(forgeV2), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = _deployERC1155();
        vm.label(token, "MockERC1155");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256 tokenID = 1;
        uint256 amount = 10 ether;
        bytes memory data = "Transfer ERC1155 token";
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        // charge token to forge
        vm.prank(OWNER);
        MockERC1155(token).forceMint(FORGE, tokenID, amount);

        bytes memory recipientSig;

        {
            bytes32 recipientStructHash = keccak256(
                abi.encode(ERC1155_TRANSFER_TYPE_HASH_V2, token, tokenID, amount, keccak256(data), nonce, deadline)
            );
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_TRANSFER_TYPE_HASH_V2, ACCOUNT.addr, keccak256(recipientSig)));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // expect emit
        vm.expectEmit();
        emit IERC1155.TransferSingle(address(FORGE), address(FORGE), ACCOUNT.addr, tokenID, amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Transferred(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenID, amount);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE).transferERC1155(ACCOUNT.addr, token, tokenID, amount, data, deadline, recipientSig, validatorSig);
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
            facetAddress: address(forgeV2), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = _deployERC1155();
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
        vm.prank(OWNER);
        MockERC1155(token).forceMintBatch(FORGE, tokenIDs, amounts);

        bytes memory recipientSig;

        {
            bytes32 recipientStructHash = keccak256(
                abi.encode(
                    ERC1155_TRANSFER_BATCH_TYPE_HASH_V2,
                    token,
                    keccak256(abi.encodePacked(tokenIDs)),
                    keccak256(abi.encodePacked(amounts)),
                    keccak256(data),
                    nonce,
                    deadline
                )
            );
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash = keccak256(
                abi.encode(ERC1155_VALIDATOR_TRANSFER_BATCH_TYPE_HASH_V2, ACCOUNT.addr, keccak256(recipientSig))
            );
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // expect emit
        vm.expectEmit();
        emit IERC1155.TransferBatch(address(FORGE), address(FORGE), ACCOUNT.addr, tokenIDs, amounts);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Transferred(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenIDs[0], amounts[0]);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Transferred(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenIDs[1], amounts[1]);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE)
            .transferERC1155Batch(ACCOUNT.addr, token, tokenIDs, amounts, data, deadline, recipientSig, validatorSig);
        assertEq(amounts[0], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[0]), "Transfered tokenID mismatch");
        assertEq(amounts[1], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[1]), "Transfered tokenID mismatch");
    }

    function test_burn_erc20_v2() external {
        // set erc20Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV2.burnERC20.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV2), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = _deployERC20();
        vm.label(token, "MockERC20");

        // validator signature
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);

        // charge token to account
        vm.prank(OWNER);
        MockERC20(token).forceMint(ACCOUNT.addr, amount);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC20(token).approve(FORGE, amount);

        bytes memory fromSig;
        {
            bytes32 fromStructHash =
                keccak256(abi.encode(ERC20_BURN_TYPE_HASH_V2, token, amount, address(0), 0, nonce, deadline));
            bytes32 fromHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, fromStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, fromHash);
            fromSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC20_VALIDATOR_BURN_TYPE_HASH_V2, ACCOUNT.addr, keccak256(fromSig)));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // expect emit
        vm.expectEmit();
        emit IERC20.Transfer(ACCOUNT.addr, address(0), amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC20Burned(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE).burnERC20(ACCOUNT.addr, token, amount, address(0), 0, deadline, fromSig, validatorSig);
        assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "Burned amount mismatch");
    }

    function test_burn_erc721_v2() external {
        // set erc721Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV2.burnERC721.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV2), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc721
        address token = _deployERC721();
        vm.label(token, "MockERC721");

        // validator signature
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);

        // charge token to account
        vm.prank(OWNER);
        MockERC721(token).forceMint(ACCOUNT.addr, tokenID, "");
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
                keccak256(abi.encode(ERC721_VALIDATOR_BURN_TYPE_HASH_V2, ACCOUNT.addr, keccak256(fromSig)));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // expect emit
        vm.expectEmit();
        emit IERC721.Transfer(ACCOUNT.addr, address(0), tokenID);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC721Burned(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenID);

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
            facetAddress: address(forgeV2), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = _deployERC1155();
        vm.label(token, "MockERC1155");

        // validator signature
        uint256 tokenID = 1;
        uint256 amount = 10 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);

        // charge token to account
        vm.prank(OWNER);
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
                keccak256(abi.encode(ERC1155_VALIDATOR_BURN_TYPE_HASH_V2, ACCOUNT.addr, keccak256(fromSig)));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // expect emit
        vm.expectEmit();
        emit IERC1155.TransferSingle(address(FORGE), ACCOUNT.addr, address(0), tokenID, amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Burned(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenID, amount);

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
            facetAddress: address(forgeV2), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = _deployERC1155();
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
        vm.prank(OWNER);
        MockERC1155(token).forceMintBatch(ACCOUNT.addr, tokenIDs, amounts);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC1155(token).setApprovalForAll(FORGE, true);

        bytes memory fromSig;

        {
            bytes32 fromStructHash = keccak256(
                abi.encode(
                    ERC1155_BURN_BATCH_TYPE_HASH_V2,
                    token,
                    keccak256(abi.encodePacked(tokenIDs)),
                    keccak256(abi.encodePacked(amounts)),
                    nonce,
                    deadline
                )
            );
            bytes32 fromHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, fromStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, fromHash);
            fromSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_BURN_BATCH_TYPE_HASH_V2, ACCOUNT.addr, keccak256(fromSig)));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // expect emit
        vm.expectEmit();
        emit IERC1155.TransferBatch(address(FORGE), ACCOUNT.addr, address(0), tokenIDs, amounts);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Burned(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenIDs[0], amounts[0]);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Burned(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenIDs[1], amounts[1]);

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
            facetAddress: address(forgeV3), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = _deployERC20();
        vm.label(token, "MockERC20");

        // validator signature

        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);
        bytes32 structHash =
            keccak256(abi.encode(ERC20_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, amount, address(0), 0, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // expect emit
        vm.expectEmit();
        emit IERC20.Transfer(address(0), ACCOUNT.addr, amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC20Minted(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);

        // send transaction
        ForgeV3(FORGE).mintERC20(ACCOUNT.addr, token, amount, address(0), 0, deadline, abi.encodePacked(r, s, v));
        assertEq(amount, MockERC20(token).balanceOf(ACCOUNT.addr), "Minted amount mismatch");
    }

    function test_mint_erc721_v3() external {
        // set erc721Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV3.mintERC721.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV3), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc721
        address token = _deployERC721();
        vm.label(token, "MockERC721");

        // validator signature

        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);
        bytes memory data = "test_mint_erc721_v3";
        bytes32 structHash = keccak256(
            abi.encode(ERC721_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, keccak256(data), nonce, deadline)
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // expect emit
        vm.expectEmit();
        emit IERC721.Transfer(address(0), ACCOUNT.addr, tokenID);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC721Minted(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenID);

        // send transaction
        ForgeV3(FORGE).mintERC721(ACCOUNT.addr, token, tokenID, data, deadline, abi.encodePacked(r, s, v));
        assertEq(ACCOUNT.addr, MockERC721(token).ownerOf(tokenID), "Minted tokenID mismatch");
    }

    function test_mint_erc1155_v3() external {
        // set erc1155Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV3.mintERC1155.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV3), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = _deployERC1155();
        vm.label(token, "MockERC1155");

        // validator signature

        uint256 tokenID = 1;
        uint256 amount = 10 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);
        bytes memory data = "Minting ERC1155 token";

        bytes32 structHash = keccak256(
            abi.encode(
                ERC1155_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, amount, keccak256(data), nonce, deadline
            )
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // expect emit
        vm.expectEmit();
        emit IERC1155.TransferSingle(address(FORGE), address(0), ACCOUNT.addr, tokenID, amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Minted(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenID, amount);

        // send transaction
        ForgeV3(FORGE).mintERC1155(ACCOUNT.addr, token, tokenID, amount, data, deadline, abi.encodePacked(r, s, v));
        assertEq(amount, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenID), "Minted tokenID mismatch");
    }

    function test_mint_batch_erc1155_v3() external {
        // set erc1155Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV3.mintERC1155Batch.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV3), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = _deployERC1155();
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
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // expect emit
        vm.expectEmit();
        emit IERC1155.TransferBatch(address(FORGE), address(0), ACCOUNT.addr, tokenIDs, amounts);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Minted(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenIDs[0], amounts[0]);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Minted(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenIDs[1], amounts[1]);

        // send transaction
        ForgeV3(FORGE)
            .mintERC1155Batch(ACCOUNT.addr, token, tokenIDs, amounts, data, deadline, abi.encodePacked(r, s, v));
        assertEq(amounts[0], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[0]), "Minted tokenID mismatch");
        assertEq(amounts[1], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[1]), "Minted tokenID mismatch");
    }

    function test_transfer_erc20_v3() external {
        // set erc20Transfer
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV3.transferERC20.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV3), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = _deployERC20();
        vm.label(token, "MockERC20");

        // validator signature
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);
        // charge token to forge
        vm.prank(OWNER);
        MockERC20(token).forceMint(FORGE, amount);

        bytes32 structHash = keccak256(
            abi.encode(ERC20_TRANSFER_TYPE_HASH_V1, ACCOUNT.addr, token, amount, address(0), 0, nonce, deadline)
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // expect emit
        vm.expectEmit();
        emit IERC20.Transfer(address(FORGE), ACCOUNT.addr, amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC20Transferred(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);

        // send transaction
        ForgeV3(FORGE).transferERC20(ACCOUNT.addr, token, amount, address(0), 0, deadline, abi.encodePacked(r, s, v));
        assertEq(amount, MockERC20(token).balanceOf(ACCOUNT.addr), "Transfered amount mismatch");
    }

    function test_transfer_erc721_v3() external {
        // set erc721Transfer
        bytes4[] memory functionSelectors = new bytes4[](2);
        functionSelectors[0] = forgeV3.transferERC721.selector;
        functionSelectors[1] = forgeV3.onERC721Received.selector;

        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV3), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc721
        address token = _deployERC721();
        vm.label(token, "MockERC721");

        // validator signature
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);
        bytes memory data = "Transfer ERC721 token";

        // charge token to forge
        vm.prank(OWNER);
        MockERC721(token).forceMint(FORGE, tokenID, "");

        bytes32 structHash = keccak256(
            abi.encode(ERC721_TRANSFER_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, keccak256(data), nonce, deadline)
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // expect emit
        vm.expectEmit();
        emit IERC721.Transfer(address(FORGE), ACCOUNT.addr, tokenID);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC721Transferred(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenID);

        // send transaction
        ForgeV3(FORGE).transferERC721(ACCOUNT.addr, token, tokenID, data, deadline, abi.encodePacked(r, s, v));
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
            facetAddress: address(forgeV3), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = _deployERC1155();
        vm.label(token, "MockERC1155");

        // validator signature
        uint256 tokenID = 1;
        uint256 amount = 10 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);
        bytes memory data = "Transfer ERC1155 token";

        // charge token to forge
        vm.prank(OWNER);
        MockERC1155(token).forceMint(FORGE, tokenID, amount);

        bytes32 structHash = keccak256(
            abi.encode(
                ERC1155_TRANSFER_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, amount, keccak256(data), nonce, deadline
            )
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // expect emit
        vm.expectEmit();
        emit IERC1155.TransferSingle(address(FORGE), address(FORGE), ACCOUNT.addr, tokenID, amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Transferred(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenID, amount);

        // send transaction
        ForgeV3(FORGE).transferERC1155(ACCOUNT.addr, token, tokenID, amount, data, deadline, abi.encodePacked(r, s, v));
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
            facetAddress: address(forgeV3), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc1155
        address token = _deployERC1155();
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
        vm.prank(OWNER);
        MockERC1155(token).forceMintBatch(FORGE, tokenIDs, amounts);

        bytes32 structHash = keccak256(
            abi.encode(
                ERC1155_TRANSFER_BATCH_TYPE_HASH_V1,
                ACCOUNT.addr,
                token,
                keccak256(abi.encodePacked(tokenIDs)),
                keccak256(abi.encodePacked(amounts)),
                keccak256(data),
                nonce,
                deadline
            )
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // expect emit
        vm.expectEmit();
        emit IERC1155.TransferBatch(address(FORGE), address(FORGE), ACCOUNT.addr, tokenIDs, amounts);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Transferred(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenIDs[0], amounts[0]);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC1155Transferred(SERVICE_NAME, uuid, ACCOUNT.addr, token, tokenIDs[1], amounts[1]);

        // send transaction
        ForgeV3(FORGE)
            .transferERC1155Batch(ACCOUNT.addr, token, tokenIDs, amounts, data, deadline, abi.encodePacked(r, s, v));
        assertEq(amounts[0], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[0]), "Transfered tokenID mismatch");
        assertEq(amounts[1], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[1]), "Transfered tokenID mismatch");
    }

    function test_burn_erc20_permit_v1() external {
        // set erc20Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV1.burnERC20Permit.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = _deployERC20();
        vm.label(token, "MockERC20");

        // validator signature
        uint256 uuid = _calcUUID(0);
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        // charge token to account
        vm.prank(OWNER);
        MockERC20(token).forceMint(ACCOUNT.addr, amount);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC20(token).approve(FORGE, amount);

        bytes memory validatorSig;
        {
            bytes32 structHash = keccak256(
                abi.encode(ERC20_BURN_TYPE_HASH_V1, ACCOUNT.addr, token, amount, address(0), 0, nonce, deadline)
            );
            bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // make permit signature
        bytes memory permitSig;
        {
            uint256 permitNonce = IERC20Permit(token).nonces(ACCOUNT.addr);
            bytes32 structHash =
                keccak256(abi.encode(PERMIT_TYPE_HASH, ACCOUNT.addr, address(FORGE), amount, permitNonce, deadline));

            bytes32 domainSeparator = IERC20Permit(token).DOMAIN_SEPARATOR();
            bytes32 hash = MessageHashUtils.toTypedDataHash(domainSeparator, structHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, hash);
            permitSig = abi.encodePacked(r, s, v);
        }

        // expect emit
        vm.expectEmit(true, true, true, true, token);
        emit IERC20.Approval(ACCOUNT.addr, address(FORGE), amount);
        vm.expectEmit(true, true, true, true, token);
        emit IERC20.Transfer(ACCOUNT.addr, address(0), amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC20Burned(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).burnERC20Permit(token, amount, address(0), 0, deadline, validatorSig, permitSig);
        assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "Burned amount mismatch");
    }

    function test_burn_erc20_permit_v2() external {
        // set erc20Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV2.burnERC20Permit.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV2), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = _deployERC20();
        vm.label(token, "MockERC20");

        // validator signature
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);

        // charge token to account
        vm.prank(OWNER);
        MockERC20(token).forceMint(ACCOUNT.addr, amount);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC20(token).approve(FORGE, amount);

        bytes memory fromSig;
        {
            bytes32 fromStructHash =
                keccak256(abi.encode(ERC20_BURN_TYPE_HASH_V2, token, amount, address(0), 0, nonce, deadline));
            bytes32 fromHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, fromStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, fromHash);
            fromSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC20_VALIDATOR_BURN_TYPE_HASH_V2, ACCOUNT.addr, keccak256(fromSig)));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // make permit signature
        bytes memory permitSig;
        {
            uint256 permitNonce = IERC20Permit(token).nonces(ACCOUNT.addr);
            bytes32 structHash =
                keccak256(abi.encode(PERMIT_TYPE_HASH, ACCOUNT.addr, address(FORGE), amount, permitNonce, deadline));

            bytes32 domainSeparator = IERC20Permit(token).DOMAIN_SEPARATOR();
            bytes32 hash = MessageHashUtils.toTypedDataHash(domainSeparator, structHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, hash);
            permitSig = abi.encodePacked(r, s, v);
        }

        // expect emit
        vm.expectEmit(true, true, true, true, token);
        emit IERC20.Approval(ACCOUNT.addr, address(FORGE), amount);
        vm.expectEmit(true, true, true, true, token);
        emit IERC20.Transfer(ACCOUNT.addr, address(0), amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC20Burned(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE)
            .burnERC20Permit(ACCOUNT.addr, token, amount, address(0), 0, deadline, fromSig, validatorSig, permitSig);
        assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "Burned amount mismatch");
    }

    function test_burn_erc20_permit_v3() external {
        // set erc20Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV3.burnERC20Permit.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV3), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = _deployERC20();
        vm.label(token, "MockERC20");

        // validator signature
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);

        // charge token to account
        vm.prank(OWNER);
        MockERC20(token).forceMint(ACCOUNT.addr, amount);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC20(token).approve(FORGE, amount);

        bytes memory validatorSig;
        {
            // validator signature
            bytes32 structHash = keccak256(
                abi.encode(ERC20_BURN_TYPE_HASH_V1, ACCOUNT.addr, token, amount, address(0), 0, nonce, deadline)
            );
            bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);
            validatorSig = abi.encodePacked(r, s, v);
        }
        // make permit signature
        bytes memory permitSig;
        {
            uint256 permitNonce = IERC20Permit(token).nonces(ACCOUNT.addr);
            bytes32 structHash =
                keccak256(abi.encode(PERMIT_TYPE_HASH, ACCOUNT.addr, address(FORGE), amount, permitNonce, deadline));

            bytes32 domainSeparator = IERC20Permit(token).DOMAIN_SEPARATOR();
            bytes32 hash = MessageHashUtils.toTypedDataHash(domainSeparator, structHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, hash);
            permitSig = abi.encodePacked(r, s, v);
        }

        // expect emit
        vm.expectEmit(true, true, true, true, token);
        emit IERC20.Approval(ACCOUNT.addr, address(FORGE), amount);
        vm.expectEmit(true, true, true, true, token);
        emit IERC20.Transfer(ACCOUNT.addr, address(0), amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC20Burned(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);

        // send transaction
        ForgeV3(FORGE).burnERC20Permit(ACCOUNT.addr, token, amount, address(0), 0, deadline, validatorSig, permitSig);
        assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "Burned amount mismatch");
    }

    function test_transfer_from_erc20_v1() external {
        // set erc20TransferFrom
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV1.transferFromERC20.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = _deployERC20();
        vm.label(token, "MockERC20");

        // validator signature
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);

        // charge token to account and approve to forge
        vm.prank(OWNER);
        MockERC20(token).forceMint(ACCOUNT.addr, amount);
        vm.prank(ACCOUNT.addr);
        MockERC20(token).approve(FORGE, amount);

        bytes32 structHash = keccak256(
            abi.encode(ERC20_TRANSFER_FROM_TYPE_HASH_V1, ACCOUNT.addr, token, amount, address(0), 0, nonce, deadline)
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // expect emit
        vm.expectEmit();
        emit IERC20.Transfer(ACCOUNT.addr, address(FORGE), amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC20TransferredFrom(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).transferFromERC20(token, amount, address(0), 0, deadline, abi.encodePacked(r, s, v));
        assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "Account balance should be 0");
        assertEq(amount, MockERC20(token).balanceOf(FORGE), "Forge balance should be amount");
    }

    function test_transfer_from_erc20_v2() external {
        // set erc20TransferFrom
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV2.transferFromERC20.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV2), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = _deployERC20();
        vm.label(token, "MockERC20");

        // validator signature
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);

        // charge token to account and approve to forge
        vm.prank(OWNER);
        MockERC20(token).forceMint(ACCOUNT.addr, amount);
        vm.prank(ACCOUNT.addr);
        MockERC20(token).approve(FORGE, amount);

        bytes memory fromSig;
        {
            bytes32 fromStructHash =
                keccak256(abi.encode(ERC20_TRANSFER_FROM_TYPE_HASH_V2, token, amount, address(0), 0, nonce, deadline));
            bytes32 fromHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, fromStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, fromHash);
            fromSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC20_VALIDATOR_TRANSFER_FROM_TYPE_HASH_V2, ACCOUNT.addr, keccak256(fromSig)));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // expect emit
        vm.expectEmit();
        emit IERC20.Transfer(ACCOUNT.addr, address(FORGE), amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC20TransferredFrom(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE).transferFromERC20(ACCOUNT.addr, token, amount, address(0), 0, deadline, fromSig, validatorSig);
        assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "Account balance should be 0");
        assertEq(amount, MockERC20(token).balanceOf(FORGE), "Forge balance should be amount");
    }

    function test_transfer_from_erc20_permit_v1() external {
        // set erc20TransferFromPermit
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV1.transferFromERC20Permit.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV1), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = _deployERC20();
        vm.label(token, "MockERC20");

        // validator signature
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);

        // charge token to account
        vm.prank(OWNER);
        MockERC20(token).forceMint(ACCOUNT.addr, amount);

        bytes memory validatorSig;
        {
            bytes32 structHash = keccak256(
                abi.encode(
                    ERC20_TRANSFER_FROM_TYPE_HASH_V1, ACCOUNT.addr, token, amount, address(0), 0, nonce, deadline
                )
            );
            bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // make permit signature
        bytes memory permitSig;
        {
            uint256 permitNonce = IERC20Permit(token).nonces(ACCOUNT.addr);
            bytes32 structHash =
                keccak256(abi.encode(PERMIT_TYPE_HASH, ACCOUNT.addr, address(FORGE), amount, permitNonce, deadline));

            bytes32 domainSeparator = IERC20Permit(token).DOMAIN_SEPARATOR();
            bytes32 hash = MessageHashUtils.toTypedDataHash(domainSeparator, structHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, hash);
            permitSig = abi.encodePacked(r, s, v);
        }

        // expect emit
        vm.expectEmit(true, true, true, true, token);
        emit IERC20.Approval(ACCOUNT.addr, address(FORGE), amount);
        vm.expectEmit();
        emit IERC20.Transfer(ACCOUNT.addr, address(FORGE), amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC20TransferredFrom(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).transferFromERC20Permit(token, amount, address(0), 0, deadline, validatorSig, permitSig);
        assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "Account balance should be 0");
        assertEq(amount, MockERC20(token).balanceOf(FORGE), "Forge balance should be amount");
    }

    function test_transfer_from_erc20_permit_v2() external {
        // set erc20TransferFromPermit
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV2.transferFromERC20Permit.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV2), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = _deployERC20();
        vm.label(token, "MockERC20");

        // validator signature
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);

        // charge token to account
        vm.prank(OWNER);
        MockERC20(token).forceMint(ACCOUNT.addr, amount);

        bytes memory fromSig;
        {
            bytes32 fromStructHash =
                keccak256(abi.encode(ERC20_TRANSFER_FROM_TYPE_HASH_V2, token, amount, address(0), 0, nonce, deadline));
            bytes32 fromHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, fromStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, fromHash);
            fromSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC20_VALIDATOR_TRANSFER_FROM_TYPE_HASH_V2, ACCOUNT.addr, keccak256(fromSig)));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // make permit signature
        bytes memory permitSig;
        {
            uint256 permitNonce = IERC20Permit(token).nonces(ACCOUNT.addr);
            bytes32 structHash =
                keccak256(abi.encode(PERMIT_TYPE_HASH, ACCOUNT.addr, address(FORGE), amount, permitNonce, deadline));

            bytes32 domainSeparator = IERC20Permit(token).DOMAIN_SEPARATOR();
            bytes32 hash = MessageHashUtils.toTypedDataHash(domainSeparator, structHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, hash);
            permitSig = abi.encodePacked(r, s, v);
        }

        // expect emit
        vm.expectEmit(true, true, true, true, token);
        emit IERC20.Approval(ACCOUNT.addr, address(FORGE), amount);
        vm.expectEmit();
        emit IERC20.Transfer(ACCOUNT.addr, address(FORGE), amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC20TransferredFrom(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ForgeV2(FORGE)
            .transferFromERC20Permit(
                ACCOUNT.addr, token, amount, address(0), 0, deadline, fromSig, validatorSig, permitSig
            );
        assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "Account balance should be 0");
        assertEq(amount, MockERC20(token).balanceOf(FORGE), "Forge balance should be amount");
    }

    function test_transfer_from_erc20_permit_v3() external {
        // set erc20TransferFromPermit
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = forgeV3.transferFromERC20Permit.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(forgeV3), action: IDiamondCut.FacetCutAction.Add, functionSelectors: functionSelectors
        });
        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = _deployERC20();
        vm.label(token, "MockERC20");

        // validator signature
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);

        // charge token to account
        vm.prank(OWNER);
        MockERC20(token).forceMint(ACCOUNT.addr, amount);

        bytes memory validatorSig;
        {
            bytes32 structHash = keccak256(
                abi.encode(
                    ERC20_TRANSFER_FROM_TYPE_HASH_V3, ACCOUNT.addr, token, amount, address(0), 0, nonce, deadline
                )
            );
            bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // make permit signature
        bytes memory permitSig;
        {
            uint256 permitNonce = IERC20Permit(token).nonces(ACCOUNT.addr);
            bytes32 structHash =
                keccak256(abi.encode(PERMIT_TYPE_HASH, ACCOUNT.addr, address(FORGE), amount, permitNonce, deadline));

            bytes32 domainSeparator = IERC20Permit(token).DOMAIN_SEPARATOR();
            bytes32 hash = MessageHashUtils.toTypedDataHash(domainSeparator, structHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, hash);
            permitSig = abi.encodePacked(r, s, v);
        }

        // expect emit
        vm.expectEmit(true, true, true, true, token);
        emit IERC20.Approval(ACCOUNT.addr, address(FORGE), amount);
        vm.expectEmit();
        emit IERC20.Transfer(ACCOUNT.addr, address(FORGE), amount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC20TransferredFrom(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);

        // send transaction
        ForgeV3(FORGE)
            .transferFromERC20Permit(ACCOUNT.addr, token, amount, address(0), 0, deadline, validatorSig, permitSig);
        assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "Account balance should be 0");
        assertEq(amount, MockERC20(token).balanceOf(FORGE), "Forge balance should be amount");
    }

    address private constant feeRecipient = address(bytes20("FeeRecipient"));
    uint256 private constant feeBPS = 100; // 1%

    function test_erc20_fee_v1() external {
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        {
            // set v1 funcs
            bytes4[] memory functionSelectors = new bytes4[](6);
            functionSelectors[0] = forgeV1.mintERC20.selector;
            functionSelectors[1] = forgeV1.transferERC20.selector;
            functionSelectors[2] = forgeV1.transferFromERC20.selector;
            functionSelectors[3] = forgeV1.transferFromERC20Permit.selector;
            functionSelectors[4] = forgeV1.burnERC20.selector;
            functionSelectors[5] = forgeV1.burnERC20Permit.selector;

            addCuts[0] = IDiamondCut.FacetCut({
                facetAddress: address(forgeV1),
                action: IDiamondCut.FacetCutAction.Add,
                functionSelectors: functionSelectors
            });
        }

        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = _deployERC20();
        vm.label(token, "MockERC20");

        // set test variables
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 expectedFee = (amount * feeBPS) / 10_000;
        uint256 expectValue = amount - expectedFee;
        // mint
        {
            uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
            uint256 uuid = _calcUUID(nonce);

            // expect emit
            vm.expectEmit();
            emit IERC20.Transfer(address(0), ACCOUNT.addr, expectValue);
            vm.expectEmit();
            emit IERC20.Transfer(address(0), feeRecipient, expectedFee);
            vm.expectEmit();
            emit ForgeFactory.ERC20Minted(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);
            vm.expectEmit();
            emit ForgeFactory.ERC20FeeCollected(uuid, feeRecipient, token, expectedFee);

            // send transaction
            bytes32 structHash = keccak256(
                abi.encode(ERC20_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, amount, feeRecipient, feeBPS, nonce, deadline)
            );
            bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

            vm.prank(ACCOUNT.addr);
            ForgeV1(FORGE).mintERC20(token, amount, feeRecipient, feeBPS, deadline, abi.encodePacked(r, s, v));
            assertEq(expectValue, MockERC20(token).balanceOf(ACCOUNT.addr), "Minted amount mismatch");
        }
        vm.prank(OWNER);
        MockERC20(token).forceMint(ACCOUNT.addr, expectedFee); // charge token to forge
        // burnFrom
        {
            vm.prank(ACCOUNT.addr);
            IERC20(token).approve(FORGE, amount);

            uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
            uint256 uuid = _calcUUID(nonce);

            // expect emit
            vm.expectEmit();
            emit IERC20.Transfer(ACCOUNT.addr, address(0), expectValue);
            vm.expectEmit();
            emit IERC20.Transfer(ACCOUNT.addr, feeRecipient, expectedFee);
            vm.expectEmit();
            emit ForgeFactory.ERC20Burned(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);
            vm.expectEmit();
            emit ForgeFactory.ERC20FeeCollected(uuid, feeRecipient, token, expectedFee);

            // send transaction
            bytes32 structHash = keccak256(
                abi.encode(ERC20_BURN_TYPE_HASH_V1, ACCOUNT.addr, token, amount, feeRecipient, feeBPS, nonce, deadline)
            );
            bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

            vm.prank(ACCOUNT.addr);
            ForgeV1(FORGE).burnERC20(token, amount, feeRecipient, feeBPS, deadline, abi.encodePacked(r, s, v));
            assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "Burned amount mismatch");
        }
        vm.prank(OWNER);
        MockERC20(token).forceMint(ACCOUNT.addr, amount);
        // burnFromPermit
        {
            uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
            uint256 uuid = _calcUUID(nonce);

            // expect emit
            vm.expectEmit();
            emit IERC20.Approval(ACCOUNT.addr, address(FORGE), amount);
            vm.expectEmit();
            emit IERC20.Transfer(ACCOUNT.addr, address(0), expectValue);
            vm.expectEmit();
            emit IERC20.Transfer(ACCOUNT.addr, feeRecipient, expectedFee);
            vm.expectEmit();
            emit ForgeFactory.ERC20Burned(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);
            vm.expectEmit();
            emit ForgeFactory.ERC20FeeCollected(uuid, feeRecipient, token, expectedFee);

            bytes memory validatorSig;
            {
                bytes32 structHash = keccak256(
                    abi.encode(
                        ERC20_BURN_TYPE_HASH_V1, ACCOUNT.addr, token, amount, feeRecipient, feeBPS, nonce, deadline
                    )
                );
                bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);
                validatorSig = abi.encodePacked(r, s, v);
            }

            // make permit signature
            bytes memory permitSig;
            {
                uint256 permitNonce = IERC20Permit(token).nonces(ACCOUNT.addr);
                bytes32 structHash = keccak256(
                    abi.encode(PERMIT_TYPE_HASH, ACCOUNT.addr, address(FORGE), amount, permitNonce, deadline)
                );

                bytes32 domainSeparator = IERC20Permit(token).DOMAIN_SEPARATOR();
                bytes32 hash = MessageHashUtils.toTypedDataHash(domainSeparator, structHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, hash);
                permitSig = abi.encodePacked(r, s, v);
            }

            vm.prank(ACCOUNT.addr);
            ForgeV1(FORGE).burnERC20Permit(token, amount, feeRecipient, feeBPS, deadline, validatorSig, permitSig);
            assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "Burned amount mismatch");
        }

        vm.prank(OWNER);
        MockERC20(token).forceMint(FORGE, amount);
        // transfer
        {
            uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
            uint256 uuid = _calcUUID(nonce);

            // charge token to forge
            vm.prank(OWNER);
            MockERC20(token).forceMint(FORGE, amount);

            bytes32 structHash = keccak256(
                abi.encode(
                    ERC20_TRANSFER_TYPE_HASH_V1, ACCOUNT.addr, token, amount, feeRecipient, feeBPS, nonce, deadline
                )
            );
            bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

            // expect emit
            vm.expectEmit();
            emit IERC20.Transfer(FORGE, ACCOUNT.addr, expectValue);
            vm.expectEmit();
            emit IERC20.Transfer(FORGE, feeRecipient, expectedFee);
            vm.expectEmit();
            emit ForgeFactory.ERC20Transferred(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);
            vm.expectEmit();
            emit ForgeFactory.ERC20FeeCollected(uuid, feeRecipient, token, expectedFee);

            // send transaction
            vm.prank(ACCOUNT.addr);
            ForgeV1(FORGE).transferERC20(token, amount, feeRecipient, feeBPS, deadline, abi.encodePacked(r, s, v));
            assertEq(expectValue, MockERC20(token).balanceOf(ACCOUNT.addr), "Transfered amount mismatch");
        }
        vm.prank(OWNER);
        MockERC20(token).forceMint(ACCOUNT.addr, expectedFee);
        // transferFrom
        {
            vm.prank(ACCOUNT.addr);
            IERC20(token).approve(FORGE, amount);

            uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
            uint256 uuid = _calcUUID(nonce);

            bytes32 structHash = keccak256(
                abi.encode(
                    ERC20_TRANSFER_FROM_TYPE_HASH_V1, ACCOUNT.addr, token, amount, feeRecipient, feeBPS, nonce, deadline
                )
            );
            bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

            // expect emit
            vm.expectEmit();
            emit IERC20.Transfer(ACCOUNT.addr, address(FORGE), amount);
            vm.expectEmit();
            emit IERC20.Transfer(address(FORGE), feeRecipient, expectedFee);
            vm.expectEmit();
            emit ForgeFactory.ERC20TransferredFrom(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);
            vm.expectEmit();
            emit ForgeFactory.ERC20FeeCollected(uuid, feeRecipient, token, expectedFee);

            // send transaction
            vm.prank(ACCOUNT.addr);
            ForgeV1(FORGE).transferFromERC20(token, amount, feeRecipient, feeBPS, deadline, abi.encodePacked(r, s, v));
            assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "TransferFrom amount mismatch");
        }
        vm.prank(OWNER);
        MockERC20(token).forceMint(ACCOUNT.addr, amount);
        // transferFromPermit
        {
            uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
            uint256 uuid = _calcUUID(nonce);

            // expect emit
            vm.expectEmit(true, true, true, true, token);
            emit IERC20.Approval(ACCOUNT.addr, address(FORGE), amount);
            vm.expectEmit();
            emit IERC20.Transfer(ACCOUNT.addr, address(FORGE), amount);
            vm.expectEmit();
            emit IERC20.Transfer(address(FORGE), feeRecipient, expectedFee);
            vm.expectEmit();
            emit ForgeFactory.ERC20TransferredFrom(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);
            vm.expectEmit();
            emit ForgeFactory.ERC20FeeCollected(uuid, feeRecipient, token, expectedFee);

            bytes memory validatorSig;
            {
                bytes32 structHash = keccak256(
                    abi.encode(
                        ERC20_TRANSFER_FROM_TYPE_HASH_V1,
                        ACCOUNT.addr,
                        token,
                        amount,
                        feeRecipient,
                        feeBPS,
                        nonce,
                        deadline
                    )
                );
                bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);
                validatorSig = abi.encodePacked(r, s, v);
            }

            // make permit signature
            bytes memory permitSig;
            {
                uint256 permitNonce = IERC20Permit(token).nonces(ACCOUNT.addr);
                bytes32 structHash = keccak256(
                    abi.encode(PERMIT_TYPE_HASH, ACCOUNT.addr, address(FORGE), amount, permitNonce, deadline)
                );

                bytes32 domainSeparator = IERC20Permit(token).DOMAIN_SEPARATOR();
                bytes32 hash = MessageHashUtils.toTypedDataHash(domainSeparator, structHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, hash);
                permitSig = abi.encodePacked(r, s, v);
            }

            vm.prank(ACCOUNT.addr);
            ForgeV1(FORGE)
                .transferFromERC20Permit(token, amount, feeRecipient, feeBPS, deadline, validatorSig, permitSig);
            assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "TransferFrom Permit amount mismatch");
        }

        assertEq(expectedFee * 6, IERC20(token).balanceOf(feeRecipient), "Fee recipient balance mismatch");
    }

    function test_erc20_fee_v2() external {
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        {
            // set v2 funcs
            bytes4[] memory functionSelectors = new bytes4[](6);
            functionSelectors[0] = forgeV2.mintERC20.selector;
            functionSelectors[1] = forgeV2.transferERC20.selector;
            functionSelectors[2] = forgeV2.transferFromERC20.selector;
            functionSelectors[3] = forgeV2.transferFromERC20Permit.selector;
            functionSelectors[4] = forgeV2.burnERC20.selector;
            functionSelectors[5] = forgeV2.burnERC20Permit.selector;

            addCuts[0] = IDiamondCut.FacetCut({
                facetAddress: address(forgeV2),
                action: IDiamondCut.FacetCutAction.Add,
                functionSelectors: functionSelectors
            });
        }

        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = _deployERC20();
        vm.label(token, "MockERC20");

        // set test variables
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 expectedFee = (amount * feeBPS) / 10_000;
        uint256 expectValue = amount - expectedFee;
        // mint
        {
            uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
            uint256 uuid = _calcUUID(nonce);

            // expect emit
            vm.expectEmit();
            emit IERC20.Transfer(address(0), ACCOUNT.addr, expectValue);
            vm.expectEmit();
            emit IERC20.Transfer(address(0), feeRecipient, expectedFee);
            vm.expectEmit();
            emit ForgeFactory.ERC20Minted(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);
            vm.expectEmit();
            emit ForgeFactory.ERC20FeeCollected(uuid, feeRecipient, token, expectedFee);
            bytes memory recipientSig;
            {
                bytes32 recipientStructHash = keccak256(
                    abi.encode(ERC20_MINT_TYPE_HASH_V2, token, amount, feeRecipient, feeBPS, nonce, deadline)
                );
                bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
                recipientSig = abi.encodePacked(r, s, v);
            }
            bytes memory validatorSig;
            {
                bytes32 validatorStructHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_MINT_TYPE_HASH_V2, ACCOUNT.addr, keccak256(recipientSig)));
                bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
                validatorSig = abi.encodePacked(r, s, v);
            }
            // send transaction
            vm.prank(ACCOUNT.addr);
            ForgeV2(FORGE)
                .mintERC20(ACCOUNT.addr, token, amount, feeRecipient, feeBPS, deadline, recipientSig, validatorSig);
            assertEq(expectValue, MockERC20(token).balanceOf(ACCOUNT.addr), "Minted amount mismatch");
        }
        vm.prank(OWNER);
        MockERC20(token).forceMint(ACCOUNT.addr, expectedFee); // charge token to
        // burnFrom
        {
            vm.prank(ACCOUNT.addr);
            IERC20(token).approve(FORGE, amount);

            uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
            uint256 uuid = _calcUUID(nonce);

            // expect emit
            vm.expectEmit();
            emit IERC20.Transfer(ACCOUNT.addr, address(0), expectValue);
            vm.expectEmit();
            emit IERC20.Transfer(ACCOUNT.addr, feeRecipient, expectedFee);
            vm.expectEmit();
            emit ForgeFactory.ERC20Burned(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);
            vm.expectEmit();
            emit ForgeFactory.ERC20FeeCollected(uuid, feeRecipient, token, expectedFee);
            bytes memory recipientSig;
            {
                bytes32 recipientStructHash = keccak256(
                    abi.encode(ERC20_BURN_TYPE_HASH_V2, token, amount, feeRecipient, feeBPS, nonce, deadline)
                );
                bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
                recipientSig = abi.encodePacked(r, s, v);
            }
            bytes memory validatorSig;
            {
                bytes32 validatorStructHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_BURN_TYPE_HASH_V2, ACCOUNT.addr, keccak256(recipientSig)));
                bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
                validatorSig = abi.encodePacked(r, s, v);
            }

            // send transaction
            ForgeV2(FORGE)
                .burnERC20(ACCOUNT.addr, token, amount, feeRecipient, feeBPS, deadline, recipientSig, validatorSig);
            assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "Burned amount mismatch");
        }
        vm.prank(OWNER);
        MockERC20(token).forceMint(ACCOUNT.addr, amount);
        // burnFromPermit
        {
            uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
            uint256 uuid = _calcUUID(nonce);
            // expect emit
            vm.expectEmit();
            emit IERC20.Approval(ACCOUNT.addr, address(FORGE), amount);
            vm.expectEmit();
            emit IERC20.Transfer(ACCOUNT.addr, address(0), expectValue);
            vm.expectEmit();
            emit IERC20.Transfer(ACCOUNT.addr, feeRecipient, expectedFee);
            vm.expectEmit();
            emit ForgeFactory.ERC20Burned(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);
            vm.expectEmit();
            emit ForgeFactory.ERC20FeeCollected(uuid, feeRecipient, token, expectedFee);
            bytes memory recipientSig;
            {
                bytes32 recipientStructHash = keccak256(
                    abi.encode(ERC20_BURN_TYPE_HASH_V2, token, amount, feeRecipient, feeBPS, nonce, deadline)
                );
                bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
                recipientSig = abi.encodePacked(r, s, v);
            }
            bytes memory validatorSig;
            {
                bytes32 validatorStructHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_BURN_TYPE_HASH_V2, ACCOUNT.addr, keccak256(recipientSig)));
                bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
                validatorSig = abi.encodePacked(r, s, v);
            }
            // make permit signature
            bytes memory permitSig;
            {
                uint256 permitNonce = IERC20Permit(token).nonces(ACCOUNT.addr);
                bytes32 structHash = keccak256(
                    abi.encode(PERMIT_TYPE_HASH, ACCOUNT.addr, address(FORGE), amount, permitNonce, deadline)
                );

                bytes32 domainSeparator = IERC20Permit(token).DOMAIN_SEPARATOR();
                bytes32 hash = MessageHashUtils.toTypedDataHash(domainSeparator, structHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, hash);
                permitSig = abi.encodePacked(r, s, v);
            }
            // send transaction
            ForgeV2(FORGE)
                .burnERC20Permit(
                    ACCOUNT.addr, token, amount, feeRecipient, feeBPS, deadline, recipientSig, validatorSig, permitSig
                );
            assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "Burned amount mismatch");
        }

        vm.prank(OWNER);
        MockERC20(token).forceMint(FORGE, amount);
        // transfer
        {
            uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
            uint256 uuid = _calcUUID(nonce);

            // expect emit
            vm.expectEmit();
            emit IERC20.Transfer(FORGE, ACCOUNT.addr, expectValue);
            vm.expectEmit();
            emit IERC20.Transfer(FORGE, feeRecipient, expectedFee);
            vm.expectEmit();
            emit ForgeFactory.ERC20Transferred(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);
            vm.expectEmit();
            emit ForgeFactory.ERC20FeeCollected(uuid, feeRecipient, token, expectedFee);

            bytes memory recipientSig;
            {
                bytes32 recipientStructHash = keccak256(
                    abi.encode(ERC20_TRANSFER_TYPE_HASH_V2, token, amount, feeRecipient, feeBPS, nonce, deadline)
                );
                bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
                recipientSig = abi.encodePacked(r, s, v);
            }
            bytes memory validatorSig;
            {
                bytes32 validatorStructHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_TRANSFER_TYPE_HASH_V2, ACCOUNT.addr, keccak256(recipientSig)));
                bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
                validatorSig = abi.encodePacked(r, s, v);
            }
            // send transaction
            ForgeV2(FORGE)
                .transferERC20(ACCOUNT.addr, token, amount, feeRecipient, feeBPS, deadline, recipientSig, validatorSig);
            assertEq(expectValue, MockERC20(token).balanceOf(ACCOUNT.addr), "Transfer amount mismatch");
        }
        vm.prank(OWNER);
        MockERC20(token).forceMint(ACCOUNT.addr, expectedFee);
        // transferFrom
        {
            vm.prank(ACCOUNT.addr);
            IERC20(token).approve(FORGE, amount);

            uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
            uint256 uuid = _calcUUID(nonce);

            // expect emit
            vm.expectEmit();
            emit IERC20.Transfer(ACCOUNT.addr, address(FORGE), amount);
            vm.expectEmit();
            emit IERC20.Transfer(address(FORGE), feeRecipient, expectedFee);
            vm.expectEmit();
            emit ForgeFactory.ERC20TransferredFrom(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);
            vm.expectEmit();
            emit ForgeFactory.ERC20FeeCollected(uuid, feeRecipient, token, expectedFee);

            bytes memory fromSig;
            {
                bytes32 fromStructHash = keccak256(
                    abi.encode(ERC20_TRANSFER_FROM_TYPE_HASH_V2, token, amount, feeRecipient, feeBPS, nonce, deadline)
                );
                bytes32 fromHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, fromStructHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, fromHash);
                fromSig = abi.encodePacked(r, s, v);
            }
            bytes memory validatorSig;
            {
                bytes32 validatorStructHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_TRANSFER_FROM_TYPE_HASH_V2, ACCOUNT.addr, keccak256(fromSig)));
                bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
                validatorSig = abi.encodePacked(r, s, v);
            }

            // send transaction
            ForgeV2(FORGE)
                .transferFromERC20(ACCOUNT.addr, token, amount, feeRecipient, feeBPS, deadline, fromSig, validatorSig);
            assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "TransferFrom amount mismatch");
        }
        vm.prank(OWNER);
        MockERC20(token).forceMint(ACCOUNT.addr, amount);
        // transferFromPermit
        {
            uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
            uint256 uuid = _calcUUID(nonce);

            // expect emit
            vm.expectEmit(true, true, true, true, token);
            emit IERC20.Approval(ACCOUNT.addr, address(FORGE), amount);
            vm.expectEmit();
            emit IERC20.Transfer(ACCOUNT.addr, address(FORGE), amount);
            vm.expectEmit();
            emit IERC20.Transfer(address(FORGE), feeRecipient, expectedFee);
            vm.expectEmit();
            emit ForgeFactory.ERC20TransferredFrom(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);
            vm.expectEmit();
            emit ForgeFactory.ERC20FeeCollected(uuid, feeRecipient, token, expectedFee);

            bytes memory fromSig;
            {
                bytes32 fromStructHash = keccak256(
                    abi.encode(ERC20_TRANSFER_FROM_TYPE_HASH_V2, token, amount, feeRecipient, feeBPS, nonce, deadline)
                );
                bytes32 fromHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, fromStructHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, fromHash);
                fromSig = abi.encodePacked(r, s, v);
            }
            bytes memory validatorSig;
            {
                bytes32 validatorStructHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_TRANSFER_FROM_TYPE_HASH_V2, ACCOUNT.addr, keccak256(fromSig)));
                bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
                validatorSig = abi.encodePacked(r, s, v);
            }

            // make permit signature
            bytes memory permitSig;
            {
                uint256 permitNonce = IERC20Permit(token).nonces(ACCOUNT.addr);
                bytes32 structHash = keccak256(
                    abi.encode(PERMIT_TYPE_HASH, ACCOUNT.addr, address(FORGE), amount, permitNonce, deadline)
                );

                bytes32 domainSeparator = IERC20Permit(token).DOMAIN_SEPARATOR();
                bytes32 hash = MessageHashUtils.toTypedDataHash(domainSeparator, structHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, hash);
                permitSig = abi.encodePacked(r, s, v);
            }
            vm.prank(ACCOUNT.addr);
            ForgeV2(FORGE)
                .transferFromERC20Permit(
                    ACCOUNT.addr, token, amount, feeRecipient, feeBPS, deadline, fromSig, validatorSig, permitSig
                );
            assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "TransferFrom Permit amount mismatch");
        }

        assertEq(expectedFee * 6, IERC20(token).balanceOf(feeRecipient), "Fee recipient balance mismatch");
    }

    function test_erc20_fee_v3() external {
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        {
            // set v3 funcs
            bytes4[] memory functionSelectors = new bytes4[](4);
            functionSelectors[0] = forgeV3.mintERC20.selector;
            functionSelectors[1] = forgeV3.transferERC20.selector;
            functionSelectors[2] = forgeV3.transferFromERC20Permit.selector;
            functionSelectors[3] = forgeV3.burnERC20Permit.selector;

            addCuts[0] = IDiamondCut.FacetCut({
                facetAddress: address(forgeV3),
                action: IDiamondCut.FacetCutAction.Add,
                functionSelectors: functionSelectors
            });
        }

        // deploy forge
        vm.prank(OWNER);
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        vm.label(FORGE, "Forge");
        bytes32 DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        // deploy mock erc20
        address token = _deployERC20();
        vm.label(token, "MockERC20");

        // set test variables
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 expectedFee = (amount * feeBPS) / 10_000;
        uint256 expectValue = amount - expectedFee;
        // mint
        {
            uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
            uint256 uuid = _calcUUID(nonce);

            // expect emit
            vm.expectEmit();
            emit IERC20.Transfer(address(0), ACCOUNT.addr, expectValue);
            vm.expectEmit();
            emit IERC20.Transfer(address(0), feeRecipient, expectedFee);
            vm.expectEmit();
            emit ForgeFactory.ERC20Minted(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);
            vm.expectEmit();
            emit ForgeFactory.ERC20FeeCollected(uuid, feeRecipient, token, expectedFee);

            bytes memory validatorSig;
            {
                bytes32 structHash = keccak256(
                    abi.encode(
                        ERC20_MINT_TYPE_HASH_V3, ACCOUNT.addr, token, amount, feeRecipient, feeBPS, nonce, deadline
                    )
                );
                bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);
                validatorSig = abi.encodePacked(r, s, v);
            }

            // send transaction
            ForgeV3(FORGE).mintERC20(ACCOUNT.addr, token, amount, feeRecipient, feeBPS, deadline, validatorSig);
            assertEq(expectValue, MockERC20(token).balanceOf(ACCOUNT.addr), "Minted amount mismatch");
        }
        vm.prank(OWNER);
        MockERC20(token).forceMint(ACCOUNT.addr, expectedFee); // charge token to
        // burnFromPermit
        {
            vm.prank(ACCOUNT.addr);
            IERC20(token).approve(FORGE, amount);

            uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
            uint256 uuid = _calcUUID(nonce);

            // expect emit
            vm.expectEmit();
            emit IERC20.Transfer(ACCOUNT.addr, address(0), expectValue);
            vm.expectEmit();
            emit IERC20.Transfer(ACCOUNT.addr, feeRecipient, expectedFee);
            vm.expectEmit();
            emit ForgeFactory.ERC20Burned(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);
            vm.expectEmit();
            emit ForgeFactory.ERC20FeeCollected(uuid, feeRecipient, token, expectedFee);

            bytes memory validatorSig;
            {
                bytes32 structHash = keccak256(
                    abi.encode(
                        ERC20_BURN_TYPE_HASH_V3, ACCOUNT.addr, token, amount, feeRecipient, feeBPS, nonce, deadline
                    )
                );
                bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);
                validatorSig = abi.encodePacked(r, s, v);
            }

            // make permit signature
            bytes memory permitSig;
            {
                uint256 permitNonce = IERC20Permit(token).nonces(ACCOUNT.addr);
                bytes32 structHash = keccak256(
                    abi.encode(PERMIT_TYPE_HASH, ACCOUNT.addr, address(FORGE), amount, permitNonce, deadline)
                );

                bytes32 domainSeparator = IERC20Permit(token).DOMAIN_SEPARATOR();
                bytes32 hash = MessageHashUtils.toTypedDataHash(domainSeparator, structHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, hash);
                permitSig = abi.encodePacked(r, s, v);
            }

            // send transaction
            ForgeV3(FORGE)
                .burnERC20Permit(ACCOUNT.addr, token, amount, feeRecipient, feeBPS, deadline, validatorSig, permitSig);
            assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "Burned amount mismatch");
        }
        vm.prank(OWNER);
        MockERC20(token).forceMint(FORGE, amount);
        // transfer
        {
            uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
            uint256 uuid = _calcUUID(nonce);

            // expect emit
            vm.expectEmit();
            emit IERC20.Transfer(FORGE, ACCOUNT.addr, expectValue);
            vm.expectEmit();
            emit IERC20.Transfer(FORGE, feeRecipient, expectedFee);
            vm.expectEmit();
            emit ForgeFactory.ERC20Transferred(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);
            vm.expectEmit();
            emit ForgeFactory.ERC20FeeCollected(uuid, feeRecipient, token, expectedFee);

            bytes memory validatorSig;
            {
                bytes32 structHash = keccak256(
                    abi.encode(
                        ERC20_TRANSFER_TYPE_HASH_V3, ACCOUNT.addr, token, amount, feeRecipient, feeBPS, nonce, deadline
                    )
                );
                bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);
                validatorSig = abi.encodePacked(r, s, v);
            }

            // send transaction
            ForgeV3(FORGE).transferERC20(ACCOUNT.addr, token, amount, feeRecipient, feeBPS, deadline, validatorSig);
            assertEq(expectValue, MockERC20(token).balanceOf(ACCOUNT.addr), "Transfer amount mismatch");
        }
        vm.prank(OWNER);
        MockERC20(token).forceMint(ACCOUNT.addr, expectedFee);
        // transferFromPermit
        {
            uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
            uint256 uuid = _calcUUID(nonce);

            // expect emit
            vm.expectEmit();
            emit IERC20.Approval(ACCOUNT.addr, address(FORGE), amount);
            vm.expectEmit();
            emit IERC20.Transfer(ACCOUNT.addr, address(FORGE), amount);
            vm.expectEmit();
            emit IERC20.Transfer(address(FORGE), feeRecipient, expectedFee);
            vm.expectEmit();
            emit ForgeFactory.ERC20TransferredFrom(SERVICE_NAME, uuid, ACCOUNT.addr, token, amount);
            vm.expectEmit();
            emit ForgeFactory.ERC20FeeCollected(uuid, feeRecipient, token, expectedFee);

            bytes memory validatorSig;
            {
                bytes32 structHash = keccak256(
                    abi.encode(
                        ERC20_TRANSFER_FROM_TYPE_HASH_V3,
                        ACCOUNT.addr,
                        token,
                        amount,
                        feeRecipient,
                        feeBPS,
                        nonce,
                        deadline
                    )
                );
                bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);
                validatorSig = abi.encodePacked(r, s, v);
            }

            // make permit signature
            bytes memory permitSig;
            {
                uint256 permitNonce = IERC20Permit(token).nonces(ACCOUNT.addr);
                bytes32 structHash = keccak256(
                    abi.encode(PERMIT_TYPE_HASH, ACCOUNT.addr, address(FORGE), amount, permitNonce, deadline)
                );

                bytes32 domainSeparator = IERC20Permit(token).DOMAIN_SEPARATOR();
                bytes32 hash = MessageHashUtils.toTypedDataHash(domainSeparator, structHash);
                (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, hash);
                permitSig = abi.encodePacked(r, s, v);
            }

            // send transaction
            ForgeV3(FORGE)
                .transferFromERC20Permit(
                    ACCOUNT.addr, token, amount, feeRecipient, feeBPS, deadline, validatorSig, permitSig
                );
            assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "TransferFrom Permit amount mismatch");
        }

        assertEq(expectedFee * 4, IERC20(token).balanceOf(feeRecipient), "Fee recipient balance mismatch");
    }

    function _calcUUID(uint256 nonce) internal view returns (uint256) {
        return uint256(keccak256(abi.encode(FORGE, ACCOUNT.addr, nonce)));
    }
    /*
    function test_print_struct_hash() external {
        //        console.logBytes32(ERC20_VALIDATOR_MINT_TYPE_HASH_V2);
        address token = 0xE7199f0cBd9114A23f6B44cCf20CE09D0D561a96;
        uint256[] memory tokenIDs = new uint256[](2);
        tokenIDs[0] = 1;
        tokenIDs[1] = 2;
        uint256[] memory amounts = new uint256[](2);
        amounts[0] = 1;
        amounts[1] = 2;

        uint256 nonce = 11;
        uint256 deadline = 1751943489;
        {
            bytes32 domainSeparator = keccak256(
                abi.encode(
                    keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"),
                    keccak256(bytes("MockService")),
                    keccak256(bytes("1")),
                    612044,
                    0xa55D9fFf44bF394bDB42b0A20D7Fcb2090e424d2
                )
            );
            bytes32 recipientStructHash = keccak256(
                abi.encode(
                    ERC1155_MINT_BATCH_TYPE_HASH_V2,
                    token,
                    keccak256(abi.encodePacked(tokenIDs)), // <- 여기가 핵심
                    keccak256(abi.encodePacked(amounts)), // <- 여기도
                    nonce,
                    deadline
                )
            );
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(domainSeparator, recipientStructHash);
            console.logBytes32(keccak256(abi.encodePacked(tokenIDs)));
            console.logBytes(abi.encodePacked(tokenIDs));
        }
    }
    */
}
