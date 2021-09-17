package main

import "fmt"

func main() {
	// create a pointer to a variable whose value is the zero-value of an integer (0)
	i := new(int)

	// print the address of this variable whose value is the zero-value of an integer (0)
	fmt.Println(i)  // some address

	// print the value of this variable whose value is the zero-value of an integer (0)
	fmt.Println(*i)

	// make value of this variable 1 without reassigning
	makeOne(i)
	fmt.Println("i is now 1: ", *i)

	// create a new integer variable whose value is 1, by passing an anonymous integer variable into on maker
	one := oneMaker(new(int))

	// get the incremented value of an integer without passing in the variable explicitly into a function
	// the variable should not be modified
	fmt.Println("one is: ", one)
	fmt.Println("one + 1 is: ", incrementor(&one))
	fmt.Println("but one is still: ", one)

}

func incrementor(i *int) int {
	return *i + 1
}

func oneMaker(p *int) int {
	*p++
	return *p
}

func makeOne(p *int) {
	*p++
}