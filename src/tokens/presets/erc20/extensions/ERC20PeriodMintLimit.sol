// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {SafeCast} from "@openzeppelin-contracts-5.3.0/utils/math/SafeCast.sol";

import {ERC20Base} from "../ERC20Base.sol";
import {PeriodManager} from "../../../../libraries/PeriodManager.sol";

abstract contract ERC20PeriodMintLimit is ERC20Base {
    using PeriodManager for PeriodManager.PeriodConfig;

    error ERC20PeriodMintLimit__InvalidLength();
    error ERC20PeriodMintLimit__InvalidLimitData();
    error ERC20PeriodMintLimit__ExceedsPeriodLimit(uint256 requested, uint256 available);

    event PeriodStarted(uint256 indexed periodStartBlock, uint256 availableCapacity);
    event MintLimitUpdated(uint256 oldLimits, uint256 newLimits);

    /// @custom:storage-location erc7201:cross.storage.forge.erc20.ERC20PeriodMintLimit
    struct ERC20PeriodMintLimitStorage {
        PeriodManager.PeriodConfig period;
        uint256 periodStartBlock; // The block number when the current period started
        uint256 limit; // The maximum amount that can be minted in a period
        uint256 periodCapacity; // The remaining capacity for the current period
    }

    // keccak256(abi.encode(uint256(keccak256("cross.storage.forge.erc20.ERC20PeriodMintLimit")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant ERC20PeriodMintLimitStorageLocation =
        0x02a11301204be66f32a67781798c285672cc036f4b261d5c0050a8b03927de00;

    function _getERC20PeriodMintLimitStorage() private pure returns (ERC20PeriodMintLimitStorage storage $) {
        assembly {
            $.slot := ERC20PeriodMintLimitStorageLocation
        }
        return $;
    }

    function __ERC20PeriodMintLimit_init(uint256 duration, int256 offsetSeconds, uint256 limit)
        internal
        onlyInitializing
    {
        if (duration == 0 || limit == 0) {
            revert TokenBase__NullInput("limit or period");
        }

        ERC20PeriodMintLimitStorage storage $ = _getERC20PeriodMintLimitStorage();
        $.period = PeriodManager.PeriodConfig(SafeCast.toUint128(duration), SafeCast.toInt128(offsetSeconds));
        $.limit = limit;
    }

    function mint(address to, uint256 amount) public virtual override {
        ERC20PeriodMintLimitStorage storage $ = _getERC20PeriodMintLimitStorage();

        uint256 periodCapacity = $.periodCapacity;
        {
            uint256 currentPeriodStartBlock = $.period.getCurrentPeriodStart();
            // Check if the period has started
            if (currentPeriodStartBlock != $.periodStartBlock) {
                // Initialize the period start block if not set
                $.periodStartBlock = currentPeriodStartBlock;
                periodCapacity = $.limit;

                emit PeriodStarted(currentPeriodStartBlock, periodCapacity);
            }
        }

        // Check available capacity
        if (periodCapacity < amount) {
            revert ERC20PeriodMintLimit__ExceedsPeriodLimit(amount, periodCapacity);
        }
        // Update available capacity
        unchecked {
            $.periodCapacity = periodCapacity - amount;
        }

        // Mint the tokens
        super.mint(to, amount);
    }

    function periodConfig() external view returns (uint256, int256) {
        PeriodManager.PeriodConfig storage period = _getERC20PeriodMintLimitStorage().period;
        return (period.duration, period.offsetSeconds);
    }

    function maxMintPerPeriod() external view returns (uint256) {
        return _getERC20PeriodMintLimitStorage().limit;
    }

    function availableMintCapacity() external view returns (uint256) {
        ERC20PeriodMintLimitStorage storage $ = _getERC20PeriodMintLimitStorage();
        return $.period.isNewPeriod($.periodStartBlock) ? $.limit : $.periodCapacity;
    }

    function periodStartBlock() public view returns (uint256) {
        return _getERC20PeriodMintLimitStorage().period.getCurrentPeriodStart();
    }

    function updateMintLimit(uint256 newLimit) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (newLimit == 0) revert TokenBase__NullInput("newLimit");

        ERC20PeriodMintLimitStorage storage $ = _getERC20PeriodMintLimitStorage();
        // Update limit
        emit MintLimitUpdated($.limit, newLimit);
        $.limit = newLimit;
    }
}
