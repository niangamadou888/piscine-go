package piscine

func LastRune(s string) rune {
	tabS := []rune(s)
	arune := tabS[(len(s))-1]
	return arune
}
