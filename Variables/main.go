package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 

	/*
		Declaracion de variables (Var).
		
		- Tipos de datos:
			- Int: Tipo de dato entero con signo.
			- uint: Tipo de dato entero sin signo.
			- string: Tipo de dato para cadenas de texto.
			- bool: Tipo de dato booleano (true/false).
	*/

	var myIntVar int // Declaracion de variable entera
	myIntVar = -42
	fmt.Println("My Integer Variable:", myIntVar)

	var myUintVar uint // Declaracion de variable entera sin signo
	myUintVar = 42
	fmt.Println("My Unsigned Integer Variable:", myUintVar)

	var myStringVar string // Declaracion de variable string
	myStringVar = "Hola, mundo!"
	fmt.Println("My String Variable:", myStringVar)

	var myBoolVar bool // Declaracion de variable booleana
	myBoolVar = true
	fmt.Println("My Boolean Variable:", myBoolVar)

	intVar := 22
	fmt.Printf("type: %d, type: %T\n", intVar, intVar)
	intStrvar := "22"
	fmt.Printf("type: %s, type: %T\n", intStrvar, intStrvar)

	intval1, err := strconv.ParseInt("1234", 0, 64)
	fmt.Println(err) // err es nil si no hubo error
	fmt.Printf("type: %T, value: %d\n", intval1, intval1)

	intval2, err := strconv.ParseInt("aa12", 0, 64)
	fmt.Println(err)
	fmt.Printf("type: %T, value: %d\n", intval2, intval2)

floatVar1, _ := strconv.ParseFloat("3.14159", 64) // Ignorando el error con _
	fmt.Printf("type: %T, value: %f\n", floatVar1, floatVar1)
}