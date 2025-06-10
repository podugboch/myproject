package main

import "fmt"

func main() {
	xyz := [2]string{"big", "tasty"}
	abc := [3]string{"apple", "orange", "banana"}
	for i := 0; i < len(xyz); i++ {
		for j := 0; j < len(abc); j++ {
			fmt.Println(xyz[i], abc[j])
		}
	}

}
