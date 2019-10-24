package piscine

func SplitWhiteSpaces(str string) []string {
	strAsRune := []rune(str)
	len := 0
	size := 0
	for i := range strAsRune {
		len = i + 1
	}
	if len > 0 && CheckAlpN(strAsRune[0]) == true {
		size++
		for i := 0; i < len; i++ {
			if strAsRune[i] == ' ' || strAsRune[i] == '\n' || strAsRune[i] == '\t' {
				if i+1 < len {
					if CheckAlpN(strAsRune[i+1]) == true {
						size++
					}
				}
			}
		}
	}
	answer := make([]string, size)
	word := ""
	index := 0
	if len > 0 && CheckAlpN(strAsRune[0]) == true {
		for i := 0; i < len; i++ {
			word += string(strAsRune[i])
			if strAsRune[i] == ' ' || strAsRune[i] == '\n' || strAsRune[i] == '\t' {
				if i+1 < len {
					if CheckAlpN(strAsRune[i+1]) == true {
						answer[index] = word
						index++
						word = ""
					}
				}
			}
		}
		answer[size-1] = word
	}
	return answer
}

func CheckAlpN(r rune) bool {
	if r >= 'A' && r <= 'Z' ||
		r >= 'a' && r <= 'z' ||
		r >= '0' && r <= '9' {
		return true
	}
	return false
}
