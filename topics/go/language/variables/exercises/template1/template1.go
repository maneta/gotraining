// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

import "fmt"

// main is the entry point for the application.
func main() {

	// Declare variables that are set to their zero value.
	var name string
	var age int
	var legal bool

	// Display the value of those variables.
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(legal)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	nameInt := "John"
	ageInt := 20
	legalInt := true
	// Display the value of those variables.
	fmt.Println("nameInt:", nameInt)
	fmt.Println("ageInt:", ageInt)
	fmt.Println("legalInt:", legalInt)

	//var pi float32
	// Perform a type conversion.
	pi := float32(3.14)

	// Display the value of that variable.
	fmt.Println("pi:", pi)
}
