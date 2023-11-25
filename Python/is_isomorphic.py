"""
Given two strings s and t, determine if they are isomorphic. 
Two strings are isomorphic if there is a one-to-one mapping 
possible for every character of the first string to every 
character of the second string.

Example:

> is_isomorphic('abb', 'cdd')
> true // 'a' maps to 'c' and 'b' maps to 'd'

> is_isomorphic('cassidy', '1234567')
> false // 's' cannot have a mapping to both '3' and '4'

> is_isomorphic('cass', '1233')
> true
"""


def is_isomorphic(s: str, t: str) -> bool:
    mapDict = {}
    for i, _ in enumerate(s):
        if s[i] not in mapDict.keys():
            mapDict[s[i]] = t[i]
        else:
            if mapDict[s[i]] != t[i]:
                return False

    return True


print(is_isomorphic("abb", "cdd"))
print(is_isomorphic("cassidy", "1234567"))
print(is_isomorphic("cass", "1233"))
