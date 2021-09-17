package main

import (
	"fmt"
	"os"
)

func main() {
	cmd := os.Args[1]
	args := []string{os.Args[0]}
	os.Args = append(args, os.Args[2:]...)

	switch cmd {
	case "1":
		ex1_1()
	case "2":
		ex1_2()
	default:
		fmt.Println("select valid option: \"1\", \"2\"")
	}

}
