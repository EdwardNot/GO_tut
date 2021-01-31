package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	// firstname  string
	// secondname string
	// city       string
	// gender     string
	// age        int

	firstname, lastname, city, gender string
	age                               int
}

func (p Person) greet() string {
	return "Hello, " + p.firstname + ". Your age is" +
		strconv.Itoa(p.age)
}

func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastname = spouseLastName
	}
}

func main() {
	person1 :=
		Person{
			firstname: "Josh",
			lastname:  "Doe",
			city:      "Bender",
			gender:    "m",
			age:       20,
		}

	person2 := Person{"Karen", "Boo", "Tender", "f", 20}

	fmt.Println(person1)
	fmt.Println(person2.firstname)
	person1.age = 30

	fmt.Println(person2.greet())
}
