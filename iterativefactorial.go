package piscine

func IterativeFactorial(nb int) int {
	if nb < 20000 {
		result := 1
		for i := 1; i < nb+1; i++ {
			result *= i
		}
		return result
	} else {
		return 0
	}

}