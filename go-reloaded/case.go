package main

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var caseInstruction = regexp.MustCompile(`^\((up|low|cap)(?:,(\d+))?\)$`)

func changeCase(words []string) []string {
	result := make([]string, 0, len(words))

	for _, word := range words {
		match := caseInstruction.FindStringSubmatch(word)
		if match == nil {
			result = append(result, word)
			continue
		}

		count := 1
		if match[2] != "" {
			n, err := strconv.Atoi(match[2])
			if err != nil {
				n = len(result)
			}
			count = n
		}

		applyCase(result, count, match[1])
	}

	return result
}

func applyCase(words []string, count int, action string) {
	for i := len(words) - 1; i >= 0 && count > 0; i-- {
		if !hasLetterOrDigit(words[i]) {
			continue
		}

		switch action {
		case "up":
			words[i] = strings.ToUpper(words[i])
		case "low":
			words[i] = strings.ToLower(words[i])
		case "cap":
			words[i] = capitalize(words[i])
		}
		count--
	}
}

func hasLetterOrDigit(word string) bool {
	for _, r := range word {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func capitalize(word string) string {
	runes := []rune(strings.ToLower(word))

	for i, r := range runes {
		if unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			break
		}
	}

	return string(runes)
}
