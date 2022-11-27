package piscine

func SplitWhiteSpaces1(s string) []string {
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

func Split(str, charset string) []string {
	ln := 0

	for i := range charset {
		ln = i + 1
	}

	ln2 := 0
	for i := range str {
		ln2 = i + 1
	}

	for i := 0; i < ln2-ln; i++ {
		if str[i:i+ln] == charset {
			str = str[:i] + " " + str[i+ln:]
			ln2 -= ln
		}
	}

	return SplitWhiteSpaces1(str)
}
