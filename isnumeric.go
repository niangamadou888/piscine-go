package piscine

func IsNumeric(s string) bool {
	bool := true

	for _, value := range s {
		if value >= '0' && value <= '9' {
			bool = bool
		} else {
			return false
		}
	}
	return bool
}
