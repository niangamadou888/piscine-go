package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, val := range args {
		if val == "01" || val == "galaxy" || val == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
			// fmt.Println("\n")
		}
	}
}
