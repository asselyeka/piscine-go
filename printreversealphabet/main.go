package main

import "fmt"

func main() {

	var i = 'z'
	for i >= 'a' {
		fmt.Print(string(i))
		i--
	}
	fmt.Println()
}
