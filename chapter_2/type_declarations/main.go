package main

import "fmt"

type Farenheight float64
type Celcius float64
type Kelvin int
type thing string

func main() {
	var F Farenheight
	var C Celcius
	var K Kelvin
	var t thing

	fmt.Println(F == 0) // true
	fmt.Println(C == 0) // true
	fmt.Println(F) // 0.0 --> 0
	fmt.Println(C) // 0.0 --> 0
	fmt.Println(K) // 0
	// fmt.Println(C - F)  compile error, type mismatch
	fmt.Println(F == Farenheight(C))
	fmt.Println(C == Celcius(F))
	fmt.Println(K == Kelvin(F)) // error --> true
	fmt.Println(t == thing(K)) // error --> false
	fmt.Println(thing(K)) // "0" --> " " or blank
	// fmt.Println(Kelvin(t))  compile time error
}