package calc

import "errors"

// Add additionne deux nombres et retourne le résultat
func Add(a, b float64) float64 {
	return a + b
}

// Subtract soustrait deux nombres et retourne le résultat
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply multiplie deux nombres et retourne le résultat
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide divise deux nombres et retourne le résultat et une erreur en cas de division par zéro
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division par zéro")
	}
	return a / b, nil
}
