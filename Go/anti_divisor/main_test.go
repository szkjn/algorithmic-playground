package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/szkjn/algorithmic-playground/Go/test_utils"
)

type TestCase struct {
	N        int   `json:"input"`
	Expected []int `json:"expected"`
}

func TestAntiDivisors(t *testing.T) {
	testCases, err := test_utils.LoadTestData("../../test_data/anti_divisor_test_data.json")
	if err != nil {
		t.Fatalf("Failed to load test data: %v", err)
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("n=%d", tc.N), func(t *testing.T) {
			got := antiDivisors(tc.N)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Errorf("antiDivisors(%d) = %v; want %v", tc.N, got, tc.Expected)
			}
		})
	}
}
