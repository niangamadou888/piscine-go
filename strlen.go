package piscine

func StrLen(s string) int {
	a := 0
	for i := 0; i < len(s); i++ {
		if byte(s[i]) < 129 {
			a++
		}
	}
	b := len(s)
	if a != b {
		a++
	}
	return a
}
