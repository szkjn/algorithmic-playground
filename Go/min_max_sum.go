package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func miniMaxSum(arr []int32) {
	// Write your code here
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

	fmt.Printf("%d %d\n", minSum, maxSum)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter numbers separated by spaces:")
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

	miniMaxSum(arr)
}
