package piscine

func Index(s string, toFind string) int {
	strAsByte := []byte(s)
	toFindByte := []byte(toFind)
	indexFind := -1
	for index, letter := range strAsByte {
		if letter == toFindByte[0] {
			j := 0
			count := 0
			for i := index; i < StrLen(toFind); i++ {
				if strAsByte[i] == toFindByte[j] {
					count++
				}
				j++
			}
			if count == j {
				indexFind = index
				break
			}
		}
	}
	return indexFind
}
