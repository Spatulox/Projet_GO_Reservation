package lechauve

import (
	. "Projet_GO_Reservation/src"
	"fmt"
	"strconv"
	"time"
)

var optionReserv int

func ReservationsMenu() {
	for {
		menuReserv()
		switch optionReserv {
		case 1:

			listReservations(nil)
		case 2:

			createReservation()
		case 3:

			cancelReservation()
		case 4:

			updateReservation(nil, nil)
		case 5:

			Println("Retour menu principal")
			return
		}
		if retourMenuReserv() == 2 {
			return
		}
	}
}

//
// ------------------------------------------------------------------------------------------------ //
//

func listReservations(condition *string) []map[string]interface{} {

	var bdd Db
	result, err := bdd.SelectDB(RESERVATIONS, []string{"*"}, condition)

	if err != nil || result == nil {
		Log.Error("Erreur lors de la lecture de la Base de donnée", err)
		return nil
	}
	Println("------------------------------")
	Println("-------- RESERVATIONS --------")
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
	var bdd Db
	// Select all the room

	result, err := bdd.SelectDB(SALLES, []string{"*"}, nil)

	if err != nil || result == nil {
		Log.Error("Impossible de lister les salles :/")
		return
	}

	idMin := result[0]["id_salle"].(int64)
	idMax := result[len(result)-1]["id_salle"].(int64)

	var salle int64
	for {
		fmt.Printf("Veuillez saisir une salle entre %d et %d : ", idMin, idMax)

		fmt.Scanln(&salle)
		if idMin > salle || salle > idMax {
			continue
		} else {
			fmt.Printf("Vouc avez choisit la salle %d\n", salle)
			break
		}
	}

	var departureDate time.Time
	var departureTime time.Time
	var err1, err2 error

	for {
		// Get the date
		var departureDateStr string
		fmt.Print("Entrez la date de départ (format yyyy-mm-dd): ")
		fmt.Scanln(&departureDateStr)

		departureDate, err1 = time.Parse("2006-01-02", departureDateStr)
		if err1 != nil {
			fmt.Println("Erreur de saisie de la date :", err1)
			continue
		}

		// Date du jour
		today := time.Now().Format("2006-01-02")
		todayDate, err2 := time.Parse("2006-01-02", today)
		if err2 != nil {
			fmt.Println("Erreur lors de la récupération de la date du jour :", err2)
			continue
		}

		// Comparer les dates
		if departureDate.Before(todayDate) || departureDate.Equal(todayDate) {
			Println("La date de départ doit être supérieure à la date du jour.")
			continue
		}

		break
	}

	for {
		// Get the hour
		var departureTimeStr string
		fmt.Print("Entrez l'heure de départ (format 15:04): ")
		fmt.Scanln(&departureTimeStr)

		departureTime, err2 = time.Parse("15:04", departureTimeStr)
		if err2 != nil {
			fmt.Println("Erreur de saisie de l'heure :", err2)
			continue
		}
		break
	}

	departureDateTime := departureDate.Format("2006-01-02") + " " + departureTime.Format("15:04:00")

	fmt.Println("Date et heure de départ : ", departureDateTime)
	Println("Toutes les vérifications ont été effectuée, ajout d'une nouvelle réservation")

	// Selectionne dans la BDD pour savoir si y'a quelque chose enrgistré a cette date/heure et dans la salle
	var tmp = "horaire = '" + departureDateTime + "'"
	result, err = bdd.SelectDB(RESERVATIONS, []string{"id_reservation"}, &tmp)

	if err != nil {
		Log.Error("Impossible de vérifier si il existe déjà une reservation a cette date")
		return
	}

	// If y'a déjà une réservation ce jour
	if len(result) > 0 {
		tmp = fmt.Sprintf("id_reservation = %d", result[0]["id_reservation"].(int64))

		result, err = bdd.SelectDB(SALLES, []string{"id_salle"}, &tmp)

		// Il y'a déjà une reservation ce jour et dans cette salle
		if err != nil || result != nil {
			fmt.Printf("Il existe déjà une reservation a cette date %s et dans cette salle %d\n", departureDateTime, salle)
			Println("------------------------------")
			return
		}
	}

	// Insertion des données
	bdd.InsertDB(RESERVATIONS, []string{"horaire", "id_etat"}, []string{departureDateTime, "4"})

	// Select the line with the MAX(id)
	tmp = "id_reservation = (SELECT MAX(id_reservation) FROM " + RESERVATIONS + ")"
	result = listReservations(&tmp)

	if result == nil {
		Log.Error("Impossible de sélectionner la dernière réservation rentrée")
		return
	}

	horaire := fmt.Sprintf("%d", result[0]["id_reservation"].(int64))
	bdd.InsertDB(RESERVER, []string{"id_salle", "id_reservation"}, []string{fmt.Sprintf("%d", salle), horaire})
}

//
// ------------------------------------------------------------------------------------------------ //
//

func cancelReservation(choix ...int) {
	reservation := listReservations(nil)

	var option int
	var maxIdReservation, minIdReservation int64
	minIdReservation = reservation[0]["id_reservation"].(int64)
	maxIdReservation = reservation[len(reservation)-1]["id_reservation"].(int64)

	if choix != nil && len(choix) > 0 {
		option = choix[0]
	} else {
		for {
			Println("Quelle réservation voulez-vous annuler ?\n(-1 pour revenir au menu)\nChoix:")

			_, err := fmt.Scanln(&option)

			if err != nil {
				Println("Erreur de saisie. Veuillez saisir un numéro valide.")
				continue
			}
			if option == -1 {
				return
			}
			if option < 1 || int64(option) > maxIdReservation {
				fmt.Printf("Option invalide. Veuillez choisir une option entre %d et %d\n", minIdReservation, maxIdReservation)
				continue
			}

			f := false
			for _, m := range reservation {
				if (m["id_reservation"]) == int64(option) {
					f = true
					break
				}
			}
			if f == false {
				Println("Cette réservation n'existe pas\n")
				continue
			}
			break
		}
	}

	// Delete from DATABASE
	var bdd Db

	tmp := fmt.Sprintf("id_reservation=%v", option)
	bdd.DeleteDB(RESERVER, &tmp)
	bdd.DeleteDB(RESERVATIONS, &tmp)

}

//
// ------------------------------------------------------------------------------------------------ //
//

func updateReservation(state *int, idReservation *int) {
	var bdd Db

	result, err := bdd.SelectDB(ETAT, []string{"*"}, nil)

	if err != nil || result == nil {
		Log.Error("Impossible de récupérer les états possible dans la Base de donnée")
		return
	}

	var newState int64

	// Ask the user for the state
	if state == nil {
		Println("--------------------------------------------------")
		Println("Choisisser un nouveau etats pour votre reservation")
		Println("--------------------------------------------------")
		for _, m := range result {
			fmt.Println(m["id_etat"], m["nom_etat"])
		}
		Println("--------------------------------------------------")

		idMin := result[0]["id_etat"].(int64)
		idMax := result[len(result)-1]["id_etat"].(int64)

		for {
			fmt.Printf("Vous devez choisir un etat entre %d, et %d : ", idMin, idMax)
			fmt.Scanln(&newState)
			if newState < idMin || newState > idMax {
				continue
			}
			exist := false

			// check if the state exist in the DB
			for _, m := range result {
				if m["id_etat"].(int64) == newState {
					exist = true
					break
				}
			}
			if exist == true {
				break
			}
		}
	} else {

		//newState, err = strconv.ParseInt(*state, 10, 64)
		newState = int64(*state)
		if err != nil {
			Log.Error("Erreur de conversionde l'état de string vers int64 :", err)
			return
		}
	}

	// Ask the user for the id_reservation
	var idReserv int
	if idReservation == nil {
		result = listReservations(nil)

		var maxIdReservation, minIdReservation int64
		minIdReservation = result[0]["id_reservation"].(int64)
		maxIdReservation = result[len(result)-1]["id_reservation"].(int64)

		for {
			fmt.Printf("Pour quelle réservation voulez-vous changer l'état ? : ")

			_, err := fmt.Scanln(&idReserv)

			if err != nil {
				Println("Erreur de saisie. Veuillez saisir un numéro valide.")
				continue
			}
			if idReserv == -1 {
				return
			}
			if idReserv < 1 || int64(idReserv) > maxIdReservation {
				fmt.Printf("Option invalide. Veuillez choisir une option entre %d et %d\n", minIdReservation, maxIdReservation)
				continue
			}

			f := false
			for _, m := range result {
				if (m["id_reservation"]) == int64(idReserv) {
					f = true
					break
				}
			}
			if f == false {
				Println("Cette réservation n'existe pas\n")
				continue
			}
			break

		}
	} else {
		idReserv = *idReservation
	}

	tmp := fmt.Sprintf("id_reservation = %d", idReserv)
	stateStr := strconv.FormatInt(newState, 10)
	bdd.UpdateDB(RESERVATIONS, []string{"id_etat"}, []string{stateStr}, &tmp)

	fmt.Printf("Etat changé pour %s pour la réservation %d\n\n", stateStr, idReserv)
	return
}

//
// ------------------------------------------------------------------------------------------------ //
//

func menuReserv() {
	for {
		Println("-----------------------------------------------------\nMenu Réservation\n-----------------------------------------------------")
		Println("1. Lister les reservations\n2. Créer une réservation\n3. Annuler une réservation\n4. Mettre à jour une reservation\n5. Menu Principal\nChoisissez une option :")
		_, err := fmt.Scanln(&optionReserv)
		if err != nil {
			Println("Erreur de saisie. Veuillez saisir un numéro valide.")
			continue
		}
		if optionReserv < 1 || optionReserv > 5 {
			Println("Option invalide. Veuillez choisir une option entre 1 et 5.")
			continue
		}
		break
	}
}

func retourMenuReserv() int64 {
	var choix int
	Println("1. Retourner au menu reservation\n2. Menu principal\nChoisissez une option :")
	fmt.Scanln(&choix)

	switch choix {
	case 1:
		// Rien à faire ici, le programme reviendra automatiquement à la boucle principale
	case 2:
		Println("Retour au menu principal!\n\n")
		return 2
	default:
		Println("Option invalide, veuillez réessayer.")
		retourMenuReserv()
	}

	return 1
}
