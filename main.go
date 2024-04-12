package main

import (
	// Import functions
	. "Projet_GO_Reservation/src"
	"fmt"
	"os"
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
			Log("Au revoir!")
			return
		}
		retourMenu()
	}

	// Exemple de comment utiliser la fonction

	/*var bdd Db

	result, err := bdd.SelectDB("RESERVATIONS", []string{"id_reservation"}, nil, true)
	if err != nil {
		// Gérer l'erreur
		fmt.Println("Erreur :", err)
		return
	}

	firstMap := result[0]
	horaire := firstMap["horaire"]
	id_etat := firstMap["id_etat"]
	id_reservation := firstMap["id_reservation"]

	fmt.Println("Horaire:", horaire)
	fmt.Println("ID Etat:", id_etat)
	fmt.Println("ID Réservation:", id_reservation)
	*/

}

func menu() {
	for {
		Println("-----------------------------------------------------\nBienvenue dans le Service de Réservation en Ligne\n-----------------------------------------------------\n")
		Println("1. Lister les salles disponibles\n2. Créer une réservation\n3. Annuler une réservation\n4. Visualiser les réservations\n5. Quitter\nChoisissez une option :")
		_, err := fmt.Scanln(&option)
		if err != nil {
			Println("Erreur de saisie. Veuillez saisir un numéro valide.")
			continue
		}
		if option < 1 || option > 5 {
			Println("Option invalide. Veuillez choisir une option entre 1 et 5.")
			continue
		}
		break
	}
}

func retourMenu() {
	var choix int
	Println("1. Retourner au menu principal\n2. Quitter\nChoisissez une option :")
	fmt.Scanln(&choix)
	switch choix {
	case 1:
		// Rien à faire ici, le programme reviendra automatiquement à la boucle principale
	case 2:
		Println("Au revoir!")
		os.Exit(0)
	default:
		Println("Option invalide, veuillez réessayer.")
		retourMenu()
	}
}
