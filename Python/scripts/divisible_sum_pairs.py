"""
Write a function to count pairs in an array whose sum is divisible by a given integer.
Given an array ar of integers and a positive integer k, the function should determine
the number of pairs (i, j) where i < j and ar[i] + ar[j] is divisible by k. The input
includes the array length n, array ar, and divisor k. The function returns the count
of such pairs.
"""


def divisible_sum_pairs(n: int, k: int, arr: list) -> int:
    res = 0
    for i in range(n):
        for j in range(i + 1, n):
            if (arr[i] + arr[j]) % k == 0:
                res += 1

    return res
