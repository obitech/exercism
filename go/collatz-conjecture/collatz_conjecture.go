// Package collatzconjecture calculates the Collatz Conjecture.
package collatzconjecture

import (
	"fmt"
)

func collatz(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return 1 + collatz(n/2)
	}
	if n%2 != 0 {
		return 1 + collatz(3*n+1)
	}
	return 0
}

// CollatzConjecture calculates the number of steps taken in the Collatz
// Conjecture until 1 is reached.
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("invalid argument: %d", n)
	}
	return collatz(n), nil
}
