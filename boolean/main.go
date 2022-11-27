package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func even(nbr int) int {
	a := 1
	if nbr%2 == 0 {
		a = 0
	}
	return a
}

func isEven(nbr int) bool {
	if even(nbr) == 1 {
		return false
	} else {
		return true
	}
}

func main() {
	args := os.Args[1:]
	lengthOfArg := len(args)

	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	if isEven(lengthOfArg) == true {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
