package mypackage

import "fmt"

// CarPublic Car with public access
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

// PrintMessage imprimir un mensaje
func PrintMessage(message string) {
	fmt.Println(message)
}

// PrintMessage imprimir un mensaje privada
func printMessagel(message string) {
	fmt.Println(message)
}
