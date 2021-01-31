package main

import "fmt"

func main() {
	var fruitArray [2]string

	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"

	fmt.Println(fruitArray)
	fmt.Println(len(fruitArray))

	vegArray := []string{"Potato", "Tomato", "Pepper"} // OR vegArray := [2]string{"Potato", "Tomato"}
	fmt.Println(vegArray[1:2])

}
