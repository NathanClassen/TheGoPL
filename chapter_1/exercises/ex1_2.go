// Ex1.2 prints the index and value of each cmmd line arg, on a separate line.
package main

import (
	"fmt"
	"os"
)

func ex1_2() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, ": "+arg)
	}
}
