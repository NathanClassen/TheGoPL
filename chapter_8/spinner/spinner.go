package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		go spinner() starts a new goroutine, executing the spinner function
		The spinner() func creates an infinite loop, wherein the characters "\", "|", "/" are printed
			over and over, giving illusion of a spinner
	*/
	go spinner(100 * time.Millisecond)
	const n = 48
	/*
		 The main routine calls fib(n) to calculate the fibonacci
	*/
	fibN := fib(n)
	/*
		We did not start another goroutine to calculate fib, so the mainroutine waits for fib(n) to return.
		One it returns, the result is printed and the main function returns.
		When the main function returns, all go routines are abruptly terminated and the program exits.
	*/
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x - 1) + fib(x - 2)
}