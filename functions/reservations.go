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
		case 4:

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

		tmp := fmt.Sprintf("id_etat=%v", idEtat)
		etatResult, err := bdd.SelectDB(ETAT, []string{"nom_etat"}, &tmp)

		// Print
		Println("------------------------------")
		fmt.Println("ID Réservation:", idReservation)
		fmt.Println("Horaire:", horaire)

		if err != nil {
			Log.Error("Impossible de récupérer l'état de la réservation")
			fmt.Println("ID Etat:", idEtat)
		} else {
			fmt.Println("Etat : ", etatResult[0]["nom_etat"])
		}

	}
	Println("------------------------------")

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
		Println("-----------------------------------------------------\nMenu Réservation\n-----------------------------------------------------")
		Println("1. Lister les reservations\n2. Créer une réservation\n3. Annuler une réservation\n4. Quitter\nChoisissez une option :")
		_, err := fmt.Scanln(&option)
		if err != nil {
			Println("Erreur de saisie. Veuillez saisir un numéro valide.")
			continue
		}
		if option < 1 || option > 4 {
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
		Println("Retour au menu principal!\n\n")
		return 2
	default:
		Println("Option invalide, veuillez réessayer.")
		retourMenu()
	}

	return 1
}
