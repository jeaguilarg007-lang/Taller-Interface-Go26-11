package main

import (
	"fmt"
)

type InformacionProducto interface {
	MostrarInfo() string
}

type ProductoBase struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

func (p ProductoBase) MostrarInfo() string {
	return fmt.Sprintf(
		"ID: %d | Nombre: %s | Precio: $%.2f | Cantidad: %d",
		p.ID, p.Nombre, p.Precio, p.Cantidad,
	)
}

type Electronico struct {
	ProductoBase
	GarantiaMeses int
}

func NuevoProductoElectronico(
	id int,
	nombre string,
	precio float64,
	cantidad int,
	garantia int,
) Electronico {
	return Electronico{
		ProductoBase: ProductoBase{
			ID:       id,
			Nombre:   nombre,
			Precio:   precio,
			Cantidad: cantidad,
		},
		GarantiaMeses: garantia,
	}
}

func (e Electronico) MostrarInfo() string {
	return fmt.Sprintf(
		"%s | Garantía: %d meses",
		e.ProductoBase.MostrarInfo(),
		e.GarantiaMeses,
	)
}

type Alimento struct {
	ProductoBase
	FechaExpiracion string
}

func NuevoProductoAlimento(
	id int,
	nombre string,
	precio float64,
	cantidad int,
	fechaExp string,
) Alimento {
	return Alimento{
		ProductoBase: ProductoBase{
			ID:       id,
			Nombre:   nombre,
			Precio:   precio,
			Cantidad: cantidad,
		},
		FechaExpiracion: fechaExp,
	}
}

func (a Alimento) MostrarInfo() string {
	return fmt.Sprintf(
		"%s | Fecha de expiración: %s",
		a.ProductoBase.MostrarInfo(),
		a.FechaExpiracion,
	)
}

type Inventario struct {
	Productos []InformacionProducto
}

func NuevoInventario() Inventario {
	return Inventario{
		Productos: []InformacionProducto{},
	}
}

func (i *Inventario) AgregarProducto(p InformacionProducto) {
	i.Productos = append(i.Productos, p)
}

func (i Inventario) ListarProductos() {
	fmt.Println("\n===== INVENTARIO DE PRODUCTOS =====")

	if len(i.Productos) == 0 {
		fmt.Println("El inventario está vacío.")
		return
	}

	for _, producto := range i.Productos {
		fmt.Println(producto.MostrarInfo())
	}
}

func main() {
	inventario := NuevoInventario()

	televisor := NuevoProductoElectronico(
		1, "Televisor Samsung", 450.50, 10, 24,
	)

	computadora := NuevoProductoElectronico(
		2, "Laptop Lenovo", 850.00, 5, 12,
	)

	leche := NuevoProductoAlimento(
		3, "Leche", 1.25, 30, "2026-12-15",
	)

	arroz := NuevoProductoAlimento(
		4, "Arroz", 2.50, 50, "2027-05-20",
	)

	inventario.AgregarProducto(televisor)
	inventario.AgregarProducto(computadora)
	inventario.AgregarProducto(leche)
	inventario.AgregarProducto(arroz)

	inventario.ListarProductos()
}
