// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std-1.9.7/src/Test.sol";
import {Vm} from "forge-std-1.9.7/src/Vm.sol";

import {IDiamondCut} from "diamond-3-hardhat-1.0.0/contracts/interfaces/IDiamondCut.sol";
import {IDiamondLoupe} from "diamond-3-hardhat-1.0.0/contracts/interfaces/IDiamondLoupe.sol";
import {IERC173} from "diamond-3-hardhat-1.0.0/contracts/interfaces/IERC173.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts-5.3.0/proxy/ERC1967/ERC1967Proxy.sol";
import {MessageHashUtils} from "@openzeppelin-contracts-5.3.0/utils/cryptography/MessageHashUtils.sol";

import {ForgeProxyCode} from "../src/ForgeProxy.sol";
import {Diamond3Facet} from "../src/Diamond3Facet.sol";
import {ERC20ForgeFacet} from "../src/ERC20Forge.sol";
import {ERC721ForgeFacet} from "../src/ERC721Forge.sol";
import {ERC1155ForgeFacet} from "../src/ERC1155Forge.sol";
import {TokenForgeFactory} from "../src/TokenForgeFactory.sol";
import "../src/BaseForge.sol";

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
    ERC20ForgeFacet public eRC20ForgeFacet;
    ERC721ForgeFacet public eRC721ForgeFacet;
    ERC1155ForgeFacet public eRC1155ForgeFacet;
    TokenForgeFactory public tokenForgeFactory;
    address public FORGE;

    function setUp() public {
        vm.label(OWNER, "owner");
        vm.startPrank(OWNER);
        // deploy contracts
        // deploy faucets
        diamond3Facet = new Diamond3Facet();
        baseForgeFacet = new BaseForgeFacet();
        eRC20ForgeFacet = new ERC20ForgeFacet();
        eRC721ForgeFacet = new ERC721ForgeFacet();
        eRC1155ForgeFacet = new ERC1155ForgeFacet();
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

    function test_mint_erc20() external {
        // set erc20Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = eRC20ForgeFacet.mintERC20.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC20ForgeFacet),
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
        uint256 uuid = 0;
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        bytes32 structHash =
            keccak256(abi.encode(ERC20_MINT_TYPEHASH, uuid, ACCOUNT.addr, token, amount, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC20ForgeFacet(FORGE).mintERC20(uuid, token, amount, deadline, abi.encodePacked(r, s, v));
        assertEq(amount, MockERC20(token).balanceOf(ACCOUNT.addr), "Minted amount mismatch");
    }

    function test_mintto_erc20() external {
        // set erc20Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = eRC20ForgeFacet.mintERC20To.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC20ForgeFacet),
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
        uint256 uuid = 0;
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        bytes memory recipientSig;

        {
            bytes32 recipientStructHash = keccak256(abi.encode(ERC20_MINTTO_TYPEHASH, token, amount, nonce, deadline));
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC20_VALIDATOR_MINTTO_TYPEHASH, uuid, ACCOUNT.addr, recipientSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC20ForgeFacet(FORGE).mintERC20To(uuid, ACCOUNT.addr, token, amount, deadline, recipientSig, validatorSig);
        assertEq(amount, MockERC20(token).balanceOf(ACCOUNT.addr), "Minted amount mismatch");
    }

    function test_mint_erc721() external {
        // set erc721Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = eRC721ForgeFacet.mintERC721.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC721ForgeFacet),
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
        uint256 uuid = 0;
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        bytes32 structHash =
            keccak256(abi.encode(ERC721_MINT_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC721ForgeFacet(FORGE).mintERC721(uuid, token, tokenID, deadline, abi.encodePacked(r, s, v));
        assertEq(ACCOUNT.addr, MockERC721(token).ownerOf(tokenID), "Minted tokenID mismatch");
    }

    function test_mintto_erc721() external {
        // set erc721Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = eRC721ForgeFacet.mintERC721To.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC721ForgeFacet),
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
        uint256 uuid = 0;
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        bytes memory recipientSig;

        {
            bytes32 recipientStructHash = keccak256(abi.encode(ERC721_MINTTO_TYPEHASH, token, tokenID, nonce, deadline));
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC721_VALIDATOR_MINTTO_TYPEHASH, uuid, ACCOUNT.addr, recipientSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC721ForgeFacet(FORGE).mintERC721To(uuid, ACCOUNT.addr, token, tokenID, deadline, recipientSig, validatorSig);
        assertEq(ACCOUNT.addr, MockERC721(token).ownerOf(tokenID), "Minted tokenID mismatch");
    }

    function test_mint_erc1155() external {
        // set erc1155Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = eRC1155ForgeFacet.mintERC1155.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC1155ForgeFacet),
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
        uint256 uuid = 0;
        uint256 tokenID = 1;
        uint256 amount = 10 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        bytes memory data = "Minting ERC1155 token";

        bytes32 structHash =
            keccak256(abi.encode(ERC1155_MINT_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, amount, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC1155ForgeFacet(FORGE).mintERC1155(uuid, token, tokenID, amount, data, deadline, abi.encodePacked(r, s, v));
        assertEq(amount, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenID), "Minted tokenID mismatch");
    }

    function test_mint_batch_erc1155() external {
        // set erc1155Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = eRC1155ForgeFacet.mintERC1155Batch.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC1155ForgeFacet),
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
        uint256 uuid = 0;
        uint256[] memory tokenIDs = new uint256[](2);
        tokenIDs[0] = 1;
        tokenIDs[1] = 2;
        uint256[] memory amounts = new uint256[](2);
        amounts[0] = 10 ether;
        amounts[1] = 20 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        bytes memory data = "Minting ERC1155 token";

        bytes32 structHash = keccak256(
            abi.encode(ERC1155_MINT_BATCH_TYPEHASH, uuid, ACCOUNT.addr, token, tokenIDs, amounts, nonce, deadline)
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC1155ForgeFacet(FORGE).mintERC1155Batch(
            uuid, token, tokenIDs, amounts, data, deadline, abi.encodePacked(r, s, v)
        );

        assertEq(amounts[0], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[0]), "Minted tokenID mismatch");
        assertEq(amounts[1], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[1]), "Minted tokenID mismatch");
    }

    function test_mintto_erc1155() external {
        // set erc1155Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = eRC1155ForgeFacet.mintERC1155To.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC1155ForgeFacet),
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
        uint256 uuid = 0;
        uint256 tokenID = 1;
        uint256 amount = 10 ether;
        bytes memory data = "Minting ERC1155 token";
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        bytes memory recipientSig;

        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC1155_MINTTO_TYPEHASH, token, tokenID, amount, nonce, deadline));
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_MINTTO_TYPEHASH, uuid, ACCOUNT.addr, recipientSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC1155ForgeFacet(FORGE).mintERC1155To(
            uuid, ACCOUNT.addr, token, tokenID, amount, data, deadline, recipientSig, validatorSig
        );
        assertEq(amount, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenID), "Minted tokenID mismatch");
    }

    function test_mintto_batch_erc1155() external {
        // set erc1155Mint
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = eRC1155ForgeFacet.mintERC1155ToBatch.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC1155ForgeFacet),
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
        uint256 uuid = 0;
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
                keccak256(abi.encode(ERC1155_MINTTO_BATCH_TYPEHASH, token, tokenIDs, amounts, nonce, deadline));
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_MINTTO_BATCH_TYPEHASH, uuid, ACCOUNT.addr, recipientSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC1155ForgeFacet(FORGE).mintERC1155ToBatch(
            uuid, ACCOUNT.addr, token, tokenIDs, amounts, data, deadline, recipientSig, validatorSig
        );
        assertEq(amounts[0], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[0]), "Minted tokenID mismatch");
        assertEq(amounts[1], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[1]), "Minted tokenID mismatch");
    }

    function test_transfer_erc20() external {
        // set erc20Transfer
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = eRC20ForgeFacet.transferERC20.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC20ForgeFacet),
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
        uint256 uuid = 0;
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        // charge token to forge
        MockERC20(token).forceMint(FORGE, amount);

        bytes32 structHash =
            keccak256(abi.encode(ERC20_TRANSFER_TYPEHASH, uuid, ACCOUNT.addr, token, amount, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC20ForgeFacet(FORGE).transferERC20(uuid, token, amount, deadline, abi.encodePacked(r, s, v));
        assertEq(amount, MockERC20(token).balanceOf(ACCOUNT.addr), "Transfered amount mismatch");
    }

    function test_transferto_erc20() external {
        // set erc20Transfer
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = eRC20ForgeFacet.transferERC20To.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC20ForgeFacet),
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
        uint256 uuid = 0;
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        MockERC20(token).forceMint(FORGE, amount);

        bytes memory recipientSig;

        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC20_TRANSFERTO_TYPEHASH, token, amount, nonce, deadline));
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC20_VALIDATOR_TRANSFERTO_TYPEHASH, uuid, ACCOUNT.addr, recipientSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC20ForgeFacet(FORGE).transferERC20To(uuid, ACCOUNT.addr, token, amount, deadline, recipientSig, validatorSig);
        assertEq(amount, MockERC20(token).balanceOf(ACCOUNT.addr), "Transfered amount mismatch");
    }

    function test_transfer_erc721() external {
        // set erc721Transfer
        bytes4[] memory functionSelectors = new bytes4[](2);
        functionSelectors[0] = eRC721ForgeFacet.transferERC721.selector;
        functionSelectors[1] = eRC721ForgeFacet.onERC721Received.selector;

        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC721ForgeFacet),
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
        uint256 uuid = 0;
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        bytes memory data = "Transfer ERC721 token";

        // charge token to forge
        MockERC721(token).forceMint(FORGE, tokenID);

        bytes32 structHash =
            keccak256(abi.encode(ERC721_TRANSFER_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC721ForgeFacet(FORGE).transferERC721(uuid, token, tokenID, data, deadline, abi.encodePacked(r, s, v));
        assertEq(ACCOUNT.addr, MockERC721(token).ownerOf(tokenID), "Transfered tokenID mismatch");
    }

    function test_transferto_erc721() external {
        // set erc721Transfer
        bytes4[] memory functionSelectors = new bytes4[](2);
        functionSelectors[0] = eRC721ForgeFacet.transferERC721To.selector;
        functionSelectors[1] = eRC721ForgeFacet.onERC721Received.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC721ForgeFacet),
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
        uint256 uuid = 0;
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        bytes memory data = "Transfer ERC721 token";

        // charge token to forge
        MockERC721(token).forceMint(FORGE, tokenID);

        bytes memory recipientSig;

        {
            bytes32 recipientStructHash =
                keccak256(abi.encode(ERC721_TRANSFERTO_TYPEHASH, token, tokenID, nonce, deadline));
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC721_VALIDATOR_TRANSFERTO_TYPEHASH, uuid, ACCOUNT.addr, recipientSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC721ForgeFacet(FORGE).transferERC721To(
            uuid, ACCOUNT.addr, token, tokenID, data, deadline, recipientSig, validatorSig
        );
        assertEq(ACCOUNT.addr, MockERC721(token).ownerOf(tokenID), "Transfered tokenID mismatch");
    }

    function test_transfer_erc1155() external {
        // set erc1155Transfer
        bytes4[] memory functionSelectors = new bytes4[](3);
        functionSelectors[0] = eRC1155ForgeFacet.transferERC1155.selector;
        functionSelectors[1] = eRC1155ForgeFacet.onERC1155Received.selector;
        functionSelectors[2] = eRC1155ForgeFacet.onERC1155BatchReceived.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC1155ForgeFacet),
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
        uint256 uuid = 0;
        uint256 tokenID = 1;
        uint256 amount = 10 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        bytes memory data = "Transfer ERC1155 token";

        // charge token to forge
        MockERC1155(token).forceMint(FORGE, tokenID, amount);

        bytes32 structHash = keccak256(
            abi.encode(ERC1155_TRANSFER_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC1155ForgeFacet(FORGE).transferERC1155(
            uuid, token, tokenID, amount, data, deadline, abi.encodePacked(r, s, v)
        );
        assertEq(amount, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenID), "Transfered tokenID mismatch");
    }

    function test_transfer_batch_erc1155() external {
        // set erc1155Transfer
        bytes4[] memory functionSelectors = new bytes4[](3);
        functionSelectors[0] = eRC1155ForgeFacet.transferERC1155Batch.selector;
        functionSelectors[1] = eRC1155ForgeFacet.onERC1155Received.selector;
        functionSelectors[2] = eRC1155ForgeFacet.onERC1155BatchReceived.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC1155ForgeFacet),
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
        uint256 uuid = 0;
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
            abi.encode(ERC1155_TRANSFER_BATCH_TYPEHASH, uuid, ACCOUNT.addr, token, tokenIDs, amounts, nonce, deadline)
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC1155ForgeFacet(FORGE).transferERC1155Batch(
            uuid, token, tokenIDs, amounts, data, deadline, abi.encodePacked(r, s, v)
        );

        assertEq(amounts[0], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[0]), "Transfered tokenID mismatch");
        assertEq(amounts[1], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[1]), "Transfered tokenID mismatch");
    }

    function test_transferto_erc1155() external {
        // set erc1155Transfer
        bytes4[] memory functionSelectors = new bytes4[](3);
        functionSelectors[0] = eRC1155ForgeFacet.transferERC1155To.selector;
        functionSelectors[1] = eRC1155ForgeFacet.onERC1155Received.selector;
        functionSelectors[2] = eRC1155ForgeFacet.onERC1155BatchReceived.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC1155ForgeFacet),
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
        uint256 uuid = 0;
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
                keccak256(abi.encode(ERC1155_TRANSFERTO_TYPEHASH, token, tokenID, amount, nonce, deadline));
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_TRANSFERTO_TYPEHASH, uuid, ACCOUNT.addr, recipientSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC1155ForgeFacet(FORGE).transferERC1155To(
            uuid, ACCOUNT.addr, token, tokenID, amount, data, deadline, recipientSig, validatorSig
        );
        assertEq(amount, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenID), "Transfered tokenID mismatch");
    }

    function test_transferto_batch_erc1155() external {
        // set erc1155Transfer
        bytes4[] memory functionSelectors = new bytes4[](3);
        functionSelectors[0] = eRC1155ForgeFacet.transferERC1155ToBatch.selector;
        functionSelectors[1] = eRC1155ForgeFacet.onERC1155Received.selector;
        functionSelectors[2] = eRC1155ForgeFacet.onERC1155BatchReceived.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC1155ForgeFacet),
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
        uint256 uuid = 0;
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
                keccak256(abi.encode(ERC1155_TRANSFERTO_BATCH_TYPEHASH, token, tokenIDs, amounts, nonce, deadline));
            bytes32 recipientHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, recipientStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, recipientHash);
            recipientSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_TRANSFERTO_BATCH_TYPEHASH, uuid, ACCOUNT.addr, recipientSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC1155ForgeFacet(FORGE).transferERC1155ToBatch(
            uuid, ACCOUNT.addr, token, tokenIDs, amounts, data, deadline, recipientSig, validatorSig
        );
        assertEq(amounts[0], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[0]), "Transfered tokenID mismatch");
        assertEq(amounts[1], MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[1]), "Transfered tokenID mismatch");
    }

    function test_burn_erc20() external {
        // set erc20Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = eRC20ForgeFacet.burnERC20.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC20ForgeFacet),
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
        uint256 uuid = 0;
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        // charge token to account
        MockERC20(token).forceMint(ACCOUNT.addr, amount);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC20(token).approve(FORGE, amount);

        bytes32 structHash =
            keccak256(abi.encode(ERC20_BURN_TYPEHASH, uuid, ACCOUNT.addr, token, amount, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC20ForgeFacet(FORGE).burnERC20(uuid, token, amount, deadline, abi.encodePacked(r, s, v));
        assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "Burned amount mismatch");
    }

    function test_burnfrom_erc20() external {
        // set erc20Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = eRC20ForgeFacet.burnERC20From.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC20ForgeFacet),
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
        uint256 uuid = 0;
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        // charge token to account
        MockERC20(token).forceMint(ACCOUNT.addr, amount);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC20(token).approve(FORGE, amount);

        bytes memory fromSig;
        {
            bytes32 fromStructHash = keccak256(abi.encode(ERC20_BURNFROM_TYPEHASH, token, amount, nonce, deadline));
            bytes32 fromHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, fromStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, fromHash);
            fromSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC20_VALIDATOR_BURNFROM_TYPEHASH, uuid, ACCOUNT.addr, fromSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC20ForgeFacet(FORGE).burnERC20From(uuid, ACCOUNT.addr, token, amount, deadline, fromSig, validatorSig);
        assertEq(0, MockERC20(token).balanceOf(ACCOUNT.addr), "Burned amount mismatch");
    }

    function test_burn_erc721() external {
        // set erc721Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = eRC721ForgeFacet.burnERC721.selector;

        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC721ForgeFacet),
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
        uint256 uuid = 0;
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        // charge token to account
        MockERC721(token).forceMint(ACCOUNT.addr, tokenID);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC721(token).approve(FORGE, tokenID);

        bytes32 structHash =
            keccak256(abi.encode(ERC721_BURN_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC721ForgeFacet(FORGE).burnERC721(uuid, token, tokenID, deadline, abi.encodePacked(r, s, v));
        vm.expectRevert(abi.encodeWithSignature("ERC721NonexistentToken(uint256)", tokenID));
        MockERC721(token).ownerOf(tokenID);
    }

    function test_burnfrom_erc721() external {
        // set erc721Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = eRC721ForgeFacet.burnERC721From.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC721ForgeFacet),
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
        uint256 uuid = 0;
        uint256 tokenID = 1;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        // charge token to account
        MockERC721(token).forceMint(ACCOUNT.addr, tokenID);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC721(token).approve(FORGE, tokenID);

        bytes memory fromSig;

        {
            bytes32 fromStructHash = keccak256(abi.encode(ERC721_BURNFROM_TYPEHASH, token, tokenID, nonce, deadline));
            bytes32 fromHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, fromStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, fromHash);
            fromSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC721_VALIDATOR_BURNFROM_TYPEHASH, uuid, ACCOUNT.addr, fromSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC721ForgeFacet(FORGE).burnERC721From(uuid, ACCOUNT.addr, token, tokenID, deadline, fromSig, validatorSig);
        vm.expectRevert(abi.encodeWithSignature("ERC721NonexistentToken(uint256)", tokenID));
        MockERC721(token).ownerOf(tokenID);
    }

    function test_burn_erc1155() external {
        // set erc1155Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = eRC1155ForgeFacet.burnERC1155.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC1155ForgeFacet),
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
        uint256 uuid = 0;
        uint256 tokenID = 1;
        uint256 amount = 10 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        // charge token to acount
        MockERC1155(token).forceMint(ACCOUNT.addr, tokenID, amount);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC1155(token).setApprovalForAll(FORGE, true);

        bytes32 structHash =
            keccak256(abi.encode(ERC1155_BURN_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, amount, nonce, deadline));
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC1155ForgeFacet(FORGE).burnERC1155(uuid, token, tokenID, amount, deadline, abi.encodePacked(r, s, v));
        assertEq(0, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenID), "Burned tokenID mismatch");
    }

    function test_burn_batch_erc1155() external {
        // set erc1155Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = eRC1155ForgeFacet.burnERC1155Batch.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC1155ForgeFacet),
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
        uint256 uuid = 0;
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
            abi.encode(ERC1155_BURN_BATCH_TYPEHASH, uuid, ACCOUNT.addr, token, tokenIDs, amounts, nonce, deadline)
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC1155ForgeFacet(FORGE).burnERC1155Batch(uuid, token, tokenIDs, amounts, deadline, abi.encodePacked(r, s, v));

        assertEq(0, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[0]), "Burned tokenID mismatch");
        assertEq(0, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[1]), "Burned tokenID mismatch");
    }

    function test_burnfrom_erc1155() external {
        // set erc1155Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = eRC1155ForgeFacet.burnERC1155From.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC1155ForgeFacet),
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
        uint256 uuid = 0;
        uint256 tokenID = 1;
        uint256 amount = 10 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        // charge token to account
        MockERC1155(token).forceMint(ACCOUNT.addr, tokenID, amount);
        // approve to forge
        vm.prank(ACCOUNT.addr);
        MockERC1155(token).setApprovalForAll(FORGE, true);

        bytes memory fromSig;
        {
            bytes32 fromStructHash =
                keccak256(abi.encode(ERC1155_BURNFROM_TYPEHASH, token, tokenID, amount, nonce, deadline));
            bytes32 fromHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, fromStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, fromHash);
            fromSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_BURNFROM_TYPEHASH, uuid, ACCOUNT.addr, fromSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC1155ForgeFacet(FORGE).burnERC1155From(
            uuid, ACCOUNT.addr, token, tokenID, amount, deadline, fromSig, validatorSig
        );
        assertEq(0, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenID), "Burned tokenID mismatch");
    }

    function test_burnfrom_batch_erc1155() external {
        // set erc1155Burn
        bytes4[] memory functionSelectors = new bytes4[](1);
        functionSelectors[0] = eRC1155ForgeFacet.burnERC1155FromBatch.selector;
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](1);
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC1155ForgeFacet),
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
        uint256 uuid = 0;
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

        bytes memory fromSig;

        {
            bytes32 fromStructHash =
                keccak256(abi.encode(ERC1155_BURNFROM_BATCH_TYPEHASH, token, tokenIDs, amounts, nonce, deadline));
            bytes32 fromHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, fromStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(ACCOUNT, fromHash);
            fromSig = abi.encodePacked(r, s, v);
        }
        bytes memory validatorSig;
        {
            bytes32 validatorStructHash =
                keccak256(abi.encode(ERC1155_VALIDATOR_BURNFROM_BATCH_TYPEHASH, uuid, ACCOUNT.addr, fromSig));
            bytes32 validatorHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, validatorStructHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, validatorHash);
            validatorSig = abi.encodePacked(r, s, v);
        }

        // send transaction
        vm.prank(ACCOUNT.addr);
        ERC1155ForgeFacet(FORGE).burnERC1155FromBatch(
            uuid, ACCOUNT.addr, token, tokenIDs, amounts, deadline, fromSig, validatorSig
        );
        assertEq(0, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[0]), "Burned tokenID mismatch");
        assertEq(0, MockERC1155(token).balanceOf(ACCOUNT.addr, tokenIDs[1]), "Burned tokenID mismatch");
    }
}
