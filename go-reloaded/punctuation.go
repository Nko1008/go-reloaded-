
package main

import (
	"strings"
	"unicode"
)

func isPunctuation(r rune) bool {
	return strings.ContainsRune(".,!?:;", r)
}

func fixPunctuation(text string) string {
	runes := []rune(text)
	out := make([]rune, 0, len(runes)+8)

	for i := 0; i < len(runes); {
		if !isPunctuation(runes[i]) {
			out = append(out, runes[i])
			i++
			continue
		}

		end := i
		for end < len(runes) && isPunctuation(runes[end]) {
			end++
		}

		if end-i == 1 && strings.ContainsRune(".,:", runes[i]) &&
			i > 0 && end < len(runes) &&
			unicode.IsDigit(runes[i-1]) && unicode.IsDigit(runes[end]) {
			out = append(out, runes[i])
			i = end
			continue
		}

		for len(out) > 0 && out[len(out)-1] == ' ' {
			out = out[:len(out)-1]
		}
		out = append(out, runes[i:end]...)

		next := end
		for next < len(runes) && runes[next] == ' ' {
			next++
		}
		if next < len(runes) && (next > end || runes[next] != '\'') {
			out = append(out, ' ')
		}

		i = next
	}

	return string(out)
}
