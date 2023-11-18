/*
Julius Caesar protected his confidential information by encrypting it using a cipher. Caesar's cipher shifts each letter by a number of letters. If the shift takes you past the end of the alphabet, just rotate back to the front of the alphabet. In the case of a rotation by 3, w, x, y and z would map to z, a, b and c.

Original alphabet:      abcdefghijklmnopqrstuvwxyz
Alphabet rotated +3:    defghijklmnopqrstuvwxyzabc

Complete the caesarCipher function in the editor below.

Parameters:
- string s: clear text
- int k: the alphabet rotation factor
Returns:
- string: the encrypted string
Input Format
- The first line contains the unencrypted string, s.
- The second line contains k, the number of letters to rotate the alphabet by.

Constraints
- 0 <= k <= 100
- s is a valid ASCII string without any spaces.

Sample Input
	> middle-Outz
	> 2

Sample Output
	> okffng-Qwvb
*/

package main

import (
	"fmt"
	"strings"
)

func caesarCipher(s string, k int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	res := ""

	for _, char := range s {

		str := string(char)
		isUpper := false
		var idx int

		if strings.Contains(alphabet, strings.ToLower(str)) {
			if strings.ToUpper(str) == str {
				isUpper = true
			}
			idx = int(strings.Index(alphabet, strings.ToLower(str)))
			moduloIdx := (idx + k) % len(alphabet)
			cipheredLetter := string(alphabet[moduloIdx])

			if isUpper {
				res += strings.ToUpper(cipheredLetter)
			} else {
				res += cipheredLetter
			}
		} else {
			res += str
		}
	}
	return res
}

func main() {
	var delta int
	var input string
	fmt.Println("Provide string to cipher:")
	fmt.Scanf("%s\n", &input)
	fmt.Println("Provide cipher's delta (int):")
	fmt.Scanf("%d\n", &delta)

	res := caesarCipher(input, delta)
	fmt.Println(res)
}
