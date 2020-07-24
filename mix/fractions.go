package mix

import (
	"strconv"
)

// Fractions compares fractions with numerators and denominators from two arrays with identical length.
// It returns the max occurrence of fractions with the same GCD without computing the fractions.
func Fractions(X []int, Y []int) int {
	frac := make(map[string]int)
	for i, num := range X {
		// compute GCD for numerator and denominator
		g := gcd(num, Y[i])
		num /= g
		Y[i] /= g
		s := strconv.Itoa(num) + "/" + strconv.Itoa(Y[i])
		if frac[s] < 1 {
			frac[s] = 1
		} else {
			frac[s]++
		}
	}

	// retrieve the max occurrence among the fractions with same GCD
	currentMax := 0
	for value := range frac {
		if currentMax < frac[value] {
			currentMax = frac[value]
		}
	}
	return currentMax
}

// greatest common divisor (GCD) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
