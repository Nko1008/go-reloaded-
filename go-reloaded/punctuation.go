package main

import "strings"

func fixPunctuation(text string) string {
	words := strings.Fields(text)
	result := []string{}

	for i := 0; i < len(words); i++ {
		word := words[i]

		if isPunctuation(word) {
			if len(result) > 0 {
				result[len(result)-1] += word
			}
		} else {
			result = append(result, word)
		}
	}

	return strings.Join(result, " ")
}

func isPunctuation(word string) bool {
	if word == "..." || word == "!?" || word == "?!" {
		return true
	}

	return word == "." ||
		word == "," ||
		word == "!" ||
		word == "?" ||
		word == ":" ||
		word == ";"
}
