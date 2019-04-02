package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	isMale    bool
	age       int
}

func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " " + strconv.Itoa(p.age)
}

func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	if !p.isMale {
		p.lastName = spouseLastName
	}
}

func main() {
	personOne := Person{
		firstName: "Foo",
		lastName:  "Bar",
		city:      "Grass Valley",
		isMale:    false,
		age:       31}
	fmt.Println(personOne)

	personTwo := Person{"Foo", "Bar", "Grass Valley", true, 31}
	fmt.Println(personTwo)

	fmt.Println(personOne.greet())
	personOne.hasBirthday()
	fmt.Println(personOne.greet())
	personOne.getMarried("Newlywed")
	fmt.Println(personOne.greet())
}
