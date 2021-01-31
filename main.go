package main

import "fmt"

// func main() {
// 	fmt.Print("Hello World!")
// }

func main() {
	var x string = "x"

	fmt.Println(x)

	var name = "Brad"
	var age = 37

	fmt.Println(name, age)

	const isCool = true

	othername := "Josh"

	fmt.Printf("%T\n", othername)

	phone, email := "8-800-88008", "some@email.com"

	fmt.Println(phone, email)
}
