package piscine

func BasicAtoi(s string) int {
	num := 0
	for _, digitChar := range s {
		digit := int(digitChar - '0') // Convert rune to integer
		num = num*10 + digit
	}
	return num
}

