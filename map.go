package piscine

func Map(f func(int) bool, a []int) []bool {
	length := len(a)
	myArray := make([]bool, length)

	for i := range a {
		myArray[i] = f(a[i])
	}
	return myArray
}
