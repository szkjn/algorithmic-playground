"""
Interview question of the week by Cassidy Williams (issue #274) :
Given a list of strings arr, and a max size n, return a new
list where the strings (from left to right) are joined together
with a space, so that each new string is less than or equal to
the max size.
"""


def combine_strings(arr: list, n: int) -> list:
    res = []
    new_str = ""

    while len(arr) > 0:
        if new_str == "":
            new_str = arr[0]
            arr.pop(0)
            if len(arr) == 0:
                break

        if len(new_str) + len(arr[0]) < n:
            new_str += " " + arr[0]
            arr.pop(0)
            if len(arr) == 0:
                break

        if len(new_str) == n or (len(new_str) < n and len(new_str) + len(arr[0]) >= n):
            res.append(new_str)
            new_str = ""

        if len(arr) == 1:
            res.append(arr[0])
            arr.pop(0)
            if len(arr) == 0:
                break

    if len(new_str) > 0:
        res.append(new_str)

    return res
