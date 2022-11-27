package piscine

func f(a, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return +1
	} else {
		return -1
	}
}

func IsSorted(f func(a, b int) int, a []int) bool {
	croissant := true
	decroissant := true

	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) < 0 {
			croissant = false
		}
	}

	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) > 0 {
			decroissant = false
		}
	}

	return croissant || decroissant
}
