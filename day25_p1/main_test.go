package main

import (
	"testing"
)

func TestLoop(t *testing.T) {
	cardKey := 5764801
	doorKey := 17807724

	_, doorLoop := findLoopSize(cardKey, doorKey, 7)

	k := findEncryptionKey(1, doorLoop, cardKey)
	expectedKey := 14897079

	if k != expectedKey {
		t.Errorf("Got %d, expected %d", k, expectedKey)
	}
}
