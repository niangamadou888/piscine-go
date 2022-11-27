package piscine

func MakeRange(min, max int) []int {
	value := max - min
	var tabA []int
	if max < min {
		value = 0
		tabA = (nil)
	}

	if min < max {
		tabA = make([]int, value)
		for i := 0; i < value; i++ {
			tabA[i] = min + i
		}
		return tabA
	}
	return tabA
}
