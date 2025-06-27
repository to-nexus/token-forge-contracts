// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IBaseForgeFacet {
    function initialize(string memory service_, address validator_) external;
    function setValidator(address validator_) external;
}
