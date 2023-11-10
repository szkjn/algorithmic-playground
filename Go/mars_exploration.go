/*
Letters in some of the SOS messages are altered by cosmic radiation during
transmission. Given the signal received by Earth as a string, , determine
how many letters of the SOS message have been changed by radiation.

Parameter:
- string s: the string as received on Earth

Returns:
- int: the number of letters changed during transmission

Input Format:
- There is one line of input: a single string, s.

Constraints:
- 1 <= length of s <= 99
- length of s module 3 = 0
- s will contain only uppercase English letters, ascii[A-Z].

Examples:
- Input: SOSSPSSQSSOR
- Output: 3

- Input: SOSSOT
- Output: 1

- Input: SOSSOSSOS
- Output: 0
*/

package main

import (
	"fmt"
	"strings"
)

func marsExploration(s string) int32 {
	res := 0
	countMsg := len(s) / 3
	originalMessage := strings.Repeat("SOS", countMsg)

	for idx, _ := range s {
		if s[idx] != originalMessage[idx] {
			res++
		}
	}

	return int32(res)
}

func main() {
	var s string
	fmt.Scanf("%s\n", &s)
	result := marsExploration(s)
	fmt.Printf("%d\n", result)
}
