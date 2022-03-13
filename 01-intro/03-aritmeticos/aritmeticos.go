package main

import "fmt"

func main() {
	a := 20
	b := 10

	// Suma
	result := a + b
	fmt.Println("Suma: ", result)

	// Resta
	result = a - b
	fmt.Println("Resta: ", result)

	// Multiplicacion
	result = a * b
	fmt.Println("Multi: ", result)

	// Division
	result = a / b
	fmt.Println("Division: ", result)

	/* Division entre enteros con resultado flotante
	Hay que tener en cuenta que si el resultado es un flotante
	pero los factores son enteros, no vamos a obtener la parte decimal.

	hay que castearlos primero a flotantes para que el resultado sea flotante
	*/
	var div float64 = 3.0 / 2.0
	fmt.Println("Division flotante: ", div)

	// Modulo
	result = 3 % 2
	fmt.Println("Modulo: ", result)

}
