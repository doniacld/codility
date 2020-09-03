package mix

import (
	"strconv"
)

// Fractions compares fractions with numerators and denominators from two arrays with identical length.
// It returns the max occurrence of fractions with the same GCD without computing the fractions.
func Fractions(X []int, Y []int) int {
	// store fractions into a map of fractions and their occurrences
	frac := make(map[string]int)
	for i, num := range X {
		// compute GCD for numerator and denominator
		g := gcd(num, Y[i])
		num /= g
		Y[i] /= g
		// store into a string the fraction on the GCD
		s := strconv.Itoa(num) + "/" + strconv.Itoa(Y[i])
		// in case it is not seen, set 1
		if frac[s] < 1 {
			frac[s] = 1
		} else { // if seen, just increment the occurrences
			frac[s]++
		}
	}

	// retrieve the max occurrence among the fractions with same GCD
	currentMax := 0
	for _, occurrence := range frac {
		if currentMax < occurrence {
			currentMax = occurrence
		}
	}
	return currentMax
}

// greatest common divisor (GCD) via Euclidean algorithm
func gcd(a, b int) int {
	// compute the modulo of numerator by denominator until we obtain 0
	for b != 0 {
		// store b in a tmp variable
		tmp := b
		// compute the modulo and store in the denominator
		b = a % b
		// store the last value of tmp
		a = tmp
	}
	return a
}
