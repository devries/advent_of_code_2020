package utils

// Permutations of an array of Integers
// Uses Heap's Algorithm (thanks wikipedia)
func IntPermutations(a []int) <-chan []int {
	ch := make(chan []int)
	go func() {
		k := len(a)
		intPermutationsRecursor(k, a, ch)
		close(ch)
	}()

	return ch
}

func intPermutationsRecursor(k int, a []int, ch chan<- []int) {
	if k == 1 {
		output := make([]int, len(a))
		copy(output, a)
		ch <- output
	} else {
		intPermutationsRecursor(k-1, a, ch)

		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				a[i], a[k-1] = a[k-1], a[i]
			} else {
				a[0], a[k-1] = a[k-1], a[0]
			}
			intPermutationsRecursor(k-1, a, ch)
		}
	}
}
