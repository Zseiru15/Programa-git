package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	fmt.Println("Texto en la funcion main")
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	err := pdf.OutputFileAndClose("hello.pdf")
	fmt.Println(err)
}