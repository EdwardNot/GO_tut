package main

import "fmt"

func main() {
	x := 5
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if x > y {
		fmt.Printf("%d is less than %d\n", y, x)
	} else {
		fmt.Printf("%d is equal to %d\n", x, y)
	}

	if x == 5 && y == 10 {
		fmt.Printf("all is correct")
	} else if x != 5 || y != 10 {
		fmt.Printf("something is wrong")
	}

	color := "red"

	switch color {
	case "red":
		fmt.Printf("color is red")
	case "blue":
		fmt.Printf("color is blue")
	default:
		fmt.Printf("color is not red or blue")
	}
}
