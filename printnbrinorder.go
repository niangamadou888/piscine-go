package piscine

import (
	"github.com/01-edu/z01"
)

func SortIntegerTable1(table []int) {
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

func PrintNbrInOrder(n int) {
	var tab []int
	if n == 0 {
		tab = append(tab, 0)
	}
	for i := 0; n > 0; i++ {
		tab = append(tab, n%10)
		n /= 10
	}
	SortIntegerTable1(tab)

	for i := range tab {
		z01.PrintRune('0' + rune(tab[i]))
	}
}
