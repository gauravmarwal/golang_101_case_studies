package main

import (
	"fmt"
)

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	s := "Hello World"
	describe(s)
	strt := struct {
		name string
	}{
		name: "Gaurav Marwal",
	}
	describe(strt)
}
