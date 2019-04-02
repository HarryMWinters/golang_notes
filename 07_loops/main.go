package main

import "fmt"

func main() {
	// Loops
	i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	for i <= 100 {
		if i%25 == 0 {
			fmt.Println("Fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("Fizz")
		} else if i%3 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
		i++
	}
}
