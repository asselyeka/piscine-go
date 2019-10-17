package piscine

import "fmt"

func PrintStr(str string) {

	for _, v := range str {
		fmt.Print(string(v))
	}
	fmt.Println()
}
