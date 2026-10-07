package main

import "strings"

func fixApostrophes(text string) string {
	words := strings.Fields(text)
	result := []string{}

	inQuote := false

	for _, word := range words {
		if strings.HasPrefix(word, "'") {
			inQuote = true
		}

		if inQuote && len(result) > 0 && !strings.HasPrefix(word, "'") {
			result = append(result, word)
		} else {
			result = append(result, word)
		}

		if strings.HasSuffix(word, "'") {
			inQuote = false
		}
	}

	return strings.Join(result, " ")
}
