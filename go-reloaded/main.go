package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}

	text := string(content)

	text = convertHex(text)
	text = convertBin(text)
	text = changeCase(text)
	text = fixPunctuation(text)
	text = fixApostrophes(text)
	text = fixArticles(text)

	err = os.WriteFile(outputFile, []byte(text), 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier :", err)
		return
	}

	fmt.Println("Fichier créé avec succès !")
}
