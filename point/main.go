package main

import (
	"github.com/01-edu/z01"
)

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

type point struct {
	x int
	y int
}

func PrintNbr(r int) {
	a := '0'
	b := '0'

	for i := 0; i < r/10; i++ {
		a++
	}
	z01.PrintRune(a)
	for i := 0; i < r%10; i++ {
		b++
	}
	z01.PrintRune(b)
}

func main() {
	points := point{}

	// nbr := points.x

	setPoint(&points)

	// fmt.Printf("x = %d, y = %d\n", points.x, points.y)

	// for _, value := range sentence {

	subchain1 := "x = "
	subchain2 := ", y = "
	for _, val := range subchain1 {
		z01.PrintRune(val)
	}
	PrintNbr(points.x)
	for _, val := range subchain2 {
		z01.PrintRune(val)
	}
	PrintNbr(points.y)
	z01.PrintRune('\n')
	/*z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	PrintNbr(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	PrintNbr(points.y)
	z01.PrintRune('\n')*/
}
