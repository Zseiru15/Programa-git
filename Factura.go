package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"time"
)

func datos_cliente() string {
	var cliente string
	fmt.Println("Ingrese el nombre del cliente: ")
	fmt.Scan(&cliente)
	return cliente
}

func compras() ([]string, []float64, float64) {
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
	return productos, precios, valor_total
}

func Factura() (string, []string, []float64, float64) {
	cliente := datos_cliente()
	productos, precios, valor_total := compras()
	return cliente, productos, precios, valor_total
}

func generarPDF(cliente string, productos []string, precios []float64, valor_total float64) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Factura")
	pdf.Ln(10) // Salto de línea

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Cliente: %s", cliente))
	pdf.Ln(10)

	pdf.Cell(40, 10, "Productos:")
	pdf.Ln(10)

	for i, producto := range productos {
		pdf.Cell(0, 10, fmt.Sprintf("%s - $%.2f", producto, precios[i]))
		pdf.Ln(10)
	}

	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Valor total de la compra: $%.2f", valor_total))

	err := pdf.OutputFileAndClose("FacturaGO.pdf")
	if err != nil {
		fmt.Println("Error al generar el PDF:", err)
	}
}

func main() {
	t := time.Now()
	cliente, productos, precios, valor_total := Factura()
	fmt.Println("Fecha y hora: ", t.Format("2006-01-02 15:04:05"))
	fmt.Println("Generando PDF...")
	generarPDF(cliente, productos, precios, valor_total)
	fmt.Println("PDF generado con éxito: Prueba.pdf")
}
