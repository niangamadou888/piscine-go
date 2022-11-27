package piscine

func Join(strs []string, sep string) string {
	var ret string
	for k, str := range strs {
		if k > 0 && k <= len(strs)-1 {
			ret += sep
			ret += str
		} else {
			ret += str
		}
	}
	return ret
}
