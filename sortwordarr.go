package piscine

func SortWordArr(array []string) {
	size := 0
	for range array {
		size++
	}

	for i := size; i > 0; i-- {
		for j := 1; j < i; j++ {
			if array[j-1] > array[j] {
				reserv := array[j-1]
				array[j-1] = array[j]
				array[j] = reserv
			}
		}
	}
}
