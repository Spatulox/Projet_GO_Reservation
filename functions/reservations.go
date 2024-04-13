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
			listReservationsByRoom(nil)
		case 3:

			createReservation(nil, nil)
		case 4:

			cancelReservation()
		case 5:

			updateReservation(nil, nil)
		case 6:

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
	// Condition can be nil
	result, err := bdd.SelectDB(RESERVATIONS, []string{"*"}, condition)

	if err != nil || result == nil {
		Log.Error("Erreur lors de la lecture de la Base de donnée", err)
		return nil
	}

	printReservations(result)

	return result

}

//
// ------------------------------------------------------------------------------------------------ //
//

func listReservationsByRoom(salle *int) []map[string]interface{} {

	var bdd Db
	// Condition can be nil

	var tmp string
	var result []map[string]interface{}
	var err error

	if salle != nil {
		tmp = fmt.Sprintf("id_salle=%d", *salle)
	} else {
		// Ask for the user for the room
		result = GetAllSalle()

		maxIdSalle := result[len(result)-1]["id_salle"].(int64)
		minIdSalle := result[0]["id_salle"].(int64)
		var choix int64

		for {
			fmt.Printf("Choisisser une salle via son ID (entre %d et %d) : ", minIdSalle, maxIdSalle)
			fmt.Scanln(&choix)
			fmt.Println(choix)

			leBool := false

			for _, r := range result {
				if r["id_salle"] == choix {
					leBool = true
					break
				}
			}

			if choix > minIdSalle && choix < maxIdSalle {
				break
			}

			if leBool == false {
				continue
			}
			break
		}

		tmp = fmt.Sprintf("id_salle=%d", choix)
	}

	result, err = bdd.SelectDB(RESERVER, []string{"*"}, &tmp)

	if err != nil || result == nil {
		Log.Error("Erreur lors de la lecture de la Base de donnée", err)
		return nil
	}

	concatCondition := ""

	for _, r := range result {
		if r["id_reservation"] == result[0]["id_reservation"] {
			concatCondition = fmt.Sprintf("id_reservation=%d", r["id_reservation"])
		} else {
			concatCondition = fmt.Sprintf("%s OR id_reservation=%d", concatCondition, r["id_reservation"])
		}

	}

	result, err = bdd.SelectDB(RESERVATIONS, []string{"*"}, &concatCondition)

	printReservations(result)
	return result

}

//
// ------------------------------------------------------------------------------------------------ //
//

func createReservation(salle *int64, departure *string) bool {
	var bdd Db
	// Select all the room

	result, err := bdd.SelectDB(SALLES, []string{"*"}, nil)

	if err != nil || result == nil {
		Log.Error("Impossible de lister les salles :/")
		return false
	}

	idMin := result[0]["id_salle"].(int64)
	idMax := result[len(result)-1]["id_salle"].(int64)

	if salle == nil {
		var newSalle int64
		for {
			fmt.Printf("Veuillez saisir une salle entre %d et %d : ", idMin, idMax)

			fmt.Scanln(&salle)
			if idMin > newSalle || newSalle > idMax {
				continue
			} else {
				fmt.Printf("Vouc avez choisit la salle %d\n", newSalle)
				break
			}
		}

		*salle = newSalle
	}

	if departure == nil {

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

		*departure = departureDateTime
	} else {
		_, err := time.Parse("2006-01-02 15:04:05", *departure)
		if err != nil {
			Log.Error("Erreur mauvais format de date", err)
			return false
		}
	}

	fmt.Println("Date et heure de départ : ", departure)

	if isRoomAvailable(departure, salle) == false {
		return false
	}

	Println("Toutes les vérifications ont été effectuée, ajout d'une nouvelle réservation")

	// Insertion des données
	bdd.InsertDB(RESERVATIONS, []string{"horaire", "id_etat"}, []string{*departure, "4"})

	// Select the line with the MAX(id)
	var tmp = "id_reservation = (SELECT MAX(id_reservation) FROM " + RESERVATIONS + ")"
	result = listReservations(&tmp)

	if result == nil {
		Log.Error("Impossible de sélectionner la dernière réservation rentrée")
		return false
	}

	horaire := fmt.Sprintf("%d", result[0]["id_reservation"].(int64))
	bdd.InsertDB(RESERVER, []string{"id_salle", "id_reservation"}, []string{fmt.Sprintf("%d", salle), horaire})

	return true
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
			Log.Error("Erreur de conversion de l'état de string vers int64 :", err)
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

func isRoomAvailable(departureDateTime *string, salle *int64) bool {

	// Selectionne dans la BDD pour savoir si y'a quelque chose enrgistré a cette date/heure et dans la salle
	var tmp = "horaire = '" + *departureDateTime + "'"
	result, err := bdd.SelectDB(RESERVATIONS, []string{"id_reservation"}, &tmp)

	if err != nil {
		Log.Error("Impossible de vérifier si il existe déjà une reservation a cette date")
		return false
	}

	// If y'a déjà une réservation ce jour
	if len(result) > 0 {
		tmp = fmt.Sprintf("id_reservation = %d", result[0]["id_reservation"].(int64))

		result, err = bdd.SelectDB(RESERVER, []string{"id_salle"}, &tmp)

		// Il y'a déjà une reservation ce jour et dans cette salle
		if err != nil || result != nil {
			fmt.Printf("Il existe déjà une reservation à cette date %s et dans cette salle %d\n", *departureDateTime, *salle)
			Println("------------------------------")
			return false
		}
	}

	return true
}

//
// ------------------------------------------------------------------------------------------------ //
//

func printReservations(result []map[string]interface{}) {
	Println("------------------------------")
	Println("-------- RESERVATIONS --------")
	for _, sResult := range result {

		horaire := sResult["horaire"]
		idEtat := sResult["id_etat"]
		idReservation := sResult["id_reservation"]

		tmp := fmt.Sprintf("id_etat=%v", idEtat)
		etatResult, err := bdd.SelectDB(ETAT, []string{"nom_etat"}, &tmp)

		tmp = fmt.Sprintf("id_reservation=%v", idReservation)
		idSalleResult, err := bdd.SelectDB(RESERVER, []string{"id_salle"}, &tmp)

		tmp = fmt.Sprintf("id_salle=%v", idSalleResult[0]["id_salle"])
		sallesResult, err := bdd.SelectDB(SALLES, []string{"nom", "place"}, &tmp)

		var salleName string
		var sallePlace int64
		if len(sallesResult) > 0 {
			salleName = sallesResult[0]["nom"].(string)
			sallePlace = sallesResult[0]["place"].(int64)
		} else {
			Log.Error("Aucune salle trouvée")
			salleName = "N/A"
			sallePlace = -1
		}

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

		fmt.Println("Nom Salle :", salleName)
		fmt.Println("Place Salle :", sallePlace)

	}
	Println("------------------------------")
}

//
// ------------------------------------------------------------------------------------------------ //
//

func menuReserv() {
	for {
		Println("-----------------------------------------------------\nMenu Réservation\n-----------------------------------------------------")
		Println("1. Lister les reservations\n2. Lister les reservations par salles\n3. Créer une réservation\n4. Annuler une réservation\n5. Mettre à jour une reservation\n6. Menu Principal\nChoisissez une option :")
		_, err := fmt.Scanln(&optionReserv)
		if err != nil {
			Println("Erreur de saisie. Veuillez saisir un numéro valide.")
			continue
		}
		if optionReserv < 1 || optionReserv > 6 {
			Println("Option invalide. Veuillez choisir une option entre 1 et 6.")
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
