// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC20Base, ERC20Capable} from "./ERC20Capable.sol";

contract ERC20PeriodsMintLimit is ERC20Capable {
    error ERC20PeriodsMintLimit__InvalidLength();
    error ERC20PeriodsMintLimit__InvalidLimitData(uint256 index);
    error ERC20PeriodsMintLimit__ExceedsPeriodLimit(uint256 period, uint256 requested, uint256 available);

    event PeriodStarted(uint256 indexed periodStartBlock, uint256 availableCapacity);
    event MintLimitUpdated(uint256[] oldLimits, uint256[] newLimits);

    /// @custom:storage-location erc7201:cross.storage.forge.erc20.ERC20PeriodsMintLimit
    struct ERC20PeriodsMintLimitStorage {
        uint256 length;
        uint256[] limits;
        uint256[] periods;
        uint256[] periodStartBlocks;
        uint256[] periodCapacities;
    }

    // keccak256(abi.encode(uint256(keccak256("cross.storage.forge.erc20.ERC20PeriodsMintLimit")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant ERC20PeriodsMintLimitStorageLocation =
        0x0f070392f17d5f958cc1ac31867dabecfc5c9758b4a419a200803226d7155d00;

    function _getERC20PeriodsMintLimitStorage() private pure returns (ERC20PeriodsMintLimitStorage storage $) {
        assembly {
            $.slot := ERC20PeriodsMintLimitStorageLocation
        }
    }

    function __ERC20PeriodsMintLimit_init(uint256[] memory periods, uint256[] memory limits)
        internal
        onlyInitializing
    {
        uint256 length = limits.length;
        if (length == 0 || length != periods.length) {
            revert ERC20PeriodsMintLimit__InvalidLength();
        }
        uint256 minPeriod = 0;
        uint256 minLimit = 0;
        for (uint256 i = 0; i < length;) {
            (uint256 period, uint256 limit) = (periods[i], limits[i]);
            if (period == 0 || limit == 0) {
                revert TokenBase__NullInput("limits or periods");
            }
            if (period <= minPeriod || limit <= minLimit) {
                revert ERC20PeriodsMintLimit__InvalidLimitData(i);
            }

            (minPeriod, minLimit) = (periods[i], limits[i]);
            unchecked {
                ++i;
            }
        }

        ERC20PeriodsMintLimitStorage storage $ = _getERC20PeriodsMintLimitStorage();
        $.length = length;
        $.limits = limits;
        $.periods = periods;
    }

    function mint(address to, uint256 amount) public virtual override {
        ERC20PeriodsMintLimitStorage storage $ = _getERC20PeriodsMintLimitStorage();

        uint256[] storage _periodStartBlocks = $.periodStartBlocks;
        uint256[] storage _periodCapacities = $.periodCapacities;

        uint256[] memory currentPeriodStartBlocks = periodStartBlocks();
        uint256 length = $.length;
        for (uint256 i = 0; i < length;) {
            (uint256 periodCapacity, uint256 periodStartBlock, uint256 currentPeriodStartBlock) =
                (_periodCapacities[i], _periodStartBlocks[i], currentPeriodStartBlocks[i]);

            // Check if the period has started
            if (periodStartBlock != currentPeriodStartBlock) {
                // Initialize the period start block if not set
                uint256 limit = $.limits[i];
                _periodStartBlocks[i] = currentPeriodStartBlock;
                periodCapacity = limit;

                emit PeriodStarted(currentPeriodStartBlock, limit);
            }
            // Check available capacity
            if (amount > periodCapacity) {
                revert ERC20PeriodsMintLimit__ExceedsPeriodLimit($.periods[i], amount, periodCapacity);
            }
            // Update available capacity
            unchecked {
                _periodCapacities[i] = (periodCapacity - amount);
                ++i;
            }
        }

        // Mint the tokens
        ERC20Base.mint(to, amount);
    }

    function periodBlocks() external view returns (uint256[] memory) {
        return _getERC20PeriodsMintLimitStorage().periods;
    }

    function maxMintPerPeriod() external view returns (uint256[] memory) {
        return _getERC20PeriodsMintLimitStorage().limits;
    }

    function availableMintCapacities() external view returns (uint256[] memory) {
        ERC20PeriodsMintLimitStorage storage $ = _getERC20PeriodsMintLimitStorage();
        uint256[] memory _currentPeriodStartBlock = periodStartBlocks();
        uint256[] memory _periodStartBlock = $.periodStartBlocks;

        uint256 length = $.length;
        uint256[] memory capacities = new uint256[](length);
        for (uint256 i = 0; i < length;) {
            uint256 _periodStart = _periodStartBlock[i];
            if (_periodStart == _currentPeriodStartBlock[i]) {
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

    function periodStartBlocks() public view returns (uint256[] memory) {
        ERC20PeriodsMintLimitStorage storage $ = _getERC20PeriodsMintLimitStorage();
        uint256[] memory _periods = $.periods;

        uint256 length = _periods.length;
        uint256[] memory startBlocks = new uint256[](length);
        uint256 _currentBlock = block.number;
        for (uint256 i = 0; i < length;) {
            uint256 _period = _periods[i];
            uint256 _block;
            unchecked {
                _block = _currentBlock - (_currentBlock & _period);
            }
            startBlocks[i] = _block;
            unchecked {
                ++i;
            }
        }
        return startBlocks;
    }

    function updateMintLimits(uint256[] calldata newLimits) external onlyOwner {
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
            if (newLimit <= minLimit) {
                revert ERC20PeriodsMintLimit__InvalidLimitData(i);
            }
            unchecked {
                ++i;
            }
        }

        // Update limit
        emit MintLimitUpdated($.limits, newLimits);
        $.limits = newLimits;
    }
}
