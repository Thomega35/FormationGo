package main

import (
	"fmt"
	"sujet2/mathutils"
)

func main() {
	fmt.Println("Hello, Go!")
	var x, y int = 10, 5

	// Utilisation de la fonction d'addition du package mathutils
	sum := mathutils.Add(x, y)
	fmt.Printf("La somme de %d et %d est %d\n", x, y, sum)

	// Tentative d'utilisation de la fonction soustraction (non exportée)
	// Cela provoquera une erreur si décommenté
	// diff := mathutils.subtract(x, y)
	// fmt.Printf("La différence entre %d et %d est %d\n", x, y, diff)

	// Afficher les 10 premiers nombres et vérifier s'ils sont pairs ou impairs
	for i := 0; i < 10; i++ {
		if mathutils.IsEven(i) {
			fmt.Printf("%d est pair\n", i)
		} else {
			fmt.Printf("%d est impair\n", i)
		}
	}

	// Parcourir un tableau avec range et afficher chaque élément avec son index
	numbers := []int{3, 5, 7, 9, 11}
	for index, value := range numbers {
		fmt.Printf("Index %d: %d\n", index, value)
	}

	// Utilisation de la fonction Divide pour obtenir le quotient et le reste
	quotient, remainder := mathutils.Divide(17, 3)
	fmt.Printf("Le quotient de 17 divisé par 3 est %d et le reste est %d\n", quotient, remainder)

	// Utilisation de la fonction variadique Sum pour calculer la somme de plusieurs entiers
	variadicSum := mathutils.Sum(1, 2, 3, 4, 5)
	fmt.Printf("La somme des nombres 1, 2, 3, 4, 5 est %d\n", variadicSum)

	// Déclaration d'un pointeur et modification de la valeur d'une variable
	var value int = 42
	fmt.Printf("Avant modification : %d\n", value)
	modifyValue(&value)
	fmt.Printf("Après modification : %d\n", value)
}

// Fonction pour modifier la valeur d'une variable en utilisant un pointeur
func modifyValue(val *int) {
	*val = 100
}
