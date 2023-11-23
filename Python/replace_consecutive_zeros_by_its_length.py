"""
Given a string of any length which contains only digits from 0 to 9, 
replace each consecutive run of the digit 0 with its length.

Example:

> replaceZeros('1234500362000440')
> 1234523623441

> replaceZeros('123450036200044')
> 123452362344

> replaceZeros('000000000000')
> 12

> replaceZeros('123456789')
> 123456789
"""


def replaceZeros(digits: str) -> str:

    res = ""
    count = 0

    for i, digit in enumerate(digits):
        if digit == "0":
            count += 1
            if i == len(digits) - 1:
                res += str(count)
        else:
            if count > 0:
                res += str(count)
                count = 0
            res += digit

    return res
