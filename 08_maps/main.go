package main

import "fmt"

func main() {
	emails := make(map[string]string)
	//fmt.Printf("%T\n", emails)

	emails["harry"] = "hmw@foo.com"
	emails["harry2"] = "hmw@baz.com"

	delete(emails, "harry2")

	roles := map[string]string{"harry": "Baller", "Shem": "Player"}
	fmt.Println(emails)
	fmt.Println(roles)
}
