package main

import "fmt"

func main() {
	nombre := "Matias"
	edad := 32

	fmt.Printf("Hola, %s y su edad es %d\n", nombre, edad)

	// si no sabemos q tipo de datos vamos a imprimir, ponemos %v
	// fmt.Printf("Hola, %v y su edad es %v\n", nombre, edad)

	mensaje := fmt.Sprintf("Hola, %v y su edad es %v\n", nombre, edad)

	fmt.Println(mensaje)

	// con %T imprimimos el tipo del dato
	fmt.Printf("nombre: %T \n", nombre)

	// Ingresar un dato
	fmt.Print("Ingrese otro nombre: ")
	fmt.Scanln(&nombre)

	fmt.Println("Otro nombre: ", nombre)
}
