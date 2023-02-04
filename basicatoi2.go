package piscine

func contains(elems []rune, v rune) bool {
	for i := 0; i < len(elems); i++ {
		if elems[i] == v {
			return true
		}
	}
	return false
}

func BasicAtoi2(s string) int {
	digit := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	arrayStr := []rune(s)
	n := len(s)
	ans := 0

	for i := 0; i < len(arrayStr); i++ {
		if !contains(digit, arrayStr[i]) {
			return 0
		}
	}

	for i := 0; i < n; i++ {
		if arrayStr[i] < '0' || arrayStr[i] > '9' {
			return ans
		} else {
			ans *= 10
			ans += int(arrayStr[i]) - '0'
		}
	}
	return ans
}
