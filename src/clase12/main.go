package main

import "fmt"

func main() {
	//Dicts
	m := make(map[string]int)
	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Iterate map

	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontr valor
	value, ok := m["Jose"]
	fmt.Println(value, ok)
}
