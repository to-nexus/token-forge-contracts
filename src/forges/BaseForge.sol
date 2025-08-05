// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {Initializable} from "@openzeppelin-contracts-upgradeable-5.3.0/proxy/utils/Initializable.sol";
import {ContextUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/utils/ContextUpgradeable.sol";
import {EIP712Upgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/utils/cryptography/EIP712Upgradeable.sol";
import {NoncesUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/utils/NoncesUpgradeable.sol";
import {IERC20Permit} from "@openzeppelin-contracts-5.3.0/token/ERC20/extensions/IERC20Permit.sol";
import {ECDSA} from "@openzeppelin-contracts-5.3.0/utils/cryptography/ECDSA.sol";

import {TokenType, IForgeFactoryAlert} from "../interfaces/IForgeFactory.sol";

abstract contract BaseForge is Initializable, ContextUpgradeable, EIP712Upgradeable, NoncesUpgradeable {
    error BaseForge__ZeroAddress();
    error BaseForge__ECDSAInvalidValidatorSignature();
    error BaseForge__ECDSAInvalidPermitSignature();
    error BaseForge__ExpiredSignature(uint256 deadline);
    error BaseForge__InvalidFeeData(address feeRecipient, uint256 feeBPS);

    event ValidatorUpdated(address indexed validator);

    /// @custom:storage-location erc7201:cross.storage.forge.BaseForge
    struct BaseForgeStorage {
        address _factory;
        address _validator;
    }

    // keccak256(abi.encode(uint256(keccak256("cross.storage.forge.BaseForge")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant BASE_FORGE_STORAGE_LOCATION =
        0x2c5210729a867e7ab740fad20902d38fd0116d2f4c79ff3d49f62658db3eca00;

    function _getBaseForgeStorage() private pure returns (BaseForgeStorage storage $) {
        assembly {
            $.slot := BASE_FORGE_STORAGE_LOCATION
        }
    }

    modifier checkDeadline(uint256 deadline) {
        _verifyDeadline(deadline);
        _;
    }

    modifier erc20Permit(address from, address token, uint256 value, uint256 deadline, bytes memory sig) {
        _erc20Permit(token, from, value, deadline, sig);
        _;
    }

    function __BaseForge_init(bytes32 service_, address validator_) internal onlyInitializing {
        __EIP712_init(string(abi.encodePacked(service_)), "1");
        __Nonces_init();
        __BaseForge_init_unchained(validator_);
    }

    function __BaseForge_init_unchained(address validator_) internal onlyInitializing {
        if (validator_ == address(0)) revert BaseForge__ZeroAddress();

        BaseForgeStorage storage $ = _getBaseForgeStorage();
        $._factory = _msgSender();
        _setValidator(validator_);
    }

    function DOMAIN_SEPARATOR() external view returns (bytes32) {
        return _domainSeparatorV4();
    }

    function validator() public view returns (address) {
        return _getBaseForgeStorage()._validator;
    }

    function _verifyValidatorSignature(bytes32 hash, bytes memory sig) internal view {
        address _validator = validator();
        if (_validator == address(0)) revert BaseForge__ZeroAddress(); // not initialized

        address recover = ECDSA.recover(hash, sig);
        if (_validator != recover) revert BaseForge__ECDSAInvalidValidatorSignature();
    }

    function _verifyDeadline(uint256 deadline) internal view {
        if (block.timestamp > deadline) revert BaseForge__ExpiredSignature(deadline);
    }

    function _alertMintToFactory(TokenType tokenType, uint256 uuid, address token, bytes memory data) internal {
        IForgeFactoryAlert factory = IForgeFactoryAlert(_getBaseForgeStorage()._factory);
        factory.alertMint(tokenType, uuid, token, data);
    }

    function _alertTransferToFactory(TokenType tokenType, uint256 uuid, address token, bytes memory data) internal {
        IForgeFactoryAlert factory = IForgeFactoryAlert(_getBaseForgeStorage()._factory);
        factory.alertTransfer(tokenType, uuid, token, data);
    }

    function _alertTransferFromToFactory(TokenType tokenType, uint256 uuid, address token, bytes memory data)
        internal
    {
        IForgeFactoryAlert factory = IForgeFactoryAlert(_getBaseForgeStorage()._factory);
        factory.alertTransferFrom(tokenType, uuid, token, data);
    }

    function _alertBurnToFactory(TokenType tokenType, uint256 uuid, address token, bytes memory data) internal {
        IForgeFactoryAlert factory = IForgeFactoryAlert(_getBaseForgeStorage()._factory);
        factory.alertBurn(tokenType, uuid, token, data);
    }

    function _setValidator(address newValidator) internal {
        if (newValidator == address(0)) revert BaseForge__ZeroAddress();

        BaseForgeStorage storage $ = _getBaseForgeStorage();
        $._validator = newValidator;
        emit ValidatorUpdated(newValidator);
    }

    function _calcUUID(address user, uint256 nonce) internal view returns (uint256) {
        return uint256(keccak256(abi.encode(address(this), user, nonce)));
    }

    function _erc20Permit(address token, address from, uint256 value, uint256 deadline, bytes memory sig) private {
        if (token == address(0)) revert BaseForge__ZeroAddress();
        if (sig.length != 65) revert BaseForge__ECDSAInvalidPermitSignature();

        bytes32 r;
        bytes32 s;
        uint8 v;
        // ecrecover takes the signature parameters, and the only way to get them
        // currently is to use assembly.
        assembly ("memory-safe") {
            r := mload(add(sig, 0x20))
            s := mload(add(sig, 0x40))
            v := byte(0, mload(add(sig, 0x60)))
        }
        IERC20Permit(token).permit(from, address(this), value, deadline, v, r, s);
    }

    function _erc20CalcFee(address feeRecipient, uint256 feeBPS, uint256 amount)
        internal
        pure
        returns (uint256 fee, uint256 value)
    {
        if (feeBPS == 0) return (0, amount); // No fee to collect
        if (feeRecipient == address(0) || feeBPS > 10_000) revert BaseForge__InvalidFeeData(feeRecipient, feeBPS);

        // Calculate the fee based on the amount and feeBPS
        unchecked {
            fee = (amount * feeBPS) / 10_000;
            value = amount - fee;
        }
    }
}

import {IDiamondCut, LibDiamond} from "diamond-3-hardhat-1.0.0/libraries/LibDiamond.sol";
import {IDefaultDiamondCut} from "../interfaces/IDefaultDiamondCut.sol";
import {IBaseForgeFacet} from "../interfaces/IBaseForgeFacet.sol";

contract BaseForgeFacet is IDefaultDiamondCut, IBaseForgeFacet, BaseForge {
    bytes4[] public BASEFORGE_FACET_FUNCTIONS;

    constructor() {
        BASEFORGE_FACET_FUNCTIONS = [
            EIP712Upgradeable.eip712Domain.selector,
            NoncesUpgradeable.nonces.selector,
            BaseForge.DOMAIN_SEPARATOR.selector,
            BaseForge.validator.selector
            // IBaseForgeFacet.setValidator.selector
        ];
    }

    function initialize(bytes32 service_, address validator_) external override initializer {
        __BaseForge_init(service_, validator_);
    }

    function defaultDiamondFacetCut() external view override returns (IDiamondCut.FacetCut memory facetCut) {
        return IDiamondCut.FacetCut({
            facetAddress: address(this),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: BASEFORGE_FACET_FUNCTIONS
        });
    }

    function setValidator(address validator_) external override {
        LibDiamond.enforceIsContractOwner();
        _setValidator(validator_);
    }
}
