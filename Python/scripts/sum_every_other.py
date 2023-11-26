"""
Interview question of the week by Cassidy Williams (issue #282) :

Given a number, sum every second digit in that number.
Example:
> sumEveryOther(548915381)
> 26 4+9+5+8
> sumEveryOther(10)
> 0
> sumEveryOther(1010.11)
> 1 // 0+0+1
"""


def sum_every_other(n: float) -> int:
    n = str(n).replace(".", "")
    result = [int(n[char]) for char in range(1, len(n), 2)]
    return sum(result)
