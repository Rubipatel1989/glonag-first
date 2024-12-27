package main

import "fmt"

func main() {
	studentGrades := make(map[string]int)

	studentGrades["Prince"] = 34
	studentGrades["Vinay"] = 30
	studentGrades["Bob"] = 34
	studentGrades["Cikago"] = 12
	studentGrades["Rahul"] = 90

	fmt.Println("Marks of Bob is :", studentGrades["Bob"])
	studentGrades["Bob"] = 95
	fmt.Println("New Marks of Bob is :", studentGrades["Bob"])
	delete(studentGrades, "Bob")
	fmt.Println("Marks", studentGrades)

	for index, value := range studentGrades {
		// fmt.Printf("Key is %s and marks is %s\n", index, value)
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}
}
