//  wikipedia example
// Interface nada mais é que um conjunto de métodos.

package main

import (
	"fmt"
	"math"
)

type geometria interface {
	area() float64
	perim() float64
}

type quadrado struct {
	largura, altura float64
}

type círculo struct {
	raio float64
}

func (q quadrado) area() float64 {
	return q.largura * q.altura
}

func (q quadrado) perim() float64 {
	return 2*q.largura + 2*q.altura
}

func (c círculo) area() float64 {
	return math.Pi * c.raio * c.raio
}

func (c círculo) perim() float64 {
	return 2 * math.Pi * c.raio
}

func medir(g geometria) {
	fmt.Println(g)         // Printa os valores da forma geometrica recebida ex: quadrado = altura e largura
	fmt.Println(g.area())  // Printa a area da forma recebida
	fmt.Println(g.perim()) // Printa o perimetro da forma recebida
}
func main() {
	q := quadrado{largura: 3, altura: 4}
	c := círculo{raio: 5}
	medir(q)
	medir(c)
}
