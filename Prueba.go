package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"time"
)

func datos_cliente(){
	var cliente string
	var cantidad int
	fmt.Println("Ingrese el nombre del cliente: ")
	fmt.Scan(&cliente)
	fmt.Println("Ingrese la cantidad del productos: ")
	fmt.Scan(&cantidad)
	fmt.Println("Nombre del cliente: ", cliente)
}

func compras() {
	var opcion string
	var producto string
	var productos []string
	var precio float64
	var cant_producto float64
	var valor float64
	var valor_total float64
	fmt.Println("Ingrese el nombre del producto: ")
	fmt.Scan(&producto)
	productos = append(productos, producto)
	fmt.Println("Ingrese el precio del producto: ")
    fmt.Scan(&precio)
	fmt.Println("Ingrese la cantidad del producto: ")
    fmt.Scan(&cant_producto)
	valor = precio * cant_producto
	valor_total = valor + valor_total
	fmt.Println("Desea agregar otro producto? (s/n): ")
	fmt.Scan(&opcion)
	switch opcion {
		case "s":
			compras()
		default:
			fmt.Println("Sus productos son: ", productos)
			fmt.Println("Valor total de la compra: $", valor_total)
			fmt.Println("Gracias por su compra!")
	}
}

func Factura(){
	datos_cliente()
	compras()
}

func main() {
	Factura()
	t := time.Now()
	fmt.Println("Fecha y hora: ", t.Format("2006-01-02 15:04:05"))
	fmt.Println("Generando PDF...");
	pdf := gofpdf.New("P", "mm", "A4", "");
	pdf.AddPage();
	pdf.SetFont("Arial", "B", 16);
	pdf.Cell(40, 10, "Factura()");
	err := pdf.OutputFileAndClose("Factura.pdf");
	fmt.Println(err);
}