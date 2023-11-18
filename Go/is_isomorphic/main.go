/*
Given two strings s and t, determine if they are isomorphic.
Two strings are isomorphic if there is a one-to-one mapping
possible for every character of the first string to every
character of the second string.

Example:

> isIsomorphic('abb', 'cdd')
> true // 'a' maps to 'c' and 'b' maps to 'd'

> isIsomorphic('cassidy', '1234567')
> false // 's' cannot have a mapping to both '3' and '4'

> isIsomorphic('cass', '1233')
> true
*/

package main

import "fmt"

func isIsomorphic(s, t string) bool {
	m := make(map[string]string)

	for i := 0; i < len(s); i++ {
		sChar := fmt.Sprintf("%c", s[i])
		tChar := fmt.Sprintf("%c", t[i])
		if _, ok := m[sChar]; !ok {
			m[sChar] = tChar
		} else if m[sChar] != tChar {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("abb", "cdd"))
	fmt.Println(isIsomorphic("cassidy", "1234567"))
	fmt.Println(isIsomorphic("cass", "1233f"))
}
