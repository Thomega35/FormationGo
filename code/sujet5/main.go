package main

import (
	"fmt"
	"os"
)

func main() {
	// Définir le nom du fichier
	fileName := "example.txt"

	// Lire et afficher le contenu du fichier
	readFile(fileName)

	// Ajouter du texte supplémentaire au fichier
	appendToFile(fileName, "Texte supplémentaire ajouté.\n")

	// Relire et afficher le contenu du fichier
	readFile(fileName)

	// Expérimenter avec la création et la suppression de fichiers
	createAndDeleteFile("tempfile.txt")
}

// Fonction pour lire un fichier et afficher son contenu
func readFile(fileName string) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	fmt.Println("Contenu du fichier:", string(content))
}

// Fonction pour ajouter du texte à un fichier
func appendToFile(fileName, text string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(text); err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier:", err)
	}
}

// Fonction pour créer et supprimer un fichier
func createAndDeleteFile(fileName string) {
	// Créer un fichier
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier:", err)
		return
	}
	f.WriteString("Ceci est un fichier temporaire.\n")
	f.Close()

	fmt.Println("Fichier créé:", fileName)

	// Supprimer le fichier
	err = os.Remove(fileName)
	if err != nil {
		fmt.Println("Erreur lors de la suppression du fichier:", err)
		return
	}

	fmt.Println("Fichier supprimé:", fileName)
}
