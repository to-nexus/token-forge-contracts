// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std-1.9.7/Test.sol";
import {Vm} from "forge-std-1.9.7/Vm.sol";

import {IDiamondCut} from "diamond-3-hardhat-1.0.0/interfaces/IDiamondCut.sol";
import {IDiamondLoupe} from "diamond-3-hardhat-1.0.0/interfaces/IDiamondLoupe.sol";
import {IERC173} from "diamond-3-hardhat-1.0.0/interfaces/IERC173.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts-5.3.0/proxy/ERC1967/ERC1967Proxy.sol";
import {MessageHashUtils} from "@openzeppelin-contracts-5.3.0/utils/cryptography/MessageHashUtils.sol";
import {ShortString, ShortStrings} from "@openzeppelin-contracts-5.3.0/utils/ShortStrings.sol";

import {ForgeProxyCode} from "../src/forges/ForgeProxy.sol";
import {Diamond3Facet} from "../src/forges/Diamond3Facet.sol";
import {ForgeFactory} from "../src/forges/ForgeFactory.sol";
import "../src/forges/BaseForge.sol";
import "../src/forges/ForgeV1.sol";
import "../src/forges/ForgeV2.sol";
import "../src/forges/ForgeV3.sol";

import {ITokenFactory, TokenFactory} from "../src/tokens/TokenFactory.sol";

import "./mock/StructHash.sol";

import {ERC20Mintable} from "../src/tokens/presets/erc20/ERC20Mintable.sol";
import {ERC20Fixed} from "../src/tokens/presets/erc20/ERC20Fixed.sol";
import {ERC20Capped} from "../src/tokens/presets/erc20/ERC20Capped.sol";

contract TestTokenPresets is Test {
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
    bytes32 public DOMAIN_SEPARATOR;

    TokenFactory public tokenFactory;

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
            address[] memory nullImpls;
            address tokenFactoryImpl = address(new TokenFactory());
            address tokenFactoryProxy = address(
                new ERC1967Proxy(
                    tokenFactoryImpl, abi.encodeCall(TokenFactory.initialize, (OWNER, nullImpls, nullImpls, nullImpls))
                )
            );
            tokenFactory = TokenFactory(tokenFactoryProxy);
        }
        vm.stopPrank();
    }

    function test_erc20_mintable() external {
        ERC20Mintable logic = new ERC20Mintable();
        address[] memory erc20Impls = new address[](1);
        erc20Impls[0] = address(logic);
        vm.prank(OWNER);
        tokenFactory.setPresetLogics(ITokenFactory.TokenType.ERC20, erc20Impls, true);
        vm.prank(OWNER);
        ERC20Mintable erc20 =
            ERC20Mintable(tokenFactory.deployERC20(SERVICE_OWNER, "ERC20Mintable", "ERC20M", 18, 0, "", address(logic)));
        assertEq(erc20.balanceOf(SERVICE_OWNER), 0, "Initial supply should be 0");

        {
            address[] memory forges = new address[](1);
            forges[0] = FORGE;
            vm.prank(SERVICE_OWNER);
            erc20.setForges(forges, true);
        }
        // check minting
        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce;
        {
            // v1
            // validator signature
            nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
            uint256 uuid = _calcUUID(nonce);

            bytes32 structHash = keccak256(
                abi.encode(
                    ERC20_MINT_TYPE_HASH_V1, ACCOUNT.addr, address(erc20), amount, address(0), 0, nonce, deadline
                )
            );
            bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

            // expect emit
            vm.expectEmit();
            emit IERC20.Transfer(address(0), ACCOUNT.addr, amount);
            vm.expectEmit(true, true, true, true, address(forgeFactory));
            emit ForgeFactory.ERC20Minted(SERVICE_NAME, uuid, ACCOUNT.addr, address(erc20), amount);

            // send transaction
            vm.prank(ACCOUNT.addr);
            ForgeV1(FORGE).mintERC20(address(erc20), amount, address(0), 0, deadline, abi.encodePacked(r, s, v));
            assertEq(amount, erc20.balanceOf(ACCOUNT.addr), "Minted amount mismatch");
        }
        {
            // v2
            nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
            uint256 uuid = _calcUUID(nonce);

            bytes memory recipientSig;

            {
                bytes32 recipientStructHash = keccak256(
                    abi.encode(ERC20_MINT_TYPE_HASH_V2, address(erc20), amount, address(0), 0, nonce, deadline)
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

            // expect emit
            vm.expectEmit();
            emit IERC20.Transfer(address(0), ACCOUNT.addr, amount);
            vm.expectEmit(true, true, true, true, address(forgeFactory));
            emit ForgeFactory.ERC20Minted(SERVICE_NAME, uuid, ACCOUNT.addr, address(erc20), amount);

            // send transaction
            ForgeV2(FORGE).mintERC20(
                ACCOUNT.addr, address(erc20), amount, address(0), 0, deadline, recipientSig, validatorSig
            );
            assertEq(amount * 2, erc20.balanceOf(ACCOUNT.addr), "Minted amount mismatch");
        }
        {
            // v3
            nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
            uint256 uuid = _calcUUID(nonce);
            bytes32 structHash = keccak256(
                abi.encode(
                    ERC20_MINT_TYPE_HASH_V1, ACCOUNT.addr, address(erc20), amount, address(0), 0, nonce, deadline
                )
            );
            bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

            // expect emit
            vm.expectEmit();
            emit IERC20.Transfer(address(0), ACCOUNT.addr, amount);
            vm.expectEmit(true, true, true, true, address(forgeFactory));
            emit ForgeFactory.ERC20Minted(SERVICE_NAME, uuid, ACCOUNT.addr, address(erc20), amount);

            // send transaction
            ForgeV3(FORGE).mintERC20(
                ACCOUNT.addr, address(erc20), amount, address(0), 0, deadline, abi.encodePacked(r, s, v)
            );
            assertEq(amount * 3, erc20.balanceOf(ACCOUNT.addr), "Minted amount mismatch");
        }

        // check burn able
        vm.prank(ACCOUNT.addr);
        // require approve self
        erc20.approve(ACCOUNT.addr, amount);
        vm.prank(ACCOUNT.addr);
        erc20.burnFrom(ACCOUNT.addr, amount);
        assertEq(erc20.balanceOf(ACCOUNT.addr), amount * 2, "Burned amount mismatch");
    }

    function test_erc20_fixed() external {
        ERC20Fixed logic = new ERC20Fixed();
        address[] memory erc20Impls = new address[](1);
        erc20Impls[0] = address(logic);

        vm.prank(OWNER);
        tokenFactory.setPresetLogics(ITokenFactory.TokenType.ERC20, erc20Impls, true);

        // initialSupply cannot be 0
        vm.expectRevert(abi.encodeWithSignature("TokenBase__NullInput(bytes32)", bytes32("initialSupply")));
        vm.prank(OWNER);
        ERC20Fixed erc20 =
            ERC20Fixed(tokenFactory.deployERC20(SERVICE_OWNER, "ERC20Fixed", "ERC20F", 18, 0, "", address(logic)));

        vm.prank(OWNER);
        erc20 =
            ERC20Fixed(tokenFactory.deployERC20(SERVICE_OWNER, "ERC20Fixed", "ERC20F", 18, 1000e18, "", address(logic)));
        assertEq(erc20.balanceOf(SERVICE_OWNER), 1000e18, "Initial supply should be 1000e18");

        address[] memory forges = new address[](1);
        forges[0] = FORGE;
        vm.prank(SERVICE_OWNER);
        erc20.setForges(forges, true);

        uint256 amount = 100 ether;
        uint256 deadline = block.timestamp + 30; // 30 seconds deadline
        uint256 nonce;
        // check minting revert
        // v1
        {
            // validator signature
            nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

            bytes32 structHash = keccak256(
                abi.encode(
                    ERC20_MINT_TYPE_HASH_V1, ACCOUNT.addr, address(erc20), amount, address(0), 0, nonce, deadline
                )
            );
            bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

            // send transaction
            vm.prank(ACCOUNT.addr);
            vm.expectRevert(abi.encodeWithSignature("ERC20Fixed__MintingNotAllowed()"));
            ForgeV1(FORGE).mintERC20(address(erc20), amount, address(0), 0, deadline, abi.encodePacked(r, s, v));
            assertEq(0, erc20.balanceOf(ACCOUNT.addr), "Minted amount mismatch");
        }
        // v2
        {
            nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

            bytes memory recipientSig;

            {
                bytes32 recipientStructHash = keccak256(
                    abi.encode(ERC20_MINT_TYPE_HASH_V2, address(erc20), amount, address(0), 0, nonce, deadline)
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
            vm.expectRevert(abi.encodeWithSignature("ERC20Fixed__MintingNotAllowed()"));
            ForgeV2(FORGE).mintERC20(
                ACCOUNT.addr, address(erc20), amount, address(0), 0, deadline, recipientSig, validatorSig
            );
            assertEq(0, erc20.balanceOf(ACCOUNT.addr), "Minted amount mismatch");
        }
        // v3
        {
            nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

            bytes32 structHash = keccak256(
                abi.encode(
                    ERC20_MINT_TYPE_HASH_V1, ACCOUNT.addr, address(erc20), amount, address(0), 0, nonce, deadline
                )
            );
            bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

            // send transaction
            vm.expectRevert(abi.encodeWithSignature("ERC20Fixed__MintingNotAllowed()"));
            ForgeV3(FORGE).mintERC20(
                ACCOUNT.addr, address(erc20), amount, address(0), 0, deadline, abi.encodePacked(r, s, v)
            );
            assertEq(0, erc20.balanceOf(ACCOUNT.addr), "Minted amount mismatch");
        }

        // check burn revert
        vm.prank(SERVICE_OWNER);
        erc20.transfer(ACCOUNT.addr, amount);
        vm.prank(ACCOUNT.addr);
        // require approve self
        erc20.approve(ACCOUNT.addr, amount);
        vm.prank(ACCOUNT.addr);
        vm.expectRevert(abi.encodeWithSignature("ERC20Fixed__BurningNotAllowed()"));
        erc20.burnFrom(ACCOUNT.addr, amount);
        assertEq(amount, erc20.balanceOf(ACCOUNT.addr), "Burned amount mismatch");
    }

    function test_erc20_capped() external {
        ERC20Capped logic = new ERC20Capped();
        address[] memory erc20Impls = new address[](1);
        erc20Impls[0] = address(logic);

        vm.prank(OWNER);
        tokenFactory.setPresetLogics(ITokenFactory.TokenType.ERC20, erc20Impls, true);

        uint256 initialSupply = 1000e18;
        uint256 cap = 2000e18;
        bytes memory capData = abi.encode(cap);

        vm.prank(OWNER);
        ERC20Capped erc20 = ERC20Capped(
            tokenFactory.deployERC20(SERVICE_OWNER, "ERC20Capped", "ERC20C", 18, initialSupply, capData, address(logic))
        );

        assertEq(erc20.balanceOf(SERVICE_OWNER), initialSupply, "Initial supply should match");
        assertEq(erc20.cap(), cap, "Cap should match");
        assertEq(erc20.remainingSupply(), cap - initialSupply, "Remaining supply should match");
        assertEq(erc20.isCapReached(), false, "Cap should not be reached initially");

        {
            address[] memory forges = new address[](1);
            forges[0] = FORGE;
            vm.prank(SERVICE_OWNER);
            erc20.setForges(forges, true);
        }

        uint256 mintAmount = 500e18; // This should work (total: 1500e18 < 2000e18)
        uint256 deadline = block.timestamp + 30;
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);
        uint256 uuid = _calcUUID(nonce);

        bytes32 structHash = keccak256(
            abi.encode(
                ERC20_MINT_TYPE_HASH_V1, ACCOUNT.addr, address(erc20), mintAmount, address(0), 0, nonce, deadline
            )
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        vm.expectEmit();
        emit IERC20.Transfer(address(0), ACCOUNT.addr, mintAmount);
        vm.expectEmit(true, true, true, true, address(forgeFactory));
        emit ForgeFactory.ERC20Minted(SERVICE_NAME, uuid, ACCOUNT.addr, address(erc20), mintAmount);

        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).mintERC20(address(erc20), mintAmount, address(0), 0, deadline, abi.encodePacked(r, s, v));

        assertEq(erc20.balanceOf(ACCOUNT.addr), mintAmount, "Minted amount should match");
        assertEq(erc20.totalSupply(), initialSupply + mintAmount, "Total supply should match");
    }

    function test_erc20_capped_mint_exceed_cap() external {
        ERC20Capped erc20 = _setupCappedToken();
        _setupCappedTokenForges(erc20);

        // First mint to get closer to cap
        _mintCappedToken(erc20, 500e18);

        // Try to mint amount that would exceed cap
        uint256 excessAmount = 600e18; // This would make total: 1500e18 + 600e18 = 2100e18 > 2000e18
        uint256 deadline = block.timestamp + 30;

        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        bytes32 structHash = keccak256(
            abi.encode(
                ERC20_MINT_TYPE_HASH_V1, ACCOUNT.addr, address(erc20), excessAmount, address(0), 0, nonce, deadline
            )
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        vm.prank(ACCOUNT.addr);
        vm.expectRevert(abi.encodeWithSignature("ERC20ExceededCap(uint256,uint256)", 1500e18 + excessAmount, 2000e18));
        ForgeV1(FORGE).mintERC20(address(erc20), excessAmount, address(0), 0, deadline, abi.encodePacked(r, s, v));
    }

    function test_erc20_capped_mint_exact_remaining() external {
        ERC20Capped erc20 = _setupCappedToken();
        _setupCappedTokenForges(erc20);

        // First mint to get closer to cap
        _mintCappedToken(erc20, 500e18);

        // Mint exact remaining amount
        uint256 remainingAmount = erc20.remainingSupply();
        uint256 deadline = block.timestamp + 30;

        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        bytes32 structHash = keccak256(
            abi.encode(
                ERC20_MINT_TYPE_HASH_V1, ACCOUNT.addr, address(erc20), remainingAmount, address(0), 0, nonce, deadline
            )
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).mintERC20(address(erc20), remainingAmount, address(0), 0, deadline, abi.encodePacked(r, s, v));

        assertEq(erc20.totalSupply(), 2000e18, "Total supply should equal cap");
        assertEq(erc20.remainingSupply(), 0, "Remaining supply should be 0");
        assertEq(erc20.isCapReached(), true, "Cap should be reached");
    }

    function test_erc20_capped_mint_when_cap_reached() external {
        ERC20Capped erc20 = _setupCappedToken();
        _setupCappedTokenForges(erc20);

        // Mint to reach cap
        _mintCappedToken(erc20, 1000e18); // Total becomes 2000e18 (cap reached)

        // Try to mint when cap is reached
        uint256 deadline = block.timestamp + 30;
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        bytes32 structHash = keccak256(
            abi.encode(ERC20_MINT_TYPE_HASH_V1, ACCOUNT.addr, address(erc20), 1, address(0), 0, nonce, deadline)
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        vm.prank(ACCOUNT.addr);
        vm.expectRevert(abi.encodeWithSignature("ERC20ExceededCap(uint256,uint256)", 2000e18 + 1, 2000e18));
        ForgeV1(FORGE).mintERC20(address(erc20), 1, address(0), 0, deadline, abi.encodePacked(r, s, v));
    }

    // Helper functions
    function _setupCappedToken() internal returns (ERC20Capped) {
        ERC20Capped logic = new ERC20Capped();
        address[] memory erc20Impls = new address[](1);
        erc20Impls[0] = address(logic);

        vm.prank(OWNER);
        tokenFactory.setPresetLogics(ITokenFactory.TokenType.ERC20, erc20Impls, true);

        uint256 initialSupply = 1000e18;
        uint256 cap = 2000e18;
        bytes memory capData = abi.encode(cap);

        vm.prank(OWNER);
        return ERC20Capped(
            tokenFactory.deployERC20(SERVICE_OWNER, "ERC20Capped", "ERC20C", 18, initialSupply, capData, address(logic))
        );
    }

    function _setupCappedTokenForges(ERC20Capped erc20) internal {
        address[] memory forges = new address[](1);
        forges[0] = FORGE;
        vm.prank(SERVICE_OWNER);
        erc20.setForges(forges, true);
    }

    function _mintCappedToken(ERC20Capped erc20, uint256 amount) internal {
        uint256 deadline = block.timestamp + 30;
        uint256 nonce = NoncesUpgradeable(FORGE).nonces(ACCOUNT.addr);

        bytes32 structHash = keccak256(
            abi.encode(ERC20_MINT_TYPE_HASH_V1, ACCOUNT.addr, address(erc20), amount, address(0), 0, nonce, deadline)
        );
        bytes32 hash = MessageHashUtils.toTypedDataHash(DOMAIN_SEPARATOR, structHash);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VALIDATOR, hash);

        vm.prank(ACCOUNT.addr);
        ForgeV1(FORGE).mintERC20(address(erc20), amount, address(0), 0, deadline, abi.encodePacked(r, s, v));
    }

    function _calcUUID(uint256 nonce) internal view returns (uint256) {
        return uint256(keccak256(abi.encode(FORGE, ACCOUNT.addr, nonce)));
    }
}
