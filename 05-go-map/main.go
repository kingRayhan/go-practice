package main

import "fmt"

func main() {
	// var passed_students map[string]bool
	var passed_students = make(map[string]bool)

	passed_students["John"] = true
	passed_students["Mary"] = true
	passed_students["Peter"] = true
	passed_students["Paul"] = true
	passed_students["Mary"] = false

	if passed_students["John"] {
		fmt.Println("John passed the exam")
	} else {
		fmt.Println("John failed the exam")
	}

	fmt.Println("---------- Map Iteration ----------")
	for name, isPassed := range passed_students {
		if isPassed {
			fmt.Print(name + " has passed the exam.\n")
		} else {
			fmt.Print(name + " hasn't passed the exam.\n")
		}
	}

	fmt.Println("---------- Delete key ----------")
	delete(passed_students, "Paul")
	for name, isPassed := range passed_students {
		if isPassed {
			fmt.Print(name + " has passed the exam.\n")
		} else {
			fmt.Print(name + " hasn't passed the exam.\n")
		}
	}

}
