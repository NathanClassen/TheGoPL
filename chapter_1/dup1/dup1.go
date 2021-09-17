// Dup1 prints the text of each line that appears more than
// once in the stdin, preceded by it's count
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// fileInput()
	userInput()
}

func dup1(reader io.Reader) {
	counts := make(map[string]int)
	input := bufio.NewScanner(reader)
	for input.Scan() {
		line := input.Text()
		if line != "" {
			counts[input.Text()]++
		}
	}
	// ignore errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d: %s\n", n, line)
		}
	}
}

func fileInput() {
	file, _ := os.Open("./sample.txt")
	dup1(file)
}

func userInput() {
	dup1(os.Stdin)
}