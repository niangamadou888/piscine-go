package piscine

func AlphaCount(s string) int {
	counter := 0
	tabS := []rune(s)

	for i := 0; i <= len(s)-1; i++ {
		if tabS[i] >= 'a' && tabS[i] <= 'z' || tabS[i] >= 'A' && tabS[i] <= 'Z' {
			counter++
		}
	}
	return counter
}
