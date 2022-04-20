package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es Palindromo")

	} else {
		fmt.Println("No es Palindromo")

	}
}

func main() {
	slice := []string{"hola", "que", "hace"}

	for i := range slice {
		fmt.Println(i)
	}

	// Palíndromo
	minus := strings.ToLower("ama")
	isPalindromo(minus)
}
