package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 50 {
		return 0
	} else {
		result := 1
		for i := 1; i < nb+1; i++ {
			result *= i
		}
		return result
	}

}
