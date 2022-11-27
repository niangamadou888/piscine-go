package piscine

func Range(start, end int) []int {
	var arrSize int = 1
	var tmpArr = []int{start}

	for start != end {
		if start < end {
			start++
		} else {
			start--
		}

		arrSize++

		tmpArr = append(tmpArr, start)
	}
	sliceOfnumbers := make([]int, arrSize)

	for i := range tmpArr {
		sliceOfnumbers[i] = tmpArr[i]
	}
	return sliceOfnumbers
}
