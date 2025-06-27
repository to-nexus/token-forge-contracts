// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IDiamondCut} from "diamond-3-hardhat-1.0.0/contracts/interfaces/IDiamondCut.sol";

interface IDefaultDiamondCut {
    function defaultDiamondFacetCut() external view returns (IDiamondCut.FacetCut memory facetCut);
}
