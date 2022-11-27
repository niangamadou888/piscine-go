package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi8(s string) int {
	num := 0
	mult := 1

	i := 0

	if len(s) == 0 {
		return 0
	}

	if s[0] == '-' {
		mult = -1
		i++
	} else if s[0] == '+' {
		mult = 1
		i++
	}

	for i < len(s) {
		num = num*10 + int(s[i]-48)
		if (int(s[i])-48) > 9 || (int(s[i])-48) < 0 {
			return 0
		}
		i++
	}

	return num * mult
}

func PrintBits() {
	arg := os.Args
	base := "01"

	if len(arg) == 2 {
		nbr := Atoi8(os.Args[1])
		var res string
		for nbr > 0 {
			digit := base[nbr%(2)]
			res = string(digit) + res
			nbr = nbr / 2
		}

		for len(res) != 8 {
			res = "0" + res
		}

		for _, result := range res {
			z01.PrintRune(result)
		}

	}
	z01.PrintRune('\n')
}
