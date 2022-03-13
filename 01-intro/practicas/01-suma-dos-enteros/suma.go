package main

import (
	"fmt"
)

func suma(a, b int) int {
	return a + b
}

func main() {
	var num1, num2 int

	fmt.Print("Ingrese primer numero: ")
	fmt.Scanln(&num1)

	fmt.Print("Ingrese segundo numero: ")
	fmt.Scanln(&num2)

	resultado := suma(num1, num2)

	fmt.Println("La suma es: ", resultado)
}
