package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: go run . input.txt output.txt")
		os.Exit(1)
	}

	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erreur lors de la lecture du fichier :", err)
		os.Exit(1)
	}

	result := processText(string(content))

	if err := os.WriteFile(os.Args[2], []byte(result), 0644); err != nil {
		fmt.Fprintln(os.Stderr, "Erreur lors de l'écriture du fichier :", err)
		os.Exit(1)
	}
}
