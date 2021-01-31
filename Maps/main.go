package main

import "fmt"

func main() {
	emails := make(map[string]string)

	emails["Josh"] = "josh@email.com"
	emails["Karen"] = "karen@email.com"

	fmt.Println(emails["Josh"])

	delete(emails, "Karen")

	phones := map[string]string{"Josh": "+1", "Karen": "+2"}
}
