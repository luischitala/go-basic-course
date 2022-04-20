package main

import "fmt"

func main() {
	// modulo := 5 % 2
	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")

	}
	// Switch sin condición
	value := 100
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 100")
	default:
		fmt.Println("No condición")
	}
}
