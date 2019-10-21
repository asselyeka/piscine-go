package piscine

func IsPrime1(nb int) bool {
	if nb < 2 {
		return false
	}
	for i := nb - 1; i > 1; i-- {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	result := nb
	for i := nb; i < nb+50; i++ {
		if IsPrime1(i) {
			result = i
			break
		}
	}
	return result
}
