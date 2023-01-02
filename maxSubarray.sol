// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

contract main {
    function sum(int256[] memory arr) private pure returns (int256) {
        int256 res = 0;
        for (uint256 i = 0; i < arr.length; i++) {
            res += arr[i];
        }
        return res;
    }

    function maxSubarray(int256[] calldata arr, uint256 n)
        public
        pure
        returns (int256[] memory)
    {
        int256[] memory subarray;
        for (uint256 i = 0; i < arr.length - n + 1; i++) {
            int256[] memory current = arr[i:i + n];
            if (sum(subarray) < sum(current)) {
                subarray = current;
            }
        }
        return subarray;
    }
}
