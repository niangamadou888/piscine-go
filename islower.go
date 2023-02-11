package piscine

func IsLower(s string) bool {
	bool := true
	for _, value := range s {
		if value >= 'a' && value <= 'z' {
			bool = bool
		} else {
			bool = false
		}
	}
	return bool
}
