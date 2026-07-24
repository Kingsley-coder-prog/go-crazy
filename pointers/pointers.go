package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int // pointer variable declaration
	agePointer = &age   // pointer variable

	fmt.Println("Age: ", *agePointer) // prints the value of age variable

	// adultYears := getAdultYears(age)
	// fmt.Println("Adult years: ", adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}
