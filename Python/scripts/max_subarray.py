"""
Interview question of the week by Cassidy Williams (issue #281) :

Given an array of integers arr and an integer n, return a subarray of
arr of length n where the sum is the largest. Make sure you maintain
the order of the original array, and if n is greater than arr.length,
you can choose what you want to return.
"""


def max_subarray(arr: list, n: int) -> list:
    subarray = []

    for i in range(len(arr) - n + 1):
        if sum(subarray) < sum(arr[i : i + n]):
            subarray = arr[i : i + n]

    return subarray
