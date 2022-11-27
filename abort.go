package piscine

func Abort(a, b, c, d, e int) int {
	var tab []int
	tab = append(tab, a, b, c, d, e)

	SortIntegerTable(tab)
	return tab[len(tab)/2]
}
