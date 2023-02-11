package piscine

func ConcatParams(args []string) string {
	var result string
	for i, value := range args {
		result += value
		if i < len(args)-1 {
			result += "\n"
		}
	}
	return result
}
