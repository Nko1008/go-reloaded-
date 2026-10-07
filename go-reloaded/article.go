package main

import (
	"strings"
	"unicode"
)

func fixArticles(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {
		if strings.ToLower(words[i]) == "a" {
			nextWord := strings.Trim(words[i+1], ".,!?;:")

			if len(nextWord) > 0 {
				firstLetter := unicode.ToLower(rune(nextWord[0]))

				if strings.ContainsRune("aeiouh", firstLetter) {
					if words[i] == "A" {
						words[i] = "An"
					} else {
						words[i] = "an"
					}
				}
			}
		}
	}

	return strings.Join(words, " ")
}
