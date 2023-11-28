/*
Write a function to count pairs in an array whose sum is divisible by a given integer.
Given an array ar of integers and a positive integer k, the function should determine
the number of pairs (i, j) where i < j and ar[i] + ar[j] is divisible by k. The input
includes the array length n, array ar, and divisor k. The function returns the count
of such pairs.
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

func divisibleSumPairs(n, k int32, ar []int32) int32 {
	var res int32
	for i := 0; i < int(n); i++ {
		for j := i + 1; j < int(n); j++ {
			if (ar[i]+ar[j])%k == 0 {
				res++
			}
		}
	}
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	arTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var ar []int32

	for i := 0; i < int(n); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := divisibleSumPairs(n, k, ar)
	fmt.Println(result)
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
