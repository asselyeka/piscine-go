package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	l := LenLen(base)
	if CheckBase(base) == false || l < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
	} else if nbr > 2000000000 || nbr < -2000000000 {
		nbrAsStr := ItoaR(nbr)
		nbrAsRune := []rune(nbrAsStr)
		for _, num := range nbrAsRune {
			z01.PrintRune(num)
		}
	} else {
		if nbr < 0 {
			z01.PrintRune('-')
			nbr *= -1
		}
		str := ""
		for nbr >= l {
			str += string(base[nbr%l])
			nbr = nbr / l
		}
		str += string(base[nbr])
		str = RevStr(str)
		for _, letter := range str {
			z01.PrintRune(letter)
		}
	}
}
func CheckBase(base string) bool {
	for index, letter := range base {
		for i := index + 1; i < LenLen(base); i++ {
			if letter == '-' || letter == '+' || letter == rune(base[i]) {
				return false
			}
		}
	}
	return true
}
func RevStr(str string) string {
	var rev string
	l := LenLen(str)
	for i := l - 1; i >= 0; i-- {
		rev += string(str[i])
	}
	return rev
}

func LenLen(str string) int {
	count := 0
	a := []rune(str)
	for index := range a {
		count = index + 1
	}
	return count
}
func ItorR(n int) string {
	switch n {
	case 0:
		return "0"
	case 1:
		return "1"
	case 2:
		return "2"
	case 3:
		return "3"
	case 4:
		return "4"
	case 5:
		return "5"
	case 6:
		return "6"
	case 7:
		return "7"
	case 8:
		return "8"
	default:
		return "9"
	}
}

func ItoaR(n int) string {
	size := 0
	numAsStr := ""
	num := n
	for num != 0 {
		num /= 10
		size++
	}
	for i := 0; i < size; i++ {
		numAsStr += ItorR(n % 10)
		n /= 10
	}
	return RevStr(numAsStr)
}
