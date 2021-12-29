package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["idfc1"] = 1
	m["idfc2"] = 2

	fmt.Println("Complete Map:", m)

	v1 := m["idfc1"]
	fmt.Println("idfc1: ", v1)

	// This prints the length
	fmt.Println("len:", len(m))

	//built in func called delelte can delete the value with key
	delete(m, "idfc2")
	fmt.Println("Map:", m)

	//The optional second return value when getting a value from a map indicates if the key was present in the map.
	//This can be used to disambiguate between missing keys and keys with zero values like 0
	_, prs := m["idfc2"]
	fmt.Println("Prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("Map:", n)
}
