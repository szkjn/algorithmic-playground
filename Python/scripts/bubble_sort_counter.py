"""
Implement a Bubble Sort algorithm to sort an array in ascending order. The program 
should count the total number of swaps needed to sort the array. After sorting, it 
prints the number of swaps, the first element, and the last element of the sorted 
array.

Input Format:
- Line 1: Integer n representing the number of elements in the array.
- Line 2: n space-separated integers representing the array elements.

Output Format:
- Line 1: "Array is sorted in <nbrSwaps> swaps."
- Line 2: "First Element: <minVal>"
- Line 3: "Last Element: <maxVal>"
"""
import math
import os
import random
import re
import sys


if __name__ == "__main__":
    n = int(input().strip())

    arr = list(map(int, input().rstrip().split()))

    count = 0

    for i in range(len(arr)):
        for j in range(len(arr) - 1):
            if arr[j] > arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
                count += 1

    print(f"Array is sorted in {count} swaps.")
    print(f"First Element: {min(arr)}")
    print(f"Last Element: {max(arr)}")
