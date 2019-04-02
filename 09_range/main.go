package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 4, 6, 645}
	fmt.Println(ids)

	sum := 0
	for _, id := range ids {
		fmt.Printf("%d\n", id)
		sum += id
	}
	fmt.Println(sum)

	myMap := map[string]int{"foo": 2, "bar": 3, "baz": 5}

	for k, v := range myMap {
		fmt.Println(k, " ", v)
	}
}
