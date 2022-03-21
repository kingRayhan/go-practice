package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "cherry"}
	fruits = append(fruits, "orange", "grape")
	fruits = append(fruits, "pear")
	fmt.Println("Length of array: ", len(fruits))
	fmt.Println("Capacity of array: ", cap(fruits))

	fmt.Println("Full array ", fruits)
	fmt.Println("---------- Array range practice ----------")
	fmt.Println("Array of [1:]", fruits[1:])
	fmt.Println("Array of [2:]", fruits[2:])
	fmt.Println("Array of [:3]", fruits[:3])
	fmt.Println("Array of [:4]", fruits[:4])
	fmt.Println("Array of [:]", fruits[:])

	fmt.Println("---------- MAKE ----------")
}
