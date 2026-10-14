package main

import (
	"strings"
	"unicode"
)

func fixArticles(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {
		if words[i] != "a" && words[i] != "A" {
			continue
		}

		if startsWithVowelOrH(words[i+1]) {
			if words[i] == "A" {
				words[i] = "An"
			} else {
				words[i] = "an"
			}
		}
	}

	return strings.Join(words, " ")
}

func startsWithVowelOrH(word string) bool {
	for _, r := range word {
		if r == '\'' || r == '"' || r == '(' {
			continue
		}
		return strings.ContainsRune("aeiouh", unicode.ToLower(r))
	}
	return false
}x
