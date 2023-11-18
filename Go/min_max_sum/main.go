/*
Given five positive integers, find the minimum and maximum values that can be calculated
by summing exactly four of the five integers. Then print the respective minimum and
maximum values as a single line of two space-separated long integers.

Example:
arr = [1, 3, 5, 7, 9]
The minimum sum is 1 + 3 + 5 + 7 = 16 and the maximum sum is 3 + 5 + 7 + 9 = 24.
The function prints: 16 24

Function Parameters:
- arr: an array of 5 integers

Input Format:
- A single line of five space-separated integers.

Output Format:
- Print two space-separated long integers denoting the respective minimum and maximum
values that can be calculated by summing exactly four of the five integers.
(The output can be greater than a 32 bit integer.)
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func miniMaxSum(arr []int32) (int, int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	minSum := 0
	maxSum := 0

	for i := 0; i < 4; i++ {
		minSum += int(arr[i])
	}
	for i := 1; i < 5; i++ {
		maxSum += int(arr[i])
	}

	return minSum, maxSum
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter 5 numbers separated by spaces:")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	input = strings.TrimSpace(input)
	substrings := strings.Split(input, " ")

	var arr []int32

	for _, s := range substrings {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Error while converting string to int:", err)
			return
		}
		arr = append(arr, int32(num))
	}

	fmt.Println(miniMaxSum(arr))
}
