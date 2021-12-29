package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// Open our jsonFile
	jsonFile, err := os.Open("accountsnew.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened accountsnew.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	for _, value := range result {
		fmt.Println(result, " \n")
		for key, v := range value.(map[string]interface{}) {
			fmt.Println(key, " ", v, "\n")
			for _, v1 := range v.(map[string]interface{}) {
				fmt.Println(v1)
			}
		}
	}
}
