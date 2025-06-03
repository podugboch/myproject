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

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("Height:", p.height)
	fmt.Println("Occupation:", p.occupation)
	fmt.Println("Married:", p.isMarried)
	fmt.Println("Student:", p.isStudent)
	fmt.Println("Nationality:", p.nationality)
	fmt.Println("State:", p.state)
}

func main() {
	Profile := createProfile("Patrick Odugboche", 29, 5.9, "Programmer", true, false, "Nigerian", "Kwararafa")
	viewProfile(Profile)
}
