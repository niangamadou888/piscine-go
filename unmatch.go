package piscine

func Unmatch(arr []int) int {
	var ter int
	for _, v := range arr {
		ter = 0
		for _, z := range arr {
			if v == z {
				ter++
			}
		}
		if ter%2 != 0 {
			return v
		}

	}
	return -1
}
