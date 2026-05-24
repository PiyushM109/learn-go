package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

type Address struct {
	Street string
	City   string
	Zip    string
}

type Employee struct {
	Person  //embedded
	Address Address
	Company string
	Salary  float64
}

func main() {
	fmt.Println("Struct tut..")

	var p Person

	p.Name = "Piyush"
	p.Age = 24
	p.Email = "example@gmail.com"

	fmt.Println(p)

	//Nested struct

	e := Employee{
		Person:  Person{Name: "Piyush", Age: 25},
		Address: Address{City: "Pune", Zip: "411001"},
		Company: "Example.inc",
		Salary:  1500,
	}

	fmt.Println(e)

	t := &Person{Name: "Raj", Age: 24}

	fmt.Print(t.Name)
}
