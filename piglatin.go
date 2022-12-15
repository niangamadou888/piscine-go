package main

import (
	"os"

	"github.com/01-edu/z01"
)

func ContainsRune(str string, r rune) bool {
	for _, s := range str {
		if s == r {
			return true
		}
	}
	return false
}

func translate(s string) {
	if len(s) < 2 {
		for _, val := range s {
			z01.PrintRune(val)
		}
		return
	}

	vowels := "aeiou"
	if ContainsRune(vowels, rune(s[0])) {
		for _, val3 := range s {
			z01.PrintRune(val3)
		}
		z01.PrintRune('a')
		z01.PrintRune('y')
		return
	}
	for _, val1 := range s[1:] {
		z01.PrintRune(val1)
	}
	/*for _,val2:= range s[0]{
		z01.PrintRune(val2)
	}*/
	z01.PrintRune(rune(s[0]))
	z01.PrintRune('a')
	z01.PrintRune('y')
}

func main() {
	args := os.Args[1:]

	translate(args[0])
}
