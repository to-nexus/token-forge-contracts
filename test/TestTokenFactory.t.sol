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

import {ITokenFactory} from "../src/interfaces/ITokenFactory.sol";
import {TokenFactory} from "../src/tokens/TokenFactory.sol";
import {MockERC20} from "./mock/MockERC20.sol";
import {MockERC721} from "./mock/MockERC721.sol";
import {MockERC1155} from "./mock/MockERC1155.sol";

import "./mock/StructHash.sol";

import {ERC165Upgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/utils/introspection/ERC165Upgradeable.sol";
import {ERC20Upgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC20/ERC20Upgradeable.sol";
import {ERC721Upgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC721/ERC721Upgradeable.sol";
import {ERC1155Upgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/token/ERC1155/ERC1155Upgradeable.sol";

contract ERC20 is ERC20Upgradeable, ERC165Upgradeable {
    function initialize(string memory name, string memory symbol) external initializer {
        __ERC20_init(name, symbol);
        __ERC165_init();
    }
}

contract ERC721 is ERC721Upgradeable {
    function initialize(string memory name, string memory symbol) external initializer {
        __ERC721_init(name, symbol);
        __ERC165_init();
    }
}

contract ERC1155 is ERC1155Upgradeable {
    function initialize(string memory uri) external initializer {
        __ERC1155_init(uri);
        __ERC165_init();
    }
}

contract TestTokenFactory is Test {
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
            bytes4[] memory v3Selectors = new bytes4[](9);
            v3Selectors[0] = forgeV3.mintERC20.selector;
            v3Selectors[1] = forgeV3.mintERC721.selector;
            v3Selectors[2] = forgeV3.mintERC1155.selector;
            v3Selectors[3] = forgeV3.mintERC1155Batch.selector;

            v3Selectors[4] = forgeV3.transferERC20.selector;
            v3Selectors[5] = forgeV3.transferERC721.selector;
            v3Selectors[6] = forgeV3.transferERC1155.selector;
            v3Selectors[7] = forgeV3.transferERC1155Batch.selector;

            v3Selectors[8] = forgeV3.burnERC20Permit.selector;

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
        vm.stopPrank();
    }

    function test_token_factory_invalid_logic_address() external {
        vm.startPrank(OWNER);
        address[] memory logicAddresses = new address[](1);
        logicAddresses[0] = address(0);
        vm.expectRevert(abi.encodeWithSignature("TokenFactory__ZeroAddress(bytes32)", bytes32("impls")));
        tokenFactory.setPresetLogics(ITokenFactory.TokenType.ERC20, logicAddresses, true);
        vm.expectRevert(abi.encodeWithSignature("TokenFactory__ZeroAddress(bytes32)", bytes32("impls")));
        tokenFactory.setPresetLogics(ITokenFactory.TokenType.ERC721, logicAddresses, true);
        vm.expectRevert(abi.encodeWithSignature("TokenFactory__ZeroAddress(bytes32)", bytes32("impls")));
        tokenFactory.setPresetLogics(ITokenFactory.TokenType.ERC1155, logicAddresses, true);

        address invalid = address(new ERC20());
        logicAddresses[0] = invalid;
        vm.expectRevert(abi.encodeWithSignature("TokenFactory__InvalidLogic(uint8,address)", 0, invalid));
        tokenFactory.setPresetLogics(ITokenFactory.TokenType.ERC20, logicAddresses, true);

        invalid = address(new ERC721());
        logicAddresses[0] = invalid;
        vm.expectRevert(abi.encodeWithSignature("TokenFactory__InvalidLogic(uint8,address)", 1, invalid));
        tokenFactory.setPresetLogics(ITokenFactory.TokenType.ERC721, logicAddresses, true);

        invalid = address(new ERC1155());
        logicAddresses[0] = invalid;
        vm.expectRevert(abi.encodeWithSignature("TokenFactory__InvalidLogic(uint8,address)", 2, invalid));
        tokenFactory.setPresetLogics(ITokenFactory.TokenType.ERC1155, logicAddresses, true);

        vm.stopPrank();
    }

    function test_token_factory_forge_factory() external view {
        assertEq(address(forgeFactory), tokenFactory.getForgeFactory());
    }

    function test_token_factory_preset_logics() external {
        address[] memory erc20Impls = tokenFactory.getPresetLogics(ITokenFactory.TokenType.ERC20);
        address[] memory erc721Impls = tokenFactory.getPresetLogics(ITokenFactory.TokenType.ERC721);
        address[] memory erc1155Impls = tokenFactory.getPresetLogics(ITokenFactory.TokenType.ERC1155);

        assertEq(1, erc20Impls.length);
        assertEq(mockERC20Impl, erc20Impls[0]);
        assertEq(1, erc721Impls.length);
        assertEq(mockERC721Impl, erc721Impls[0]);
        assertEq(1, erc1155Impls.length);
        assertEq(mockERC1155Impl, erc1155Impls[0]);

        vm.startPrank(OWNER);
        address newMockERC20 = address(new MockERC20());
        address[] memory newErc20Impls = new address[](1);
        newErc20Impls[0] = mockERC20Impl;
        tokenFactory.setPresetLogics(ITokenFactory.TokenType.ERC20, newErc20Impls, true);
        erc20Impls = tokenFactory.getPresetLogics(ITokenFactory.TokenType.ERC20);
        assertEq(1, erc20Impls.length);

        newErc20Impls[0] = newMockERC20;
        tokenFactory.setPresetLogics(ITokenFactory.TokenType.ERC20, newErc20Impls, true);
        erc20Impls = tokenFactory.getPresetLogics(ITokenFactory.TokenType.ERC20);
        assertEq(2, erc20Impls.length);

        tokenFactory.setPresetLogics(ITokenFactory.TokenType.ERC20, newErc20Impls, false);
        erc20Impls = tokenFactory.getPresetLogics(ITokenFactory.TokenType.ERC20);
        assertEq(1, erc20Impls.length);

        newErc20Impls[0] = mockERC20Impl;
        tokenFactory.setPresetLogics(ITokenFactory.TokenType.ERC20, newErc20Impls, false);
        erc20Impls = tokenFactory.getPresetLogics(ITokenFactory.TokenType.ERC20);
        assertEq(0, erc20Impls.length);

        address newMockERC721 = address(new MockERC721());
        address[] memory newErc721Impls = new address[](1);
        newErc721Impls[0] = mockERC721Impl;
        tokenFactory.setPresetLogics(ITokenFactory.TokenType.ERC721, newErc721Impls, true);
        erc721Impls = tokenFactory.getPresetLogics(ITokenFactory.TokenType.ERC721);
        assertEq(1, erc721Impls.length);

        newErc721Impls[0] = newMockERC721;
        tokenFactory.setPresetLogics(ITokenFactory.TokenType.ERC721, newErc721Impls, true);
        erc721Impls = tokenFactory.getPresetLogics(ITokenFactory.TokenType.ERC721);
        assertEq(2, erc721Impls.length);

        tokenFactory.setPresetLogics(ITokenFactory.TokenType.ERC721, newErc721Impls, false);
        erc721Impls = tokenFactory.getPresetLogics(ITokenFactory.TokenType.ERC721);
        assertEq(1, erc721Impls.length);

        newErc721Impls[0] = mockERC721Impl;
        tokenFactory.setPresetLogics(ITokenFactory.TokenType.ERC721, newErc721Impls, false);
        erc721Impls = tokenFactory.getPresetLogics(ITokenFactory.TokenType.ERC721);
        assertEq(0, erc721Impls.length);

        address newMockERC1155 = address(new MockERC1155());
        address[] memory newErc1155Impls = new address[](1);
        newErc1155Impls[0] = mockERC1155Impl;
        tokenFactory.setPresetLogics(ITokenFactory.TokenType.ERC1155, newErc1155Impls, true);
        erc1155Impls = tokenFactory.getPresetLogics(ITokenFactory.TokenType.ERC1155);
        assertEq(1, erc1155Impls.length);

        newErc1155Impls[0] = newMockERC1155;
        tokenFactory.setPresetLogics(ITokenFactory.TokenType.ERC1155, newErc1155Impls, true);
        erc1155Impls = tokenFactory.getPresetLogics(ITokenFactory.TokenType.ERC1155);
        assertEq(2, erc1155Impls.length);

        tokenFactory.setPresetLogics(ITokenFactory.TokenType.ERC1155, newErc1155Impls, false);
        erc1155Impls = tokenFactory.getPresetLogics(ITokenFactory.TokenType.ERC1155);
        assertEq(1, erc1155Impls.length);

        newErc1155Impls[0] = mockERC1155Impl;
        tokenFactory.setPresetLogics(ITokenFactory.TokenType.ERC1155, newErc1155Impls, false);
        erc1155Impls = tokenFactory.getPresetLogics(ITokenFactory.TokenType.ERC1155);
        assertEq(0, erc1155Impls.length);

        vm.stopPrank();
    }

    function test_token_factory_service_tokens() external {
        assertEq(0, tokenFactory.getServiceTokens(SERVICE_NAME, ITokenFactory.TokenType.ERC20).length);
        assertEq(0, tokenFactory.getServiceTokens(SERVICE_NAME, ITokenFactory.TokenType.ERC721).length);
        assertEq(0, tokenFactory.getServiceTokens(SERVICE_NAME, ITokenFactory.TokenType.ERC1155).length);

        vm.prank(OWNER);
        tokenFactory.deployERC20(OWNER, SERVICE_NAME, "TestToken", "TTK", 18, 0, mockERC20Impl);
        assertEq(1, tokenFactory.getServiceTokens(SERVICE_NAME, ITokenFactory.TokenType.ERC20).length);
        assertEq(0, tokenFactory.getServiceTokens(SERVICE_NAME, ITokenFactory.TokenType.ERC721).length);
        assertEq(0, tokenFactory.getServiceTokens(SERVICE_NAME, ITokenFactory.TokenType.ERC1155).length);

        vm.prank(OWNER);
        tokenFactory.deployERC721(OWNER, SERVICE_NAME, "TestNFT", "TNFT", "https://xxx.yyy.zzz/", mockERC721Impl);
        assertEq(1, tokenFactory.getServiceTokens(SERVICE_NAME, ITokenFactory.TokenType.ERC20).length);
        assertEq(1, tokenFactory.getServiceTokens(SERVICE_NAME, ITokenFactory.TokenType.ERC721).length);
        assertEq(0, tokenFactory.getServiceTokens(SERVICE_NAME, ITokenFactory.TokenType.ERC1155).length);

        vm.prank(OWNER);
        tokenFactory.deployERC1155(OWNER, SERVICE_NAME, "https://xxx.yyy.zzz/", mockERC1155Impl);
        assertEq(1, tokenFactory.getServiceTokens(SERVICE_NAME, ITokenFactory.TokenType.ERC20).length);
        assertEq(1, tokenFactory.getServiceTokens(SERVICE_NAME, ITokenFactory.TokenType.ERC721).length);
        assertEq(1, tokenFactory.getServiceTokens(SERVICE_NAME, ITokenFactory.TokenType.ERC1155).length);

        address[] memory erc20Tokens = tokenFactory.getServiceTokens(SERVICE_NAME, ITokenFactory.TokenType.ERC20);
        address[] memory erc721Tokens = tokenFactory.getServiceTokens(SERVICE_NAME, ITokenFactory.TokenType.ERC721);
        address[] memory erc1155Tokens = tokenFactory.getServiceTokens(SERVICE_NAME, ITokenFactory.TokenType.ERC1155);
        assertNotEq(erc20Tokens[0], address(0));
        assertNotEq(erc721Tokens[0], address(0));
        assertNotEq(erc1155Tokens[0], address(0));

        assertNotEq(erc20Tokens[0], erc721Tokens[0]);
        assertNotEq(erc1155Tokens[0], erc721Tokens[0]);
        assertNotEq(erc1155Tokens[0], erc20Tokens[0]);
    }

    function test_token_factory_invalid_service_name() external {
        string memory invalidServiceName = "InvalidService";

        vm.startPrank(OWNER);
        vm.expectRevert(abi.encodeWithSignature("TokenFactory__InvalidService(string)", invalidServiceName));
        tokenFactory.deployERC20(OWNER, invalidServiceName, "TestToken", "TTK", 18, 0, mockERC20Impl);

        vm.expectRevert(abi.encodeWithSignature("TokenFactory__InvalidService(string)", invalidServiceName));
        tokenFactory.deployERC721(OWNER, invalidServiceName, "TestNFT", "TNFT", "https://xxx.yyy.zzz/", mockERC721Impl);

        vm.expectRevert(abi.encodeWithSignature("TokenFactory__InvalidService(string)", invalidServiceName));
        tokenFactory.deployERC1155(OWNER, invalidServiceName, "https://xxx.yyy.zzz/", mockERC1155Impl);
        vm.stopPrank();
    }
}
