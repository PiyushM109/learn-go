package main

import "fmt"

// Package-level variables (only var, no :=)
var appName = "Go Mastery"
var version = 1

func main() {
	// TODO 1: Declare name using var with explicit type string
	var name string = "Piyush"
	fmt.Println(name)

	// TODO 2: Declare age using var with type inferred
	var age = 24
	fmt.Println(age)

	// TODO 3: Declare height using :=
	height := 5.8
	fmt.Println(height)

	// TODO 4: Declare isStudent using var, let it be zero value
	var isStudent bool
	fmt.Println(isStudent)

	// TODO 5: Declare x, y, z in one line using :=
	x, y, z := 10, 9, 2
	fmt.Println(x, y, z)

	// TODO 6: Use a var block for host, port, debug
	var (
		host  = "localhost"
		port  = 8080
		debug = false
	)
	fmt.Println(host, port, debug)

}
