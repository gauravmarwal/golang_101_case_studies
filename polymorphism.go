package main

import "fmt"

// Interface
type bank interface {
	branch_id() int
	name() string
}

// Structure 1
type branch1 struct {
	b1_id  int
	name_1 string
}

func (b1 branch1) branch_id() int {
	return b1.b1_id
}

func (b1 branch1) name() string {
	return b1.name_1
}

// Structure 2
type branch2 struct {
	b2_id  int
	name_2 string
}

// Methods of employee interface are
// implemented by the team2 structure
func (b2 branch2) branch_id() int {
	return b2.b2_id
}

func (b2 branch2) name() string {
	return b2.name_2
}

func allbranches(i []bank) {

	for _, ele := range i {

		fmt.Print("\nBranch Id:", ele.branch_id())
		fmt.Print("\nBranch Name: ", ele.name())

	}
}

// Main function
func main() {

	res1 := branch1{b1_id: 20,
		name_1: "Akola"}

	res2 := branch2{b2_id: 35,
		name_2: "Mumbai"}

	final := []bank{res1, res2}
	allbranches(final)

}
