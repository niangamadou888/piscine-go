package piscine

func Compact(ptr *[]string) int {
	tab := *ptr
	count := 0

	for _, val := range *ptr {
		if val != "" {
			tab[count] = val
			count++
		}
	}
	*ptr = tab[0:count]

	return count
}
