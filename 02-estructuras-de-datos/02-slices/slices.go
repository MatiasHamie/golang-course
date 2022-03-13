package main

import "fmt"

func main() {
	// Slices: SON MUTABLES
	numeros := []int{1, 2, 3}

	fmt.Println(numeros, len(numeros))

	// agregar datos
	numeros = append(numeros, 4)
	numeros = append(numeros, 5)

	fmt.Println(numeros, len(numeros))

	// sub slice
	subSlice := numeros[:2]

	numeros[0] = 100

	fmt.Println(subSlice)
	fmt.Println(numeros)

	meses := []string{"Enero", "Febrero", "Marzo"}

	fmt.Printf("Longitud: %v, Capacidad: %v, %p\n", len(meses), cap(meses), meses)

	meses = append(meses, "Abril")
	fmt.Printf("Longitud: %v, Capacidad: %v, %p\n", len(meses), cap(meses), meses)

}
