"""
Interview question of the week by Cassidy Williams (issue #260) :

Given a list, swap every two adjacent nodes. Something to think 
about: How would your code change if this were a linked list, 
versus an array?
"""

def swap_pairs(arr: list) -> list:

    results = []

    assert len(arr) % 2 == 0

    while len(arr) > 0:
        results.append(arr.pop(1))
        results.append(arr.pop(0))

    print(results)

swap_pairs([1, 2, 3, 4])
swap_pairs([])
