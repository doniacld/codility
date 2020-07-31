package brackets

import "strings"

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
	var isOk int
	bracket := make([]string, 0)

	a := strings.Split(S, "")
	for _, b := range a {
		if b == bracketOpen || b == hookOpen || b == curlyOpen {

		}
	}
	return isOk
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
