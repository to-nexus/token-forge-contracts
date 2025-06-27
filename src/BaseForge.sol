// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {Initializable} from "@openzeppelin-contracts-upgradeable-5.3.0/proxy/utils/Initializable.sol";
import {ContextUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/utils/ContextUpgradeable.sol";
import {EIP712Upgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/utils/cryptography/EIP712Upgradeable.sol";
import {NoncesUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/utils/NoncesUpgradeable.sol";
import {ECDSA} from "@openzeppelin-contracts-5.3.0/utils/cryptography/ECDSA.sol";

import {TokenType, ITokenForgeFactoryAlert} from "./interfaces/ITokenForgeFactory.sol";

abstract contract BaseForge is Initializable, ContextUpgradeable, EIP712Upgradeable, NoncesUpgradeable {
    error BaseForge__ZeroAddress();
    error BaseForge__ECDSAInvalidValidatorSignature();
    error BaseForge__ExpiredSignature(uint256 deadline);

    event ValidatorUpdated(address indexed validator);

    /// @custom:storage-location erc7201:cross.storage.BaseForge
    struct BaseForgeStorage {
        address _factory;
        address _validator;
    }

    // keccak256(abi.encode(uint256(keccak256("cross.storage.BaseForge")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant TOKEN_FORGE_FACTORY_STORAGE_LOCATION =
        0x5ac6ce4fd3cf7358c9a2cbdc90756f82e10994295c639f27faf15b59f70a0600;

    function _getBaseForgeStorage() private pure returns (BaseForgeStorage storage $) {
        assembly {
            $.slot := TOKEN_FORGE_FACTORY_STORAGE_LOCATION
        }
    }

    modifier checkDeadline(uint256 deadline) {
        _verifyDeadline(deadline);
        _;
    }

    function __BaseForge_init(string memory service_, address validator_) internal onlyInitializing {
        __EIP712_init(service_, "1");
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
        ITokenForgeFactoryAlert factory = ITokenForgeFactoryAlert(_getBaseForgeStorage()._factory);
        factory.alertMint(tokenType, uuid, token, data);
    }

    function _alertTransferToFactory(TokenType tokenType, uint256 uuid, address token, bytes memory data) internal {
        ITokenForgeFactoryAlert factory = ITokenForgeFactoryAlert(_getBaseForgeStorage()._factory);
        factory.alertTransfer(tokenType, uuid, token, data);
    }

    function _alertBurnToFactory(TokenType tokenType, uint256 uuid, address token, bytes memory data) internal {
        ITokenForgeFactoryAlert factory = ITokenForgeFactoryAlert(_getBaseForgeStorage()._factory);
        factory.alertBurn(tokenType, uuid, token, data);
    }

    function _setValidator(address newValidator) internal {
        if (newValidator == address(0)) revert BaseForge__ZeroAddress();

        BaseForgeStorage storage $ = _getBaseForgeStorage();
        $._validator = newValidator;
        emit ValidatorUpdated(newValidator);
    }
}

import {IDiamondCut, LibDiamond} from "diamond-3-hardhat-1.0.0/contracts/libraries/LibDiamond.sol";
import {IDefaultDiamondCut} from "./interfaces/IDefaultDiamondCut.sol";
import {IBaseForgeFacet} from "./interfaces/IBaseForgeFacet.sol";

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

    function initialize(string memory service_, address validator_) external override initializer {
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
