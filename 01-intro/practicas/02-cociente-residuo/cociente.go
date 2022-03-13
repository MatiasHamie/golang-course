package main

import (
	"fmt"
)

func cociente(a, b int) int {
	return a / b
}

func resto(a, b int) int {
	return a % b
}

func main() {
	var num1, num2 int

	fmt.Print("Ingrese el primer numero: ")
	fmt.Scanln(&num1)

	fmt.Print("Ingrese el segundo numero: ")
	fmt.Scanln(&num2)

	fmt.Println("El cociente es: ", cociente(num1, num2))
	fmt.Println("El resto es: ", resto(num1, num2))

}
