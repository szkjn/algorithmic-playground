"""
Write a function that converts a base-10 integer n to its binary representation
and then finds the maximum number of consecutive 1s in this binary form. The
task is to parse the binary string of n and count the longest sequence of
consecutive 1s.

The input is a single integer n, representing the number in base-10.
The output should be a base-10 integer indicating the longest run of consecutive
1s in n's binary representation.

For example, given the input 125, its binary form is 1111101. The longest
sequence of consecutive 1s is 5, so the output is 5.
"""


def count_max_consecutive_ones(n: str) -> int:
    max_count = 0
    tmp_count = n.split("0")
    for el in tmp_count:
        if len(el) > max_count:
            max_count = len(el)
    return max_count
