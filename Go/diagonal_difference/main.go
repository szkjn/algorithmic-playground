/*
Given a square matrix, calculate the absolute difference between the sums of its diagonals.

For example:
1 2 3
4 5 6
9 8 9
The left-to-right diagonal = 1 + 5 + 9 = 15.
The right to left diagonal = 3 + 5 + 9 = 17.
Their absolute difference is |15 - 17| = 2.

Parameter:
- int arr[n][m]: an array of integers

Return:
- int: the absolute diagonal difference

Input Format:
- The first line contains a single integer, n, the number of rows and columns in the square matrix .
- Each of the next n lines describes a row, arr[i], and consists of n space-separated integers arr[i][j].

Constraints:
- -100 <= arr[i][j] <= 100
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

func diagonalDifference(arr [][]int32) int32 {

	var firstDiag, secondDiag int32 = 0, 0

	for i := 0; i < len(arr); i++ {
		firstDiag += arr[i][i]
		secondDiag += arr[i][len(arr)-i-1]
	}

	res := firstDiag - secondDiag
	if res < 0 {
		return -res
	}
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	fmt.Println("Enter dimension of the square (int):")
	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)
	fmt.Printf("Enter %d rows of %d numbers separated by a space:\n", n, n)

	var arr [][]int32
	for i := 0; i < int(n); i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(n) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := diagonalDifference(arr)
	fmt.Printf("Results: %v\n", result)
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
