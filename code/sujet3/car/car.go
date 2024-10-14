package car

import "fmt"

// Car représente une voiture avec sa marque, son modèle et son année
type Car struct {
	Make  string
	Model string
	Year  int
}

// DisplayInfo affiche les informations de la voiture
func (c Car) DisplayInfo() {
	fmt.Printf("Voiture: %s %s, Année: %d\n", c.Make, c.Model, c.Year)
}
