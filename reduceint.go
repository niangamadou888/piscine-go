package main

import (
	"github.com/01-edu/z01"
)

func main() {

}

func Itoa8(n int) string {
	num := n
	ch := ""
	if n < 0 {
		num = -n
		ch = "-"
	}
	digits := "0123456789"
	if num < 10 {
		return ch + digits[num:num+1]
	}
	return ch + Itoa(num/10) + digits[num%10:num%10+1]
}

func ReduceInt(f func(int, int) int, arr []int) {
	res := f(arr[0], arr[1])
	s := Itoa8(res)
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
