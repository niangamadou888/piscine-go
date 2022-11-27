package piscine

func ReverseBits(oct byte) byte {
	var r byte
	a := 8
	for a > 0 {
		r = (r << 1) | (oct & 1)
		oct >>= 1
		a--
	}
	return r
}
