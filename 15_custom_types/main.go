package main

import "fmt"

type myInt int

var a myInt = 3

func main() {
	var c myInt = 12
	fmt.Printf("Var c is a %T\n", c)

	var d = int(c)
	fmt.Printf("Var d is a %T\n", d)

	var e = myInt(d)
	fmt.Printf("Var e is a %T\n", e)
	fmt.Println("Program complete.")
}
