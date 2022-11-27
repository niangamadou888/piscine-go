package main

import (
	"fmt"
	"os"
)

func SortIntegerTable(table []rune) {
	i := 1
	for i < len(table) {
		if table[i-1] > table[i] {
			tmp := table[i-1]
			tmp1 := table[i]
			table[i-1] = tmp1
			table[i] = tmp
			i = 1
		} else {
			i++
		}
	}
}

func help() {
	fmt.Printf("--insert\n  -i\n	 This flag inserts the string into the string passed as argument.\n--order\n  -o\n	 This flag will behave like a boolean, if it is called it will order the argument.")
}

func main() {
	args := os.Args[1:]
	var arr []rune

	var str []string
	var st string

	var returnvalue string
	var space string

	if len(args) >= 1 {
		// je gere les espaces telque
		for _, value := range args {
			for _, val := range value {
				if val == ' ' {
					space += string(val)
				}
				if val == '=' {
					arr = append(arr, ' ')
				} else {
					arr = append(arr, rune(val))
				}
			}
			arr = append(arr, ' ')
		}

		// je met ma commande dans un tableau de type string
		for i, value := range arr {
			if value == ' ' || i == len(arr)-1 {

				str = append(str, st)
				st = ""
			} else {
				st += string(value)
			}
		}

		isauthercommande := false
		for i := 0; i < len(str); i++ {
			value := str[i]

			if value == "--insert" || value == "-i" {
				for j := len(str) - 1; j >= 0; j-- {
					val := str[j]

					if len(val) != 0 && val[0] != '-' {
						returnvalue += val
					}
				}
				isauthercommande = true
			}
			if value == "-o" || value == "--order" {
				if isauthercommande {
					to := []rune(returnvalue)
					SortIntegerTable(to)
					returnvalue = string(to)

				} else {
					for j := len(str) - 1; j >= 0; j-- {
						val := str[j]
						if len(val) != 0 && val[0] != '-' {
							returnvalue += val
						}
					}
					to := []rune(returnvalue)
					SortIntegerTable(to)
					returnvalue = string(to)

				}
			}
			if value == "-h" || value == "--help" {
				help()
			}

			iscommandhere := false
			for _, val := range str {
				if val == "-o" || val == "--order" || val == "-i" || val == "--insert" || val == "-h" || val == "--help" {
					iscommandhere = true
				}
			}
			if !iscommandhere {
				returnvalue = value
			}

		}

	} else {
		help()
	}
	finlal := space + returnvalue

	fmt.Println(finlal)
}
