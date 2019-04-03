package main

import "fmt"

var w string
var x int
var y string
var z bool

func main() {
	fmt.Print(x, " ", y, " ", z, "\n")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println("\nAssiging values...")
	x := 42
	y := "James Bond"
	z := true

	fmt.Print(x, " ", y, " ", z, "\n")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println("Now with .Sprint")
	w = fmt.Sprintf("%[1]d %[2]s %[3]t", x, y, z)
	fmt.Println("w == ", w)
	fmt.Print("Program finished.\n")
}
