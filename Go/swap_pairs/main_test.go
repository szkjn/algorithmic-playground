package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	testCases := []struct {
		arr      []int
		expected []int
	}{
		{arr: []int{1, 2, 3, 4}, expected: []int{2, 1, 4, 3}},
		{arr: []int{}, expected: []int{}},
		{arr: []int{9, 22, 8}, expected: []int{22, 9, 8}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("arr=%v", tc.arr), func(t *testing.T) {
			got := swapPairs(tc.arr)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("swapPairs(%v) = %v; want %v", tc.arr, got, tc.expected)
			}
		})
	}
}
