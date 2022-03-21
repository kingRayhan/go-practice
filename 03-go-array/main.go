package main

import "fmt"

func main() {
	var fruits = [5]string{"apple", "banana", "cherry"}
	fmt.Println(fruits)
	fmt.Println("Length of array: ", len(fruits))

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println("Length of array: ", len(numbers))
}
