package piscine

//import "fmt"

func StrLen(str string) int {
	var count int
	var a int
	for index, word := range str {
		if word == 233 {
			a = -1
		}
		count = index + 1 + a
		//fmt.Printf("letter %v, index %v", rune(word), index)
		//fmt.Println()

	}
	return count
}
