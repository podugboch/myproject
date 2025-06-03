package main

import "fmt"

type Profile struct {
	name        string
	age         int
	height      float64
	occupation  string
	isMarried   bool
	isStudent   bool
	nationality string
	state       string
}

func createProfile(name string, age int, height float64, occupation string, isMarried bool, isStudent bool, nationality string, state string) Profile {
	return Profile{name, age, height, occupation, isMarried, isStudent, nationality, state}
}

func viewProfile(p Profile) {

	fmt.Printf("%-15s %s\n", "Name:", p.name)
	fmt.Printf("%-15s %d\n", "Age:", p.age)
	fmt.Printf("%-15s %.1f\n", "Height:", p.height)
	fmt.Printf("%-15s %s\n", "Occupation:", p.occupation)
	fmt.Printf("%-15s %v\n", "Married:", p.isMarried)
	fmt.Printf("%-15s %v\n", "Student:", p.isStudent)
	fmt.Printf("%-15s %s\n", "Nationality:", p.nationality)
	fmt.Printf("%-15s %s\n", "State:", p.state)
}

func main() {
	Profile := createProfile("Patrick Sean Odugboche", 29, 5.9, "Programmer", true, false, "Nigerian", "Kwararafa")
	viewProfile(Profile)
}
