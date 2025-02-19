package cap13

import (
	"fmt"
)

type Polygon interface {
	area() float64
}

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (c Circle) area() float64 {
	return 2 * 3.14 * c.radius //Math.Pi
}

func infoPolygon(p Polygon) float64 {
	return p.area()
}

func ex05() {

	square := Square{5}
	circle := Circle{2}

	fmt.Println("A área do quadrado é:", infoPolygon(square))
	fmt.Println("A área do círculo é:", infoPolygon(circle))

}
