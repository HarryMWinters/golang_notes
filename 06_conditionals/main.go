package main

import "fmt"

func main() {
	x := 51
	y := 51

	if x < y {
		fmt.Printf(
			"%d is less than %d.\n", x, y)
	} else if !(x <= y) {
		fmt.Printf(
			"%d is NOT less then %d\n", x, y)
	} else {
		fmt.Printf(
			"%d is equal to %d\n", x, y)
	}

	color := "bue"

	switch color {
	case "red":
		fmt.Println("Color is red.")
	case "blue":
		fmt.Println("Color is blue.")
	default:
		fmt.Println("Color is neither red nor blue.")
	}
}
