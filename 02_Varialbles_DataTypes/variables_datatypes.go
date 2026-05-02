package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Variable Declartion")

	// var name string = "Piyush"
	// var age = 24
	// var language string
	// language = "Marathi"

	// fmt.Println(name, age, language)

	// fullName := "Piyush More"

	// fmt.Println(fullName)

	// fmt.Println("Zero Values in the GO")

	// var i int
	// var f float64
	// var b bool
	// var s string

	// fmt.Printf("int:     %d\n", i)
	// fmt.Printf("float64: %f\n", f)
	// fmt.Printf("bool:    %t\n", b)
	// fmt.Printf("string:  %q\n", s)

	// fmt.Println("Declare multiple variables")
	// var x, y, z int = 10, 11, 12

	// fmt.Println(x, y, z)

	// p, q, r := 1, 2, 3

	// fmt.Println(p, q, r)

	// fmt.Println("Swap two numbers")

	// a, k := 12, 13

	// k, a = a, k

	// fmt.Println(a, k)

	// fmt.Println("Type conversion")

	// var i int = 40
	// var f float64 = float64(i)
	// var u uint = uint(f)
	// fmt.Println(u)

	n := 109

	s := strconv.Itoa(n)

	fmt.Println(s)

	n2, err := strconv.Atoi(s)

	if err == nil {
		fmt.Println(n2)
	}

	fmt.Printf("Type of s: %T, n2 %T", s, n2)

}
