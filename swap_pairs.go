package main

import "fmt"

func swapPairs(arr []int) []int {

	results := make([]int, 0, len(arr))
	i := 0

	for ; i < len(arr)-1; i += 2 {
		results = append(results, arr[i+1], arr[i])
	}

	if len(arr)%2 != 0 {
		results = append(results, arr[i])
	}

	return results
}

func main() {
	fmt.Println(swapPairs([]int{1, 2, 3, 4}))
	fmt.Println(swapPairs([]int{}))
	fmt.Println(swapPairs([]int{9, 22, 8}))
}
