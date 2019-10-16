package main

import "github.com/01-edu/z01"

func main() {

	i := 0
	for i < 26 {
		z01.PrintRune(rune(122 - i))
		i++
	}
	z01.PrintRune('\n')
}
