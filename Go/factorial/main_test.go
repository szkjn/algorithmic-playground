package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type TestCase struct {
	N        int32 `json:"n"`
	Expected int32 `json:"expected"`
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
func TestFactorial(t *testing.T) {
	testCases, err := loadTestData("../../test_data/factorial.json")
	if err != nil {
		t.Fatalf("Failed to load test data: %v", err)
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("n=%d", tc.N), func(t *testing.T) {
			got := factorial(tc.N)
			if got != tc.Expected {
				t.Errorf("factorial(%d) = %d; want %d", tc.N, got, tc.Expected)
			}
		})
	}
}
