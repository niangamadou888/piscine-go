package piscine

func Any(f func(string) bool, a []string) bool {
	boolean := false

	for _, val := range a {
		if f(val) {
			boolean = true
		}
	}

	return boolean
}
