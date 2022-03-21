package main

import (
	"fmt"
	"time"
)

// https://stackoverflow.com/a/20531179/3705299

func main() {
	fmt.Println("Welcome to time study of golang")
	currentTime := time.Now()
	fmt.Println("Current time is: ", currentTime)
	fmt.Print("Enter a time in format: " + currentTime.Format("Mon 01 2006 02 - 15:04:05 pm"))

	// Custom date object
	custom := time.Date(2048, time.March, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("Custom date object: ", custom)
	fmt.Println("Custom day: ", custom.Format("Monday"))
}
