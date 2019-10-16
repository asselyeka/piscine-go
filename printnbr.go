package piscine

import (
	"github.com/01-edu/z01"
)

func IntToRune(num int) (runenum rune) {
	switch num {
	case 0:
		runenum = '0'
	case 1:
		runenum = '1'
	case 2:
		runenum = '2'
	case 3:
		runenum = '3'
	case 4:
		runenum = '4'
	case 5:
		runenum = '5'
	case 6:
		runenum = '6'
	case 7:
		runenum = '7'
	case 8:
		runenum = '8'
	case 9:
		runenum = '9'
	}
	return runenum
}

func PrintNbr(n int) {

	if n == 0 {
		z01.PrintRune(IntToRune(n))
	} else {

		var num int = 0
		var minus rune

		if n < 0 {
			minus = '-'
			n *= (-1)
		}

		var reverseNum int = n % 10
		n = n / 10
		for i := 1; n != 0; i++ {
			num = n % 10
			n = n / 10
			reverseNum = reverseNum*10 + num
		}

		if minus == '-' {
			z01.PrintRune('-')
		}
		for i := 1; reverseNum != 0; i++ {

			num = reverseNum % 10
			z01.PrintRune(IntToRune(num))
			reverseNum = reverseNum / 10
		}

	}

}
