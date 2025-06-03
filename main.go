package main

import "fmt"

func main() {
	// Declairing variables like age, name, height, state, etc.
	name := "Ejembi"
	var age int = 23
	var height float64 = 5.9
	var occupation string = "Soldier"
	var isMarried bool = true
	isStudent := false
	var nationality string = "Nigerian"
	state := "Kwararafa"

	fmt.Println(name, age, height, occupation, isMarried, isStudent, nationality, state)
	// function greet
	greet("Ejembi")
	// function addNumbers. it helps to sum up two integers or an int + float64, etc.
	result := addNumbers(24, 24.5)
	fmt.Println(result)
}

// function greet, with the concatination of the message and name
func greet(name string) {
	message := "Abole, " + name + "?"
	fmt.Println(message)
}

// sum of float64 of y and z resulting in an integer
func addNumbers(y float64, z float64) int {
	return int(y + z)
}
