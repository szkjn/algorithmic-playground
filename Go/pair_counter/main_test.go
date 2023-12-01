package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type TestCase struct {
	N        int32   `json:"n"`
	Arr      []int32 `json:"arr"`
	Expected int32   `json:"expected"`
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

func TestPairCounter(t *testing.T) {
	testCases, err := loadTestData("../../test_data/pair_counter.json")
	if err != nil {
		t.Fatalf("Failed to load test data: %v", err)
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("n=%d, arr=%v", tc.N, tc.Arr), func(t *testing.T) {
			got := pairCounter(tc.N, tc.Arr)
			if got != tc.Expected {
				t.Errorf("pairCounter(%d, %v) = %d; want %d", tc.N, tc.Arr, got, tc.Expected)
			}
		})
	}
}
