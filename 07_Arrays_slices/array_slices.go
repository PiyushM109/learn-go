package main

import "fmt"

func main() {
	fmt.Println("Arrays and Slices")
	var primes [3]int
	primes[0] = 2
	primes[1] = 3
	primes[2] = 5

	//Slices from array

	arr := [7]int{1, 2, 3, 4, 5, 6, 7}

	s := arr[1:4]
	fmt.Println(s)
	fmt.Println(len(s), cap(s))
	s[2] = 99
	fmt.Println(arr, s)
}
