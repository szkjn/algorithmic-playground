package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type TestCase struct {
	n        int32   `json:"n"`
	k        int32   `json:"k"`
	arr      []int32 `json:"arr"`
	expected int32   `json:"expected"`
}

func loadTestData(filePath string) ([]TestCase, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var testCases []TestCase
	err = json.NewDecoder(file).Decode(&testCases)
	if err != nil {
		return nil, err
	}

	return testCases, nil
}

func TestDivisibleSumPairs(t *testing.T) {
	testCases, err := loadTestData("../../test_data/divisible_sum_pairs.json")
	if err != nil {
		t.Fatalf("Failed to load test data: %v", err)
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("n=%d,k=%d,arr=%v", tc.n, tc.k, tc.arr), func(t *testing.T) {
			got := divisibleSumPairs(tc.n, tc.k, tc.arr)
			if got != tc.expected {
				t.Errorf("divisibleSumPairs(%d, %d, %v) = %d; want %d", tc.n, tc.k, tc.arr, got, tc.expected)
			}
		})
	}
}
