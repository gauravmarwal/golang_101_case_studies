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

func printInterface(trv revenueType) error {
	fmt.Println("Initial Value of interface is: ", trv)
	fmt.Println("Initial Type of interface is: ", trv)
	return nil
}

func returnN(bank idfcBank) (int, string, error) {
	result := "Hello"
	return bank.rv, result, nil
}

func main() {
	var trv revenueType
	var bank idfcBank
	printInterface(trv) // Initial value and type of a interface is <nil>

	fmt.Println(returnN((bank)))

	trv = idfcBank{100, "google"}
	fmt.Println("Revenue Type: ", trv.profitLoss())
}
