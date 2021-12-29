package main

import "fmt"

// Interface 1
type BankDetails interface {
	details()
}

// Interface 2
type BankPartners interface {
	partners()
}

// Interface 3

// Interface 3 embedded with
// interface 1 and 2
type FinalDetails interface {
	BankDetails
	BankPartners
}

// Structure
type bank struct {
	name    string
	branch  string
	partner string
}

// Implementing method of
// the interface 1
func (a bank) details() {

	fmt.Print("Bank Name: ", a.name)
	fmt.Print("\nBranch: ", a.branch)
}

// Implementing method
// of the interface 2
func (a bank) partners() {
	fmt.Print("\nPartner: ", a.partner)
}

// Main value
func main() {

	// Assigning values
	// to the structure
	values := bank{
		name:    "IDFC First Bank",
		branch:  "Akola",
		partner: "Gaurav",
	}

	// Accessing the methods of
	// the interface 1 and 2
	// Using FinalDetails interface
	var f FinalDetails = values
	f.details()
	f.partners()
}
