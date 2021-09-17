package main

import "fmt"

func main() {
	pointers()
}

func pointers() {
	x := 1
	p := &x  // p is of type *int;  so could declare as var p *int = &x

	a := x   // a is assigned to the --value-- at &x. a DOES NOT reference the value at the address of x. So affecting a will not affect x

	fmt.Println(&x) // address of x
	fmt.Println(p)  // address of x
	y := &p
	fmt.Println(y)  // address of p
	fmt.Println(*y) // address of x
	fmt.Println(*p) // 1
	incrP(p)
	fmt.Println(x)  // 2
	incr(x)         // here, we pass the value of x. So incr() recieves the value of 2 (and int) and adds 1 to that value.
					// 		It does not however add 1 to the very value at &x
	fmt.Println(*p) // 2
	fmt.Println(x)  // 2
	fmt.Println(a)  // 1

}

func incr(i int) {
	i++
}

func incrP(i *int) {
	*i++
}