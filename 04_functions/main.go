package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name + "\n"
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	x := getSum(3, 5)
	fmt.Println(x)
}
