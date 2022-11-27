package piscine

func IsPrintable(s string) bool {
	// tabA := []rune(s)
	// lenA := len(tabA) - 2
	bool := true
	for _, value := range s {
		if value >= 32 {
			bool = bool
		} else {
			bool = false
		}
	}
	return bool
}
