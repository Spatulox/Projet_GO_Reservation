package lechauve

import (
	. "Projet_GO_Reservation/src"
	"fmt"
)

var option int

func ReservationsMenu() {
	for {
		menu()
		switch option {
		case 1:

			listReservations()
		case 2:

			createReservation()
		case 3:

			cancelReservation()
		case 5:

			Println("Retour menu principal")
			return
		}
		if retourMenu() == 2 {
			return
		}
	}
}

//
// ------------------------------------------------------------------------------------------------ //
//

func listReservations() []map[string]interface{} {

	var bdd Db

	result, err := bdd.SelectDB(RESERVATIONS, []string{"*"}, nil, true)

	if err != nil {
		Log.Error("Erreur lors de la lecture de la Base de donnée", err)
	}

	for _, sResult := range result {
		horaire := sResult["horaire"]
		idEtat := sResult["id_etat"]
		idReservation := sResult["id_reservation"]

		Println("------------------------------")
		fmt.Println("ID Réservation:", idReservation)
		fmt.Println("Horaire:", horaire)
		fmt.Println("ID Etat:", idEtat)
	}

	return result

}

//
// ------------------------------------------------------------------------------------------------ //
//

func createReservation() {

}

//
// ------------------------------------------------------------------------------------------------ //
//

func cancelReservation() {

}

//
// ------------------------------------------------------------------------------------------------ //
//

func menu() {
	for {
		Println("-----------------------------------------------------\nMenu Réservation\n-----------------------------------------------------\n")
		Println("1. Lister les reservations\n2. Créer une réservation\n3. Modifier une réservation\n4. Annuler une réservation\n5. Quitter\nChoisissez une option :")
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

func retourMenu() int64 {
	var choix int
	Println("1. Retourner au menu reservation\n2. Quitter\nChoisissez une option :")
	fmt.Scanln(&choix)
	switch choix {
	case 1:
		// Rien à faire ici, le programme reviendra automatiquement à la boucle principale
	case 2:
		Println("Retour au menu principal!")
		return 2
	default:
		Println("Option invalide, veuillez réessayer.")
		retourMenu()
	}

	return 1
}
