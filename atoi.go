package piscine

func Atoi(s string) int {
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
