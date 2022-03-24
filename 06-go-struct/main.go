package main

import (
	"fmt"
)

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func main() {
	user1 := User{
		Name:   "Rayhan",
		Email:  "rayhan.dev.bd@gmail.com",
		Age:    25,
		Status: true,
	}

	fmt.Println(user1)
	fmt.Printf("%+v\n", user1)
	fmt.Println("Name:", user1.Name)
	fmt.Println("Email:", user1.Email)

	// 5 users of struct
	users := []User{
		{
			Name:  "Rayhan",
			Email: "rayhan@go.dev",
		},
		{
			Name:  "Nibbi",
			Email: "nibbi@go.dev",
		},
	}

	fmt.Println(users)

	// append a new user
	users = append(
		users,
		User{Name: "Nibbi2", Email: "nibbi2@go.dev", Age: 26},
		User{Name: "Nibbi3", Email: "nibbi2@go.dev", Age: 27},
		User{Name: "Nibbi4", Email: "nibbi2@go.dev", Age: 28},
		User{Name: "Nibbi5", Email: "nibbi2@go.dev", Age: 29},
		User{Name: "Nibbi5", Email: "nibbi2@go.dev", Age: 29},
	)

	for index := range users {
		fmt.Println(users[index].Name)
	}

}
