package piscine

import (
	"github.com/01-edu/z01"
)

func verifyDone(arr []rune) bool {
	if arr[len(arr)-1] == '9' {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i]+1 != arr[i+1] {
				return false
			}
		}
		return true
	}
	return false
}

func recursiveAdd(arr []rune) []rune {
	if arr[len(arr)-1] == '9' {
		x := '0'
		arr = arr[:len(arr)-1]
		value := recursiveAdd(arr)
		value = append(value, x)
		return value
	}
	arr[len(arr)-1]++
	return arr

}

func isLast(arr []rune) bool {
	if verifyDone(arr) {
		return true
	}
	return false
}

func PrintCombN(n int) {
	r := make([]rune, n)
	a := '0'
	for i := 0; i < n; i++ {
		r[i] = a
		a++
	}
	for !verifyDone(r) {
		for j := 0; j < len(r); j++ {
			z01.PrintRune(r[j])
		}
		z01.PrintRune(',')
		z01.PrintRune(' ')
		r = recursiveAdd(r)
	}
	for j := 0; j < len(r); j++ {
		z01.PrintRune(r[j])
	}

	z01.PrintRune(10)
}
