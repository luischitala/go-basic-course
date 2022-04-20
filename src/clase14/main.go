package main

import (
	pk "curso_golang_platzi/src/mypackage"
	"fmt"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	// var myOtherCar pk.carPrivate
	// fmt.Println(myOtherCar)

	pk.PrintMessage("Hola Platzi")
	pk.printMessagel("Hola Platzi")
}
