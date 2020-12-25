package main

import "fmt"

func main() {
	cardKey := 10441485
	doorKey := 1004920

	_, doorLoop := findLoopSize(cardKey, doorKey, 7)

	k := findEncryptionKey(1, doorLoop, cardKey)

	fmt.Println(k)
}

func findLoopSize(a int, b int, subject int) (int, int) {
	aFound := false
	bFound := false
	var aLoop int
	var bLoop int

	i := 0
	v := 1
	for {
		v *= subject
		v %= 20201227
		i++

		if v == a {
			aLoop = i
			aFound = true
		}

		if v == b {
			bLoop = i
			bFound = true
		}

		if aFound && bFound {
			return aLoop, bLoop
		}
	}
}

func findEncryptionKey(start int, nLoop int, subject int) int {
	result := start

	for i := 0; i < nLoop; i++ {
		result *= subject
		result %= 20201227
	}

	return result
}
