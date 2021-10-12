package main

import "fmt"

func operation(inputFunc func(num1 int, num2 int) int) func(num3 int) int {
	newFunc := func(num3 int) int {
		return inputFunc(2, 3) * num3
	}
	return newFunc
}
func counterGenerator() func() {
	count := 0
	return func() {
		fmt.Println(count)
		count++
	}
}
func main() {
	fmt.Println("Hello!")
	multiplier := func(num1 int, num2 int) int {
		return num1 * num2
	}
	f := operation(multiplier)
	fmt.Println(f(4))
	counter := counterGenerator()
	counter2 := counterGenerator()
	counter()
	counter()
	counter()
	counter2()
	counter2()
}
