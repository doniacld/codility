package mix

import (
	"fmt"
	"regexp"
	"strings"
)

// Find minimum number of words in a sentence, given a document, using white-space as separator for words and (!, ?, .)
// as the separator for sentence. There can be multiple white-spaces between two words.
// CountNumberWords returns the number of words
func CountNumberWords(file string) int {

	// open a file
	// data, err := ioutil.ReadFile(file)
	// check error
	// s := string(data)
	// store the file in a string
	fileString := "I am the one blabla. Coucou! titi ? toto? bibi ! Caca  ."
	f := strings.Split(fileString, ".")
	f = strings.Split(strings.Join(f, " "), "!")
	f = strings.Split(strings.Join(f, " "), ".")
	f = strings.Split(strings.Join(f, " "), "?")
	fmt.Println(f)
	res := computeWordsInSentence(strings.Join(f, " "))
	return res
}

func isEndChar(s string) bool {
	if s == "!" || s == "?" || s == "." {
		return true
	}
	return false
}

func containsEndChar(s string) bool {
	if strings.Contains(s, "!") || strings.Contains(s, "?") || strings.Contains(s, ".") {
		return true
	}
	return false
}

func computeWordsInSentence(fileString string) int {
	re, err := regexp.Compile(`[\S]+`)
	if err != nil {
		return -1
	}

	results := re.FindAllString(fileString, -1)
	return len(results)
}
