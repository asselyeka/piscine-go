package piscine

func FirstRune(s string) rune {
	word := []byte(s)
	return rune(word[0])
}
