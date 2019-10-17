package piscine

func StrLen(str string) int {
	var count int
	for index := range str {
		count = index + 1
	}
	return count
}
