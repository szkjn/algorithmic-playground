/*
Given a base-10 integer, n, convert it to binary (base-2). Then find and print the
base-10 integer denoting the maximum number of consecutive 1's in n's binary
representation.

Example:
- The binary representation of 125 is 1111101. In base-10, there are 5 and 1
consecutive ones in two groups. Print the maximum, 5.

Input Format:
- A single integer, n.

Output Format:
- Print a single base-10 integer that denotes the maximum number of consecutive 1's
in the binary representation of n.

Sample Input 1:
5
Sample Output 1:
1
Sample Input 2:
13
Sample Output 2:
2
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func countMaxConsecutiveOnes(n string) int {
	maxCount := 0
	tmpCount := strings.Split(n, "0")
	for _, el := range tmpCount {
		if len(el) > maxCount {
			maxCount = len(el)
		}
	}
	return maxCount
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int64(nTemp)
	nBase2 := strconv.FormatInt(n, 2)
	fmt.Println(nBase2)

	fmt.Println(countMaxConsecutiveOnes(nBase2))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
