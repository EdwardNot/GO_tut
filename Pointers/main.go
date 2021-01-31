package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(*b) //a value
	fmt.Println(*&a)

	*b = 10
	fmt.Println(a) // >>10
}
