// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

contract main {
    function swapPairs(int256[] calldata arr)
        public
        pure
        returns (int256[] memory)
    {
        require(arr.length % 2 == 0, "Error: length of array must be even.");
        int256[] memory results = new int256[](arr.length);

        for (uint256 i = 0; i < arr.length; i += 2) {
            int256 firstEl = arr[i];
            int256 secondEl = arr[i + 1];
            results[i] = (secondEl);
            results[i + 1] = (firstEl);
        }

        return results;
    }
}
