package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// fmt.Println("Control Flow")

	// x := 18

	// if x > 18 {
	// 	fmt.Println("You can vote")
	// } else if x == 18 {
	// 	fmt.Println("You have to register")
	// } else {
	// 	fmt.Println("Can not vote")
	// }

	// if remainder := x % 2; remainder == 0 {
	// 	fmt.Println("passed")
	// } else {
	// 	fmt.Println("Failed")
	// }

	// fmt.Println("Switch Cases")

	day := time.Now().Weekday()
	fmt.Println(day)
	fmt.Println(time.Now())
	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	case time.Monday:
		fmt.Println("Monday Blues")
	case time.Friday:
		fmt.Println("Fiday hai friday B.....")
	default:
		fmt.Println("Weekday Boring")
	}
	fmt.Println(runtime.GOOS)
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Running on macOS")
	case "linux":
		fmt.Println("Running on Linux")
	default:
		fmt.Printf("Running on %s\n", os)
	}
}
