package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	// tabA := []rune(a)
	for _, value := range a {
		for _, val := range value {
			z01.PrintRune(val)
		}
		z01.PrintRune('\n')
	}
}
