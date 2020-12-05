package main

import (
	"testing"
)

func TestSeatId(t *testing.T) {
	tests := []struct {
		BoardingPass string
		SeatId       int
	}{
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}

	for _, test := range tests {
		result, err := getSeatId(test.BoardingPass)
		if err != nil {
			t.Errorf("encoundered error: %s", err)
		}
		if result != test.SeatId {
			t.Errorf("For %s expected %d got %d", test.BoardingPass, test.SeatId, result)
		}
	}
}
