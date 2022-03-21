package main

import (
	"fmt"
	"sort"
)

func removeItemByIndex(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

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
	scores := make([]int, 3)
	scores[0] = 100
	scores[1] = 90
	scores[2] = 80
	scores = append(scores, 70, 75, 65, 95)
	fmt.Println("Scores: ", scores)

	fmt.Println("---------- SORT ----------")
	fmt.Println("Before sort: ", scores)
	fmt.Println("Is sorted: ", sort.IntsAreSorted(scores))
	sort.Ints(scores)
	fmt.Println("Sorted scores: ", scores)
	fmt.Println("Is sorted: ", sort.IntsAreSorted(scores))
	// sort descending order
	sort.Sort(sort.Reverse(sort.IntSlice(scores)))
	fmt.Println("Reverse sorted scores: ", scores)

	fmt.Println("---------- Remove a item ----------")
	fmt.Println("Before remove: ", scores)
	fmt.Println("After remove [2]: ", removeItemByIndex(scores, 2))
	fmt.Println("After remove [1]: ", removeItemByIndex(scores, 1))

}
