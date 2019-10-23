package main

import (
	"os"

	"github.com/01-edu/z01"

	piscine ".."
)

func main() {
	arguments := os.Args
	len := 0
	table := []string(arguments[1:])
	for index := range table {
		len = index + 1
	}
	if len > 0 {
		if table[0] == "--upper" {
			tableUpp := []string(table[1:])
			for _, num := range tableUpp {
				if piscine.Atoi(num) >= 1 && piscine.Atoi(num) <= 26 {
					z01.PrintRune(rune(piscine.Atoi(num) + 64))
				} else {
					z01.PrintRune(' ')
				}
			}
		} else {
			for _, num := range table {
				if piscine.Atoi(num) >= 1 && piscine.Atoi(num) <= 26 {
					z01.PrintRune(rune(piscine.Atoi(num) + 96))
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}
