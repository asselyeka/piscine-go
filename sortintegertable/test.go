package main

import (
	"fmt"

	piscine ".."
)

func main() {
	s := []int{12, 10, 7, 6, 5, 4, 3, 2, 1, 0}
	a := []int{-81, -13, 25, -67}
	b := []int{42, 82, -62, -19}
	c := []int{68, 65, -50, -38}
	d := []int{89, 28, 31, -81}
	piscine.SortIntegerTable(s)
	piscine.SortIntegerTable(a)
	piscine.SortIntegerTable(b)
	piscine.SortIntegerTable(c)
	piscine.SortIntegerTable(d)
	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

//	sortintegertable_test.go:27: SortIntegerTable([-81 -13 25 -67]), table == [-81 -13 25 -67] instead of [-81 -67 -13 25]
//    sortintegertable_test.go:27: SortIntegerTable([42 82 -62 -19]), table == [-62 42 82 -19] instead of [-62 -19 42 82]
//    sortintegertable_test.go:27: SortIntegerTable([68 65 -50 -38]), table == [-50 65 68 -38] instead of [-50 -38 65 68]
//    sortintegertable_test.go:27: SortIntegerTable([89 28 31 -81]), table == [28 31 89 -81] instead of [-81 28 31 89]
