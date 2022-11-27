package piscine

func Rot14(s string) string {
	arrayRune := []rune(s)
	var result string

	for i := 0; i < len(s); i++ {
		if arrayRune[i] >= 'a' && arrayRune[i] <= 'z' {
			if arrayRune[i] >= 'm' {
				arrayRune[i] -= 12
			} else if arrayRune[i] >= 'n' {
				arrayRune[i] -= 13
			} else {
				arrayRune[i] += 14
			}
		} else if arrayRune[i] >= 'A' && arrayRune[i] <= 'Z' {
			if arrayRune[i] >= 'M' {
				arrayRune[i] -= 12
			} else if arrayRune[i] >= 'N' {
				arrayRune[i] -= 13
			} else {
				arrayRune[i] += 14
			}
		}
		result += string(arrayRune[i])
	}
	return result
}
