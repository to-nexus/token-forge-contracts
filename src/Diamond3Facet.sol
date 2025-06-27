// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {IDiamondCut} from "diamond-3-hardhat-1.0.0/contracts/interfaces/IDiamondCut.sol";
import {DiamondCutFacet} from "diamond-3-hardhat-1.0.0/contracts/facets/DiamondCutFacet.sol";
import {DiamondLoupeFacet} from "diamond-3-hardhat-1.0.0/contracts/facets/DiamondLoupeFacet.sol";
import {OwnershipFacet} from "diamond-3-hardhat-1.0.0/contracts/facets/OwnershipFacet.sol";

import {IDefaultDiamondCut} from "./interfaces/IDefaultDiamondCut.sol";

contract Diamond3Facet is IDefaultDiamondCut, DiamondCutFacet, DiamondLoupeFacet, OwnershipFacet {
    bytes4[] public DIAMOND3_FACET_FUNCTIONS;

    constructor() {
        DIAMOND3_FACET_FUNCTIONS = [
            DiamondCutFacet.diamondCut.selector,
            DiamondLoupeFacet.facets.selector,
            DiamondLoupeFacet.facetFunctionSelectors.selector,
            DiamondLoupeFacet.facetAddresses.selector,
            DiamondLoupeFacet.facetAddress.selector,
            // OwnershipFacet.transferOwnership.selector,
            OwnershipFacet.owner.selector
        ];
    }

    function defaultDiamondFacetCut() external view returns (IDiamondCut.FacetCut memory) {
        return IDiamondCut.FacetCut({
            facetAddress: address(this),
            action: IDiamondCut.FacetCutAction.Add,
            functionSelectors: DIAMOND3_FACET_FUNCTIONS
        });
    }
}
