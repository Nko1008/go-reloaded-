
package main

import (
	"regexp"
	"strings"
)

var instructionPattern = regexp.MustCompile(`\(\s*(hex|bin|up|low|cap)\s*(?:,\s*(\d+)\s*)?\)`)

func processText(text string) string {
	lines := strings.Split(text, "\n")

	for i, line := range lines {
		hasCR := strings.HasSuffix(line, "\r")
		line = processLine(strings.TrimSuffix(line, "\r"))
		if hasCR {
			line += "\r"
		}
		lines[i] = line
	}

	return strings.Join(lines, "\n")
}

func processLine(line string) string {
	words := tokenize(line)

	words = convertHex(words)
	words = convertBin(words)
	words = changeCase(words)

	text := strings.Join(words, " ")

	text = fixApostrophes(text)
	text = fixPunctuation(text)
	text = fixArticles(text)

	return text
}

func tokenize(line string) []string {
	spaced := instructionPattern.ReplaceAllStringFunc(line, func(match string) string {
		parts := instructionPattern.FindStringSubmatch(match)
		if parts[2] == "" {
			return " (" + parts[1] + ") "
		}
		return " (" + parts[1] + "," + parts[2] + ") "
	})

	return strings.Fields(spaced)
}
