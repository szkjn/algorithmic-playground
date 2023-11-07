/*
Interview question of the week by Cassidy Williams (issue #282) :

Given a number, sum every second digit in that number.
Example:
> sumEveryOther(548915381)
> 26 4+9+5+8
> sumEveryOther(10)
> 0
> sumEveryOther(1010.11)
> 1 // 0+0+1
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumEveryOther(n float64) int {
	numberStr := strings.Replace(strconv.FormatFloat(n, 'f', -1, 64), ".", "", -1)
	sum := 0

	for i := 1; i < len(numberStr); i += 2 {
		digit, err := strconv.Atoi(string(numberStr[i]))
		if err != nil {
			fmt.Println("Error converting digit to int:", err)
			return 0
		}
		sum += digit
	}
	return sum
}

func main() {
	fmt.Println(sumEveryOther(548915381)) // Should print the sum of every other digit starting from the second
	fmt.Println(sumEveryOther(10))        // Should print 0, as there's only one digit after the first
	fmt.Println(sumEveryOther(1010.11))   // Should print the sum of 0 and 1 which are after the decimal point
}
