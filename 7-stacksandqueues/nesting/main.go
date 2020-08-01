package nesting

import (
	"strings"
)

const (
	openBracket  = "("
)

func Nesting(S string) int {
	if len(S)%2 != 0 {
		return 0
	}

	s := strings.Split(S, "")

	stack := make([]string, 0)
	for _, value := range s {
		switch {
		case isOpenBracket(value):
			stack = append(stack, value)
		case len(stack) > 0 && !isOpenBracket(value):
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

func isOpenBracket(value string) bool {
	return value == openBracket
}
