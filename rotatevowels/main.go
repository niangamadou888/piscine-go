package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(value rune) bool {
	if value == 'a' || value == 'o' || value == 'u' || value == 'i' || value == 'e' || value == 'A' || value == 'O' || value == 'U' || value == 'I' || value == 'E' {
		return true
	}
	return false
}

func main() {
	args := os.Args[1:]
	var voyelles []rune
	var vowel []rune
	var str string
	if len(args) >= 1 {
		for indx, arguments := range args {
			str += arguments
			if indx != len(args)-1 {
				str += " "
			}
			for _, value := range arguments {
				if isVowel(value) {
					voyelles = append(voyelles, value)
				}
			}
		}

		i := len(voyelles) - 1
		for _, value := range str {
			if isVowel(value) {
				vowel = append(vowel, voyelles[i])
				i--
			} else {
				vowel = append(vowel, value)
			}
		}

		for _, value := range vowel {
			z01.PrintRune(value)
		}

	}
	z01.PrintRune('\n')
}
