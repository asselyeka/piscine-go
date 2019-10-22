package main

import (
	"fmt"

	piscine ".."
)

func main() {
	fmt.Println(piscine.Index("Hello!", "l"))
	fmt.Println(piscine.Index("Salut!", "alu"))
	fmt.Println(piscine.Index("Ola!", "hOl"))
	fmt.Println(piscine.Index("<bm;crAQt~&$6", ""))
	fmt.Println(piscine.Index("7i:\\J@V;)l-i>", ">*K<V?.sr<QM-"))
	fmt.Println(piscine.Index("Ola!", "hOl"))
}
