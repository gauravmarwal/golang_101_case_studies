package main

import "fmt"

func main() {
	// Creating a map of string slices.
	m := map[string][]string{
		"akola":  {"cotton", "wheat"},
		"nagpur": {"rice"},
	}

	res := append(m["akola"], "corn")
	fmt.Println(res)

	// Add a key for fish.
	m["mumbai"] = []string{"vada", "pav"}

	// Print slice at key.
	fmt.Println(m["mumbai"])

	// Loop over string slice at key.
	for i := range m["mumbai"] {
		fmt.Println(i, m["mumbai"][i])
	}
}
