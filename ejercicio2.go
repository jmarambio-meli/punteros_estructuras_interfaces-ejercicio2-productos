package main

import "fmt"

const (
	small  = "SMALL"
	medium = "MEDIUM"
	big    = "BIG"
)

type Producto interface {
	Precio() float64
}

//-------

type Small struct {
	costo float64
}

func (s Small) Precio() float64 {
	return s.costo
}

//--------

type Medium struct {
	costo float64
}

func (m Medium) Precio() float64 {
	return m.costo * 1.3
}

//--------

type Big struct {
	costo float64
}

func (b Big) Precio() float64 {
	return b.costo*1.6 + 2500
}

func details(p Producto) {
	fmt.Println(p.Precio())
}

func Factory(productoTipo string, costo float64) Producto {
	switch productoTipo {
	case small:
		return Small{costo}
	case medium:
		return Medium{costo}
	case big:
		return Big{costo}
	}
	return nil
}

func main() {
	producto_pequeno := Factory(small, 1000)
	details(producto_pequeno)

	producto_mediano := Factory(medium, 1000)
	details(producto_mediano)

	producto_grande := Factory(big, 1000)
	details(producto_grande)
}
