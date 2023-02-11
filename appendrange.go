package piscine

func AppendRange(min, max int) []int {
	var tabA []int

	if min < max {
		for i := min; i < max; {
			tabA = append(tabA, i)
			i++
		}
	}
	if max > min {
		return tabA
	}
	return tabA
}
