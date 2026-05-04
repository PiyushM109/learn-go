package main

import "fmt"

func main() {
	fmt.Println("Loops")

	// for i := 0; i < 6; i++ {
	// 	fmt.Println(i)
	// }

	// n := 1
	// for n < 20 {
	// 	fmt.Println(n)
	// 	n = n * 2
	// }

	fruits := []string{"Piyush", "Girish", "More"}

	for i, v := range fruits {
		fmt.Println(i, v)
	}

}
