package main

import "fmt"

func main() {
	// fmt.Println("Arrays and Slices")
	// var primes [3]int
	// primes[0] = 2
	// primes[1] = 3
	// primes[2] = 5

	// //Slices from array

	// arr := [7]int{1, 2, 3, 4, 5, 6, 7}

	// s := arr[1:4]
	// fmt.Println(s)
	// fmt.Println(len(s), cap(s))
	// s[2] = 99
	// fmt.Println(arr, s)

	// s := make([]int, 0)
	// fmt.Println(s, len(s), cap(s))
	// s = append(s, 4)
	// fmt.Println(s, len(s), cap(s))
	// s = append(s, 8)
	// fmt.Println(s, len(s), cap(s))
	// s = append(s, 8)
	// fmt.Println(s, len(s), cap(s))
	// s = append(s, 8)
	// fmt.Println(s, len(s), cap(s))
	// s = append(s, 8)
	// fmt.Println(s, len(s), cap(s))
	// s = append(s, 8, 8, 8, 8, 8, 8, 8, 8, 8)
	// fmt.Println(s, len(s), cap(s))
	// s = append(s, 8, 8, 8, 8, 8, 8, 8, 8, 8)
	// fmt.Println(s, len(s), cap(s))
	// s = append(s, 8, 8, 8, 8, 8, 8, 8, 8, 8)
	// fmt.Println(s, len(s), cap(s))
	// s = append(s, 8, 8, 8, 8, 8, 8, 8, 8, 8)
	// fmt.Println(s, len(s), cap(s))
	// s = append(s, 8, 8, 8, 8, 8, 8, 8, 8, 8)
	// fmt.Println(s, len(s), cap(s))
	// s = append(s, 8, 8, 8, 8, 8, 8, 8, 8, 8)
	// fmt.Println(s, len(s), cap(s))
	// s = append(s, 8, 8, 8, 8, 8, 8, 8, 8, 8)
	// fmt.Println(s, len(s), cap(s))
	// s = append(s, 8, 8, 8, 8, 8, 8, 8, 8, 8)
	// fmt.Println(s, len(s), cap(s))

	fmt.Println("Stack implementation with the slices")
	// s := make([]int, 0)
	// fmt.Println(s)
	// st := stackImpl()
	// st.push(4)
	// st.push(5)
	// st.push(9)
	// st.display()
	// st.pop()
	// st.display()
}

type stack struct {
	push    func(int)
	pop     func()
	peek    func() int
	display func()
}

func stackImpl() stack {
	s := make([]int, 0)
	push := func(num int) {
		s = append(s, num)
	}

	pop := func() {
		if len(s) <= 0 {
			return
		}
		s = s[:len(s)-1]
	}

	peek := func() int {
		if len(s) <= 0 {
			return -1
		}
		return s[len(s)-1]
	}
	display := func() {
		fmt.Println(s)
	}
	return stack{
		push:    push,
		pop:     pop,
		peek:    peek,
		display: display,
	}
}
