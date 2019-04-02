package main

import (
	"fmt"
	"math"

	"github.com/HarryMWinters/go_crash_course/strutil"
)

func main() {
	var square = (math.Sqrt(2.3))
	fmt.Println(square)
	fmt.Printf("%T\n", square)
	myString := "Hello"
	fmt.Println(strutil.Reverse(myString))

}
