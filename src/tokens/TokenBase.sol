// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {UUPSUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/proxy/utils/UUPSUpgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/access/OwnableUpgradeable.sol";
import {ERC165Upgradeable} from "@openzeppelin-contracts-upgradeable-5.3.0/utils/introspection/ERC165Upgradeable.sol";
import {EnumerableSet} from "@openzeppelin-contracts-5.3.0/utils/structs/EnumerableSet.sol";

abstract contract TokenBase is UUPSUpgradeable, OwnableUpgradeable, ERC165Upgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;

    error TokenBase__NullInput(bytes32 field);
    error TokenBase__OnlyForge(address caller);

    event ForgeAdded(address indexed forge);
    event ForgeRemoved(address indexed forge);

    // keccak256(abi.encode(uint256(keccak256("cross.storage.forge.token")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant TokenBaseStorageLocation =
        0x5336721aabdf58b9f67dce2b9c89749de013c2e87912f00ce8440ff438c6c800;

    function _getForgesStorage() private pure returns (EnumerableSet.AddressSet storage $) {
        assembly {
            $.slot := TokenBaseStorageLocation
        }
    }

    function __TokenBase_init(address _owner) internal onlyInitializing {
        __TokenBase_init_unchained(_owner);
    }

    function __TokenBase_init_unchained(address _owner) internal onlyInitializing {
        __Ownable_init();
        __ERC165_init();
        transferOwnership(_owner);
    }

    modifier onlyForge() {
        if (!isForge(_msgSender())) revert TokenBase__OnlyForge(_msgSender());
        _;
    }

    function forges() external view returns (address[] memory) {
        return _getForgesStorage().values();
    }

    function forgeCount() external view returns (uint256) {
        return _getForgesStorage().length();
    }

    function forgeByIndex(uint256 index) external view returns (address) {
        return _getForgesStorage().at(index);
    }

    function isForge(address forge) public view returns (bool) {
        return _getForgesStorage().contains(forge);
    }

    function setForges(address[] calldata forges, bool add) external onlyOwner {
        EnumerableSet.AddressSet storage $ = _getForgesStorage();
        function (EnumerableSet.AddressSet storage $, address forge) fn = add ? _addForge : _removeForge;
        unchecked {
            for (uint256 i = 0; i < forges.length;) {
                fn($, forges[i]);
                ++i;
            }
        }
    }

    function _addForge(EnumerableSet.AddressSet storage $, address forge) private {
        if (forge == address(0)) revert TokenBase__NullInput("forge");
        if ($.add(forge)) {
            emit ForgeAdded(forge);
        }
    }

    function _removeForge(EnumerableSet.AddressSet storage $, address forge) private {
        if ($.remove(forge)) {
            emit ForgeRemoved(forge);
        }
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
