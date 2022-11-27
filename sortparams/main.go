package main

import (
	"os"

	"github.com/01-edu/z01"
)

func SortIntegerTable2(table []string) {
	i := 1
	for i < len(table) {
		if table[i-1] > table[i] {
			tmp := table[i]
			table[i] = table[i-1]
			table[i-1] = tmp
			i = 1
		} else {
			i++
		}
	}
}

func main() {
	arg := os.Args
	var str []string

	for i := 1; i < len(arg); i++ {
		str = append(str, arg[i])
	}
	SortIntegerTable2(str)

	for i := 0; i < len(str); i++ {
		a := str[i]

		for j := 0; j < len(a); j++ {
			z01.PrintRune(rune(a[j]))
		}
		z01.PrintRune('\n')
	}
}
