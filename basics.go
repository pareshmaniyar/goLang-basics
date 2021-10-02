package main

import "fmt"

func main() {
	// implicit
	h := 2
	// explicit
	var numa int = 23
	var x uint16 = 2500
	fmt.Println("Hello World!", h, numa, x)

	// default
	var b1 bool
	var x2 uint16
	fmt.Println("b1: ", b1, x2)
	// type of variable
	fmt.Printf("%T %T %T", x2, 10, 3000000000000)
	var store string = fmt.Sprintf("%T", b1)
	fmt.Println(store)
	fmt.Printf("%b %o %d %x", 32324, 32324, 32324, 32324)
	fmt.Println("")
	fmt.Printf("%s %q", "Batman", "Spiderman")

}
