package main

import "fmt"

func main() {
	//like while method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//Real for loop
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}
}
