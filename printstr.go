package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	length := len(str)
	for i := 0; i < length; i++ {
		z01.PrintRune(rune(str[i]))
	}
	z01.PrintRune('\n')
}
