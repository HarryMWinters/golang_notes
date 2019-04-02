package main

import "fmt"

func main() {
	// Arrays
	// var fruitArray [2]string
	fruitArray := [2]string{"Apple", "Orange"}
	// fruitArray[0] = "Kiwi"
	// fruitArray[1] = "Watermelon"
	fmt.Println("fruitArray = ", fruitArray)

	// Slices
	fruitSlice := []string{"Bannana", "Dragon fruit", "Kumkwat"}
	fmt.Println("fruitSlice = ", fruitSlice)
	fmt.Println("Program finished.")
}
