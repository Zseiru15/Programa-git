package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"time"
)

// Función para capturar el nombre del cliente
// Esta función solicita al usuario que ingrese el nombre del cliente y lo retorna.
func datos_cliente() string {
	var cliente string
	fmt.Println("Ingrese el nombre del cliente: ") // Muestra un mensaje solicitando el nombre del cliente.
	fmt.Scan(&cliente)                            // Captura el nombre ingresado por el usuario.
	return cliente                                // Retorna el nombre del cliente.
}

// Función que captura productos, precios y calcula el valor total de la compra
// Esta función permite al usuario ingresar varios productos, precios y cantidades. Calcula el total de la compra.
func compras() ([]string, []float64, []float64, float64) {
	var opcion string
	var productos []string        // Lista de nombres de productos.
	var precios []float64         // Lista de precios de productos.
	var cantidades []float64      // Lista de cantidades de productos.
	var valor_total float64       // Valor total de la compra.

	for {
		var producto string
		var precio, cantidad float64

		// Captura el nombre del producto
		fmt.Println("Ingrese el nombre del producto: ")
		fmt.Scan(&producto)
		productos = append(productos, producto) // Agrega el producto a la lista de productos.

		// Captura el precio del producto
		fmt.Println("Ingrese el precio del producto: ")
		fmt.Scan(&precio)
		precios = append(precios, precio) // Agrega el precio del producto a la lista de precios.

		// Captura la cantidad de productos
		fmt.Println("Ingrese la cantidad del producto: ")
		fmt.Scan(&cantidad)
		cantidades = append(cantidades, cantidad) // Agrega la cantidad del producto a la lista de cantidades.

		// Calcula el valor total por producto (precio * cantidad)
		valor := precio * cantidad
		valor_total += valor // Suma el valor calculado al valor total.

		// Pregunta si se desea agregar otro producto
		fmt.Println("Desea agregar otro producto? (s/n): ")
		fmt.Scan(&opcion)

		// Si no se desea agregar más productos, se rompe el ciclo.
		if opcion != "s" {
			break
		}
	}
	// Retorna los productos, precios, cantidades y el valor total de la compra.
	return productos, precios, cantidades, valor_total
}

// Función para generar la factura: llama a datos_cliente() y compras()
// Esta función combina el nombre del cliente y los datos de la compra.
func Factura() (string, []string, []float64, []float64, float64) {
	cliente := datos_cliente()                        // Obtiene el nombre del cliente.
	productos, precios, cantidades, valor_total := compras() // Obtiene los productos, precios y cantidades.
	return cliente, productos, precios, cantidades, valor_total // Retorna todos los datos obtenidos.
}

// Función para generar el PDF con tablas y estructura más detallada
// Esta función genera el archivo PDF que representa la factura con los productos y el total.
func generarPDF(cliente string, productos []string, precios, cantidades []float64, valor_total float64) {
	pdf := gofpdf.New("P", "mm", "A4", "") // Inicializa un nuevo documento PDF en tamaño A4.
	pdf.AddPage()                          // Agrega una nueva página al PDF.

	// Título
	pdf.SetFont("Arial", "B", 16)        // Establece la fuente Arial en negrita, tamaño 16.
	pdf.Cell(190, 10, "Factura")         // Escribe el título "Factura" en la parte superior del PDF.
	pdf.Ln(15)                           // Inserta un salto de línea.

	// Información de la factura
	pdf.SetFont("Arial", "", 12)         // Cambia la fuente a Arial normal, tamaño 12.
	pdf.Cell(100, 10, fmt.Sprintf("Cliente: %s", cliente)) // Escribe el nombre del cliente.
	pdf.Ln(5)                            // Inserta un salto de línea.
	pdf.Cell(100, 10, fmt.Sprintf("Fecha: %s", time.Now().Format("02/01/2006 - 15:04:05"))) // Escribe la fecha actual.
	pdf.Ln(15)                           // Inserta otro salto de línea.

	// Cabecera de la tabla
	pdf.SetFont("Arial", "B", 12)        // Cambia la fuente a Arial negrita, tamaño 12.
	// Añade las columnas de la tabla: Ref, Descripción, Cantidad, Precio, Importe.
	pdf.CellFormat(10, 7, "Ref", "1", 0, "C", false, 0, "") 
	pdf.CellFormat(80, 7, "Descripcion", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 7, "Cantidad", "1", 0, "C", false, 0, "")
	pdf.CellFormat(30, 7, "Precio", "1", 0, "C", false, 0, "")
	pdf.CellFormat(30, 7, "Importe", "1", 0, "C", false, 0, "")
	pdf.Ln(-1)                          // Salto de línea para iniciar el contenido de la tabla.

	// Datos de la tabla
	pdf.SetFont("Arial", "", 12)         // Cambia la fuente a Arial normal, tamaño 12.
	for i, producto := range productos { // Itera sobre los productos para llenar la tabla.
		// Cada producto ocupa una fila en la tabla.
		pdf.CellFormat(10, 7, fmt.Sprintf("%d", i+1), "1", 0, "C", false, 0, "") // Número de referencia.
		pdf.CellFormat(80, 7, producto, "1", 0, "", false, 0, "") // Nombre del producto.
		pdf.CellFormat(20, 7, fmt.Sprintf("%.2f", cantidades[i]), "1", 0, "C", false, 0, "") // Cantidad del producto.
		pdf.CellFormat(30, 7, fmt.Sprintf("%.2f", precios[i]), "1", 0, "C", false, 0, "") // Precio del producto.
		pdf.CellFormat(30, 7, fmt.Sprintf("%.2f", precios[i]*cantidades[i]), "1", 0, "C", false, 0, "") // Importe total.
		pdf.Ln(-1)                      // Salto de línea después de cada producto.
	}

	// Resumen final de la factura
	pdf.Ln(10)                          // Espacio antes del resumen.
	pdf.SetFont("Arial", "B", 12)        // Cambia la fuente a Arial negrita, tamaño 12.
	pdf.Cell(110, 7, "")                // Espacio vacío para alinear el texto a la derecha.
	pdf.CellFormat(40, 7, "BASE", "0", 0, "", false, 0, "") // Texto "BASE".
	pdf.CellFormat(30, 7, fmt.Sprintf("%.2f", valor_total), "1", 0, "C", false, 0, "") // Muestra el valor total (subtotal).
	pdf.Ln(10)                          // Salto de línea.
	pdf.Cell(110, 7, "")                // Espacio vacío.
	pdf.CellFormat(40, 7, "IVA", "0", 0, "", false, 0, "") // Texto "IVA".
	iva := valor_total * 0.16           // Calcula el IVA (16% del total).
	pdf.CellFormat(30, 7, fmt.Sprintf("%.2f", iva), "1", 0, "C", false, 0, "") // Muestra el valor del IVA.
	pdf.Ln(10)                          // Salto de línea.
	pdf.Cell(110, 7, "")                // Espacio vacío.
	pdf.CellFormat(40, 7, "TOTAL", "0", 0, "", false, 0, "") // Texto "TOTAL".
	pdf.CellFormat(30, 7, fmt.Sprintf("%.2f", valor_total+iva), "1", 0, "C", false, 0, "") // Muestra el total final (subtotal + IVA).

	// Guardar el archivo PDF
	err := pdf.OutputFileAndClose("PruebaGO.pdf") // Guarda el archivo PDF en el sistema con el nombre "Prueba.pdf".
	if err != nil {
		fmt.Println("Error al generar el PDF:", err) // Muestra un mensaje de error si la generación falla.
	}
}

// Función principal
// Esta función coordina la generación de la factura y la creación del PDF.
func main() {
	// Llama a la función Factura() para capturar los datos y generar la factura.
	cliente, productos, precios, cantidades, valor_total := Factura()

	// Genera el PDF con los datos obtenidos.
	fmt.Println("Generando PDF...")
	generarPDF(cliente, productos, precios, cantidades, valor_total) // Llama a la función para generar el PDF.
	fmt.Println("PDF generado con éxito: Factura.pdf") // Mensaje de confirmación de éxito.
}
