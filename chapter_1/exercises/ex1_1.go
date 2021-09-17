// Ex1.1 prints the command that invoked it and it's cmmd line args.
package main

import (
	"fmt"
	"os"
	"strings"
)

func ex1_1() {
	fmt.Println(os.Args[0] + "\n")
	fmt.Println(strings.Join(os.Args[1:], " "))
}
