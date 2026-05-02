package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	name := "Piyush"
	age := 24

	fmt.Printf("Name: %s, Age: %d\n", name, age)

	s := fmt.Sprintf("Go version %d.%d", 1, 24)
	fmt.Println(s)

	fmt.Print("No nweline here")
}
