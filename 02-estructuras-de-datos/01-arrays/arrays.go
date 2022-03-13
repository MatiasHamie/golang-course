package main

import "fmt"

func main() {
	// por defecto te crea todos con el valor por defecto
	// del tipo de dato
	var numeros [5]int
	// [0 0 0 0 0]
	fmt.Println(numeros)

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	// [10 20 30 0 0]
	fmt.Println(numeros)

	nombres := [3]string{"Alex", "Pepe", "Juan"}
	// [Alex Pepe Juan]
	fmt.Println(nombres)

	/* los ... lo que hacen es hacer que el array sea del tamaño
	igual a las variables q se estan a guardar
	*/

	colores := [...]string{
		"Rojo",
		"Verde",
		"Negro",
		"Azul",
	}
	// [Rojo Verde Negro Azul] 4
	fmt.Println(colores, len(colores))

	monedas := [...]string{
		0: "Dolar",
		2: "Soles",
		3: "Pesos",
		5: "Euro",
	}
	// [Dolar  Soles Pesos  Euro] 6
	fmt.Println(monedas, len(monedas))

	/* sub aray
	extraemos un array, diciendo desde que indice
	a que indice queremos traernos
	subMoneda := monedas[0:3]

	si queremos decir que agarre desde el inicio, no
	hace falta aclarar el indice 0
	subMoneda := monedas[:3]

	esto es lo mismo pero hasta el final
	subMoneda := monedas[3:]
	*/
	subMoneda := monedas[0:3]
	// [Dolar  Soles]
	fmt.Println(subMoneda)
}
