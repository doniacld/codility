package main

import (
	"fmt"
	"strconv"
)

/*
A message containing letters from A-Z is being encoded to numbers using the following mapping:

'A' -> 1
'B' -> 2
...
'Z' -> 26

Given a non-empty string containing only digits, determine the total number of ways to decode it.
*/

// numDecodings decodes a string and give the number of the way of decoding
// non recursive waty
func numDecodings(s string) int {
	slen := len(s)

	// early exit
	if s[0] == '0' {
		return 0
	}

	decodePos := make([]int, slen+1)
	decodePos[0], decodePos[1] = 1, 1

	for i := 2 ; i < slen+1 ; i++ {
		if s[i-1] != '0' {
			decodePos[i] += decodePos[i-1]
		}
		twoDigit, _ := strconv.Atoi(s[i-2:i])
		fmt.Println(twoDigit)
		if twoDigit >= 10 && twoDigit <=26 {
			decodePos[i] += decodePos[i-2]
			fmt.Println(decodePos)
		}
	}

	return decodePos[len(s)]
}

/*
// recursive solution
func numDecodings(s string) int {
    if len(s) == 0 { return 1 }
    if len(s) == 1 && s[0] != '0' { return 1}
    if s[0] == '0' { return 0 }
    x := numDecodings(s[1:])
    if s[:2] <= "26" {
        x = x + numDecodings(s[2:])
    }
    return x
}

*/

func main() {
	fmt.Println(numDecodings("256"))
}