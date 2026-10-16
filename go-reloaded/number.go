package main

import "strconv"

func convertBase(words []string, marker string, base int) []string {
	result := make([]string, 0, len(words))

	for _, word := range words {
		if word == marker && len(result) > 0 {
			number, err := strconv.ParseInt(result[len(result)-1], base, 64)
			if err == nil {
				result[len(result)-1] = strconv.FormatInt(number, 10)
				continue
			}
		}
		result = append(result, word)
	}

	return result
}
