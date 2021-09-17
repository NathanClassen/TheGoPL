// Dup2 reads a text file and for lines of text that appear more than once
// prints the lines as well as line cound
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		// if no files were provided to the program, then well examine Stdin instead
		countLines(os.Stdin, counts, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d: %s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, fname string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[fname + " --  [  " + input.Text() + "  ]"]++
	}
	// ignoring errors from input.Err()
}