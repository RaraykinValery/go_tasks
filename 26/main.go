package main

import (
	"fmt"
	"unicode"
)

func hasUniqSymbols(s string) bool {
	seen := make(map[rune]bool)

	for _, symbol := range s {
		if unicode.IsLetter(symbol) {
			symbol = unicode.ToLower(symbol)
		}

		if _, ok := seen[symbol]; !ok {
			seen[symbol] = true
		} else {
			return false
		}
	}

	return true
}

func main() {
	strings_set := []string{"abcd", "abCdefAaf", "aabcd", "😃emoji😃", "😊"}

	for _, s := range strings_set {
		fmt.Println(hasUniqSymbols(s))
	}
}
