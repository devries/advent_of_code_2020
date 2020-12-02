package main

import (
	"testing"
)

func TestValidator(t *testing.T) {
	tests := []struct {
		Entry string
		Valid bool
	}{
		{"1-3 a: abcde", true},
		{"1-3 b: cdefg", false},
		{"2-9 c: ccccccccc", true},
	}

	for _, test := range tests {
		result, err := validPassword(test.Entry)
		if err != nil {
			t.Errorf("Got error: %s", err)
		}

		if result != test.Valid {
			t.Errorf("For %s got %t, expected %t", test.Entry, result, test.Valid)
		}
	}
}
