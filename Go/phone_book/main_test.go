package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreatePhoneBook(t *testing.T) {
	testCases := []struct {
		arr      []string
		expected map[string]string
	}{
		{
			[]string{"toto 01234567", "lulu 55553333", "fifi 22222222"},
			map[string]string{"toto": "01234567", "lulu": "55553333", "fifi": "22222222"},
		},
		{
			[]string{"gopher 911", "gozilla 666"},
			map[string]string{"gopher": "911", "gozilla": "666"},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("arr=%v", tc.arr), func(t *testing.T) {
			got := createPhoneBook(tc.arr)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("createPhoneBook(%v) = %v; want %v", tc.arr, got, tc.expected)
			}
		})
	}
}

func TestGetPhone(t *testing.T) {
	testCases := []struct {
		n         string
		phoneBook map[string]string
		expected  string
	}{
		{
			"toto",
			map[string]string{"toto": "12345678", "fifi": "22222222"},
			"toto=12345678",
		},
		{
			"jesus",
			map[string]string{"toto": "12345678", "fifi": "22222222"},
			"Not found",
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("n=%s, phoneBook=%s", tc.n, tc.phoneBook), func(t *testing.T) {
			got := getPhone(tc.n, tc.phoneBook)
			if got != tc.expected {
				t.Errorf("getPhone(%v) = %v; want %v", tc.n, got, tc.expected)
			}
		})
	}
}
