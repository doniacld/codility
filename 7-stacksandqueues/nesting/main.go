package nesting

import (
	"strings"
)

const (
	openBracket  = "("
)

// Nesting returns 1 if the brackets are welled nested otherwise returns 0
func Nesting(S string) int {
	if len(S)%2 != 0 {
		// early exit due to odd length
		return 0
	}

	s := strings.Split(S, "")

	stack := make([]string, 0)
	for _, value := range s {
		isOpenBracket := value == openBracket
		switch {
		case isOpenBracket:
			stack = append(stack, value)
		case len(stack) > 0 && !isOpenBracket:
			stack = stack[:len(stack)-1]
		default:
			return 0
		}
	}

	if len(stack) != 0 {
		return 0
	} else {
		return 1
	}
}
