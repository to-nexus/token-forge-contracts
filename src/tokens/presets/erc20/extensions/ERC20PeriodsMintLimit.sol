// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {SafeCast} from "@openzeppelin-contracts-5.3.0/utils/math/SafeCast.sol";
import {PeriodManager} from "../../../../libraries/PeriodManager.sol";
import {ERC20Base} from "../ERC20Base.sol";

abstract contract ERC20PeriodsMintLimit is ERC20Base {
    using PeriodManager for PeriodManager.PeriodConfig;

    error ERC20PeriodsMintLimit__InvalidLength();
    error ERC20PeriodsMintLimit__InvalidLimitData(uint256 index);
    error ERC20PeriodsMintLimit__ExceedsPeriodLimit(uint256 period, uint256 requested, uint256 available);

    event PeriodStarted(uint256 indexed periodStart, uint256 availableCapacity);
    event MintLimitUpdated(uint256[] oldLimits, uint256[] newLimits);

    /// @custom:storage-location erc7201:cross.storage.forge.erc20.ERC20PeriodsMintLimit
    struct ERC20PeriodsMintLimitStorage {
        uint256 length;
        PeriodManager.PeriodConfig[] periods;
        uint256[] limits;
        uint256[] periodStartTimes;
        uint256[] periodCapacities;
    }

    // keccak256(abi.encode(uint256(keccak256("cross.storage.forge.erc20.ERC20PeriodsMintLimit")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant ERC20PeriodsMintLimitStorageLocation =
        0x70751f9b068a2d675ee078361261fbe89193a5378437f7a2680d0d1e667b9d00;

    function _getERC20PeriodsMintLimitStorage() private pure returns (ERC20PeriodsMintLimitStorage storage $) {
        assembly {
            $.slot := ERC20PeriodsMintLimitStorageLocation
        }
        return $;
    }

    function __ERC20PeriodsMintLimit_init(
        uint256[] memory durations,
        int256[] memory offsetSeconds,
        uint256[] memory limits
    ) internal onlyInitializing {
        uint256 length = limits.length;
        if (length == 0 || length != durations.length || length != offsetSeconds.length) {
            revert ERC20PeriodsMintLimit__InvalidLength();
        }

        ERC20PeriodsMintLimitStorage storage $ = _getERC20PeriodsMintLimitStorage();

        uint256 minDuration = 0;
        uint256 minLimit = 0;
        unchecked {
            for (uint256 i = 0; i < length; ++i) {
                (uint256 duration, uint256 limit) = (durations[i], limits[i]);
                if (duration == 0 || limit == 0) {
                    revert TokenBase__NullInput("limits or durations");
                }
                if (duration <= minDuration || limit <= minLimit) {
                    revert ERC20PeriodsMintLimit__InvalidLimitData(i);
                }
                $.periods.push(
                    PeriodManager.PeriodConfig(SafeCast.toUint128(duration), SafeCast.toInt128(offsetSeconds[i]))
                );
                (minDuration, minLimit) = (duration, limit);
            }
        }

        $.length = length;
        $.limits = limits;
        $.periodStartTimes = new uint256[](length);
        $.periodCapacities = new uint256[](length);
    }

    function mint(address to, uint256 amount) public virtual override {
        ERC20PeriodsMintLimitStorage storage $ = _getERC20PeriodsMintLimitStorage();

        uint256[] storage _periodStartTimes = $.periodStartTimes;
        uint256[] storage _periodCapacities = $.periodCapacities;
        uint256[] memory currentPeriodStartTimes = periodStartTimes();

        uint256 length = $.length;

        unchecked {
            for (uint256 i = 0; i < length; ++i) {
                (uint256 periodCapacity, uint256 periodStartTime, uint256 currentPeriodStartTime) =
                    (_periodCapacities[i], _periodStartTimes[i], currentPeriodStartTimes[i]);

                // Check if the period has started
                if (periodStartTime != currentPeriodStartTime) {
                    // Initialize the period start time if not set
                    uint256 limit = $.limits[i];
                    _periodStartTimes[i] = currentPeriodStartTime;
                    periodCapacity = limit;

                    emit PeriodStarted(currentPeriodStartTime, limit);
                }

                if (periodCapacity == type(uint256).max) {
                    // If the period capacity is set to max, it means no limit for this period
                    _periodCapacities[i] = periodCapacity;
                } else {
                    // Check available capacity
                    if (periodCapacity < amount) {
                        revert ERC20PeriodsMintLimit__ExceedsPeriodLimit($.periods[i].duration, amount, periodCapacity);
                    }
                    // Update available capacity
                    _periodCapacities[i] = (periodCapacity - amount);
                }
            }
        }

        // Mint the tokens
        super.mint(to, amount);
    }

    function periodConfigs() external view returns (uint256[] memory, int256[] memory) {
        ERC20PeriodsMintLimitStorage storage $ = _getERC20PeriodsMintLimitStorage();
        uint256 length = $.length;
        uint256[] memory durations = new uint256[](length);
        int256[] memory offsets = new int256[](length);
        PeriodManager.PeriodConfig[] storage periods = $.periods;
        for (uint256 i = 0; i < length;) {
            unchecked {
                durations[i] = periods[i].duration;
                offsets[i] = periods[i].offsetSeconds;
                ++i;
            }
        }
        return (durations, offsets);
    }

    function maxMintPerPeriods() external view returns (uint256[] memory) {
        return _getERC20PeriodsMintLimitStorage().limits;
    }

    function availableMintCapacities() external view returns (uint256[] memory) {
        ERC20PeriodsMintLimitStorage storage $ = _getERC20PeriodsMintLimitStorage();
        uint256[] memory _currentPeriodStartTimes = periodStartTimes();
        uint256[] memory _periodStartTimes = $.periodStartTimes;

        uint256 length = $.length;
        uint256[] memory capacities = new uint256[](length);
        for (uint256 i = 0; i < length;) {
            uint256 _periodStart = _periodStartTimes[i];
            if (_periodStart == _currentPeriodStartTimes[i]) {
                capacities[i] = $.periodCapacities[i];
            } else {
                capacities[i] = $.limits[i];
            }
            unchecked {
                ++i;
            }
        }
        return capacities;
    }

    function periodStartTimes() public view returns (uint256[] memory) {
        ERC20PeriodsMintLimitStorage storage $ = _getERC20PeriodsMintLimitStorage();
        PeriodManager.PeriodConfig[] storage _periods = $.periods;

        uint256 length = _periods.length;
        uint256[] memory startTimes = new uint256[](length);
        for (uint256 i = 0; i < length;) {
            startTimes[i] = _periods[i].getCurrentPeriodStart();
            unchecked {
                ++i;
            }
        }
        return startTimes;
    }

    function updateMintLimits(uint256[] calldata newLimits) external onlyRole(DEFAULT_ADMIN_ROLE) {
        // Validate new limit
        uint256 length = newLimits.length;
        if (length == 0) revert TokenBase__NullInput("newLimits");

        ERC20PeriodsMintLimitStorage storage $ = _getERC20PeriodsMintLimitStorage();
        if (length != $.length) {
            revert ERC20PeriodsMintLimit__InvalidLength();
        }
        uint256 minLimit = 0;
        for (uint256 i = 0; i < length;) {
            uint256 newLimit = newLimits[i];
            if (newLimit == 0) revert TokenBase__NullInput("newLimits");
            if (newLimit <= minLimit && newLimit != type(uint256).max) {
                revert ERC20PeriodsMintLimit__InvalidLimitData(i);
            }
            minLimit = newLimit;
            unchecked {
                ++i;
            }
        }

        // Update limit
        emit MintLimitUpdated($.limits, newLimits);
        $.limits = newLimits;
    }
}
