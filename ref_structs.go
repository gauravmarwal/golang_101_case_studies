package main

import (
	"fmt"
)

func main() {
	branchProfit := map[string]int{
		"Akola":  12000,
		"Nagpur": 15000,
		"Mumbai": 9000,
	}
	fmt.Println("Original Branch Profit", branchProfit)
	modified := branchProfit
	modified["Akola"] = 18000
	fmt.Println("Modified Branch Profit", branchProfit)

}
