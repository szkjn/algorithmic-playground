package main

import (
	"fmt"
	"testing"
)

func TestSubarrayDivision(t *testing.T) {
	testCases := []struct {
		s        []int32
		d        int32
		m        int32
		expected int32
	}{
		{
			[]int32{1, 2, 1, 3, 2},
			3,
			2,
			2,
		},
		{
			[]int32{1, 1, 1, 1, 1, 1},
			3,
			2,
			0,
		},
		{
			[]int32{4},
			4,
			1,
			1,
		},
		{
			[]int32{2, 3, 4, 4, 2, 1, 2, 5, 3, 4, 4, 3, 4, 1, 3, 5, 4, 5, 3, 1, 1, 5, 4, 3, 5, 3, 5, 3, 4, 4, 2, 4, 5, 2, 3, 2, 5, 3, 4, 2, 4, 3, 3, 4, 3, 5, 2, 5, 1, 3, 1, 4, 2, 2, 4, 3, 3, 3, 3, 4, 1, 1, 4, 3, 1, 5, 2, 5, 1, 3, 5, 4, 3, 3, 1, 5, 3, 3, 3, 4, 5, 2},
			26,
			8,
			16,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("s=%v, d=%d, m=%d", tc.s, tc.d, tc.m), func(t *testing.T) {
			got := subarrayDivision(tc.s, tc.d, tc.m)
			if got != tc.expected {
				t.Errorf("subarrayDivision(%v, %d, %d) = %d; want %v", tc.s, tc.d, tc.m, got, tc.expected)
			}
		})
	}
}
