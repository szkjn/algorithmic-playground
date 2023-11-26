"""
Interview question of the week by Cassidy Williams (issue #260) :

Given a list, swap every two adjacent nodes. Something to think 
about: How would your code change if this were a linked list, 
versus an array?
"""


def swap_pairs(arr: list) -> list:
    results = []
    i = 0

    while i < len(arr):
        if i + 1 < len(arr):
            results.append(arr[i + 1])
        results.append(arr[i])
        i += 2

    return results
