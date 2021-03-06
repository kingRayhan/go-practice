package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func (user User) userDetails() {
	fmt.Println("--------------- User details ---------------")
	fmt.Println("Name:", user.Name)
	fmt.Println("Email:", user.Email)
	fmt.Println("Age:", user.Age)
}

type Student struct {
	Name  string
	Age   int
	Marks []int
}

func (student Student) studentDetails() {
	fmt.Println("--------------- Student details ---------------")
	fmt.Println("Name:", student.Name)
	fmt.Println("Age:", student.Age)
	fmt.Println("Marks:", student.Marks)
}

func (student *Student) updateMarks(marks []int) {
	student.Marks = marks
}

func (student Student) getGrade() string {
	var grade string
	total := student.getAverageGrade()

	if total >= 90 {
		grade = "A"
	} else if total >= 80 {
		grade = "B"
	} else if total >= 70 {
		grade = "C"
	} else if total >= 60 {
		grade = "D"
	} else {
		grade = "F"
	}
	return grade
}

func (student Student) getAverageGrade() float32 {
	var total int
	for _, mark := range student.Marks {
		total += mark
	}
	return float32(total / len(student.Marks))
}

func (student Student) getMaxMark() int {
	var max int
	for _, mark := range student.Marks {
		if mark > max {
			max = mark
		}
	}
	return max
}

func main() {
	fmt.Println("Study go struct methods")

	user1 := User{
		Name:  "Rayhan",
		Email: "rayhan@go.dev",
		Age:   25,
	}

	user1.userDetails()

	student1 := Student{
		Name:  "Pawl",
		Age:   26,
		Marks: []int{80, 90, 100},
	}

	fmt.Println(student1)
	student1.studentDetails()
	student1.updateMarks([]int{96, 84, 38})
	student1.studentDetails()
	fmt.Println("Grade:", student1.getGrade())
	fmt.Println("Average grade:", student1.getAverageGrade())
	fmt.Println("Max mark:", student1.getMaxMark())
}
