package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	help := "--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n"
	arguments := os.Args
	str := ""
	strInsert := ""
	strForOrder := ""
	strI := ""
	strForInsert := ""
	argIndex := 0
	for indx := range arguments {
		argIndex = indx + 1
	}
	if argIndex == 1 {
		str = help
	} else {
		arg := []string(arguments[1:])
		for index, flag := range arg {
			len := 0
			for indx := range flag {
				len = indx + 1
			}
			if len > 8 {
				strTemp := flag
				strInsert = string(strTemp[:9])
			}
			if len > 2 {
				strTemp := flag
				strI = string(strTemp[:3])
			}

			if strInsert == "--insert=" {
				strTemp := flag
				for index, letter := range strTemp {
					if index > 8 {
						strForInsert += string(letter)
					}
				}
			}
			if strI == "-i=" {
				strTemp := flag
				for index, letter := range strTemp {
					if index > 2 {
						strForInsert += string(letter)
					}
				}
			}
			if flag == "--order" || flag == "-o" {
				strForOrder += string(arg[index+1])
			}
			if flag == "--help" || flag == "-h" {
				str = help
				break
			}
			if argIndex == 2 {
				str = flag
			}
		}
	}
	if strForOrder != "" || strForInsert != "" {
		str = SortStr(strForOrder + strForInsert)
	}
	for _, r := range str {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func SortStr(s string) string {
	table := []rune(s)
	size := 0
	for index := range table {
		size = index + 1
	}
	var temp rune
	for i := 0; i < size-1; i++ {
		for j := 0; j < (size - i - 1); j++ {
			if table[j] > table[j+1] {
				// меняем элементы местами
				temp = table[j]
				table[j] = table[j+1]
				table[j+1] = temp
			}
		}
	}
	return string(table)
}
