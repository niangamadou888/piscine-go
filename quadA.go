package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		if x == 1 && y == 1 {
			z01.PrintRune('o')
			z01.PrintRune('\n')
		} else if x > 1 && y == 1 {
			for i := 1; i <= x; i++ {
				if i == 1 {
					z01.PrintRune('o')
				} else if i == x {
					z01.PrintRune('o')
					z01.PrintRune('\n')
				} else {
					z01.PrintRune('-')
				}
			}
		} else if y > 1 && x == 1 {
			for b := 1; b <= y; b++ {
				if b == 1 || b == y {
					z01.PrintRune('o')
					z01.PrintRune('\n')
				} else {
					z01.PrintRune('|')
					z01.PrintRune('\n')
				}
			}
		} else if y > 2 && x > 2 {
			z01.PrintRune('o')
			for i := 1; i <= x-2; i++ {
				z01.PrintRune('-')
			}
			z01.PrintRune('o')
			z01.PrintRune('\n')

			for i := 1; i <= y-2; i++ {
				z01.PrintRune('|')
				for i := 1; i <= x-2; i++ {
					z01.PrintRune(' ')
				}
				z01.PrintRune('|')
				z01.PrintRune('\n')
			}
			z01.PrintRune('o')
			for i := 1; i <= x-2; i++ {
				z01.PrintRune('-')
			}
			z01.PrintRune('o')
			z01.PrintRune('\n')

		}
	}
}
