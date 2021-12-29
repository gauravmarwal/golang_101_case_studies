package main

import (
	"fmt"
)

type ProfitCalculator interface {
	DisplayProfit()
}

type ProfitableDaysCalculator interface {
	CalculateProfitableDays() int
}

type Bank struct {
	BankName  string
	BSA       int
	RSA       int
	TotalDays int
	LossDays  int
}

func (e Bank) DisplayProfit() {
	fmt.Print("Total Profit is: ", (e.BSA + e.RSA))
}

func (e Bank) CalculateProfitableDays() int {
	return e.TotalDays - e.LossDays
}

func main() {
	e := Bank{
		BankName:  "IDFC First Bank",
		BSA:       1000000,
		RSA:       10000,
		TotalDays: 365,
		LossDays:  1,
	}
	var s ProfitCalculator = e
	s.DisplayProfit()
	var l ProfitableDaysCalculator = e
	fmt.Println("\nProfitable Days: ", l.CalculateProfitableDays())
}
