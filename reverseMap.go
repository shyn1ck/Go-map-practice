package main

import (
	"fmt"
)

func main() {
	Map := map[string]string{
		"Alice": "111-111-111",
		"Alex":  "222-222-222",
		"Bob":   "333-333-333",
	}

	reverse := make(map[string]string)
	for key, value := range Map {
		reverse[value] = key
	}

	fmt.Println("First map:", Map)
	fmt.Println("Second map:", reverse)
}
