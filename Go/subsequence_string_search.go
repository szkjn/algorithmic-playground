/*
We say that a string contains the word hackerrank if a subsequence
of its characters spell the word hackerrank. Remember that a subsequence
maintains the order of characters selected from a sequence.

For each query, print YES on a new line if the string contains hackerrank,
otherwise, print NO.

Function Description
- Complete the hackerrankInString function in the editor below.

Parameter:
- string s: a string

Returns:
- string: YES or NO

Sample Input:

hereiamstackerrank 	// should return YES
hackerworld 		// should return NO
hhaacckkekraraannk 	// should return YES
*/

package main

import "fmt"

func hackerrankInString(s string) string {
	subStr := "hackerrank"
	idx := 0

	for i, _ := range s {

		if s[i] == subStr[idx] {
			idx++

			if idx == len(subStr) {
				return "YES"
			}
		}
	}
	return "NO"
}

func main() {
	fmt.Println(hackerrankInString("hereiamstackerrank"))
	fmt.Println(hackerrankInString("hackerworld"))
	fmt.Println(hackerrankInString("hhaacckkekraraannk"))
}
