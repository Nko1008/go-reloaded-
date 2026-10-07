package main

import (
	"strconv"
	"strings"
)

func convertBin(text string) string {
	words := strings.Fields(text)

	for i := 1; i < len(words); i++ {
		if words[i] == "(bin)" {
			number, err := strconv.ParseInt(words[i-1], 2, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(number, 10)
				words[i] = ""
			}
		}
	}

	return strings.Join(words, " ")
}
