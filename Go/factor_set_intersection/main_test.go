package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type TestCase struct {
	Input    [][]int32 `json:"input"`
	Expected int32     `json:"expected"`
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

func TestGetTotalX(t *testing.T) {
	testCases, err := loadTestData("../../test_data/factor_set_intersection.json")
	if err != nil {
		t.Fatalf("Failed to load test data: %v", err)
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("a=%v,b=%v", tc.Input[0], tc.Input[1]), func(t *testing.T) {
			got := getTotalX(tc.Input[0], tc.Input[1])
			if got != tc.Expected {
				t.Errorf("getTotalX(%v, %v) = %d; want %d", tc.Input[0], tc.Input[1], got, tc.Expected)
			}
		})
	}
}
