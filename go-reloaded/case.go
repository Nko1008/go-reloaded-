package main

import (
	"strconv"
	"strings"
	"unicode"
)

func changeCase(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		switch {
		case words[i] == "(up)":
			changePreviousWords(words, i, 1, "up")
			words[i] = ""

		case words[i] == "(low)":
			changePreviousWords(words, i, 1, "low")
			words[i] = ""

		case words[i] == "(cap)":
			changePreviousWords(words, i, 1, "cap")
			words[i] = ""

		case strings.HasPrefix(words[i], "(up,") && strings.HasSuffix(words[i], ")"):
			number := getNumber(words[i])
			changePreviousWords(words, i, number, "up")
			words[i] = ""

		case strings.HasPrefix(words[i], "(low,") && strings.HasSuffix(words[i], ")"):
			number := getNumber(words[i])
			changePreviousWords(words, i, number, "low")
			words[i] = ""

		case strings.HasPrefix(words[i], "(cap,") && strings.HasSuffix(words[i], ")"):
			number := getNumber(words[i])
			changePreviousWords(words, i, number, "cap")
			words[i] = ""
		}
	}

	return strings.Join(words, " ")
}

func changePreviousWords(words []string, position int, number int, action string) {
	start := position - number

	if start < 0 {
		start = 0
	}

	for i := start; i < position; i++ {
		switch action {
		case "up":
			words[i] = strings.ToUpper(words[i])

		case "low":
			words[i] = strings.ToLower(words[i])

		case "cap":
			words[i] = capitalize(words[i])
		}
	}
}

func getNumber(instruction string) int {
	instruction = strings.TrimPrefix(instruction, "(up,")
	instruction = strings.TrimPrefix(instruction, "(low,")
	instruction = strings.TrimPrefix(instruction, "(cap,")
	instruction = strings.TrimSuffix(instruction, ")")
	instruction = strings.TrimSpace(instruction)

	number, err := strconv.Atoi(instruction)

	if err != nil {
		return 1
	}

	return number
}

func capitalize(word string) string {
	runes := []rune(strings.ToLower(word))

	if len(runes) > 0 {
		runes[0] = unicode.ToUpper(runes[0])
	}

	return string(runes)
}
