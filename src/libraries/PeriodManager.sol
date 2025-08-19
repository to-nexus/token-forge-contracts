// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.20;

/**
 * @title PeriodManager
 * @dev Library for managing time-based periods with timezone adjustment
 */
library PeriodManager {
    error PeriodManager__InvalidDuration();

    struct PeriodConfig {
        uint128 duration; // 주기 길이 (초 단위)
        int128 offsetSeconds; // 시간 조정값 (초 단위, 음수 가능)
    }

    /**
     * @dev Get the period start time for a specific timestamp
     * @param $ Period configuration
     * @param timestamp Target timestamp
     * @return Period start timestamp for the given time
     */
    function getPeriodStartForTime(PeriodConfig storage $, uint256 timestamp) internal view returns (uint256) {
        if ($.duration == 0) revert PeriodManager__InvalidDuration();
        unchecked {
            timestamp = timestamp - (timestamp % $.duration);
        }
        if ($.offsetSeconds == 0) {
            return timestamp;
        } else if ($.offsetSeconds > 0) {
            return timestamp + uint256(int256($.offsetSeconds));
        } else {
            uint256 offset = uint256(-int256($.offsetSeconds));
            // Ensure timestamp does not underflow
            if (timestamp < offset) {
                return 0; // Underflow, return 0 or handle as needed
            }
            return timestamp - offset;
        }
    }

    /**
     * @dev Get the start time of the current period
     * @param $ Period configuration
     * @return Period start timestamp
     */
    function getCurrentPeriodStart(PeriodConfig storage $) internal view returns (uint256) {
        return getPeriodStartForTime($, block.timestamp);
    }

    /**
     * @dev Get the end time of the current period
     * @param $ Period configuration
     * @return Period end timestamp
     */
    function getPeriodEnd(PeriodConfig storage $) internal view returns (uint256) {
        return getCurrentPeriodStart($) + $.duration;
    }

    /**
     * @dev Check if it's a new period compared to the last period start
     * @param $ Period configuration
     * @param lastPeriodStart Last recorded period start time
     * @return True if it's a new period
     */
    function isNewPeriod(PeriodConfig storage $, uint256 lastPeriodStart) internal view returns (bool) {
        return getCurrentPeriodStart($) > lastPeriodStart;
    }
}
