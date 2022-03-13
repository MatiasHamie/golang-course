package main

import "fmt"

func calcularIgv(vv, igv float64) float64 {
	return igv * vv
}

func precioVenta(vv, igv float64) float64 {
	return vv + igv
}

func main() {
	var vv, igv, pv float64

	fmt.Print("Ingrese Valor de Venta: ")
	fmt.Scanln(&vv)

	igv = calcularIgv(vv, 0.18)

	pv = precioVenta(vv, igv)

	fmt.Println("==========================")
	fmt.Println("Valor de Venta:", vv)
	fmt.Println("IGV:", igv)
	fmt.Println("Precio de Venta:", pv)
	fmt.Println("==========================")
}
