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

import "./TestTokenForge.t.sol";

contract TestException is Test {
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
    bytes32 public DOMAIN_SEPARATOR;

    MockERC20 public mockERC20;
    MockERC721 public mockERC721;
    MockERC1155 public mockERC1155;

    function setUp() public {
        VALIDATOR = vm.createWallet("Validator");
        ACCOUNT = vm.createWallet("Account");

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
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](3);
        bytes4[] memory erc20FunctionSelectors = new bytes4[](6);
        erc20FunctionSelectors[0] = eRC20ForgeFacet.mintERC20.selector;
        erc20FunctionSelectors[1] = eRC20ForgeFacet.mintERC20To.selector;
        erc20FunctionSelectors[2] = eRC20ForgeFacet.transferERC20.selector;
        erc20FunctionSelectors[3] = eRC20ForgeFacet.transferERC20To.selector;
        erc20FunctionSelectors[4] = eRC20ForgeFacet.burnERC20.selector;
        erc20FunctionSelectors[5] = eRC20ForgeFacet.burnERC20From.selector;
        addCuts[0] = IDiamondCut.FacetCut({
            facetAddress: address(eRC20ForgeFacet),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: erc20FunctionSelectors
        });
        bytes4[] memory erc721FunctionSelectors = new bytes4[](7);
        erc721FunctionSelectors[0] = eRC721ForgeFacet.mintERC721.selector;
        erc721FunctionSelectors[1] = eRC721ForgeFacet.mintERC721To.selector;
        erc721FunctionSelectors[2] = eRC721ForgeFacet.transferERC721.selector;
        erc721FunctionSelectors[3] = eRC721ForgeFacet.transferERC721To.selector;
        erc721FunctionSelectors[4] = eRC721ForgeFacet.burnERC721.selector;
        erc721FunctionSelectors[5] = eRC721ForgeFacet.burnERC721From.selector;
        erc721FunctionSelectors[6] = eRC721ForgeFacet.onERC721Received.selector;
        addCuts[1] = IDiamondCut.FacetCut({
            facetAddress: address(eRC721ForgeFacet),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: erc721FunctionSelectors
        });
        bytes4[] memory erc1155FunctionSelectors = new bytes4[](14);
        erc1155FunctionSelectors[0] = eRC1155ForgeFacet.mintERC1155.selector;
        erc1155FunctionSelectors[1] = eRC1155ForgeFacet.mintERC1155To.selector;
        erc1155FunctionSelectors[2] = eRC1155ForgeFacet.transferERC1155.selector;
        erc1155FunctionSelectors[3] = eRC1155ForgeFacet.transferERC1155To.selector;
        erc1155FunctionSelectors[4] = eRC1155ForgeFacet.burnERC1155.selector;
        erc1155FunctionSelectors[5] = eRC1155ForgeFacet.burnERC1155From.selector;
        erc1155FunctionSelectors[6] = eRC1155ForgeFacet.mintERC1155Batch.selector;
        erc1155FunctionSelectors[7] = eRC1155ForgeFacet.mintERC1155ToBatch.selector;
        erc1155FunctionSelectors[8] = eRC1155ForgeFacet.transferERC1155Batch.selector;
        erc1155FunctionSelectors[9] = eRC1155ForgeFacet.transferERC1155ToBatch.selector;
        erc1155FunctionSelectors[10] = eRC1155ForgeFacet.burnERC1155Batch.selector;
        erc1155FunctionSelectors[11] = eRC1155ForgeFacet.burnERC1155FromBatch.selector;
        erc1155FunctionSelectors[12] = eRC1155ForgeFacet.onERC1155BatchReceived.selector;
        erc1155FunctionSelectors[13] = eRC1155ForgeFacet.onERC1155Received.selector;
        addCuts[2] = IDiamondCut.FacetCut({
            facetAddress: address(eRC1155ForgeFacet),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: erc1155FunctionSelectors
        });
        tokenForgeFactory = TokenForgeFactory(address(tokenForgeFactoryProxy));
        FORGE = tokenForgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);

        vm.stopPrank();
        mockERC20 = new MockERC20(FORGE);
        mockERC721 = new MockERC721(FORGE);
        mockERC1155 = new MockERC1155(FORGE);

        DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();
    }

    function test_sender_is_not_msgsender() external {
        uint256 uuid = 1;
        uint256 tokenID = 1;
        uint256[] memory tokenIDs;
        address token;
        uint256 nonce;
        uint256 amount = 100;
        uint256[] memory amounts;
        uint256 deadline = block.timestamp + 1 days;
        bytes memory data = "test_sender_is_not_msgsender";
        bytes32 structHash;
        bytes memory signature;
        // mint
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash =
                    keccak256(abi.encode(ERC20_MINT_TYPEHASH, uuid, ACCOUNT.addr, token, amount, nonce, deadline));
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ERC20ForgeFacet(FORGE).mintERC20(uuid, token, amount, deadline, signature);
                // check success
                vm.prank(ACCOUNT.addr);
                ERC20ForgeFacet(FORGE).mintERC20(uuid, token, amount, deadline, signature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash =
                    keccak256(abi.encode(ERC721_MINT_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, nonce, deadline));
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ERC721ForgeFacet(FORGE).mintERC721(uuid, token, tokenID, deadline, signature);
                // check success
                vm.prank(ACCOUNT.addr);
                ERC721ForgeFacet(FORGE).mintERC721(uuid, token, tokenID, deadline, signature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(
                    abi.encode(ERC1155_MINT_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ERC1155ForgeFacet(FORGE).mintERC1155(uuid, token, tokenID, amount, data, deadline, signature);
                // check success
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).mintERC1155(uuid, token, tokenID, amount, data, deadline, signature);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(
                    abi.encode(
                        ERC1155_MINT_BATCH_TYPEHASH, uuid, ACCOUNT.addr, token, tokenIDs, amounts, nonce, deadline
                    )
                );
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ERC1155ForgeFacet(FORGE).mintERC1155Batch(uuid, token, tokenIDs, amounts, data, deadline, signature);
                // check success
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).mintERC1155Batch(uuid, token, tokenIDs, amounts, data, deadline, signature);
            }
        }
        vm.startPrank(ACCOUNT.addr);
        mockERC20.transfer(FORGE, amount);
        mockERC721.transferFrom(ACCOUNT.addr, FORGE, tokenID);
        mockERC1155.safeTransferFrom(ACCOUNT.addr, FORGE, tokenID, amount, data);
        mockERC1155.safeBatchTransferFrom(ACCOUNT.addr, FORGE, tokenIDs, amounts, data);
        vm.stopPrank();
        // transfer
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash =
                    keccak256(abi.encode(ERC20_TRANSFER_TYPEHASH, uuid, ACCOUNT.addr, token, amount, nonce, deadline));
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ERC20ForgeFacet(FORGE).transferERC20(uuid, token, amount, deadline, signature);
                // check success
                vm.prank(ACCOUNT.addr);
                ERC20ForgeFacet(FORGE).transferERC20(uuid, token, amount, deadline, signature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash =
                    keccak256(abi.encode(ERC721_TRANSFER_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, nonce, deadline));
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ERC721ForgeFacet(FORGE).transferERC721(uuid, token, tokenID, data, deadline, signature);
                // check success
                vm.prank(ACCOUNT.addr);
                ERC721ForgeFacet(FORGE).transferERC721(uuid, token, tokenID, data, deadline, signature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(
                    abi.encode(ERC1155_TRANSFER_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ERC1155ForgeFacet(FORGE).transferERC1155(uuid, token, tokenID, amount, data, deadline, signature);
                // check success
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).transferERC1155(uuid, token, tokenID, amount, data, deadline, signature);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(
                    abi.encode(
                        ERC1155_TRANSFER_BATCH_TYPEHASH, uuid, ACCOUNT.addr, token, tokenIDs, amounts, nonce, deadline
                    )
                );
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ERC1155ForgeFacet(FORGE).transferERC1155Batch(uuid, token, tokenIDs, amounts, data, deadline, signature);
                // check success
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).transferERC1155Batch(uuid, token, tokenIDs, amounts, data, deadline, signature);
            }
        }
        vm.startPrank(ACCOUNT.addr);
        mockERC20.approve(FORGE, amount);
        mockERC721.approve(FORGE, tokenID);
        mockERC1155.setApprovalForAll(FORGE, true);
        vm.stopPrank();
        // burn
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash =
                    keccak256(abi.encode(ERC20_BURN_TYPEHASH, uuid, ACCOUNT.addr, token, amount, nonce, deadline));
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ERC20ForgeFacet(FORGE).burnERC20(uuid, token, amount, deadline, signature);
                // check success
                vm.prank(ACCOUNT.addr);
                ERC20ForgeFacet(FORGE).burnERC20(uuid, token, amount, deadline, signature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash =
                    keccak256(abi.encode(ERC721_BURN_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, nonce, deadline));
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ERC721ForgeFacet(FORGE).burnERC721(uuid, token, tokenID, deadline, signature);
                // check success
                vm.prank(ACCOUNT.addr);
                ERC721ForgeFacet(FORGE).burnERC721(uuid, token, tokenID, deadline, signature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(
                    abi.encode(ERC1155_BURN_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ERC1155ForgeFacet(FORGE).burnERC1155(uuid, token, tokenID, amount, deadline, signature);
                // check success
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).burnERC1155(uuid, token, tokenID, amount, deadline, signature);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(
                    abi.encode(
                        ERC1155_BURN_BATCH_TYPEHASH, uuid, ACCOUNT.addr, token, tokenIDs, amounts, nonce, deadline
                    )
                );
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ERC1155ForgeFacet(FORGE).burnERC1155Batch(uuid, token, tokenIDs, amounts, deadline, signature);
                // check success
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).burnERC1155Batch(uuid, token, tokenIDs, amounts, deadline, signature);
            }
        }
    }

    function test_other_account() external {
        address otherAccount = address(bytes20("OtherAccount"));
        vm.label(otherAccount, "OtherAccount");

        uint256 uuid = 1;
        uint256 tokenID = 1;
        uint256[] memory tokenIDs;
        address token;
        uint256 nonce;
        uint256 amount = 100;
        uint256[] memory amounts;
        uint256 deadline = block.timestamp + 1 days;
        bytes memory data = "test_sender_is_not_msgsender";
        bytes32 structHash;
        bytes memory accountSignature;
        bytes memory validatorSignature;
        // mintTo
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(abi.encode(ERC20_MINTTO_TYPEHASH, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_MINTTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(
                    abi.encodeWithSignature("ERC20MintForge__InvalidRecipientSignature(address)", otherAccount)
                );
                ERC20ForgeFacet(FORGE).mintERC20To(
                    uuid, otherAccount, token, amount, deadline, accountSignature, validatorSignature
                );
                // check success
                ERC20ForgeFacet(FORGE).mintERC20To(
                    uuid, ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature
                );
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(abi.encode(ERC721_MINTTO_TYPEHASH, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC721_VALIDATOR_MINTTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(
                    abi.encodeWithSignature("ERC721MintForge__InvalidRecipientSignature(address)", otherAccount)
                );
                ERC721ForgeFacet(FORGE).mintERC721To(
                    uuid, otherAccount, token, tokenID, deadline, accountSignature, validatorSignature
                );
                // check success
                ERC721ForgeFacet(FORGE).mintERC721To(
                    uuid, ACCOUNT.addr, token, tokenID, deadline, accountSignature, validatorSignature
                );
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    structHash = keccak256(abi.encode(ERC1155_MINTTO_TYPEHASH, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash =
                        keccak256(abi.encode(ERC1155_VALIDATOR_MINTTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature("ERC1155MintForge__InvalidRecipientSignature(address)", otherAccount)
                    );
                    ERC1155ForgeFacet(FORGE).mintERC1155To(
                        uuid, otherAccount, token, tokenID, amount, data, deadline, accountSignature, validatorSignature
                    );
                    // check success
                    ERC1155ForgeFacet(FORGE).mintERC1155To(
                        uuid, ACCOUNT.addr, token, tokenID, amount, data, deadline, accountSignature, validatorSignature
                    );

                    // mint batch
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    structHash =
                        keccak256(abi.encode(ERC1155_MINTTO_BATCH_TYPEHASH, token, tokenIDs, amounts, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_MINTTO_BATCH_TYPEHASH, uuid, ACCOUNT.addr, accountSignature)
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature("ERC1155MintForge__InvalidRecipientSignature(address)", otherAccount)
                    );
                    ERC1155ForgeFacet(FORGE).mintERC1155ToBatch(
                        uuid,
                        otherAccount,
                        token,
                        tokenIDs,
                        amounts,
                        data,
                        deadline,
                        accountSignature,
                        validatorSignature
                    );
                    // check success
                    ERC1155ForgeFacet(FORGE).mintERC1155ToBatch(
                        uuid,
                        ACCOUNT.addr,
                        token,
                        tokenIDs,
                        amounts,
                        data,
                        deadline,
                        accountSignature,
                        validatorSignature
                    );
                }
            }
        }
        vm.startPrank(ACCOUNT.addr);
        mockERC20.transfer(FORGE, amount);
        mockERC721.transferFrom(ACCOUNT.addr, FORGE, tokenID);
        mockERC1155.safeTransferFrom(ACCOUNT.addr, FORGE, tokenID, amount, data);
        mockERC1155.safeBatchTransferFrom(ACCOUNT.addr, FORGE, tokenIDs, amounts, data);
        vm.stopPrank();
        // transferTo
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(abi.encode(ERC20_TRANSFERTO_TYPEHASH, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_TRANSFERTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(
                    abi.encodeWithSignature("ERC20TransferForge__InvalidRecipientSignature(address)", otherAccount)
                );
                ERC20ForgeFacet(FORGE).transferERC20To(
                    uuid, otherAccount, token, amount, deadline, accountSignature, validatorSignature
                );
                // check success
                ERC20ForgeFacet(FORGE).transferERC20To(
                    uuid, ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature
                );
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(abi.encode(ERC721_TRANSFERTO_TYPEHASH, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC721_VALIDATOR_TRANSFERTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(
                    abi.encodeWithSignature("ERC721TransferForge__InvalidRecipientSignature(address)", otherAccount)
                );
                ERC721ForgeFacet(FORGE).transferERC721To(
                    uuid, otherAccount, token, tokenID, data, deadline, accountSignature, validatorSignature
                );
                // check success
                ERC721ForgeFacet(FORGE).transferERC721To(
                    uuid, ACCOUNT.addr, token, tokenID, data, deadline, accountSignature, validatorSignature
                );
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    structHash =
                        keccak256(abi.encode(ERC1155_TRANSFERTO_TYPEHASH, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_TRANSFERTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature)
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature(
                            "ERC1155TransferForge__InvalidRecipientSignature(address)", otherAccount
                        )
                    );
                    ERC1155ForgeFacet(FORGE).transferERC1155To(
                        uuid, otherAccount, token, tokenID, amount, data, deadline, accountSignature, validatorSignature
                    );
                    // check success
                    ERC1155ForgeFacet(FORGE).transferERC1155To(
                        uuid, ACCOUNT.addr, token, tokenID, amount, data, deadline, accountSignature, validatorSignature
                    );

                    // transfer batch
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    structHash = keccak256(
                        abi.encode(ERC1155_TRANSFERTO_BATCH_TYPEHASH, token, tokenIDs, amounts, nonce, deadline)
                    );
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_TRANSFERTO_BATCH_TYPEHASH, uuid, ACCOUNT.addr, accountSignature)
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature(
                            "ERC1155TransferForge__InvalidRecipientSignature(address)", otherAccount
                        )
                    );
                    ERC1155ForgeFacet(FORGE).transferERC1155ToBatch(
                        uuid,
                        otherAccount,
                        token,
                        tokenIDs,
                        amounts,
                        data,
                        deadline,
                        accountSignature,
                        validatorSignature
                    );
                    // check success
                    ERC1155ForgeFacet(FORGE).transferERC1155ToBatch(
                        uuid,
                        ACCOUNT.addr,
                        token,
                        tokenIDs,
                        amounts,
                        data,
                        deadline,
                        accountSignature,
                        validatorSignature
                    );
                }
            }
        }
        vm.startPrank(ACCOUNT.addr);
        mockERC20.approve(FORGE, amount);
        mockERC721.approve(FORGE, tokenID);
        mockERC1155.setApprovalForAll(FORGE, true);
        vm.stopPrank();
        // burnFrom
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(abi.encode(ERC20_BURNFROM_TYPEHASH, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_BURNFROM_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("ERC20BurnForge__InvalidFromSignature(address)", otherAccount));
                ERC20ForgeFacet(FORGE).burnERC20From(
                    uuid, otherAccount, token, amount, deadline, accountSignature, validatorSignature
                );
                // check success
                ERC20ForgeFacet(FORGE).burnERC20From(
                    uuid, ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature
                );
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(abi.encode(ERC721_BURNFROM_TYPEHASH, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC721_VALIDATOR_BURNFROM_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("ERC721BurnForge__InvalidFromSignature(address)", otherAccount));
                ERC721ForgeFacet(FORGE).burnERC721From(
                    uuid, otherAccount, token, tokenID, deadline, accountSignature, validatorSignature
                );
                // check success
                ERC721ForgeFacet(FORGE).burnERC721From(
                    uuid, ACCOUNT.addr, token, tokenID, deadline, accountSignature, validatorSignature
                );
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    structHash =
                        keccak256(abi.encode(ERC1155_BURNFROM_TYPEHASH, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash =
                        keccak256(abi.encode(ERC1155_VALIDATOR_BURNFROM_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature("ERC1155BurnForge__InvalidFromSignature(address)", otherAccount)
                    );
                    ERC1155ForgeFacet(FORGE).burnERC1155From(
                        uuid, otherAccount, token, tokenID, amount, deadline, accountSignature, validatorSignature
                    );
                    // check success
                    ERC1155ForgeFacet(FORGE).burnERC1155From(
                        uuid, ACCOUNT.addr, token, tokenID, amount, deadline, accountSignature, validatorSignature
                    );

                    // burn batch
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    structHash = keccak256(
                        abi.encode(ERC1155_BURNFROM_BATCH_TYPEHASH, token, tokenIDs, amounts, nonce, deadline)
                    );
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_BURNFROM_BATCH_TYPEHASH, uuid, ACCOUNT.addr, accountSignature)
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature("ERC1155BurnForge__InvalidFromSignature(address)", otherAccount)
                    );
                    ERC1155ForgeFacet(FORGE).burnERC1155FromBatch(
                        uuid, otherAccount, token, tokenIDs, amounts, deadline, accountSignature, validatorSignature
                    );
                    // check success
                    ERC1155ForgeFacet(FORGE).burnERC1155FromBatch(
                        uuid, ACCOUNT.addr, token, tokenIDs, amounts, deadline, accountSignature, validatorSignature
                    );
                }
            }
        }
    }

    function test_deadline_expired() external {
        uint256 uuid = 1;
        uint256 tokenID = 1;
        uint256[] memory tokenIDs;
        address token;
        uint256 nonce;
        uint256 amount = 100;
        uint256[] memory amounts;
        uint256 deadline = block.timestamp + 1 days; // expired deadline
        bytes memory data = "test_deadline_expired";
        bytes32 structHash;
        bytes memory accountSignature;
        bytes memory validatorSignature;

        vm.warp(deadline + 1); // move time forward to make deadline expired
        // mint
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash =
                    keccak256(abi.encode(ERC20_MINT_TYPEHASH, uuid, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ERC20ForgeFacet(FORGE).mintERC20(uuid, token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash =
                    keccak256(abi.encode(ERC721_MINT_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ERC721ForgeFacet(FORGE).mintERC721(uuid, token, tokenID, deadline, validatorSignature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(
                    abi.encode(ERC1155_MINT_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).mintERC1155(uuid, token, tokenID, amount, data, deadline, validatorSignature);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(
                    abi.encode(
                        ERC1155_MINT_BATCH_TYPEHASH, uuid, ACCOUNT.addr, token, tokenIDs, amounts, nonce, deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).mintERC1155Batch(
                    uuid, token, tokenIDs, amounts, data, deadline, validatorSignature
                );
            }
        }
        // transfer
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash =
                    keccak256(abi.encode(ERC20_TRANSFER_TYPEHASH, uuid, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ERC20ForgeFacet(FORGE).transferERC20(uuid, token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash =
                    keccak256(abi.encode(ERC721_TRANSFER_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ERC721ForgeFacet(FORGE).transferERC721(uuid, token, tokenID, data, deadline, validatorSignature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(
                    abi.encode(ERC1155_TRANSFER_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).transferERC1155(
                    uuid, token, tokenID, amount, data, deadline, validatorSignature
                );

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(
                    abi.encode(
                        ERC1155_TRANSFER_BATCH_TYPEHASH, uuid, ACCOUNT.addr, token, tokenIDs, amounts, nonce, deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).transferERC1155Batch(
                    uuid, token, tokenIDs, amounts, data, deadline, validatorSignature
                );
            }
        }
        // burn
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash =
                    keccak256(abi.encode(ERC20_BURN_TYPEHASH, uuid, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ERC20ForgeFacet(FORGE).burnERC20(uuid, token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash =
                    keccak256(abi.encode(ERC721_BURN_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ERC721ForgeFacet(FORGE).burnERC721(uuid, token, tokenID, deadline, validatorSignature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(
                    abi.encode(ERC1155_BURN_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).burnERC1155(uuid, token, tokenID, amount, deadline, validatorSignature);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(
                    abi.encode(
                        ERC1155_BURN_BATCH_TYPEHASH, uuid, ACCOUNT.addr, token, tokenIDs, amounts, nonce, deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).burnERC1155Batch(uuid, token, tokenIDs, amounts, deadline, validatorSignature);
            }
        }

        // mintTo
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(abi.encode(ERC20_MINTTO_TYPEHASH, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_MINTTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ERC20ForgeFacet(FORGE).mintERC20To(
                    uuid, ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature
                );
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(abi.encode(ERC721_MINTTO_TYPEHASH, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC721_VALIDATOR_MINTTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ERC721ForgeFacet(FORGE).mintERC721To(
                    uuid, ACCOUNT.addr, token, tokenID, deadline, accountSignature, validatorSignature
                );
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    structHash = keccak256(abi.encode(ERC1155_MINTTO_TYPEHASH, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash =
                        keccak256(abi.encode(ERC1155_VALIDATOR_MINTTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                    ERC1155ForgeFacet(FORGE).mintERC1155To(
                        uuid, ACCOUNT.addr, token, tokenID, amount, data, deadline, accountSignature, validatorSignature
                    );

                    // mint batch
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    structHash =
                        keccak256(abi.encode(ERC1155_MINTTO_BATCH_TYPEHASH, token, tokenIDs, amounts, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_MINTTO_BATCH_TYPEHASH, uuid, ACCOUNT.addr, accountSignature)
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                    ERC1155ForgeFacet(FORGE).mintERC1155ToBatch(
                        uuid,
                        ACCOUNT.addr,
                        token,
                        tokenIDs,
                        amounts,
                        data,
                        deadline,
                        accountSignature,
                        validatorSignature
                    );
                }
            }
        }
        // transferTo
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(abi.encode(ERC20_TRANSFERTO_TYPEHASH, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_TRANSFERTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ERC20ForgeFacet(FORGE).transferERC20To(
                    uuid, ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature
                );
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(abi.encode(ERC721_TRANSFERTO_TYPEHASH, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC721_VALIDATOR_TRANSFERTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ERC721ForgeFacet(FORGE).transferERC721To(
                    uuid, ACCOUNT.addr, token, tokenID, data, deadline, accountSignature, validatorSignature
                );
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    structHash =
                        keccak256(abi.encode(ERC1155_TRANSFERTO_TYPEHASH, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_TRANSFERTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature)
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                    ERC1155ForgeFacet(FORGE).transferERC1155To(
                        uuid, ACCOUNT.addr, token, tokenID, amount, data, deadline, accountSignature, validatorSignature
                    );

                    // transfer batch
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    structHash = keccak256(
                        abi.encode(ERC1155_TRANSFERTO_BATCH_TYPEHASH, token, tokenIDs, amounts, nonce, deadline)
                    );
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_TRANSFERTO_BATCH_TYPEHASH, uuid, ACCOUNT.addr, accountSignature)
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                    ERC1155ForgeFacet(FORGE).transferERC1155ToBatch(
                        uuid,
                        ACCOUNT.addr,
                        token,
                        tokenIDs,
                        amounts,
                        data,
                        deadline,
                        accountSignature,
                        validatorSignature
                    );
                }
            }
        }
        // burnFrom
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(abi.encode(ERC20_BURNFROM_TYPEHASH, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_BURNFROM_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ERC20ForgeFacet(FORGE).burnERC20From(
                    uuid, ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature
                );
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(abi.encode(ERC721_BURNFROM_TYPEHASH, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC721_VALIDATOR_BURNFROM_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ERC721ForgeFacet(FORGE).burnERC721From(
                    uuid, ACCOUNT.addr, token, tokenID, deadline, accountSignature, validatorSignature
                );
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    structHash =
                        keccak256(abi.encode(ERC1155_BURNFROM_TYPEHASH, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash =
                        keccak256(abi.encode(ERC1155_VALIDATOR_BURNFROM_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                    ERC1155ForgeFacet(FORGE).burnERC1155From(
                        uuid, ACCOUNT.addr, token, tokenID, amount, deadline, accountSignature, validatorSignature
                    );

                    // burn batch
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    structHash = keccak256(
                        abi.encode(ERC1155_BURNFROM_BATCH_TYPEHASH, token, tokenIDs, amounts, nonce, deadline)
                    );
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_BURNFROM_BATCH_TYPEHASH, uuid, ACCOUNT.addr, accountSignature)
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                    ERC1155ForgeFacet(FORGE).burnERC1155FromBatch(
                        uuid, ACCOUNT.addr, token, tokenIDs, amounts, deadline, accountSignature, validatorSignature
                    );
                }
            }
        }
    }

    function test_invalid_validator() external {
        uint256 uuid = 1;
        uint256 tokenID = 1;
        uint256[] memory tokenIDs;
        address token;
        uint256 nonce;
        uint256 amount = 100;
        uint256[] memory amounts;
        uint256 deadline = block.timestamp + 1 days;
        bytes memory data = "test_invalid_validator";
        bytes32 structHash;
        bytes memory accountSignature;
        bytes memory validatorSignature;

        VALIDATOR = vm.createWallet("OtherValidator");

        // mint
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash =
                    keccak256(abi.encode(ERC20_MINT_TYPEHASH, uuid, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC20ForgeFacet(FORGE).mintERC20(uuid, token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash =
                    keccak256(abi.encode(ERC721_MINT_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC721ForgeFacet(FORGE).mintERC721(uuid, token, tokenID, deadline, validatorSignature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(
                    abi.encode(ERC1155_MINT_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).mintERC1155(uuid, token, tokenID, amount, data, deadline, validatorSignature);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(
                    abi.encode(
                        ERC1155_MINT_BATCH_TYPEHASH, uuid, ACCOUNT.addr, token, tokenIDs, amounts, nonce, deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).mintERC1155Batch(
                    uuid, token, tokenIDs, amounts, data, deadline, validatorSignature
                );
            }
        }
        // transfer
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash =
                    keccak256(abi.encode(ERC20_TRANSFER_TYPEHASH, uuid, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC20ForgeFacet(FORGE).transferERC20(uuid, token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash =
                    keccak256(abi.encode(ERC721_TRANSFER_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC721ForgeFacet(FORGE).transferERC721(uuid, token, tokenID, data, deadline, validatorSignature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(
                    abi.encode(ERC1155_TRANSFER_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).transferERC1155(
                    uuid, token, tokenID, amount, data, deadline, validatorSignature
                );

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(
                    abi.encode(
                        ERC1155_TRANSFER_BATCH_TYPEHASH, uuid, ACCOUNT.addr, token, tokenIDs, amounts, nonce, deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).transferERC1155Batch(
                    uuid, token, tokenIDs, amounts, data, deadline, validatorSignature
                );
            }
        }
        // burn
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash =
                    keccak256(abi.encode(ERC20_BURN_TYPEHASH, uuid, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC20ForgeFacet(FORGE).burnERC20(uuid, token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash =
                    keccak256(abi.encode(ERC721_BURN_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC721ForgeFacet(FORGE).burnERC721(uuid, token, tokenID, deadline, validatorSignature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(
                    abi.encode(ERC1155_BURN_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).burnERC1155(uuid, token, tokenID, amount, deadline, validatorSignature);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(
                    abi.encode(
                        ERC1155_BURN_BATCH_TYPEHASH, uuid, ACCOUNT.addr, token, tokenIDs, amounts, nonce, deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).burnERC1155Batch(uuid, token, tokenIDs, amounts, deadline, validatorSignature);
            }
        }

        // mintTo
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(abi.encode(ERC20_MINTTO_TYPEHASH, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_MINTTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ERC20ForgeFacet(FORGE).mintERC20To(
                    uuid, ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature
                );
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(abi.encode(ERC721_MINTTO_TYPEHASH, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC721_VALIDATOR_MINTTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ERC721ForgeFacet(FORGE).mintERC721To(
                    uuid, ACCOUNT.addr, token, tokenID, deadline, accountSignature, validatorSignature
                );
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    structHash = keccak256(abi.encode(ERC1155_MINTTO_TYPEHASH, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash =
                        keccak256(abi.encode(ERC1155_VALIDATOR_MINTTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                    ERC1155ForgeFacet(FORGE).mintERC1155To(
                        uuid, ACCOUNT.addr, token, tokenID, amount, data, deadline, accountSignature, validatorSignature
                    );

                    // mint batch
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    structHash =
                        keccak256(abi.encode(ERC1155_MINTTO_BATCH_TYPEHASH, token, tokenIDs, amounts, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_MINTTO_BATCH_TYPEHASH, uuid, ACCOUNT.addr, accountSignature)
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                    ERC1155ForgeFacet(FORGE).mintERC1155ToBatch(
                        uuid,
                        ACCOUNT.addr,
                        token,
                        tokenIDs,
                        amounts,
                        data,
                        deadline,
                        accountSignature,
                        validatorSignature
                    );
                }
            }
        }
        // transferTo
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(abi.encode(ERC20_TRANSFERTO_TYPEHASH, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_TRANSFERTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ERC20ForgeFacet(FORGE).transferERC20To(
                    uuid, ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature
                );
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(abi.encode(ERC721_TRANSFERTO_TYPEHASH, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC721_VALIDATOR_TRANSFERTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ERC721ForgeFacet(FORGE).transferERC721To(
                    uuid, ACCOUNT.addr, token, tokenID, data, deadline, accountSignature, validatorSignature
                );
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    structHash =
                        keccak256(abi.encode(ERC1155_TRANSFERTO_TYPEHASH, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_TRANSFERTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature)
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                    ERC1155ForgeFacet(FORGE).transferERC1155To(
                        uuid, ACCOUNT.addr, token, tokenID, amount, data, deadline, accountSignature, validatorSignature
                    );

                    // transfer batch
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    structHash = keccak256(
                        abi.encode(ERC1155_TRANSFERTO_BATCH_TYPEHASH, token, tokenIDs, amounts, nonce, deadline)
                    );
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_TRANSFERTO_BATCH_TYPEHASH, uuid, ACCOUNT.addr, accountSignature)
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                    ERC1155ForgeFacet(FORGE).transferERC1155ToBatch(
                        uuid,
                        ACCOUNT.addr,
                        token,
                        tokenIDs,
                        amounts,
                        data,
                        deadline,
                        accountSignature,
                        validatorSignature
                    );
                }
            }
        }
        // burnFrom
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(abi.encode(ERC20_BURNFROM_TYPEHASH, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_BURNFROM_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ERC20ForgeFacet(FORGE).burnERC20From(
                    uuid, ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature
                );
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                structHash = keccak256(abi.encode(ERC721_BURNFROM_TYPEHASH, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC721_VALIDATOR_BURNFROM_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ERC721ForgeFacet(FORGE).burnERC721From(
                    uuid, ACCOUNT.addr, token, tokenID, deadline, accountSignature, validatorSignature
                );
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    structHash =
                        keccak256(abi.encode(ERC1155_BURNFROM_TYPEHASH, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash =
                        keccak256(abi.encode(ERC1155_VALIDATOR_BURNFROM_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                    ERC1155ForgeFacet(FORGE).burnERC1155From(
                        uuid, ACCOUNT.addr, token, tokenID, amount, deadline, accountSignature, validatorSignature
                    );

                    // burn batch
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    structHash = keccak256(
                        abi.encode(ERC1155_BURNFROM_BATCH_TYPEHASH, token, tokenIDs, amounts, nonce, deadline)
                    );
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_BURNFROM_BATCH_TYPEHASH, uuid, ACCOUNT.addr, accountSignature)
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                    ERC1155ForgeFacet(FORGE).burnERC1155FromBatch(
                        uuid, ACCOUNT.addr, token, tokenIDs, amounts, deadline, accountSignature, validatorSignature
                    );
                }
            }
        }
    }

    function test_dup_nonce() external {
        uint256 uuid = 1;
        uint256 tokenID = 1;
        uint256[] memory tokenIDs;
        address token;
        uint256 nonce = 0;
        uint256 amount = 100;
        uint256[] memory amounts;
        uint256 deadline = block.timestamp + 1 days;
        bytes memory data = "test_dup_nonce";
        bytes32 structHash;
        bytes memory accountSignature;
        bytes memory validatorSignature;

        // spend nonce
        {
            token = address(mockERC20);
            structHash = keccak256(abi.encode(ERC20_MINT_TYPEHASH, uuid, ACCOUNT.addr, token, amount, nonce, deadline));
            validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
            vm.prank(ACCOUNT.addr);
            ERC20ForgeFacet(FORGE).mintERC20(uuid, token, amount, deadline, validatorSignature);
        }

        // mint
        {
            // erc20
            {
                token = address(mockERC20);
                structHash =
                    keccak256(abi.encode(ERC20_MINT_TYPEHASH, uuid, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC20ForgeFacet(FORGE).mintERC20(uuid, token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                structHash =
                    keccak256(abi.encode(ERC721_MINT_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC721ForgeFacet(FORGE).mintERC721(uuid, token, tokenID, deadline, validatorSignature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                structHash = keccak256(
                    abi.encode(ERC1155_MINT_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).mintERC1155(uuid, token, tokenID, amount, data, deadline, validatorSignature);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                structHash = keccak256(
                    abi.encode(
                        ERC1155_MINT_BATCH_TYPEHASH, uuid, ACCOUNT.addr, token, tokenIDs, amounts, nonce, deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).mintERC1155Batch(
                    uuid, token, tokenIDs, amounts, data, deadline, validatorSignature
                );
            }
        }
        // transfer
        {
            // erc20
            {
                token = address(mockERC20);
                structHash =
                    keccak256(abi.encode(ERC20_TRANSFER_TYPEHASH, uuid, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC20ForgeFacet(FORGE).transferERC20(uuid, token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                structHash =
                    keccak256(abi.encode(ERC721_TRANSFER_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC721ForgeFacet(FORGE).transferERC721(uuid, token, tokenID, data, deadline, validatorSignature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                structHash = keccak256(
                    abi.encode(ERC1155_TRANSFER_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).transferERC1155(
                    uuid, token, tokenID, amount, data, deadline, validatorSignature
                );

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                structHash = keccak256(
                    abi.encode(
                        ERC1155_TRANSFER_BATCH_TYPEHASH, uuid, ACCOUNT.addr, token, tokenIDs, amounts, nonce, deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).transferERC1155Batch(
                    uuid, token, tokenIDs, amounts, data, deadline, validatorSignature
                );
            }
        }
        // burn
        {
            // erc20
            {
                token = address(mockERC20);
                structHash =
                    keccak256(abi.encode(ERC20_BURN_TYPEHASH, uuid, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC20ForgeFacet(FORGE).burnERC20(uuid, token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                structHash =
                    keccak256(abi.encode(ERC721_BURN_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC721ForgeFacet(FORGE).burnERC721(uuid, token, tokenID, deadline, validatorSignature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                structHash = keccak256(
                    abi.encode(ERC1155_BURN_TYPEHASH, uuid, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).burnERC1155(uuid, token, tokenID, amount, deadline, validatorSignature);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                structHash = keccak256(
                    abi.encode(
                        ERC1155_BURN_BATCH_TYPEHASH, uuid, ACCOUNT.addr, token, tokenIDs, amounts, nonce, deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ERC1155ForgeFacet(FORGE).burnERC1155Batch(uuid, token, tokenIDs, amounts, deadline, validatorSignature);
            }
        }

        // mintTo
        {
            // erc20
            {
                token = address(mockERC20);
                structHash = keccak256(abi.encode(ERC20_MINTTO_TYPEHASH, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_MINTTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(
                    abi.encodeWithSignature("ERC20MintForge__InvalidRecipientSignature(address)", ACCOUNT.addr)
                );
                ERC20ForgeFacet(FORGE).mintERC20To(
                    uuid, ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature
                );
            }
            // erc721
            {
                token = address(mockERC721);
                structHash = keccak256(abi.encode(ERC721_MINTTO_TYPEHASH, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC721_VALIDATOR_MINTTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(
                    abi.encodeWithSignature("ERC721MintForge__InvalidRecipientSignature(address)", ACCOUNT.addr)
                );
                ERC721ForgeFacet(FORGE).mintERC721To(
                    uuid, ACCOUNT.addr, token, tokenID, deadline, accountSignature, validatorSignature
                );
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    structHash = keccak256(abi.encode(ERC1155_MINTTO_TYPEHASH, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash =
                        keccak256(abi.encode(ERC1155_VALIDATOR_MINTTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature("ERC1155MintForge__InvalidRecipientSignature(address)", ACCOUNT.addr)
                    );
                    ERC1155ForgeFacet(FORGE).mintERC1155To(
                        uuid, ACCOUNT.addr, token, tokenID, amount, data, deadline, accountSignature, validatorSignature
                    );

                    // mint batch
                    structHash =
                        keccak256(abi.encode(ERC1155_MINTTO_BATCH_TYPEHASH, token, tokenIDs, amounts, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_MINTTO_BATCH_TYPEHASH, uuid, ACCOUNT.addr, accountSignature)
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature("ERC1155MintForge__InvalidRecipientSignature(address)", ACCOUNT.addr)
                    );
                    ERC1155ForgeFacet(FORGE).mintERC1155ToBatch(
                        uuid,
                        ACCOUNT.addr,
                        token,
                        tokenIDs,
                        amounts,
                        data,
                        deadline,
                        accountSignature,
                        validatorSignature
                    );
                }
            }
        }
        // transferTo
        {
            // erc20
            {
                token = address(mockERC20);
                structHash = keccak256(abi.encode(ERC20_TRANSFERTO_TYPEHASH, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_TRANSFERTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(
                    abi.encodeWithSignature("ERC20TransferForge__InvalidRecipientSignature(address)", ACCOUNT.addr)
                );
                ERC20ForgeFacet(FORGE).transferERC20To(
                    uuid, ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature
                );
            }
            // erc721
            {
                token = address(mockERC721);
                structHash = keccak256(abi.encode(ERC721_TRANSFERTO_TYPEHASH, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC721_VALIDATOR_TRANSFERTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(
                    abi.encodeWithSignature("ERC721TransferForge__InvalidRecipientSignature(address)", ACCOUNT.addr)
                );
                ERC721ForgeFacet(FORGE).transferERC721To(
                    uuid, ACCOUNT.addr, token, tokenID, data, deadline, accountSignature, validatorSignature
                );
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    structHash =
                        keccak256(abi.encode(ERC1155_TRANSFERTO_TYPEHASH, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_TRANSFERTO_TYPEHASH, uuid, ACCOUNT.addr, accountSignature)
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature(
                            "ERC1155TransferForge__InvalidRecipientSignature(address)", ACCOUNT.addr
                        )
                    );
                    ERC1155ForgeFacet(FORGE).transferERC1155To(
                        uuid, ACCOUNT.addr, token, tokenID, amount, data, deadline, accountSignature, validatorSignature
                    );

                    // transfer batch
                    structHash = keccak256(
                        abi.encode(ERC1155_TRANSFERTO_BATCH_TYPEHASH, token, tokenIDs, amounts, nonce, deadline)
                    );
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_TRANSFERTO_BATCH_TYPEHASH, uuid, ACCOUNT.addr, accountSignature)
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature(
                            "ERC1155TransferForge__InvalidRecipientSignature(address)", ACCOUNT.addr
                        )
                    );
                    ERC1155ForgeFacet(FORGE).transferERC1155ToBatch(
                        uuid,
                        ACCOUNT.addr,
                        token,
                        tokenIDs,
                        amounts,
                        data,
                        deadline,
                        accountSignature,
                        validatorSignature
                    );
                }
            }
        }
        // burnFrom
        {
            // erc20
            {
                token = address(mockERC20);
                structHash = keccak256(abi.encode(ERC20_BURNFROM_TYPEHASH, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_BURNFROM_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("ERC20BurnForge__InvalidFromSignature(address)", ACCOUNT.addr));
                ERC20ForgeFacet(FORGE).burnERC20From(
                    uuid, ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature
                );
            }
            // erc721
            {
                token = address(mockERC721);
                structHash = keccak256(abi.encode(ERC721_BURNFROM_TYPEHASH, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC721_VALIDATOR_BURNFROM_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("ERC721BurnForge__InvalidFromSignature(address)", ACCOUNT.addr));
                ERC721ForgeFacet(FORGE).burnERC721From(
                    uuid, ACCOUNT.addr, token, tokenID, deadline, accountSignature, validatorSignature
                );
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    structHash =
                        keccak256(abi.encode(ERC1155_BURNFROM_TYPEHASH, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash =
                        keccak256(abi.encode(ERC1155_VALIDATOR_BURNFROM_TYPEHASH, uuid, ACCOUNT.addr, accountSignature));
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature("ERC1155BurnForge__InvalidFromSignature(address)", ACCOUNT.addr)
                    );
                    ERC1155ForgeFacet(FORGE).burnERC1155From(
                        uuid, ACCOUNT.addr, token, tokenID, amount, deadline, accountSignature, validatorSignature
                    );

                    // burn batch
                    structHash = keccak256(
                        abi.encode(ERC1155_BURNFROM_BATCH_TYPEHASH, token, tokenIDs, amounts, nonce, deadline)
                    );
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_BURNFROM_BATCH_TYPEHASH, uuid, ACCOUNT.addr, accountSignature)
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature("ERC1155BurnForge__InvalidFromSignature(address)", ACCOUNT.addr)
                    );
                    ERC1155ForgeFacet(FORGE).burnERC1155FromBatch(
                        uuid, ACCOUNT.addr, token, tokenIDs, amounts, deadline, accountSignature, validatorSignature
                    );
                }
            }
        }
    }

    function _sign_with_domain_separator(bytes32 hash, Vm.Wallet memory wallet) internal returns (bytes memory) {
        bytes32 messageHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, hash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(wallet, messageHash);
        return abi.encodePacked(r, s, v);
    }
}
