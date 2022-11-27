package piscine

func RecursiveFactorial(nb int) (a int) {
	if nb > 21 {
		return 0
	}
	if nb >= 0 {
		if nb == 0 {
			a = 1
		} else {
			a = nb * RecursiveFactorial(nb-1)
		}
	} else {
		return 0
	}
	return a
}
