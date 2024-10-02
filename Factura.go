package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"time"
)

func datos_cliente() {
	var cliente string
	fmt.Println("Ingrese el nombre del cliente: ")
	fmt.Scan(&cliente)
	fmt.Println("Nombre del cliente: ", cliente)
}

func compras() {
	var opcion string
	var productos []string
	var precios []float64
	var cant_producto float64
	var valor_total float64

	for {
		var producto string
		var precio float64

		fmt.Println("Ingrese el nombre del producto: ")
		fmt.Scan(&producto)
		productos = append(productos, producto)

		fmt.Println("Ingrese el precio del producto: ")
		fmt.Scan(&precio)
		precios = append(precios, precio)


		fmt.Println("Ingrese la cantidad del producto: ")
		fmt.Scan(&cant_producto)

		valor := precio * cant_producto
		valor_total += valor

		fmt.Println("Desea agregar otro producto? (s/n): ")
		fmt.Scan(&opcion)

		if opcion != "s" {
			break
		}
	}
	fmt.Println("Sus productos son: ", productos)
	fmt.Println("los precios son: ", precios)
	fmt.Println("Valor total de la compra: $", valor_total)
	fmt.Println("Gracias por su compra!")
}

func Factura() {
	datos_cliente()
	compras()
}

func main() {
	t := time.Now()
	Factura()
	fmt.Println("Fecha y hora: ", t.Format("2006-01-02 15:04:05"))
	fmt.Println("Generando PDF...")
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Factura()")
	err := pdf.OutputFileAndClose("Factura.pdf")
	fmt.Println(err)
}