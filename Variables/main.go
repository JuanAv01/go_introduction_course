package main

import "fmt"

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
}