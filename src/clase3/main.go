package main

import "fmt"

func main() {

	//Area cuadrado
	const baseCuadrado = 10
	const baseRectangulo = 10
	const alturaRectangulo = 5
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma:", result)

	// Resta
	result = y - x
	fmt.Println("Resta:", result)

	// Multiplicación
	result = x * y
	fmt.Println("Multiplicación:", result)

	// División
	result = y / x
	fmt.Println("División:", result)

	// Módulo
	result = y % x
	fmt.Println("Módulo:", result)

	// Incremental
	x++
	fmt.Println("Incremental:", x)

	// Decremental
	x--
	fmt.Println("Decremental:", x)

	// Retos
	// Área Rectángulo
	result = baseRectangulo * alturaRectangulo
	fmt.Println("Área rectángulo:", result)
	// Área Trapecio

	// Área Círculo

}
