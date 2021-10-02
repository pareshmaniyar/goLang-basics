package main

import "fmt"

func main() {
	number := 12

	if number < 10 {
		fmt.Printf("true")
	} else {
		fmt.Printf("false")
	}
	for i := 0; i < 10; i++ {
		fmt.Println("i: %d", i)
	}
	switch number {
	case 12:
		println("12")
	case 14, 13:
		println("12")
	}
	switch {
	case number < 13:
		fmt.Printf("<13")
	case number > 3 && number > 9:
		fmt.Printf("<13")
	}
}
