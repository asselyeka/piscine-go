package piscine

func IsAlpha(str string) bool {
	sAsRune := []rune(str)
	counter := 0
	for _, letter := range sAsRune {
		if checkAlpNum(letter) {
			counter++
		}
	}
	if counter == len(str) {
		return true
	}
	return false
}

func checkAlpNum(r rune) bool {
	if r >= 'A' && r <= 'Z' ||
		r >= 'a' && r <= 'z' ||
		r >= '0' && r <= '9' {
		return true
	}
	return false
}
