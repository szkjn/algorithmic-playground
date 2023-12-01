"""
Given an array of integers where each number represents a sock's color, the 
function calculates the number of matching sock pairs. Each pair consists 
of two socks of the same color. The function takes the total number of socks 
and an array of their colors as input, and returns the total number of 
matching pairs.

Input Parameters:
- int n: Number of socks in the pile.
- int arr[n]: Array representing the colors of each sock.

Output:
- int: The total number of matching sock pairs.

Constraints:
- 1 ≤ n ≤ 100
- 1 ≤ arr[i] ≤ 100 where 0 ≤ i < n
"""

import math
import os
import random
import re
import sys


def pair_counter(n, arr):
    countMap = {}

    for el in arr:
        if el in countMap:
            countMap[el] += 1
        else:
            countMap[el] = 1

    countPairs = 0

    for el in countMap.values():
        countPairs += el // 2

    return int(countPairs)


if __name__ == "__main__":
    n = int(input().strip())
    arr = list(map(int, input().rstrip().split()))
    result = pair_counter(n, arr)
    print(str(result))
