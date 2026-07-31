package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int // pointer variable declaration
	agePointer = &age   // pointer variable

	fmt.Println("Age: ", *agePointer) // prints the value of age variable

	editAgeToAdultYears(agePointer) // passing pointer to function
	fmt.Println("Adult years: ", age)
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
