package main

import "fmt"

func main() {
	a := 5
	b := &a

	//b = 0xc0000180a0
	fmt.Println(a, " - ", b, " - ", *b)
}
