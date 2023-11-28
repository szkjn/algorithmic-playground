/*
Write a function that converts a base-10 integer n to its binary representation
and then finds the maximum number of consecutive 1s in this binary form. The
task is to parse the binary string of n and count the longest sequence of
consecutive 1s.

The input is a single integer n, representing the number in base-10.
The output should be a base-10 integer indicating the longest run of consecutive
1s in n's binary representation.

For example, given the input 125, its binary form is 1111101. The longest
sequence of consecutive 1s is 5, so the output is 5.
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
