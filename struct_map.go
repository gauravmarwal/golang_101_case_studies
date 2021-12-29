package main

import (
	"fmt"
)

type idfc struct {
	profit int
	branch string
}

func main() {
	branch1 := idfc{
		profit: 12000,
		branch: "Akola",
	}
	branch2 := idfc{
		profit: 14000,
		branch: "Mumbai",
	}
	branch3 := idfc{
		profit: 13000,
		branch: "Nagpur",
	}
	branchInfo := map[string]idfc{
		"Akola":  branch1,
		"Mumbai": branch2,
		"Nagpur": branch3,
	}

	for name, info := range branchInfo {
		fmt.Printf("Branch: %s Profit: Rs.%d \n", name, info.profit)
	}

}
