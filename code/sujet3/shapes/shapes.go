package shapes

import (
	"fmt"
	"math"
)

// Shape représente une interface avec une méthode Area
type Shape interface {
	Area() float64
}

// Circle représente un cercle avec un rayon
type Circle struct {
	Radius float64
}

// Rectangle représente un rectangle avec une largeur et une hauteur
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calcule et retourne la surface du cercle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area calcule et retourne la surface du rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func PrintArea(s Shape) {
	fmt.Printf("La surface est : %f\n", s.Area())
}
