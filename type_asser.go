// Golang program to show type
// assertions with error checking
package main

import (
	"fmt"
)

// main function
func main() {

	// an interface that has
	// an int value
	var value interface{} = 20024

	// retrieving a value
	// of type int and assigning
	// it to value1 variable
	var value1 int = value.(int)

	// printing the concrete value
	fmt.Println(value1)

	// this will test if interface
	// has string type and
	// return true if found or
	// false otherwise
	value2, test := value.(string)
	if test {

		fmt.Println("String Value found!")
		fmt.Println(value2)
	} else {

		fmt.Println("String value not found!")
	}
}
