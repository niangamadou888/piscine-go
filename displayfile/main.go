package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("quest8.txt")
	if err != nil {
		fmt.Printf("the mistake is : %v\n", err.Error())
	}
	args := os.Args[1:]
	// length := len(args)

	if len(args) == 0 {
		fmt.Println("File name missing")
	} else if len(args) > 1 {
		fmt.Println("Too many arguments")
	} else {

		arr := make([]byte, 14)

		file.Read(arr)

		fmt.Println(string(arr))
		file.Close()
	}
}
