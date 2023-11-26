"""
Create a function replace_zeros that transforms a string containing
digits (0-9) by replacing consecutive zeros with their count. For 
instance, '1234500362000440' becomes '1234523623441', replacing runs 
of zeros with their length. The function should handle any length of 
input and any combination of digits, including strings without zeros 
or consisting only of zeros.
"""


def replace_zeros(digits: str) -> str:
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
