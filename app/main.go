package main

import (
	"fmt"
)

var option int

func main() {
	for {
		menu()
		switch option {
		case 1:
			// listerSallesDisponibles()
		case 2:
			// creerReservation()
		case 3:
			// annulerReservation()
		case 4:
			// visualiserReservations()
		case 5:
			fmt.Println("Au revoir!")
			return
		}

	}
}

func menu() {
	for {
		fmt.Println("Bienvenue dans le Service de Réservation en Ligne\n-----------------------------------------------------\n")
		fmt.Println("1. Lister les salles disponibles\n2. Créer une réservation\n3. Annuler une réservation\n4. Visualiser les réservations\n5. Quitter\nChoisissez une option :")
		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Erreur de saisie. Veuillez saisir un numéro valide.")
			continue
		}
		if option < 1 || option > 5 {
			fmt.Println("Option invalide. Veuillez choisir une option entre 1 et 5.")
			continue
		}
		break
	}
}
