package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 0 {
		str := "abcdefghijklmnopqrstuvwxyz"
		var params []int
		for _, value := range args {
			result := 0
			for k := 0; k < len(value); k++ {
				if byte(value[k]) >= 48 && byte(value[k]) <= 57 {
					result = result*10 + int(value[k]-'0')
				}
			}
			params = append(params, result)
		}
		b := false
		if args[0] != "--upper" {
			b = true
		}
		for _, value := range params {
			if value >= 1 && value <= 26 {
				if args[0] == "--upper" {
					z01.PrintRune(rune(str[value-1] - 32))
				} else {
					z01.PrintRune(rune(str[value-1]))
				}
			} else {
				if b {
					z01.PrintRune(' ')
				}
				if args[0] == "--upper" {
					b = true
				}
			}
		}
		z01.PrintRune('\n')
	}
}
