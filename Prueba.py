from fpdf import FPDF
import time

# Función para capturar el nombre del cliente
def datos_cliente():
    cliente = input("Ingrese el nombre del cliente: ")
    return cliente

# Función que captura productos, precios y calcula el valor total de la compra
def compras():
    productos = []     # Lista de productos
    precios = []       # Lista de precios
    cantidades = []    # Lista de cantidades
    valor_total = 0.0  # Total de la compra

    while True:
        # Captura del producto
        producto = input("Ingrese el nombre del producto: ")
        productos.append(producto)

        # Captura del precio
        precio = float(input("Ingrese el precio del producto: "))
        precios.append(precio)

        # Captura de la cantidad
        cantidad = float(input("Ingrese la cantidad del producto: "))
        cantidades.append(cantidad)

        # Cálculo del importe
        valor_total += precio * cantidad

        # Pregunta si desea agregar otro producto
        opcion = input("Desea agregar otro producto? (s/n): ")
        if opcion.lower() != 's':
            break

    return productos, precios, cantidades, valor_total

# Función para generar la factura: captura los datos del cliente y de la compra
def factura():
    cliente = datos_cliente()
    productos, precios, cantidades, valor_total = compras()
    return cliente, productos, precios, cantidades, valor_total

# Función para generar el PDF con tablas y una estructura detallada
def generar_pdf(cliente, productos, precios, cantidades, valor_total):
    pdf = FPDF(orientation='P', unit='mm', format='A4')
    pdf.add_page()

    # Título
    pdf.set_font('Arial', 'B', 16)
    pdf.cell(190, 10, 'Factura', ln=True, align='C')

    # Información del cliente y la fecha
    pdf.set_font('Arial', '', 12)
    pdf.cell(100, 10, f'Cliente: {cliente}', ln=True)
    pdf.cell(100, 10, f'Fecha: {time.strftime("%d/%m/%Y - %h:%m:%s")}', ln=True)
    pdf.ln(10)

    # Cabecera de la tabla
    pdf.set_font('Arial', 'B', 12)
    pdf.cell(10, 7, 'Ref', 1, align='C')
    pdf.cell(80, 7, 'Descripción', 1, align='C')
    pdf.cell(20, 7, 'Cantidad', 1, align='C')
    pdf.cell(30, 7, 'Precio', 1, align='C')
    pdf.cell(30, 7, 'Importe', 1, align='C')
    pdf.ln()

    # Datos de la tabla
    pdf.set_font('Arial', '', 12)
    for i, producto in enumerate(productos):
        pdf.cell(10, 7, str(i + 1), 1, align='C')
        pdf.cell(80, 7, producto, 1)
        pdf.cell(20, 7, f'{cantidades[i]:.2f}', 1, align='C')
        pdf.cell(30, 7, f'{precios[i]:.2f}', 1, align='C')
        pdf.cell(30, 7, f'{precios[i] * cantidades[i]:.2f}', 1, align='C')
        pdf.ln()

    # Resumen de la factura
    pdf.ln(10)
    pdf.cell(110, 7, '')
    pdf.cell(40, 7, 'BASE', 0)
    pdf.cell(30, 7, f'{valor_total:.2f}', 1, align='C')
    pdf.ln(10)

    iva = valor_total * 0.16  # Asumiendo un IVA del 16%
    pdf.cell(110, 7, '')
    pdf.cell(40, 7, 'IVA', 0)
    pdf.cell(30, 7, f'{iva:.2f}', 1, align='C')
    pdf.ln(10)

    pdf.cell(110, 7, '')
    pdf.cell(40, 7, 'TOTAL', 0)
    pdf.cell(30, 7, f'{valor_total + iva:.2f}', 1, align='C')

    # Guardar el archivo PDF
    pdf.output('PruebaPY.pdf')
    print("PDF generado con éxito: PruebaPY.pdf")

# Función principal que genera la factura y el PDF
def main():
    cliente, productos, precios, cantidades, valor_total = factura()  # Captura los datos
    generar_pdf(cliente, productos, precios, cantidades, valor_total)  # Genera el PDF

# Ejecutar la función principal
if __name__ == "__main__":
    main()
