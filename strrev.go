package piscine

func StrRev(s string) string {
	byte_str := []rune(s)
	long := len(byte_str)
	world := ""
	for i := long - 1; i >= 0; i-- {
		world += string(byte_str[i])
	}
	return world
}

