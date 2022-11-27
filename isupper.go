package piscine

func IsUpper(s string) bool {
	bool := true
	for _, value := range s {
		if value >= 'A' && value <= 'Z' {
			bool = bool
		} else {
			bool = false
		}
	}
	return bool
}
