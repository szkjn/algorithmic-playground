package main

import (
	"fmt"
	"testing"
)

func TestGetTotalX(t *testing.T) {
	testCases := []struct {
		a        []int32
		b        []int32
		expected int32
	}{
		{
			[]int32{2, 4},
			[]int32{16, 32, 96},
			3,
		},
		{
			[]int32{3, 4},
			[]int32{24, 48},
			2,
		},
		{
			[]int32{1},
			[]int32{72, 48},
			8,
		},
		{
			[]int32{2, 3, 6},
			[]int32{42, 84},
			2,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("a=%v,b=%v", tc.a, tc.b), func(t *testing.T) {
			got := getTotalX(tc.a, tc.b)
			if got != tc.expected {
				t.Errorf("getTotalX(%v, %v) = %d; want %d", tc.a, tc.b, got, tc.expected)
			}
		})
	}
}
