package main

import "fmt"

func main() {

	name := "Ejembi"
	var age int = 23
	var height float64 = 5.9
	var occupation string = "Soldier"
	var isMarried bool = true
	isStudent := false
	var nationality string = "Nigerian"
	state := "Kwararafa"

	fmt.Println(name, age, height, occupation, isMarried, isStudent, nationality, state)

	greet("Ejembi")

	result := addNumbers(24, 24.5)
	fmt.Println(result)
}

func greet(name string) {
	message := "Abole, " + name + "?"
	fmt.Println(message)
}

func addNumbers(y int, z float64) float64 {
	return float64(y) + z
}
