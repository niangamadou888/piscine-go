package piscine

func NRune(s string, n int) rune {
	tabS := []rune(s)
	if n <= 0 || n > len(s) {
		return 0
	} else {
		ans := tabS[n-1]
		return ans
	}
}
