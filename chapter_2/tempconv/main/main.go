package main

import (
	"fmt"
	"flag"
	"strconv"
	tmpconv "TheGoProgrammingLanguage/chapter_2/tempconv"
)

var i = flag.String("i", "f", "input temperature type") // `i` is a pointer to the flag
var o = flag.String("o", "f", "desired outpute temperature type") // `o` is a pointer to the flag

func main() {
	flag.Parse()

	temp, err := strconv.ParseFloat(flag.Args()[0], 64)

	if err != nil {
		fmt.Println("could not convert temperature argument to float64")
	}

	if *i == "f" {
		t := tmpconv.Fahrenheit(temp)
		switch *o {
		case "c":
			fmt.Println(tmpconv.FToC(t))
		case "k":
			fmt.Println(tmpconv.FToK(tmpconv.Fahrenheit(temp)))
		default:
			fmt.Println(fmt.Sprintf("%g°F", temp))
		}
	} else if *i == "c" {
		t := tmpconv.Celsius(temp)
		switch *o {
		case "f":
			fmt.Println(tmpconv.CToF(t))
		case "k":
			fmt.Println(tmpconv.CToK(tmpconv.Celsius(temp)))
		default:
			fmt.Println(fmt.Sprintf("%g°C", temp))
		}
	} else if *i == "k" {
		t := tmpconv.Kelvin(temp)
		switch *o {
		case "c":
			fmt.Println(tmpconv.KToC(t))
		case "f":
			fmt.Println(tmpconv.KToF(tmpconv.Kelvin(temp)))
		default:
			fmt.Println(fmt.Sprintf("%g°K", temp))
		}
	}
}