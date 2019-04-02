package main

import "fmt"

func main() {

	var name string = "Harry."
	fmt.Println("Name = ", name)
	fmt.Printf("%T\n", name)
	var age = 27
	fmt.Println("Age = ", age)
	fmt.Printf("%T\n", age)
	const birthday = "March 7, 2044"
	deathday := "May 7, 2350"
	fmt.Println(birthday + " " + deathday)
}
