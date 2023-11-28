package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type TestCase struct {
	n        string `json:n`
	expected int    `json:expected`
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

func TestCountMaxConsecutiveOnes(t *testing.T) {
	testCases, err := loadTestData("../../test_data/max_consecutive_ones.json")
	if err != nil {
		t.Fatalf("Failed to load test data: %v", err)
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("n=%s", tc.n), func(t *testing.T) {
			got := countMaxConsecutiveOnes(tc.n)
			if got != tc.expected {
				t.Errorf("countMaxConsecutiveOnes(%s) = %d; want %d", tc.n, got, tc.expected)
			}
		})
	}
}
