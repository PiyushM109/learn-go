package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func greet() {
	fmt.Println("Hello, Gopher!")
}

func describe(name string, age int, active bool) string {
	status := "inactive"
	if active {
		status = "active"
	}
	return fmt.Sprintf("%s (age %d) — %s", name, age, status)
}

func stats(nums []float64) (min, max, avg float64) {
	if len(nums) == 0 {
		return 0, 0, 0
	}

	min, max = nums[0], nums[0]

	sum := 0.0

	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
		sum += n
	}

	avg = sum / float64(len(nums))

	return min, max, avg
}

func divide(a, b float64) (result float64, err error) {
	if b == 0 {
		err = fmt.Errorf("divisionn by zero")
		return
	}

	result = a / b
	return
}

func main() {
	fmt.Println(add(3, 4))
	greet()
	fmt.Println(describe("Piyush", 24, true))

	// nums := []float64{3.5, 4.0, 2.6, 8.9}

	// stats((nums))
	// min, max, avg := stats(nums)
	// fmt.Printf("Min: %.1f  Max: %.1f  Avg: %.2f\n", min, max, avg)

	res, err := divide(6, 9)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("10/3 = %.4f ", res)
	}
}
