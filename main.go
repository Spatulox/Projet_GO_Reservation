package main

import (
	. "Projet_GO_Reservation/functions"
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

			ReservationsMenu()
		case 2:

			MenuSalle()
		case 3:

			// annulerReservation()
		case 4:

			// visualiserReservations()
		case 5:

			Println("Au revoir!")
			return
		}
		//retourMenu()
	}

	// Exemple de comment utiliser la fonction

	/*
	 */

}

func menu() {
	for {
		Println("\n-----------------------------------------------------\nBienvenue dans le Service de Réservation en Ligne\n-----------------------------------------------------")
		Println("1. Menu pour les réservations\n2. Menu pour les Salles\n3. Créer une réservation\n4. Visualiser les réservations\n5. Quitter\nChoisissez une option :")
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
	Println("\n----------Retour-----------")
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

func example() {

	var bdd Db

	// Insert ----------------------------------------------------------
	bdd.InsertDB(ETAT, []string{"id_etat", "nom_etat"}, []string{"5", "Se nourrir 2fois"}, true)

	// Select ----------------------------------------------------------
	var tmpa = "id_reservation = 1"
	result, err := bdd.SelectDB(RESERVATIONS, []string{"id_reservation", "horaire", "id_etat"}, &tmpa, true)
	if err != nil {
		Log.Error("Impossible de sélectionner dans la BDD : ", err)
		return
	}

	if result == nil {
		Log.Error("Impossible de sélectionner les données")
		return
	}

	firstMap := result[0]
	horaire := firstMap["horaire"]
	idEtat := firstMap["id_etat"]
	idReservation := firstMap["id_reservation"]

	fmt.Println("Horaire:", horaire)
	fmt.Println("ID Etat:", idEtat)
	fmt.Println("ID Réservation:", idReservation)

	// Update ----------------------------------------------------------
	var tmp = "id_etat = 4"
	bdd.UpdateDB(ETAT, []string{"nom_etat"}, []string{"Coucou"}, &tmp, true)

	// Delete ----------------------------------------------------------
	var tmpb = "id_etat = 5"
	bdd.DeleteDB(ETAT, &tmpb, true)
}
