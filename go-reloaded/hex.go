package main

import (
	"strconv"
	"strings"
)

func convertHex(text string) string {
	words := strings.Fields(text)

	for i := 1; i < len(words); i++ {
		if words[i] == "(hex)" {
			number, err := strconv.ParseInt(words[i-1], 16, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(number, 10)
				words[i] = ""
			}
		}
	}

	return strings.Join(words, " ")
}
