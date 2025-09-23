// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.28;

import {IDiamondCut} from "diamond-3-hardhat-1.0.0/interfaces/IDiamondCut.sol";
import {DiamondCutFacet} from "diamond-3-hardhat-1.0.0/facets/DiamondCutFacet.sol";
import {DiamondLoupeFacet} from "diamond-3-hardhat-1.0.0/facets/DiamondLoupeFacet.sol";
import {OwnershipFacet} from "diamond-3-hardhat-1.0.0/facets/OwnershipFacet.sol";

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
            DiamondLoupeFacet.supportsInterface.selector,
            OwnershipFacet.transferOwnership.selector,
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
