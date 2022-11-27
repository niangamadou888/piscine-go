package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func pStr(s string) {
	sRune := []rune(s)
	for i := 0; i < len(sRune); i++ {
		z01.PrintRune(sRune[i])
	}
}

func main() {
	arg := os.Args[1:]
	if len(arg) > 0 {
		for i := 0; i < len(arg); i++ {
			data, err := ioutil.ReadFile(arg[i])
			if err != nil {
				errFile := "ERROR: open " + string(arg[i]) + ": no such file or directory"
				pStr(errFile)
				z01.PrintRune('\n')
				os.Exit(1)
			}
			pStr(string(data))
		}
	} else {
		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			panic(err)
		}
	}
}
