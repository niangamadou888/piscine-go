package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	ln := 0
	ln2 := 0
	ln3 := 0
	ans := ""
	st_t := map[rune]int{}
	for c := range nbr {
		ln = c + 1
	}
	for i, c := range baseFrom {
		st_t[c] = i
		ln2 = i + 1
	}
	for c := range baseTo {
		ln3 = c + 1
	}
	pw := 1
	cnt := 0
	for i := ln - 1; i >= 0; i-- {
		cnt = cnt + st_t[[]rune(nbr)[i]]*pw
		pw *= ln2
	}
	for cnt != 0 {
		ans = string(baseTo[cnt%ln3]) + ans
		cnt /= ln3
	}
	return ans
}
