package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	lim := 21
	a := 1
	for i := 1; i <= nb; i++ {
		if i > lim {
			return 0
		}
		a *= i
	}
	return a
}
