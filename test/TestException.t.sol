// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std-1.9.7/Test.sol";
import {Vm} from "forge-std-1.9.7/Vm.sol";

import {IDiamondCut} from "diamond-3-hardhat-1.0.0/interfaces/IDiamondCut.sol";
import {IDiamondLoupe} from "diamond-3-hardhat-1.0.0/interfaces/IDiamondLoupe.sol";
import {IERC173} from "diamond-3-hardhat-1.0.0/interfaces/IERC173.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts-5.3.0/proxy/ERC1967/ERC1967Proxy.sol";
import {MessageHashUtils} from "@openzeppelin-contracts-5.3.0/utils/cryptography/MessageHashUtils.sol";

import {ForgeProxyCode} from "../src/forges/ForgeProxy.sol";
import {Diamond3Facet} from "../src/forges/Diamond3Facet.sol";
import {ForgeFactory} from "../src/forges/ForgeFactory.sol";
import "../src/forges/BaseForge.sol";
import "../src/forges/ForgeV1.sol";
import "../src/forges/ForgeV2.sol";
import "../src/forges/ForgeV3.sol";

import {TokenFactory} from "../src/tokens/TokenFactory.sol";
import {MockERC20} from "./mock/MockERC20.sol";
import {MockERC721} from "./mock/MockERC721.sol";
import {MockERC1155} from "./mock/MockERC1155.sol";

import "./mock/StructHash.sol";

contract TestException is Test {
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
    ForgeFactory public forgeFactory;
    address public FORGE;
    bytes32 public DOMAIN_SEPARATOR;

    TokenFactory public tokenFactory;
    address public mockERC20Impl;
    address public mockERC721Impl;
    address public mockERC1155Impl;
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
        IDiamondCut.FacetCut[] memory addCuts = new IDiamondCut.FacetCut[](3);
        {
            bytes4[] memory v1Selectors = new bytes4[](15);
            v1Selectors[0] = forgeV1.mintERC20.selector;
            v1Selectors[1] = forgeV1.mintERC721.selector;
            v1Selectors[2] = forgeV1.mintERC1155.selector;
            v1Selectors[3] = forgeV1.mintERC1155Batch.selector;

            v1Selectors[4] = forgeV1.transferERC20.selector;
            v1Selectors[5] = forgeV1.transferERC721.selector;
            v1Selectors[6] = forgeV1.transferERC1155.selector;
            v1Selectors[7] = forgeV1.transferERC1155Batch.selector;

            v1Selectors[8] = forgeV1.burnERC20.selector;
            v1Selectors[9] = forgeV1.burnERC721.selector;
            v1Selectors[10] = forgeV1.burnERC1155.selector;
            v1Selectors[11] = forgeV1.burnERC1155Batch.selector;

            v1Selectors[12] = forgeV1.onERC721Received.selector;
            v1Selectors[13] = forgeV1.onERC1155Received.selector;
            v1Selectors[14] = forgeV1.onERC1155BatchReceived.selector;

            addCuts[0] = IDiamondCut.FacetCut({
                facetAddress: address(forgeV1),
                action: IDiamondCut.FacetCutAction.Add,
                functionSelectors: v1Selectors
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
                facetAddress: address(forgeV2),
                action: IDiamondCut.FacetCutAction.Add,
                functionSelectors: v2Selectors
            });
        }
        {
            bytes4[] memory v3Selectors = new bytes4[](12);
            v3Selectors[0] = forgeV3.mintERC20.selector;
            v3Selectors[1] = forgeV3.mintERC721.selector;
            v3Selectors[2] = forgeV3.mintERC1155.selector;
            v3Selectors[3] = forgeV3.mintERC1155Batch.selector;

            v3Selectors[4] = forgeV3.transferERC20.selector;
            v3Selectors[5] = forgeV3.transferERC721.selector;
            v3Selectors[6] = forgeV3.transferERC1155.selector;
            v3Selectors[7] = forgeV3.transferERC1155Batch.selector;

            v3Selectors[8] = forgeV3.burnERC20.selector;
            v3Selectors[9] = forgeV3.burnERC721.selector;
            v3Selectors[10] = forgeV3.burnERC1155.selector;
            v3Selectors[11] = forgeV3.burnERC1155Batch.selector;

            addCuts[2] = IDiamondCut.FacetCut({
                facetAddress: address(forgeV3),
                action: IDiamondCut.FacetCutAction.Add,
                functionSelectors: v3Selectors
            });
        }

        forgeFactory = ForgeFactory(address(forgeFactoryProxy));
        FORGE = forgeFactory.addService(SERVICE_OWNER, VALIDATOR.addr, SERVICE_NAME, addCuts);
        DOMAIN_SEPARATOR = BaseForge(FORGE).DOMAIN_SEPARATOR();

        {
            mockERC20Impl = address(new MockERC20());
            mockERC721Impl = address(new MockERC721());
            mockERC1155Impl = address(new MockERC1155());
            address[] memory erc20Impls = new address[](1);
            erc20Impls[0] = mockERC20Impl;
            address[] memory erc721Impls = new address[](1);
            erc721Impls[0] = mockERC721Impl;
            address[] memory erc1155Impls = new address[](1);
            erc1155Impls[0] = mockERC1155Impl;
            address tokenFactoryImpl = address(new TokenFactory());
            address tokenFactoryProxy = address(
                new ERC1967Proxy(
                    tokenFactoryImpl,
                    abi.encodeCall(
                        TokenFactory.initialize, (OWNER, address(forgeFactory), erc20Impls, erc721Impls, erc1155Impls)
                    )
                )
            );
            tokenFactory = TokenFactory(tokenFactoryProxy);
        }
        mockERC20 = MockERC20(tokenFactory.deployERC20(OWNER, SERVICE_NAME, "MockERC20", "M20", 18, 0, mockERC20Impl));
        mockERC721 = MockERC721(
            tokenFactory.deployERC721(OWNER, SERVICE_NAME, "MockERC721", "M721", "https://xxx.yyy.zzz/", mockERC721Impl)
        );
        mockERC1155 =
            MockERC1155(tokenFactory.deployERC1155(OWNER, SERVICE_NAME, "https://xxx.yyy.zzz/", mockERC1155Impl));

        vm.stopPrank();
    }

    function test_v1_sender_is_not_msgsender() external {
        uint256 uuid;
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
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC20_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, amount, nonce, deadline));
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV1(FORGE).mintERC20(token, amount, deadline, signature);
                // check success
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).mintERC20(token, amount, deadline, signature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC721_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, nonce, deadline));
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV1(FORGE).mintERC721(token, tokenID, deadline, signature);
                // check success
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).mintERC721(token, tokenID, deadline, signature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(ERC1155_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV1(FORGE).mintERC1155(token, tokenID, amount, deadline, signature, data);
                // check success
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).mintERC1155(token, tokenID, amount, deadline, signature, data);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(
                        ERC1155_MINT_BATCH_TYPE_HASH_V1,
                        ACCOUNT.addr,
                        token,
                        keccak256(abi.encodePacked(tokenIDs)),
                        keccak256(abi.encodePacked(amounts)),
                        nonce,
                        deadline
                    )
                );
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV1(FORGE).mintERC1155Batch(token, tokenIDs, amounts, deadline, signature, data);
                // check success
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).mintERC1155Batch(token, tokenIDs, amounts, deadline, signature, data);
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
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC20_TRANSFER_TYPE_HASH_V1, ACCOUNT.addr, token, amount, nonce, deadline));
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV1(FORGE).transferERC20(token, amount, deadline, signature);
                // check success
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).transferERC20(token, amount, deadline, signature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC721_TRANSFER_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, nonce, deadline));
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV1(FORGE).transferERC721(token, tokenID, deadline, signature, data);
                // check success
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).transferERC721(token, tokenID, deadline, signature, data);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(ERC1155_TRANSFER_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV1(FORGE).transferERC1155(token, tokenID, amount, deadline, signature, data);
                // check success
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).transferERC1155(token, tokenID, amount, deadline, signature, data);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(
                        ERC1155_TRANSFER_BATCH_TYPE_HASH_V1,
                        ACCOUNT.addr,
                        token,
                        keccak256(abi.encodePacked(tokenIDs)),
                        keccak256(abi.encodePacked(amounts)),
                        nonce,
                        deadline
                    )
                );
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV1(FORGE).transferERC1155Batch(token, tokenIDs, amounts, deadline, signature, data);
                // check success
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).transferERC1155Batch(token, tokenIDs, amounts, deadline, signature, data);
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
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC20_BURN_TYPE_HASH_V1, ACCOUNT.addr, token, amount, nonce, deadline));
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV1(FORGE).burnERC20(token, amount, deadline, signature);
                // check success
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).burnERC20(token, amount, deadline, signature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC721_BURN_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, nonce, deadline));
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV1(FORGE).burnERC721(token, tokenID, deadline, signature);
                // check success
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).burnERC721(token, tokenID, deadline, signature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(ERC1155_BURN_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV1(FORGE).burnERC1155(token, tokenID, amount, deadline, signature);
                // check success
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).burnERC1155(token, tokenID, amount, deadline, signature);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
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
                signature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV1(FORGE).burnERC1155Batch(token, tokenIDs, amounts, deadline, signature);
                // check success
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).burnERC1155Batch(token, tokenIDs, amounts, deadline, signature);
            }
        }
    }

    function test_v1_other_account() external {
        address otherAccount = address(bytes20("OtherAccount"));
        vm.label(otherAccount, "OtherAccount");

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
        uint256 uuid;
        // mintTo
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(abi.encode(ERC20_MINT_TYPE_HASH_V2, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_MINT_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature)));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("ERC20ForgeV2__InvalidAccountSignature(address)", otherAccount));
                ForgeV2(FORGE).mintERC20(otherAccount, token, amount, deadline, accountSignature, validatorSignature);
                // check success
                ForgeV2(FORGE).mintERC20(ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(abi.encode(ERC721_MINT_TYPE_HASH_V2, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC721_VALIDATOR_MINT_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature)));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(
                    abi.encodeWithSignature("ERC721ForgeV2__InvalidAccountSignature(address)", otherAccount)
                );
                ForgeV2(FORGE).mintERC721(otherAccount, token, tokenID, deadline, accountSignature, validatorSignature);
                // check success
                ForgeV2(FORGE).mintERC721(ACCOUNT.addr, token, tokenID, deadline, accountSignature, validatorSignature);
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    uuid = _calcUUID(nonce);
                    structHash =
                        keccak256(abi.encode(ERC1155_MINT_TYPE_HASH_V2, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_MINT_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature("ERC1155ForgeV2__InvalidAccountSignature(address)", otherAccount)
                    );
                    ForgeV2(FORGE).mintERC1155(
                        otherAccount, token, tokenID, amount, deadline, accountSignature, validatorSignature, data
                    );
                    // check success
                    ForgeV2(FORGE).mintERC1155(
                        ACCOUNT.addr, token, tokenID, amount, deadline, accountSignature, validatorSignature, data
                    );

                    // mint batch
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    uuid = _calcUUID(nonce);
                    structHash = keccak256(
                        abi.encode(
                            ERC1155_MINT_BATCH_TYPE_HASH_V2,
                            token,
                            keccak256(abi.encodePacked(tokenIDs)),
                            keccak256(abi.encodePacked(amounts)),
                            nonce,
                            deadline
                        )
                    );
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_MINT_BATCH_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature("ERC1155ForgeV2__InvalidAccountSignature(address)", otherAccount)
                    );
                    ForgeV2(FORGE).mintERC1155Batch(
                        otherAccount, token, tokenIDs, amounts, deadline, accountSignature, validatorSignature, data
                    );
                    // check success
                    ForgeV2(FORGE).mintERC1155Batch(
                        ACCOUNT.addr, token, tokenIDs, amounts, deadline, accountSignature, validatorSignature, data
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
                uuid = _calcUUID(nonce);
                structHash = keccak256(abi.encode(ERC20_TRANSFER_TYPE_HASH_V2, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash = keccak256(
                    abi.encode(ERC20_VALIDATOR_TRANSFER_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("ERC20ForgeV2__InvalidAccountSignature(address)", otherAccount));
                ForgeV2(FORGE).transferERC20(
                    otherAccount, token, amount, deadline, accountSignature, validatorSignature
                );
                // check success
                ForgeV2(FORGE).transferERC20(
                    ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature
                );
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(abi.encode(ERC721_TRANSFER_TYPE_HASH_V2, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash = keccak256(
                    abi.encode(ERC721_VALIDATOR_TRANSFER_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(
                    abi.encodeWithSignature("ERC721ForgeV2__InvalidAccountSignature(address)", otherAccount)
                );
                ForgeV2(FORGE).transferERC721(
                    otherAccount, token, tokenID, deadline, accountSignature, validatorSignature, data
                );
                // check success
                ForgeV2(FORGE).transferERC721(
                    ACCOUNT.addr, token, tokenID, deadline, accountSignature, validatorSignature, data
                );
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    uuid = _calcUUID(nonce);
                    structHash =
                        keccak256(abi.encode(ERC1155_TRANSFER_TYPE_HASH_V2, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_TRANSFER_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature("ERC1155ForgeV2__InvalidAccountSignature(address)", otherAccount)
                    );
                    ForgeV2(FORGE).transferERC1155(
                        otherAccount, token, tokenID, amount, deadline, accountSignature, validatorSignature, data
                    );
                    // check success
                    ForgeV2(FORGE).transferERC1155(
                        ACCOUNT.addr, token, tokenID, amount, deadline, accountSignature, validatorSignature, data
                    );

                    // transfer batch
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    uuid = _calcUUID(nonce);
                    structHash = keccak256(
                        abi.encode(
                            ERC1155_TRANSFER_BATCH_TYPE_HASH_V2,
                            token,
                            keccak256(abi.encodePacked(tokenIDs)),
                            keccak256(abi.encodePacked(amounts)),
                            nonce,
                            deadline
                        )
                    );
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(
                            ERC1155_VALIDATOR_TRANSFER_BATCH_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature)
                        )
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature("ERC1155ForgeV2__InvalidAccountSignature(address)", otherAccount)
                    );
                    ForgeV2(FORGE).transferERC1155Batch(
                        otherAccount, token, tokenIDs, amounts, deadline, accountSignature, validatorSignature, data
                    );
                    // check success
                    ForgeV2(FORGE).transferERC1155Batch(
                        ACCOUNT.addr, token, tokenIDs, amounts, deadline, accountSignature, validatorSignature, data
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
                uuid = _calcUUID(nonce);
                structHash = keccak256(abi.encode(ERC20_BURN_TYPE_HASH_V2, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_BURN_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature)));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("ERC20ForgeV2__InvalidAccountSignature(address)", otherAccount));
                ForgeV2(FORGE).burnERC20(otherAccount, token, amount, deadline, accountSignature, validatorSignature);
                // check success
                ForgeV2(FORGE).burnERC20(ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(abi.encode(ERC721_BURN_TYPE_HASH_V2, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC721_VALIDATOR_BURN_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature)));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(
                    abi.encodeWithSignature("ERC721ForgeV2__InvalidAccountSignature(address)", otherAccount)
                );
                ForgeV2(FORGE).burnERC721(otherAccount, token, tokenID, deadline, accountSignature, validatorSignature);
                // check success
                ForgeV2(FORGE).burnERC721(ACCOUNT.addr, token, tokenID, deadline, accountSignature, validatorSignature);
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    uuid = _calcUUID(nonce);
                    structHash =
                        keccak256(abi.encode(ERC1155_BURN_TYPE_HASH_V2, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_BURN_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature("ERC1155ForgeV2__InvalidAccountSignature(address)", otherAccount)
                    );
                    ForgeV2(FORGE).burnERC1155(
                        otherAccount, token, tokenID, amount, deadline, accountSignature, validatorSignature
                    );
                    // check success
                    ForgeV2(FORGE).burnERC1155(
                        ACCOUNT.addr, token, tokenID, amount, deadline, accountSignature, validatorSignature
                    );

                    // burn batch
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    uuid = _calcUUID(nonce);
                    structHash = keccak256(
                        abi.encode(
                            ERC1155_BURN_BATCH_TYPE_HASH_V2,
                            token,
                            keccak256(abi.encodePacked(tokenIDs)),
                            keccak256(abi.encodePacked(amounts)),
                            nonce,
                            deadline
                        )
                    );
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_BURN_BATCH_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature("ERC1155ForgeV2__InvalidAccountSignature(address)", otherAccount)
                    );
                    ForgeV2(FORGE).burnERC1155Batch(
                        otherAccount, token, tokenIDs, amounts, deadline, accountSignature, validatorSignature
                    );
                    // check success
                    ForgeV2(FORGE).burnERC1155Batch(
                        ACCOUNT.addr, token, tokenIDs, amounts, deadline, accountSignature, validatorSignature
                    );
                }
            }
        }
    }

    function test_deadline_expired() external {
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
        uint256 uuid;

        vm.warp(deadline + 1); // move time forward to make deadline expired
        // v1 mint
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC20_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).mintERC20(token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC721_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).mintERC721(token, tokenID, deadline, validatorSignature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(ERC1155_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).mintERC1155(token, tokenID, amount, deadline, validatorSignature, data);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(
                        ERC1155_MINT_BATCH_TYPE_HASH_V1,
                        ACCOUNT.addr,
                        token,
                        keccak256(abi.encodePacked(tokenIDs)),
                        keccak256(abi.encodePacked(amounts)),
                        nonce,
                        deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).mintERC1155Batch(token, tokenIDs, amounts, deadline, validatorSignature, data);
            }
        }
        // v1 transfer
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC20_TRANSFER_TYPE_HASH_V1, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).transferERC20(token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC721_TRANSFER_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).transferERC721(token, tokenID, deadline, validatorSignature, data);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(ERC1155_TRANSFER_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).transferERC1155(token, tokenID, amount, deadline, validatorSignature, data);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(
                        ERC1155_TRANSFER_BATCH_TYPE_HASH_V1,
                        ACCOUNT.addr,
                        token,
                        keccak256(abi.encodePacked(tokenIDs)),
                        keccak256(abi.encodePacked(amounts)),
                        nonce,
                        deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).transferERC1155Batch(token, tokenIDs, amounts, deadline, validatorSignature, data);
            }
        }
        // v1 burn
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC20_BURN_TYPE_HASH_V1, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).burnERC20(token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC721_BURN_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).burnERC721(token, tokenID, deadline, validatorSignature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(ERC1155_BURN_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).burnERC1155(token, tokenID, amount, deadline, validatorSignature);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
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
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).burnERC1155Batch(token, tokenIDs, amounts, deadline, validatorSignature);
            }
        }

        // v2 mint
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(abi.encode(ERC20_MINT_TYPE_HASH_V2, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_MINT_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature)));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ForgeV2(FORGE).mintERC20(ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(abi.encode(ERC721_MINT_TYPE_HASH_V2, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC721_VALIDATOR_MINT_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature)));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ForgeV2(FORGE).mintERC721(ACCOUNT.addr, token, tokenID, deadline, accountSignature, validatorSignature);
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    uuid = _calcUUID(nonce);
                    uuid = _calcUUID(nonce);
                    structHash =
                        keccak256(abi.encode(ERC1155_MINT_TYPE_HASH_V2, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_MINT_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                    ForgeV2(FORGE).mintERC1155(
                        ACCOUNT.addr, token, tokenID, amount, deadline, accountSignature, validatorSignature, data
                    );

                    // mint batch
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    uuid = _calcUUID(nonce);
                    uuid = _calcUUID(nonce);
                    structHash = keccak256(
                        abi.encode(
                            ERC1155_MINT_BATCH_TYPE_HASH_V2,
                            token,
                            keccak256(abi.encodePacked(tokenIDs)),
                            keccak256(abi.encodePacked(amounts)),
                            nonce,
                            deadline
                        )
                    );
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_MINT_BATCH_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                    ForgeV2(FORGE).mintERC1155Batch(
                        ACCOUNT.addr, token, tokenIDs, amounts, deadline, accountSignature, validatorSignature, data
                    );
                }
            }
        }
        // v2 transfer
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(abi.encode(ERC20_TRANSFER_TYPE_HASH_V2, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash = keccak256(
                    abi.encode(ERC20_VALIDATOR_TRANSFER_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ForgeV2(FORGE).transferERC20(
                    ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature
                );
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(abi.encode(ERC721_TRANSFER_TYPE_HASH_V2, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash = keccak256(
                    abi.encode(ERC721_VALIDATOR_TRANSFER_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ForgeV2(FORGE).transferERC721(
                    ACCOUNT.addr, token, tokenID, deadline, accountSignature, validatorSignature, data
                );
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    uuid = _calcUUID(nonce);
                    uuid = _calcUUID(nonce);
                    structHash =
                        keccak256(abi.encode(ERC1155_TRANSFER_TYPE_HASH_V2, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_TRANSFER_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                    ForgeV2(FORGE).transferERC1155(
                        ACCOUNT.addr, token, tokenID, amount, deadline, accountSignature, validatorSignature, data
                    );

                    // transfer batch
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    uuid = _calcUUID(nonce);
                    uuid = _calcUUID(nonce);
                    structHash = keccak256(
                        abi.encode(
                            ERC1155_TRANSFER_BATCH_TYPE_HASH_V2,
                            token,
                            keccak256(abi.encodePacked(tokenIDs)),
                            keccak256(abi.encodePacked(amounts)),
                            nonce,
                            deadline
                        )
                    );
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(
                            ERC1155_VALIDATOR_TRANSFER_BATCH_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature)
                        )
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                    ForgeV2(FORGE).transferERC1155Batch(
                        ACCOUNT.addr, token, tokenIDs, amounts, deadline, accountSignature, validatorSignature, data
                    );
                }
            }
        }
        // v2 burn
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(abi.encode(ERC20_BURN_TYPE_HASH_V2, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_BURN_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature)));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ForgeV2(FORGE).burnERC20(ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(abi.encode(ERC721_BURN_TYPE_HASH_V2, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC721_VALIDATOR_BURN_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature)));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ForgeV2(FORGE).burnERC721(ACCOUNT.addr, token, tokenID, deadline, accountSignature, validatorSignature);
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    uuid = _calcUUID(nonce);
                    uuid = _calcUUID(nonce);
                    structHash =
                        keccak256(abi.encode(ERC1155_BURN_TYPE_HASH_V2, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_BURN_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                    ForgeV2(FORGE).burnERC1155(
                        ACCOUNT.addr, token, tokenID, amount, deadline, accountSignature, validatorSignature
                    );

                    // burn batch
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    uuid = _calcUUID(nonce);
                    uuid = _calcUUID(nonce);
                    structHash = keccak256(
                        abi.encode(
                            ERC1155_BURN_BATCH_TYPE_HASH_V2,
                            token,
                            keccak256(abi.encodePacked(tokenIDs)),
                            keccak256(abi.encodePacked(amounts)),
                            nonce,
                            deadline
                        )
                    );
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_BURN_BATCH_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                    ForgeV2(FORGE).burnERC1155Batch(
                        ACCOUNT.addr, token, tokenIDs, amounts, deadline, accountSignature, validatorSignature
                    );
                }
            }
        }

        // v3 mint
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC20_MINT_TYPE_HASH_V3, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ForgeV3(FORGE).mintERC20(ACCOUNT.addr, token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC721_MINT_TYPE_HASH_V3, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ForgeV3(FORGE).mintERC721(ACCOUNT.addr, token, tokenID, deadline, validatorSignature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(ERC1155_MINT_TYPE_HASH_V3, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ForgeV3(FORGE).mintERC1155(ACCOUNT.addr, token, tokenID, amount, deadline, validatorSignature, data);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(
                        ERC1155_MINT_BATCH_TYPE_HASH_V3,
                        ACCOUNT.addr,
                        token,
                        keccak256(abi.encodePacked(tokenIDs)),
                        keccak256(abi.encodePacked(amounts)),
                        nonce,
                        deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ForgeV3(FORGE).mintERC1155Batch(
                    ACCOUNT.addr, token, tokenIDs, amounts, deadline, validatorSignature, data
                );
            }
        }
        // v3 transfer
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC20_TRANSFER_TYPE_HASH_V3, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ForgeV3(FORGE).transferERC20(ACCOUNT.addr, token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC721_TRANSFER_TYPE_HASH_V3, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ForgeV3(FORGE).transferERC721(ACCOUNT.addr, token, tokenID, deadline, validatorSignature, data);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(ERC1155_TRANSFER_TYPE_HASH_V3, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ForgeV3(FORGE).transferERC1155(ACCOUNT.addr, token, tokenID, amount, deadline, validatorSignature, data);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(
                        ERC1155_TRANSFER_BATCH_TYPE_HASH_V3,
                        ACCOUNT.addr,
                        token,
                        keccak256(abi.encodePacked(tokenIDs)),
                        keccak256(abi.encodePacked(amounts)),
                        nonce,
                        deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ForgeV3(FORGE).transferERC1155Batch(
                    ACCOUNT.addr, token, tokenIDs, amounts, deadline, validatorSignature, data
                );
            }
        }
        // v3 burn
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC20_BURN_TYPE_HASH_V3, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ForgeV3(FORGE).burnERC20(ACCOUNT.addr, token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC721_BURN_TYPE_HASH_V3, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ForgeV3(FORGE).burnERC721(ACCOUNT.addr, token, tokenID, deadline, validatorSignature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(ERC1155_BURN_TYPE_HASH_V3, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ForgeV3(FORGE).burnERC1155(ACCOUNT.addr, token, tokenID, amount, deadline, validatorSignature);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(
                        ERC1155_BURN_BATCH_TYPE_HASH_V3,
                        ACCOUNT.addr,
                        token,
                        keccak256(abi.encodePacked(tokenIDs)),
                        keccak256(abi.encodePacked(amounts)),
                        nonce,
                        deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ExpiredSignature(uint256)", deadline));
                ForgeV3(FORGE).burnERC1155Batch(ACCOUNT.addr, token, tokenIDs, amounts, deadline, validatorSignature);
            }
        }
    }

    function test_invalid_validator() external {
        uint256 uuid;
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

        // v1 mint
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC20_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).mintERC20(token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC721_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).mintERC721(token, tokenID, deadline, validatorSignature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(ERC1155_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).mintERC1155(token, tokenID, amount, deadline, validatorSignature, data);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(
                        ERC1155_MINT_BATCH_TYPE_HASH_V1,
                        ACCOUNT.addr,
                        token,
                        keccak256(abi.encodePacked(tokenIDs)),
                        keccak256(abi.encodePacked(amounts)),
                        nonce,
                        deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).mintERC1155Batch(token, tokenIDs, amounts, deadline, validatorSignature, data);
            }
        }
        // v1 transfer
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC20_TRANSFER_TYPE_HASH_V1, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).transferERC20(token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC721_TRANSFER_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).transferERC721(token, tokenID, deadline, validatorSignature, data);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(ERC1155_TRANSFER_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).transferERC1155(token, tokenID, amount, deadline, validatorSignature, data);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(
                        ERC1155_TRANSFER_BATCH_TYPE_HASH_V1,
                        ACCOUNT.addr,
                        token,
                        keccak256(abi.encodePacked(tokenIDs)),
                        keccak256(abi.encodePacked(amounts)),
                        nonce,
                        deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).transferERC1155Batch(token, tokenIDs, amounts, deadline, validatorSignature, data);
            }
        }
        // v1 burn
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC20_BURN_TYPE_HASH_V1, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).burnERC20(token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC721_BURN_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).burnERC721(token, tokenID, deadline, validatorSignature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(ERC1155_BURN_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).burnERC1155(token, tokenID, amount, deadline, validatorSignature);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
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
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).burnERC1155Batch(token, tokenIDs, amounts, deadline, validatorSignature);
            }
        }

        // v2 mint
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(abi.encode(ERC20_MINT_TYPE_HASH_V2, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_MINT_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature)));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV2(FORGE).mintERC20(ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(abi.encode(ERC721_MINT_TYPE_HASH_V2, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC721_VALIDATOR_MINT_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature)));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV2(FORGE).mintERC721(ACCOUNT.addr, token, tokenID, deadline, accountSignature, validatorSignature);
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    uuid = _calcUUID(nonce);
                    structHash =
                        keccak256(abi.encode(ERC1155_MINT_TYPE_HASH_V2, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_MINT_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                    ForgeV2(FORGE).mintERC1155(
                        ACCOUNT.addr, token, tokenID, amount, deadline, accountSignature, validatorSignature, data
                    );

                    // mint batch
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    uuid = _calcUUID(nonce);
                    structHash = keccak256(
                        abi.encode(
                            ERC1155_MINT_BATCH_TYPE_HASH_V2,
                            token,
                            keccak256(abi.encodePacked(tokenIDs)),
                            keccak256(abi.encodePacked(amounts)),
                            nonce,
                            deadline
                        )
                    );
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_MINT_BATCH_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                    ForgeV2(FORGE).mintERC1155Batch(
                        ACCOUNT.addr, token, tokenIDs, amounts, deadline, accountSignature, validatorSignature, data
                    );
                }
            }
        }
        // v2 transfer
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(abi.encode(ERC20_TRANSFER_TYPE_HASH_V2, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash = keccak256(
                    abi.encode(ERC20_VALIDATOR_TRANSFER_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV2(FORGE).transferERC20(
                    ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature
                );
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(abi.encode(ERC721_TRANSFER_TYPE_HASH_V2, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash = keccak256(
                    abi.encode(ERC721_VALIDATOR_TRANSFER_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV2(FORGE).transferERC721(
                    ACCOUNT.addr, token, tokenID, deadline, accountSignature, validatorSignature, data
                );
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    uuid = _calcUUID(nonce);
                    structHash =
                        keccak256(abi.encode(ERC1155_TRANSFER_TYPE_HASH_V2, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_TRANSFER_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                    ForgeV2(FORGE).transferERC1155(
                        ACCOUNT.addr, token, tokenID, amount, deadline, accountSignature, validatorSignature, data
                    );

                    // transfer batch
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    uuid = _calcUUID(nonce);
                    structHash = keccak256(
                        abi.encode(
                            ERC1155_TRANSFER_BATCH_TYPE_HASH_V2,
                            token,
                            keccak256(abi.encodePacked(tokenIDs)),
                            keccak256(abi.encodePacked(amounts)),
                            nonce,
                            deadline
                        )
                    );
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(
                            ERC1155_VALIDATOR_TRANSFER_BATCH_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature)
                        )
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                    ForgeV2(FORGE).transferERC1155Batch(
                        ACCOUNT.addr, token, tokenIDs, amounts, deadline, accountSignature, validatorSignature, data
                    );
                }
            }
        }
        // v2 burn
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(abi.encode(ERC20_BURN_TYPE_HASH_V2, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_BURN_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature)));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV2(FORGE).burnERC20(ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(abi.encode(ERC721_BURN_TYPE_HASH_V2, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC721_VALIDATOR_BURN_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature)));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV2(FORGE).burnERC721(ACCOUNT.addr, token, tokenID, deadline, accountSignature, validatorSignature);
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    uuid = _calcUUID(nonce);
                    structHash =
                        keccak256(abi.encode(ERC1155_BURN_TYPE_HASH_V2, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_BURN_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                    ForgeV2(FORGE).burnERC1155(
                        ACCOUNT.addr, token, tokenID, amount, deadline, accountSignature, validatorSignature
                    );

                    // burn batch
                    nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                    uuid = _calcUUID(nonce);
                    structHash = keccak256(
                        abi.encode(
                            ERC1155_BURN_BATCH_TYPE_HASH_V2,
                            token,
                            keccak256(abi.encodePacked(tokenIDs)),
                            keccak256(abi.encodePacked(amounts)),
                            nonce,
                            deadline
                        )
                    );
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_BURN_BATCH_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                    ForgeV2(FORGE).burnERC1155Batch(
                        ACCOUNT.addr, token, tokenIDs, amounts, deadline, accountSignature, validatorSignature
                    );
                }
            }
        }

        // v3 mint
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC20_MINT_TYPE_HASH_V3, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV3(FORGE).mintERC20(ACCOUNT.addr, token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC721_MINT_TYPE_HASH_V3, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV3(FORGE).mintERC721(ACCOUNT.addr, token, tokenID, deadline, validatorSignature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(ERC1155_MINT_TYPE_HASH_V3, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV3(FORGE).mintERC1155(ACCOUNT.addr, token, tokenID, amount, deadline, validatorSignature, data);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(
                        ERC1155_MINT_BATCH_TYPE_HASH_V3,
                        ACCOUNT.addr,
                        token,
                        keccak256(abi.encodePacked(tokenIDs)),
                        keccak256(abi.encodePacked(amounts)),
                        nonce,
                        deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV3(FORGE).mintERC1155Batch(
                    ACCOUNT.addr, token, tokenIDs, amounts, deadline, validatorSignature, data
                );
            }
        }
        // v3 transfer
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC20_TRANSFER_TYPE_HASH_V3, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV3(FORGE).transferERC20(ACCOUNT.addr, token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC721_TRANSFER_TYPE_HASH_V3, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV3(FORGE).transferERC721(ACCOUNT.addr, token, tokenID, deadline, validatorSignature, data);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(ERC1155_TRANSFER_TYPE_HASH_V3, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV3(FORGE).transferERC1155(ACCOUNT.addr, token, tokenID, amount, deadline, validatorSignature, data);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(
                        ERC1155_TRANSFER_BATCH_TYPE_HASH_V3,
                        ACCOUNT.addr,
                        token,
                        keccak256(abi.encodePacked(tokenIDs)),
                        keccak256(abi.encodePacked(amounts)),
                        nonce,
                        deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV3(FORGE).transferERC1155Batch(
                    ACCOUNT.addr, token, tokenIDs, amounts, deadline, validatorSignature, data
                );
            }
        }
        // v3 burn
        {
            // erc20
            {
                token = address(mockERC20);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC20_BURN_TYPE_HASH_V3, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV3(FORGE).burnERC20(ACCOUNT.addr, token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash =
                    keccak256(abi.encode(ERC721_BURN_TYPE_HASH_V3, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV3(FORGE).burnERC721(ACCOUNT.addr, token, tokenID, deadline, validatorSignature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(ERC1155_BURN_TYPE_HASH_V3, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV3(FORGE).burnERC1155(ACCOUNT.addr, token, tokenID, amount, deadline, validatorSignature);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                nonce = BaseForge(FORGE).nonces(ACCOUNT.addr);
                uuid = _calcUUID(nonce);
                structHash = keccak256(
                    abi.encode(
                        ERC1155_BURN_BATCH_TYPE_HASH_V3,
                        ACCOUNT.addr,
                        token,
                        keccak256(abi.encodePacked(tokenIDs)),
                        keccak256(abi.encodePacked(amounts)),
                        nonce,
                        deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV3(FORGE).burnERC1155Batch(ACCOUNT.addr, token, tokenIDs, amounts, deadline, validatorSignature);
            }
        }
    }

    function test_dup_nonce() external {
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
            structHash = keccak256(abi.encode(ERC20_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, amount, nonce, deadline));
            validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
            vm.prank(ACCOUNT.addr);
            ForgeV1(FORGE).mintERC20(token, amount, deadline, validatorSignature);
        }

        // v1 mint
        {
            // erc20
            {
                token = address(mockERC20);
                structHash =
                    keccak256(abi.encode(ERC20_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).mintERC20(token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                structHash =
                    keccak256(abi.encode(ERC721_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).mintERC721(token, tokenID, deadline, validatorSignature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                structHash = keccak256(
                    abi.encode(ERC1155_MINT_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).mintERC1155(token, tokenID, amount, deadline, validatorSignature, data);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                structHash = keccak256(
                    abi.encode(
                        ERC1155_MINT_BATCH_TYPE_HASH_V1,
                        ACCOUNT.addr,
                        token,
                        keccak256(abi.encodePacked(tokenIDs)),
                        keccak256(abi.encodePacked(amounts)),
                        nonce,
                        deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).mintERC1155Batch(token, tokenIDs, amounts, deadline, validatorSignature, data);
            }
        }
        // v1 transfer
        {
            // erc20
            {
                token = address(mockERC20);
                structHash =
                    keccak256(abi.encode(ERC20_TRANSFER_TYPE_HASH_V1, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).transferERC20(token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                structHash =
                    keccak256(abi.encode(ERC721_TRANSFER_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).transferERC721(token, tokenID, deadline, validatorSignature, data);
            }
            // erc1155
            {
                token = address(mockERC1155);
                structHash = keccak256(
                    abi.encode(ERC1155_TRANSFER_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).transferERC1155(token, tokenID, amount, deadline, validatorSignature, data);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                structHash = keccak256(
                    abi.encode(
                        ERC1155_TRANSFER_BATCH_TYPE_HASH_V1,
                        ACCOUNT.addr,
                        token,
                        keccak256(abi.encodePacked(tokenIDs)),
                        keccak256(abi.encodePacked(amounts)),
                        nonce,
                        deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).transferERC1155Batch(token, tokenIDs, amounts, deadline, validatorSignature, data);
            }
        }
        // v1 burn
        {
            // erc20
            {
                token = address(mockERC20);
                structHash =
                    keccak256(abi.encode(ERC20_BURN_TYPE_HASH_V1, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).burnERC20(token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                structHash =
                    keccak256(abi.encode(ERC721_BURN_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).burnERC721(token, tokenID, deadline, validatorSignature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                structHash = keccak256(
                    abi.encode(ERC1155_BURN_TYPE_HASH_V1, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).burnERC1155(token, tokenID, amount, deadline, validatorSignature);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                structHash = keccak256(
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
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                vm.prank(ACCOUNT.addr);
                ForgeV1(FORGE).burnERC1155Batch(token, tokenIDs, amounts, deadline, validatorSignature);
            }
        }

        // v2 mint
        {
            // erc20
            {
                token = address(mockERC20);
                structHash = keccak256(abi.encode(ERC20_MINT_TYPE_HASH_V2, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_MINT_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature)));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("ERC20ForgeV2__InvalidAccountSignature(address)", ACCOUNT.addr));
                ForgeV2(FORGE).mintERC20(ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                structHash = keccak256(abi.encode(ERC721_MINT_TYPE_HASH_V2, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC721_VALIDATOR_MINT_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature)));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(
                    abi.encodeWithSignature("ERC721ForgeV2__InvalidAccountSignature(address)", ACCOUNT.addr)
                );
                ForgeV2(FORGE).mintERC721(ACCOUNT.addr, token, tokenID, deadline, accountSignature, validatorSignature);
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    structHash =
                        keccak256(abi.encode(ERC1155_MINT_TYPE_HASH_V2, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_MINT_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature("ERC1155ForgeV2__InvalidAccountSignature(address)", ACCOUNT.addr)
                    );
                    ForgeV2(FORGE).mintERC1155(
                        ACCOUNT.addr, token, tokenID, amount, deadline, accountSignature, validatorSignature, data
                    );

                    // mint batch
                    structHash = keccak256(
                        abi.encode(ERC1155_MINT_BATCH_TYPE_HASH_V2, token, tokenIDs, amounts, nonce, deadline)
                    );
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_MINT_BATCH_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature("ERC1155ForgeV2__InvalidAccountSignature(address)", ACCOUNT.addr)
                    );
                    ForgeV2(FORGE).mintERC1155Batch(
                        ACCOUNT.addr, token, tokenIDs, amounts, deadline, accountSignature, validatorSignature, data
                    );
                }
            }
        }
        // v2 transfer
        {
            // erc20
            {
                token = address(mockERC20);
                structHash = keccak256(abi.encode(ERC20_TRANSFER_TYPE_HASH_V2, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash = keccak256(
                    abi.encode(ERC20_VALIDATOR_TRANSFER_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("ERC20ForgeV2__InvalidAccountSignature(address)", ACCOUNT.addr));
                ForgeV2(FORGE).transferERC20(
                    ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature
                );
            }
            // erc721
            {
                token = address(mockERC721);
                structHash = keccak256(abi.encode(ERC721_TRANSFER_TYPE_HASH_V2, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash = keccak256(
                    abi.encode(ERC721_VALIDATOR_TRANSFER_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(
                    abi.encodeWithSignature("ERC721ForgeV2__InvalidAccountSignature(address)", ACCOUNT.addr)
                );
                ForgeV2(FORGE).transferERC721(
                    ACCOUNT.addr, token, tokenID, deadline, accountSignature, validatorSignature, data
                );
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    structHash =
                        keccak256(abi.encode(ERC1155_TRANSFER_TYPE_HASH_V2, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_TRANSFER_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature("ERC1155ForgeV2__InvalidAccountSignature(address)", ACCOUNT.addr)
                    );
                    ForgeV2(FORGE).transferERC1155(
                        ACCOUNT.addr, token, tokenID, amount, deadline, accountSignature, validatorSignature, data
                    );

                    // transfer batch
                    structHash = keccak256(
                        abi.encode(
                            ERC1155_TRANSFER_BATCH_TYPE_HASH_V2,
                            token,
                            keccak256(abi.encodePacked(tokenIDs)),
                            keccak256(abi.encodePacked(amounts)),
                            nonce,
                            deadline
                        )
                    );
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(
                            ERC1155_VALIDATOR_TRANSFER_BATCH_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature)
                        )
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature("ERC1155ForgeV2__InvalidAccountSignature(address)", ACCOUNT.addr)
                    );
                    ForgeV2(FORGE).transferERC1155Batch(
                        ACCOUNT.addr, token, tokenIDs, amounts, deadline, accountSignature, validatorSignature, data
                    );
                }
            }
        }
        // v2 burnFrom
        {
            // erc20
            {
                token = address(mockERC20);
                structHash = keccak256(abi.encode(ERC20_BURN_TYPE_HASH_V2, token, amount, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC20_VALIDATOR_BURN_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature)));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("ERC20ForgeV2__InvalidAccountSignature(address)", ACCOUNT.addr));
                ForgeV2(FORGE).burnERC20(ACCOUNT.addr, token, amount, deadline, accountSignature, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                structHash = keccak256(abi.encode(ERC721_BURN_TYPE_HASH_V2, token, tokenID, nonce, deadline));
                accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                structHash =
                    keccak256(abi.encode(ERC721_VALIDATOR_BURN_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature)));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(
                    abi.encodeWithSignature("ERC721ForgeV2__InvalidAccountSignature(address)", ACCOUNT.addr)
                );
                ForgeV2(FORGE).burnERC721(ACCOUNT.addr, token, tokenID, deadline, accountSignature, validatorSignature);
            }
            // erc1155
            {
                {
                    token = address(mockERC1155);
                    structHash =
                        keccak256(abi.encode(ERC1155_BURN_TYPE_HASH_V2, token, tokenID, amount, nonce, deadline));
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_BURN_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature("ERC1155ForgeV2__InvalidAccountSignature(address)", ACCOUNT.addr)
                    );
                    ForgeV2(FORGE).burnERC1155(
                        ACCOUNT.addr, token, tokenID, amount, deadline, accountSignature, validatorSignature
                    );

                    // burn batch
                    structHash = keccak256(
                        abi.encode(
                            ERC1155_BURN_BATCH_TYPE_HASH_V2,
                            token,
                            keccak256(abi.encodePacked(tokenIDs)),
                            keccak256(abi.encodePacked(amounts)),
                            nonce,
                            deadline
                        )
                    );
                    accountSignature = _sign_with_domain_separator(structHash, ACCOUNT);
                    structHash = keccak256(
                        abi.encode(ERC1155_VALIDATOR_BURN_BATCH_TYPE_HASH_V2, ACCOUNT.addr, keccak256(accountSignature))
                    );
                    validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                    vm.expectRevert(
                        abi.encodeWithSignature("ERC1155ForgeV2__InvalidAccountSignature(address)", ACCOUNT.addr)
                    );
                    ForgeV2(FORGE).burnERC1155Batch(
                        ACCOUNT.addr, token, tokenIDs, amounts, deadline, accountSignature, validatorSignature
                    );
                }
            }
        }

        // v3 mint
        {
            // erc20
            {
                token = address(mockERC20);
                structHash =
                    keccak256(abi.encode(ERC20_MINT_TYPE_HASH_V3, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV3(FORGE).mintERC20(ACCOUNT.addr, token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                structHash =
                    keccak256(abi.encode(ERC721_MINT_TYPE_HASH_V3, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV3(FORGE).mintERC721(ACCOUNT.addr, token, tokenID, deadline, validatorSignature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                structHash = keccak256(
                    abi.encode(ERC1155_MINT_TYPE_HASH_V3, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV3(FORGE).mintERC1155(ACCOUNT.addr, token, tokenID, amount, deadline, validatorSignature, data);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                structHash = keccak256(
                    abi.encode(
                        ERC1155_MINT_BATCH_TYPE_HASH_V3,
                        ACCOUNT.addr,
                        token,
                        keccak256(abi.encodePacked(tokenIDs)),
                        keccak256(abi.encodePacked(amounts)),
                        nonce,
                        deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV3(FORGE).mintERC1155Batch(
                    ACCOUNT.addr, token, tokenIDs, amounts, deadline, validatorSignature, data
                );
            }
        }
        // v3 transfer
        {
            // erc20
            {
                token = address(mockERC20);
                structHash =
                    keccak256(abi.encode(ERC20_TRANSFER_TYPE_HASH_V3, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV3(FORGE).transferERC20(ACCOUNT.addr, token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                structHash =
                    keccak256(abi.encode(ERC721_TRANSFER_TYPE_HASH_V3, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV3(FORGE).transferERC721(ACCOUNT.addr, token, tokenID, deadline, validatorSignature, data);
            }
            // erc1155
            {
                token = address(mockERC1155);
                structHash = keccak256(
                    abi.encode(ERC1155_TRANSFER_TYPE_HASH_V3, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV3(FORGE).transferERC1155(ACCOUNT.addr, token, tokenID, amount, deadline, validatorSignature, data);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                structHash = keccak256(
                    abi.encode(
                        ERC1155_TRANSFER_BATCH_TYPE_HASH_V3,
                        ACCOUNT.addr,
                        token,
                        keccak256(abi.encodePacked(tokenIDs)),
                        keccak256(abi.encodePacked(amounts)),
                        nonce,
                        deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV3(FORGE).transferERC1155Batch(
                    ACCOUNT.addr, token, tokenIDs, amounts, deadline, validatorSignature, data
                );
            }
        }
        // v3 burn
        {
            // erc20
            {
                token = address(mockERC20);
                structHash =
                    keccak256(abi.encode(ERC20_BURN_TYPE_HASH_V3, ACCOUNT.addr, token, amount, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV3(FORGE).burnERC20(ACCOUNT.addr, token, amount, deadline, validatorSignature);
            }
            // erc721
            {
                token = address(mockERC721);
                structHash =
                    keccak256(abi.encode(ERC721_BURN_TYPE_HASH_V3, ACCOUNT.addr, token, tokenID, nonce, deadline));
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV3(FORGE).burnERC721(ACCOUNT.addr, token, tokenID, deadline, validatorSignature);
            }
            // erc1155
            {
                token = address(mockERC1155);
                structHash = keccak256(
                    abi.encode(ERC1155_BURN_TYPE_HASH_V3, ACCOUNT.addr, token, tokenID, amount, nonce, deadline)
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV3(FORGE).burnERC1155(ACCOUNT.addr, token, tokenID, amount, deadline, validatorSignature);

                // mint batch
                tokenIDs = new uint256[](1);
                amounts = new uint256[](1);
                tokenIDs[0] = tokenID;
                amounts[0] = amount;

                structHash = keccak256(
                    abi.encode(
                        ERC1155_BURN_BATCH_TYPE_HASH_V3,
                        ACCOUNT.addr,
                        token,
                        keccak256(abi.encodePacked(tokenIDs)),
                        keccak256(abi.encodePacked(amounts)),
                        nonce,
                        deadline
                    )
                );
                validatorSignature = _sign_with_domain_separator(structHash, VALIDATOR);
                vm.expectRevert(abi.encodeWithSignature("BaseForge__ECDSAInvalidValidatorSignature()"));
                ForgeV3(FORGE).burnERC1155Batch(ACCOUNT.addr, token, tokenIDs, amounts, deadline, validatorSignature);
            }
        }
    }

    function _sign_with_domain_separator(bytes32 hash, Vm.Wallet memory wallet) internal returns (bytes memory) {
        bytes32 messageHash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, hash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(wallet, messageHash);
        return abi.encodePacked(r, s, v);
    }

    function _calcUUID(uint256 nonce) internal view returns (uint256) {
        return uint256(keccak256(abi.encode(FORGE, ACCOUNT.addr, nonce)));
    }
}
