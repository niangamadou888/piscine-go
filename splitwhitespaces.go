package piscine

func SplitWhiteSpaces(s string) []string {
	var myString []string
	var result string

	for i := range s {
		if s[i] > ' ' {
			result += string(s[i])
		} else if result != "" {
			myString = append(myString, result)
			result = ""
		}
		if i == len(s)-1 {
			myString = append(myString, result)
		}
	}
	return myString
}
