package main

import (
	"sujet3/car"
	"sujet3/shapes"
)

func main() {
	// Créer une instance de la structure Car
	myCar := car.Car{
		Make:  "Toyota",
		Model: "Corolla",
		Year:  2021,
	}

	// Utiliser la méthode DisplayInfo pour afficher les informations de la voiture
	myCar.DisplayInfo()

	// Créer des instances de Circle et Rectangle
	myCircle := shapes.Circle{Radius: 5}
	myRectangle := shapes.Rectangle{Width: 10, Height: 20}

	// Utiliser PrintArea pour afficher les surfaces
	shapes.PrintArea(myCircle)
	shapes.PrintArea(myRectangle)
}
