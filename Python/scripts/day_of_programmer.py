"""
Create a function that calculates the date of the 256th day of a given year in 
Russia, considering the transition from the Julian to Gregorian calendar in 1918. 
Leap years in the Julian calendar are divisible by 4, while in the Gregorian 
calendar, leap years are either divisible by 400, or divisible by 4 but not by 
100. The function should return the date in the format dd.mm.yyyy.

Parameters:
- year: an integer representing the year (1700 ≤ year ≤ 2700)

Input: 
- A single integer year.

Output: 
- A string representing the date of the 256th day of year in the format dd.mm.yyyy.
"""

import math
import os
import random
import re
import sys


def day_of_programmer(year):
    if year == 1918:
        # Special case for the year of the Russian calendar switch
        return "26.09.1918"
    
    if year < 1918:
        # Julian calendar
        is_leap = year % 4 == 0
    else:
        # Gregorian calendar
        is_leap = (year % 400 == 0) or (year % 4 == 0 and year % 100 != 0)

    if is_leap:
        return f"12.09.{year}"
    else:
        return f"13.09.{year}"


if __name__ == '__main__':

    year = int(input().strip())
    result = day_of_programmer(year)
    print(result)
