package brackets

import (
	"strings"
)

const (
	bracketOpen  = "("
	bracketClose = ")"
	hookOpen     = "["
	hookClose    = "]"
	curlyOpen    = "{"
	curlyClose   = "}"
)

// Brackets returns 1 if the brackets are correctly nested otherwise 0
func Brackets(S string) int {
	if len(S)%2 != 0 {
		return 0
	}

	s := strings.Split(S, "")
	bracket := make([]string, 0)

	for _, value := range s {
		isOpen := value == bracketOpen || value == hookOpen || value == curlyOpen
		if isOpen {
			bracket = append(bracket, value)
		} else {
			if len(bracket) > 0 && isValidPair(bracket[len(bracket)-1], value) {
				bracket = bracket[:len(bracket)-1]
			} else {
				return 0
			}
		}
	}
	if len(bracket) != 0 {
		return 0
	} else {
		return 1
	}
}

func isValidPair(left, right string) bool {
	if left == bracketOpen && right == bracketClose {
		return true
	}
	if left == hookOpen && right == hookClose {
		return true
	}
	if left == curlyOpen && right == curlyClose {
		return true
	}
	return false
}
