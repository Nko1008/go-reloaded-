package main

import (
	"strings"
	"unicode"
)

func changeCase(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		switch words[i] {

		case "(up)":
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
				words[i] = ""
			}

		case "(low)":
			if i > 0 {
				words[i-1] = strings.ToLower(words[i-1])
				words[i] = ""
			}

		case "(cap)":
			if i > 0 {
				words[i-1] = capitalize(words[i-1])
				words[i] = ""
			}
		}
	}

	return strings.Join(words, " ")
}

func capitalize(word string) string {
	runes := []rune(strings.ToLower(word))

	if len(runes) > 0 {
		runes[0] = unicode.ToUpper(runes[0])
	}

	return string(runes)
}
