package main

import (
	"strconv"
	"sujet6/calc"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	// Créer l'application Fyne
	a := app.New()
	w := a.NewWindow("Calculatrice")

	// Créer les champs d'entrée et les boutons
	entry1 := widget.NewEntry()
	entry2 := widget.NewEntry()
	resultLabel := widget.NewLabel("Résultat :")

	addButton := widget.NewButton("Addition", func() {
		calculate(calc.Add, entry1, entry2, resultLabel)
	})
	subButton := widget.NewButton("Soustraction", func() {
		calculate(calc.Subtract, entry1, entry2, resultLabel)
	})
	mulButton := widget.NewButton("Multiplication", func() {
		calculate(calc.Multiply, entry1, entry2, resultLabel)
	})
	divButton := widget.NewButton("Division", func() {
		calculateWithError(calc.Divide, entry1, entry2, resultLabel)
	})

	// Disposer les widgets dans une grille
	grid := container.NewGridWithColumns(2,
		widget.NewLabel("Nombre 1:"), entry1,
		widget.NewLabel("Nombre 2:"), entry2,
		addButton, subButton,
		mulButton, divButton,
		resultLabel,
	)

	// Définir le contenu de la fenêtre
	w.SetContent(grid)

	// Redimensionner la fenêtre et lancer l'application
	w.Resize(fyne.NewSize(400, 200))
	w.ShowAndRun()
}

// Fonction pour effectuer un calcul sans erreur
func calculate(operation func(float64, float64) float64, entry1, entry2 *widget.Entry, resultLabel *widget.Label) {
	num1, err1 := strconv.ParseFloat(entry1.Text, 64)
	num2, err2 := strconv.ParseFloat(entry2.Text, 64)
	if err1 != nil || err2 != nil {
		resultLabel.SetText("Erreur : Entrée non valide")
		return
	}
	result := operation(num1, num2)
	resultLabel.SetText("Résultat : " + strconv.FormatFloat(result, 'f', -1, 64))
}

// Fonction pour effectuer un calcul avec gestion d'erreur
func calculateWithError(operation func(float64, float64) (float64, error), entry1, entry2 *widget.Entry, resultLabel *widget.Label) {
	num1, err1 := strconv.ParseFloat(entry1.Text, 64)
	num2, err2 := strconv.ParseFloat(entry2.Text, 64)
	if err1 != nil || err2 != nil {
		resultLabel.SetText("Erreur : Entrée non valide")
		return
	}
	result, err := operation(num1, num2)
	if err != nil {
		resultLabel.SetText("Erreur : " + err.Error())
		return
	}
	resultLabel.SetText("Résultat : " + strconv.FormatFloat(result, 'f', -1, 64))
}
