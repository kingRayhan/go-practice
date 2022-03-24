package main

import "fmt"

func aFunction() {
	fmt.Println("Function 1")
	defer fmt.Println("Function 2")
	fmt.Println("Function 3")
}

func main() {

	// defer statement will execute just before the return statement
	// defer statement will execute in reverse order
	defer fmt.Println("Defer line 1")
	defer fmt.Println("Defer line 2")
	defer fmt.Println("Defer line 3")

	aFunction()

	fmt.Println("Without defer 1")
	fmt.Println("Without defer 2")

}
