package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type TestCase struct {
	S        string `json:"s"`
	K        int    `json:"k"`
	Expected string `json:"expected"`
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

func TestCaesarCipher(t *testing.T) {
	testCases, err := loadTestData("../../test_data/caesar_cipher.json")
	if err != nil {
		t.Fatalf("Failed to load test data: %v", err)
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("s=%v,k=%d", tc.S, tc.K), func(t *testing.T) {
			got := caesarCipher(tc.S, tc.K)
			if got != tc.Expected {
				t.Errorf("caesarCipher(%v, %d) = %v; want %v", tc.S, tc.K, got, tc.Expected)
			}
		})
	}
}
