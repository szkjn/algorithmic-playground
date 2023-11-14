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
	"os"
	"strconv"
)

func diagonalDifference(arr [][]int32) int32 {
	var res int32
	res = 0

	fmt.Println(arr[0][0])
	fmt.Println(arr[0][1])
	fmt.Println(arr[0][2])
	fmt.Println(arr[1][0])
	fmt.Println(arr[1][1])
	fmt.Println(arr[1][2])

	// for i := 0; i < len(arr); i++ {
	// 	res += arr[i] + arr[i][i]
	// }
	return res
}

func main() {
	var n int
	fmt.Print("Number of dimensions : ")
	fmt.Scanf("%d\n", &n)

	scanner := bufio.NewScanner(os.Stdin)

	var sqMatrix [][]int32
	var sqRow []int32

	for i := 0; i < int(n); i++ {
		fmt.Printf("Row #%d: ", i)

		scanner.Scan()
		sqStr := scanner.Text()
		sqItem, err := strconv.Atoi(sqStr)
		if err != nil {
			panic(err)
		}
		sqRow = append(sqRow, int32(sqItem))
	}

	if len(sqRow) != int(n) {
		panic("Bad input.")
	}

	sqMatrix = append(sqMatrix, sqRow)

	res := diagonalDifference(sqMatrix)
	fmt.Println(res)
}
