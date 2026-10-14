package main

import (
	"strings"
	"unicode"
)

func isQuote(runes []rune, i int) bool {
	if runes[i] != '\'' {
		return false
	}

	prevLetter := i > 0 && unicode.IsLetter(runes[i-1])
	nextLetter := i+1 < len(runes) && unicode.IsLetter(runes[i+1])

	return !(prevLetter && nextLetter)
}

func fixApostrophes(text string) string {
	runes := []rune(text)

	var quotes []int
	for i := range runes {
		if isQuote(runes, i) {
			quotes = append(quotes, i)
		}
	}

	remove := make([]bool, len(runes))

	for p := 0; p+1 < len(quotes); p += 2 {
		open, closing := quotes[p], quotes[p+1]

		for i := open + 1; i < closing && unicode.IsSpace(runes[i]); i++ {
			remove[i] = true
		}
		for i := closing - 1; i > open && unicode.IsSpace(runes[i]); i-- {
			remove[i] = true
		}
	}

	var sb strings.Builder
	for i, r := range runes {
		if !remove[i] {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}
