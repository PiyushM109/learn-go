package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name     string
	ID       string
	Homework float64
	Midterm  float64
	Final    float64
}

type Classroom map[string]*Student

func GPA(s *Student) float64 {
	return s.Homework*0.3 + s.Midterm*0.3 + s.Final*0.4
}

func LatterGrade(s *Student) string {
	gpa := GPA(s)
	switch {
	case gpa >= 90 && gpa <= 100:
		return "A"
	case gpa >= 80 && gpa < 90:
		return "B"
	case gpa >= 70 && gpa < 80:
		return "C"
	case gpa >= 60 && gpa < 70:
		return "D"
	default:
		return "F"
	}
}

func seed() {
	// TO BE Implemented later
}

func classAverage(cRoom Classroom) float64 {
	var total float64 = 0
	for _, student := range cRoom {
		gpa := GPA(student)
		total += gpa
	}
	return total / float64(len(cRoom))
}

func topStudent(cRoom Classroom) Student {
	var max float64 = 0
	var top Student

	for _, student := range cRoom {
		gpa := GPA(student)

		if gpa > max {
			top = *student
		}
	}
	return top
}

func bottomStudent(cRoom Classroom) Student {
	var min float64 = math.MinInt64
	var btm Student

	for _, student := range cRoom {
		gpa := GPA(student)

		if gpa < min {
			btm = *student
		}
	}
	return btm
}

func rangeDistribution(cRoom Classroom) map[string]int {
	mp := make(map[string]int)
	mp["A"] = 0
	mp["B"] = 0
	mp["C"] = 0
	mp["D"] = 0
	mp["F"] = 0
	for _, student := range cRoom {
		grade := LatterGrade(student)
		mp[grade]++
	}
	return mp
}

func sampleData() Classroom {
	return Classroom{
		"S001": {Name: "Alice Johnson", ID: "S001", Homework: 95, Midterm: 88, Final: 92},
		"S002": {Name: "Bob Smith", ID: "S002", Homework: 72, Midterm: 65, Final: 70},
		"S003": {Name: "Carol Williams", ID: "S003", Homework: 88, Midterm: 91, Final: 87},
		"S004": {Name: "David Brown", ID: "S004", Homework: 60, Midterm: 55, Final: 58},
		"S005": {Name: "Eve Davis", ID: "S005", Homework: 98, Midterm: 95, Final: 99},
	}
}

func printReport(c Classroom) {
	// Convert map to sorted slice
	students := make([]*Student, 0, len(c))
	for _, s := range c {
		students = append(students, s)
	}

	// TODO: Sort by GPA descending using sort.Slice

	// TODO: Print the formatted table
	fmt.Println("╔═══════════════════════════════════════════════╗")
	fmt.Println("║          STUDENT GRADE REPORT                 ║")
	fmt.Println("╠═══════════════════════════════════════════════╣")
	fmt.Printf("║ %-18s %5s %5s %5s %5s %2s ║\n",
		"Student", "HW", "Mid", "Fin", "GPA", "Gr")
	fmt.Println("╠═══════════════════════════════════════════════╣")
	for _, s := range students {
		fmt.Printf("║ %-18s %5.1f %5.1f %5.1f %5.1f %2s ║\n",
			s.Name, s.Homework, s.Midterm, s.Final, GPA(s), LatterGrade(s))
	}
	fmt.Println("╚═══════════════════════════════════════════════╝")
}

func main() {
	classroom := sampleData()
	printReport(classroom)

	avg := classAverage(classroom)
	fmt.Printf("\nClass Average: %.1f\n", avg)

	top := topStudent(classroom)
	fmt.Printf("\nTop in class is %s \n", top.Name)
}
