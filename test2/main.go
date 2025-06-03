package main

import "fmt"

type Profile struct {
	Name        string
	Age         int
	Occupation  string
	Height      float64
	IsMarried   bool
	IsStudent   bool
	Nationality string
	State       string
}

func createProfile(name string, age int, occupation string, height float64, isMarried bool, isStudent bool, nationality string, state string) Profile {
	return Profile{name, age, occupation, height, isMarried, isStudent, nationality, state}
}

func viewProfile(p Profile) {

	fmt.Printf("%-15s %s\n", "name", p.Name)
	fmt.Printf("%-15s %d\n", "age", p.Age)
	fmt.Printf("%-15s %s\n", "occupation", p.Occupation)
	fmt.Printf("%-15s %.1f\n", "height", p.Height)
	fmt.Printf("%-15s %v\n", "isMaried", p.IsMarried)
	fmt.Printf("%-15s %v\n", "isStudent", p.IsStudent)
	fmt.Printf("%-15s %s\n", "nationality", p.Nationality)
	fmt.Printf("%-15s %s\n", "state", p.State)

}

func inputProfile() Profile {

	var name, occupation, nationality, state string
	var age int
	var height float64
	var isMarried, isStudent bool

	fmt.Print("Enter name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter age: ")
	fmt.Scanln(&age)

	fmt.Print("Enter height: ")
	fmt.Scanln(&height)

	fmt.Print("Enter occupation: ")
	fmt.Scanln(&occupation)

	fmt.Print("Are you married? yes/no: ")
	fmt.Scanln(&isMarried)

	fmt.Print("Are you a student? yes/no: ")
	fmt.Scanln(&isStudent)

	fmt.Print("Enter nationality: ")
	fmt.Scanln(&nationality)

	fmt.Print("Enter state: ")
	fmt.Scanln(&state)

	return createProfile(name, age, occupation, height, isMarried, isStudent, nationality, state)
}

func main() {

	profile := inputProfile()
	viewProfile(profile)

}
