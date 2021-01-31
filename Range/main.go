package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//without using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	phones := map[string]string{"Josh": "+1", "Karen": "+2"}

	for k, v := range phones {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range phones {
		fmt.Printf("Only name: %d\n", k)
	}
}
