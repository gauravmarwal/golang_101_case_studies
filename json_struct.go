package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// IMPORTANT: Struct names should have first letter capital (exported) to be compatible with JSON
type Accounts struct {
	Accounts []Account `json:"account"`
}

type Account struct {
	account_name                  string                     `json:"AccountName"`
	beneficiaries                 []string                   `json:"Beneficiaries"`
	bank_balance                  int                        `json:"BankBalance"`
	bank_status                   bool                       `json:"BankStatus"`
	bank_feeds_account_id         int                        `json:"BankFeedsAccountId"`
	code                          int                        `json:"Code"`
	default_invoice_payment_type  DefaultInvoicePaymentType  `json:"DefaultInvoicePaymentType"`
	default_purchase_payment_type DefaultPurchasePaymentType `json:"DefaultPurchasePaymentType"`
}

type DefaultInvoicePaymentType struct {
	Id   int `json:"Id"`
	Name int `json:"Name"`
}

type DefaultPurchasePaymentType struct {
	Id   int `json:"Id"`
	Name int `json:"Name"`
}

func main() {

	jsonFile, err := os.Open("accountsnew.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened JSON file")

	defer jsonFile.Close() // File isn't closed till all the above functions are executed.

	// We convert the data we've read from the file into a  byte array using ioutil.ReadAll
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Intialize our accounts array/struct
	var accounts Accounts

	// Now we unmarshall our content
	json.Unmarshal(byteValue, &accounts)

	// We iterate over every account in the array and print its details
	for index, value := range accounts.Accounts {
		fmt.Println("End User Number:", index)
		fmt.Println("AccountName: ", value.account_name)
		for _, name := range value.beneficiaries {
			fmt.Println("Beneficiaries: ", name)
		}
		fmt.Println("BankBalance: ", value.bank_balance)
		fmt.Println("BankStatus: ", value.bank_status)
		fmt.Println("Code: ", value.code)
		fmt.Println("DefaultInvoicePaymentType Id: ", value.default_invoice_payment_type.Id)
		fmt.Println("DefaultInvoicePaymentType Name: ", value.default_invoice_payment_type.Name)
		fmt.Println("DefaultPurchasePaymentType Id: ", value.default_purchase_payment_type.Id)
		fmt.Println("DefaultPurchasePaymentType Name: ", value.default_purchase_payment_type.Name)
		fmt.Println("\n****************************************\n")
	}

}
