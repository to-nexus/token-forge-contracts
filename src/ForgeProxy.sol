// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

/**
 * override constructor from diamond-3-hardhat/contracts/Diamond.sol
 */
import {LibDiamond} from "diamond-3-hardhat-1.0.0/libraries/LibDiamond.sol";
import {IDiamondCut} from "diamond-3-hardhat-1.0.0/interfaces/IDiamondCut.sol";
import {IDefaultDiamondCut} from "./interfaces/IDefaultDiamondCut.sol";
import {IBaseForgeFacet} from "./interfaces/IBaseForgeFacet.sol";

contract ForgeProxy {
    constructor(
        address owner,
        address validator,
        IDiamondCut.FacetCut[] memory addCuts,
        string memory service,
        address diamondImpl,
        address baseImpl
    ) payable {
        LibDiamond.setContractOwner(owner);
        uint256 length = addCuts.length + 2;
        // Add the diamondCut external function from the diamondCutFacet
        IDiamondCut.FacetCut[] memory cuts = new IDiamondCut.FacetCut[](length);

        cuts[0] = IDefaultDiamondCut(diamondImpl).defaultDiamondFacetCut();
        cuts[1] = IDefaultDiamondCut(baseImpl).defaultDiamondFacetCut();
        unchecked {
            for (uint256 i = 2; i < length; ++i) {
                cuts[i] = addCuts[i - 2];
            }
        }
        LibDiamond.diamondCut(cuts, baseImpl, abi.encodeCall(IBaseForgeFacet.initialize, (service, validator)));
    }

    // Find facet for function that is called and execute the
    // function if a facet is found and return any value.
    fallback() external payable {
        LibDiamond.DiamondStorage storage ds;
        bytes32 position = LibDiamond.DIAMOND_STORAGE_POSITION;
        // get diamond storage
        assembly {
            ds.slot := position
        }
        // get facet from function selector
        address facet = ds.selectorToFacetAndPosition[msg.sig].facetAddress;
        require(facet != address(0), "Diamond: Function does not exist");
        // Execute external function from facet using delegatecall and return any value.
        assembly {
            // copy function selector and any arguments
            calldatacopy(0, 0, calldatasize())
            // execute function call using the facet
            let result := delegatecall(gas(), facet, 0, calldatasize(), 0, 0)
            // get any return value
            returndatacopy(0, 0, returndatasize())
            // return any return value or error back to the caller
            switch result
            case 0 { revert(0, returndatasize()) }
            default { return(0, returndatasize()) }
        }
    }

    receive() external payable {}
}

contract ForgeProxyCode {
    function code() external pure returns (bytes memory) {
        return type(ForgeProxy).creationCode;
    }
}
