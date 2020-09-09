package main

import "fmt"

func main() {
	fmt.Println(stringsToMap([]string{"titi", "toto", "titi"}))
}

func stringsToMap(list []string) map[string]int {
	mstrings := make(map[string]int, len(list))
	for _, s := range list {
		mstrings[s]++
	}
	return mstrings
}
