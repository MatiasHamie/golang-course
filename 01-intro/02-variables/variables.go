package main

import "fmt"

func main() {
	/*
		Estas son diferentes formas de declar el tipo y value de una varibale
	*/
	var nombre1 string
	nombre1 = "Matias"

	var nombre2 string = "Roel"

	edad := 26
	fmt.Println(nombre1, nombre2, edad)

	/*
		Por defecto ya con declarar y ponerle el tipo, se le pone el valor por defecto
	*/
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	const pi = 3.141592
	fmt.Println(pi)
}
