package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Aera() float64
}

type Meserable interface {
	Perameter() float64
}

type GeoMetey interface {
	Shape
	Meserable
}

type Circular struct {
	radius float64
}
type Ractangular struct {
	heigh float64
	width float64
}

func (r Ractangular) Perameter() float64 {
	return 2 * (r.heigh + r.width)
}

func (c Circular) Aera() float64 {
	return math.Pi * c.radius * c.radius
}
func (r Ractangular) Aera() float64 {
	return r.heigh * r.width
}

func main() {
	circular := Circular{radius: 10}
	rectangular := Ractangular{heigh: 10, width: 5}
	fmt.Println(getAera(circular))
	fmt.Println(getAera(rectangular))
	fmt.Println(calculateGeoMetery(rectangular))

}

func getAera(s Shape) float64 {
	return s.Aera()
}

func calculateGeoMetery(g GeoMetey) float64 {
	return g.Perameter()
}
