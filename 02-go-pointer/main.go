package main

import "fmt"

func changeValue(x *int) {
	*x = *x * 2
}

func main() {
	fmt.Print("Pointer study\n")

	// pass by value
	// pass by reference

	var a int = 10

	// print memory address of a
	fmt.Println(&a)
	// print value of a
	fmt.Println(*&a)

	var ptr *int

	// assign memory address of a to ptr
	ptr = &a

	// print memory address of ptr
	fmt.Println(*ptr)

	*ptr = 200

	// print value of a
	fmt.Println(a)

	fmt.Print("Calling changeValue()\n")
	changeValue(&a)

	fmt.Println("a is: ", a)
}
