package mathutils

// Add additionne deux entiers et retourne le résultat.
func Add(a int, b int) int {
	return a + b
}

// subtract soustrait deux entiers et retourne le résultat.
// Notez que cette fonction est en minuscule, donc elle n'est pas exportée.
func subtract(a int, b int) int {
	return a - b
}

// IsEven vérifie si un nombre est pair.
func IsEven(n int) bool {
	return n%2 == 0
}

// Divide retourne le quotient et le reste de la division de a par b.
func Divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// Sum calcule la somme de plusieurs entiers.
func Sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
