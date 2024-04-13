package main

import (
	// Import functions
	. "Projet_GO_Reservation/functions"
	. "Projet_GO_Reservation/src"
	"fmt"
	"os"
)

var option int

func main() {

	var bdd Db

	for {
		menu()
		switch option {
		case 1:

			bdd.InsertDB(ETAT, []string{"id_etat", "nom_etat"}, []string{"5", "Se nourrir 2fois"}, nil, true)
			// listerSallesDisponibles()
		case 2:
			
			var tmpa = "id_reservation = 1"
			result, err := bdd.SelectDB(RESERVATIONS, []string{"id_reservation", "horaire", "id_etat"}, &tmpa, true)
			if err != nil {
				Log.Error("Impossible de sélectionner dans la BDD : ", err)
				return
			}

			if result == nil {
				Log.Error("Impossible de sélectionner les données")
				break
			}

			firstMap := result[0]
			horaire := firstMap["horaire"]
			id_etat := firstMap["id_etat"]
			id_reservation := firstMap["id_reservation"]

			fmt.Println("Horaire:", horaire)
			fmt.Println("ID Etat:", id_etat)
			fmt.Println("ID Réservation:", id_reservation)

			// creerReservation()
		case 3:
			var tmp = "id_etat = 4"
			bdd.UpdateDB(ETAT, []string{"nom_etat"}, []string{"Coucou"}, &tmp, true)
			// annulerReservation()
		case 4:
			var tmp = "id_etat = 5"
			bdd.DeleteDB(ETAT, &tmp, true)
			// visualiserReservations()
		case 5:

			ILog("Au revoir!")
			return
		}
		retourMenu()
	}

	// Exemple de comment utiliser la fonction

	/*
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
