package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for i := 1; i < len(arguments); i++ {
		strAsRune := []rune(arguments[i])
		for _, letter := range strAsRune {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
