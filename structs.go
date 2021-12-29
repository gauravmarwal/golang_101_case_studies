package main

import "fmt"

type revenueType interface {
	profitLoss() string // This is an interface
}

type idfcBank struct {
	rv    int
	cname string
}

type rbi struct {
	details    idfcBank // This is a nested struct
	guidelines string
}

func (m idfcBank) profitLoss() string {
	if m.rv > 0 {
		return "profitable"
	} else {
		return "loss making"
	}
}

func printEmptyStruct(a idfcBank) {
	fmt.Println("Empty Structure: ", a) // Prints {0,0} cause structs are initialized with zero
}

func printStruct(a1 idfcBank) {
	fmt.Println("Structure: ", a1)
	fmt.Println("Company Name: ", a1.cname) // We use "." to access individual elements of a struct
}

func printPointer(aptr *idfcBank) {

	fmt.Println("Pointer:", aptr)
	fmt.Println("Company Name: ", (*aptr).cname) // We deference the pointer to access the elements
	fmt.Println("Company Name: ", aptr.cname)    // IMPORTANT: Golang allows us the use without deferencing the pointer

}

func printNestedStruct(b1 rbi) {

	fmt.Println("Direct Print: ", b1)
	fmt.Println("Company Name: ", b1.details.cname)
	fmt.Println("Revenue: ", b1.details.rv)

}

func printInterface(trv revenueType) error {
	fmt.Println("Initial Value of interface is: ", trv)
	fmt.Println("Inital Type of interface is: ", trv)
	return nil
}

func returnN()

func main() {

	//structs
	//Go’s structs are typed collections of fields.
	//They’re useful for grouping data together to form records.

	var a idfcBank
	printEmptyStruct(a)

	temp := new(idfcBank) // declaration using new keyword and shorthand assignment
	printEmptyStruct(*temp)

	a1 := idfcBank{2000, "google"} // We can add values to struct using short hand assignment
	printStruct(a1)

	// Pointers to structs
	aptr := &a1
	printPointer(aptr)

	// Nested Structure
	// A struct inside a struct is called a nested struct

	b1 := rbi{
		details:    idfcBank{2000, "google"},
		guidelines: "This is a guideline",
	}

	printNestedStruct(b1)

	// Difference between structs and interfaces
	// Structs and interfaces are Go’s way of organizing methods and data handling. Where structs define the fields of an object, like a idfc's client name and revenue.
	// The interfaces define the methods; e.g. changing the name of any idfc client.

	// Interfaces
	// The interface is a custom type that is used to specify a set of one or more method signatures.

	var trv revenueType
	printInterface(trv) // Inital value and type of a interface is <nil>

	trv = idfcBank{100, "google"}
	fmt.Println("Revenue Type: ", trv.profitLoss())

	// Interfaces are of two types: 1) Static 2) Dynamic
	// The static type of an interface is the interface itself. Eg: revenueType is a interface here
	// An interface does not have a static value, rather it points to a dynamic value.

}
