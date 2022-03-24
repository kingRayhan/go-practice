package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func addAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// Anonymous function
var greet = func(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

// ------------------------ Multi return value ------------------------
func swap(a, b int) (int, int) {
	return b, a
}

func namedMultireturn() (first int, second int, third int) {
	first = 1
	second = 2
	third = 3
	return
}

func multiReturn1() (a, b, c int) {
	return 1, 2, 3
}

func multiReturn2() (a, b, c int, d, e bool, f string) {
	a = 1
	b = 2
	c = 3
	d = true
	e = false
	f = "Hello"
	return
	// or ->> return 1, 2, 3, true, false, "Hello"
}

func main() {
	fmt.Println("Study go function and methods")

	fmt.Println("add(1, 2): ", add(1, 2))
	fmt.Println("addAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10): ", addAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println("greet('Rayhan'): ", greet("Rayhan"))

	a, b := swap(1, 2)
	fmt.Println("swap(1, 2) -> a,b: ", a, b)

	a, b, c := multiReturn1()
	fmt.Println("multiReturn1() -> a,b,c: ", a, b, c)

	first, second, third := namedMultireturn()
	fmt.Println("namedMultireturn() -> first,second,third: ", first, second, third)

	a, b, c, d, e, f := multiReturn2()
	fmt.Println("multiReturn2() -> a,b,c,d,e,f: ", a, b, c, d, e, f)
}
