// -*- coding: utf-8 -*-

package lechauve

import (
	. "Projet_GO_Reservation/pkg/bdd"
	. "Projet_GO_Reservation/pkg/const"
	. "Projet_GO_Reservation/pkg/json"
	. "Projet_GO_Reservation/pkg/log"
	. "Projet_GO_Reservation/pkg/models"
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

			ListReservations(nil)
		case 2:

			ListReservationsByRoom(nil)
		case 3:

			ListReservationsByDate(nil)
		case 4:

			CreateReservation(nil, nil, nil)
		case 5:

			CancelReservation()
		case 6:

			UpdateReservation(nil, nil)
		case 7:

			DataToJson(ListReservations(nil))
		case 8:

			JsonToData(nil)
		case 9:

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

func ListReservations(condition *string, noPrintRoom ...bool) []Reservation {

	var bdd Db
	// Condition can be nil
	result, err := bdd.SelectDB(RESERVATIONS, []string{"*"}, nil, condition)

	if err != nil || result == nil {
		Log.Error("Erreur lors de la lecture de la Base de donnée", err)
		return nil
	}

	// If noPrint == nil globalement
	if len(noPrintRoom) == 0 || !noPrintRoom[0] {
		resultRes := printReservations(result)
		return resultRes
	}

	resultRes := printReservations(result, true)
	return resultRes

}

//
// ------------------------------------------------------------------------------------------------ //
//

func ListReservationsByRoom(salle *int) []Reservation {

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

	result, err = bdd.SelectDB(RESERVER, []string{"*"}, nil, &tmp)

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

	result, err = bdd.SelectDB(RESERVATIONS, []string{"*"}, nil, &concatCondition)

	resultRes := printReservations(result)
	return resultRes

}

//
// ------------------------------------------------------------------------------------------------ //
//

func ListReservationsByDate(date *string) []Reservation {

	fmt.Println(*date)
	if date != nil {
		dateTime, err := time.Parse("2006-01-02 15:04:05", *date)
		if err != nil {
			Log.Error("Erreur mauvais format de date", err)
			return nil
		}
		*date = dateTime.Format("2006-01-02 15:04:05")
	} else {
		departureDate, departureTime := getDateAndHour()
		departureDateTime := departureDate.Format("2006-01-02") + " " + departureTime.Format("15:04:00")

		*date = departureDateTime
	}

	tmp := "'" + *date + "' BETWEEN horaire_start AND horaire_end"

	result, err := bdd.SelectDB(RESERVATIONS, []string{"*"}, nil, &tmp)

	if err != nil {
		Log.Error("Impossible de récupérer les réservations par date")
		return nil
	}

	resultRes := printReservations(result)
	return resultRes
}

//
// ------------------------------------------------------------------------------------------------ //
//

func CreateReservation(salle *int64, departure *string, end *string) bool {
	var bdd Db
	// Select all the room

	result, err := bdd.SelectDB(SALLES, []string{"*"}, nil, nil)

	if err != nil || result == nil {
		Log.Error("Impossible de lister les salles :/")
		return false
	}

	idMin := result[0]["id_salle"].(int64)
	idMax := result[len(result)-1]["id_salle"].(int64)

	var newSalle int64
	if salle == nil {
		for {
			fmt.Printf("Veuillez saisir une salle entre %d et %d : ", idMin, idMax)

			fmt.Scanln(&newSalle)

			if idMin > newSalle || newSalle > idMax {
				continue
			} else {
				fmt.Printf("Vouc avez choisit la salle %d\n", newSalle)
				break
			}
		}

		salle = &newSalle
	}

	var departureDateTime string
	var endDateTime string

	if departure == nil {

		departureDate, departureTime := getDateAndHour()

		departureDateTime = departureDate.Format("2006-01-02") + " " + departureTime.Format("15:04:00")

		departure = &departureDateTime
	} else {
		fmt.Println(*departure)
		yeet, err := time.Parse("2006-01-02 15:04:05", *departure)
		if err != nil {
			fmt.Println(yeet)
			Log.Error("Erreur mauvais format de date de début", err)
			return false
		}
	}

	if end == nil {

		endDate, endTime := getDateAndHour()

		endDateTime = endDate.Format("2006-01-02") + " " + endTime.Format("15:04:00")

		end = &endDateTime
	} else {
		_, err := time.Parse("2006-01-02 15:04:05", *end)
		if err != nil {
			Log.Error("Erreur mauvais format de date de fin", err)
			return false
		}
	}

	fmt.Println("Date et heure de départ : ", *departure)
	fmt.Println("Date et heure de fin : ", *end)

	var tmp2 = "id_etat != 3"
	var tmpSalle = int(*salle)
	leBool, _ := isRoomAvailable(departure, end, &tmpSalle, &tmp2)
	if leBool == false {
		Println("Annulation de l'enregistrement d'une nouvelle réservation")
		return false
	}

	Println("Toutes les vérifications ont été effectuée, ajout d'une nouvelle réservation")

	// Insertion des données
	bdd.InsertDB(RESERVATIONS, []string{"horaire_start", "horaire_end", "id_etat"}, []string{*departure, *end, "4"})

	// Selectionne la dernière entrée avec MAX(id)
	var tmp = "id_reservation = (SELECT MAX(id_reservation) FROM " + RESERVATIONS + ")"
	resultRes := ListReservations(&tmp, true)

	if resultRes == nil {
		Log.Error("Impossible de sélectionner la dernière réservation rentrée")
		return false
	}

	horaire := fmt.Sprintf("%d", resultRes[0].IdReservation)
	tmp2 = fmt.Sprintf("%d", *salle)
	bdd.InsertDB(RESERVER, []string{"id_salle", "id_reservation"}, []string{tmp2, horaire})

	ListReservations(&tmp)

	return true
}

//
// ------------------------------------------------------------------------------------------------ //
//

func CancelReservation(choix ...int) {
	reservation := ListReservations(nil)

	var option int
	var maxIdReservation, minIdReservation int64
	minIdReservation = reservation[0].IdReservation
	maxIdReservation = reservation[len(reservation)-1].IdReservation

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
				if (m.IdReservation) == int64(option) {
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

func UpdateReservation(state *int, idReservation *int) {
	var bdd Db

	result, err := bdd.SelectDB(ETAT, []string{"*"}, nil, nil)

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
		resultRes := ListReservations(nil)

		var maxIdReservation, minIdReservation int64
		minIdReservation = resultRes[0].IdReservation
		maxIdReservation = resultRes[len(resultRes)-1].IdReservation

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
			for _, m := range resultRes {
				if (m.IdReservation) == int64(idReserv) {
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

	// Need to check if reservation with another state than "Annulé" exists at the same time (room & date/hour) to block the user
	// Get the reservation with id to retrieve date/hour
	var tmp = fmt.Sprintf("id_reservation=%d", idReserv)
	ListResult := ListReservations(&tmp, true)
	// Get the time of the reservation
	if len(ListResult) > 0 {
		horaireStart := fmt.Sprintf("%s", ListResult[0].HoraireStart)
		horaireEnd := fmt.Sprintf("%s", ListResult[0].HoraireEnd)

		var tmp = fmt.Sprintf("id_reservation=%d", idReserv)
		idSalle, err := bdd.SelectDB(RESERVER, []string{"id_salle"}, nil, &tmp)

		if err != nil {
			Log.Error("Impossible de récupérer la salle de la réservation sélectionnée", err)
			return
		}
		salle := idSalle[0]["id_salle"].(int64)
		fmt.Println(horaireStart, horaireEnd, salle)

		var leBool bool
		var length int
		var idEtat int

		var tmp2 = fmt.Sprintf("id_reservation=%d", idReserv)
		IdResult := ListReservations(&tmp2, true)

		if len(IdResult) > 0 {
			idEtat = int(IdResult[0].IdEtat)
		} else {
			Log.Error("Impossible de récupérer le statut de la réservation actuelle")
			return
		}

		tmp = "id_etat != 3"
		intSalle := int(salle)
		if idEtat == 3 {
			// Vérifier s'il y a une autre réservation sur la même plage horaire, sauf celle annulée
			_, length = isRoomAvailable(&horaireStart, &horaireEnd, &intSalle, &tmp)
			if length > 0 {
				Println("Il y a déjà une réservation sur la même plage horaire (Voir au dessus)")
				return
			}
		} else {
			// Vérifier s'il y a plus d'une réservation sur la même plage horaire, sauf celle annulée
			_, length = isRoomAvailable(&horaireStart, &horaireEnd, &intSalle, &tmp)
			if length > 1 {
				Log.Error("Ouatte da héque brau, c'est pas normal ça O.O")
				return
			}
		}

		// We check at the same date/hour, so there is always one reservation
		if leBool == false && length > 1 {
			Println("Une autre réservation à remplacé la votre. Veuillez en créer une autre dans une autre salle, ou un autre horaire de début/fin")
			return
		} else if leBool == true {
			Log.Error("On ne peut pas mettre a jour une reservation qui n'existe pas :/")
			return
		} else {
			Println("Mise a jour de l'état de la salle")
		}
	}

	tmp = fmt.Sprintf("id_reservation = %d", idReserv)
	stateStr := strconv.FormatInt(newState, 10)
	bdd.UpdateDB(RESERVATIONS, []string{"id_etat"}, []string{stateStr}, &tmp)

	fmt.Printf("Etat changé pour %s pour la réservation %d\n\n", stateStr, idReserv)
	return
}

//
// ------------------------------------------------------------------------------------------------ //
//

func isRoomAvailable(departureDateTime *string, endDateTime *string, salle *int, condition *string) (bool, int) {

	// Selectionne dans la BDD pour savoir si y'a quelque chose enregistré a cette date/heure et dans la salle
	var fin = ""
	if condition != nil {
		fin = " AND " + *condition
	}

	var tmp = "(('" + *departureDateTime + "' BETWEEN horaire_start AND horaire_end) OR ('" + *endDateTime + "' BETWEEN horaire_start AND horaire_end) OR (horaire_start BETWEEN '" + *departureDateTime + "' AND '" + *endDateTime + "') OR (horaire_end BETWEEN '" + *departureDateTime + "' AND '" + *endDateTime + "'))" + fin
	result, err := bdd.SelectDB(RESERVATIONS, []string{"*"}, nil, &tmp, true)

	if err != nil {
		Log.Error("Impossible de vérifier si il existe déjà une reservation a cette date")
		return false, 0
	}

	// If y'a déjà une réservation ce jour
	if len(result) > 1 {
		// It's not normal to have two result
		return false, len(result)
	} else if len(result) > 0 {

		for _, r := range result {
			tmp = fmt.Sprintf("id_reservation = %d", r["id_reservation"].(int64))
			result, err = bdd.SelectDB(RESERVER, []string{"id_salle"}, nil, &tmp, true)

			// Il y'a déjà une reservation ce jour et dans cette salle
			if err != nil || result != nil {
				Println("\nIl existe (déjà) une reservation à cette date et heure dans cette salle : ")
				var idreservation = fmt.Sprintf("id_reservation=%d", r["id_reservation"])
				ListReservations(&idreservation)
				return false, len(result)
			}
		}

	}

	return true, len(result)
}

//
// ------------------------------------------------------------------------------------------------ //
//

func printReservations(result []map[string]interface{}, noPrint ...bool) []Reservation {

	if len(noPrint) == 0 || !noPrint[0] {
		Println("------------------------------")
		Println("------- RESERVATION(S) -------")
	}

	reservations := make([]Reservation, 0, len(result))

	for _, sResult := range result {

		horaireDebut := sResult["horaire_start"]
		horaireFin := sResult["horaire_end"]
		idEtat := sResult["id_etat"]
		idReservation := sResult["id_reservation"]

		tmp := fmt.Sprintf("id_etat=%v", idEtat)
		etatResult, err := bdd.SelectDB(ETAT, []string{"nom_etat"}, nil, &tmp)
		nomEtat := etatResult[0]["nom_etat"].(string)

		tmp = fmt.Sprintf("id_reservation=%v", idReservation)
		idSalleResult, err := bdd.SelectDB(RESERVER, []string{"id_salle"}, nil, &tmp)

		var sallesResult = make([]map[string]interface{}, 0)
		//var err error
		var idSalleTmp int64
		if err == nil && len(idSalleResult) > 0 {
			tmp = fmt.Sprintf("id_salle=%v", idSalleResult[0]["id_salle"])
			sallesResult, err = bdd.SelectDB(SALLES, []string{"nom", "place"}, nil, &tmp)

			idSalleTmp = idSalleResult[0]["id_salle"].(int64)

		} else {
			idSalleTmp = -1
		}

		if len(noPrint) == 0 || !noPrint[0] {
			Println("------------------------------")
		}

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

		reservation := Reservation{
			HoraireStart:  horaireDebut.(string),
			HoraireEnd:    horaireFin.(string),
			NomEtat:       nomEtat,
			IdEtat:        idEtat.(int64),
			IdReservation: idReservation.(int64),
			IdSalle:       idSalleTmp,
			NomSalle:      salleName,
			PlaceSalle:    sallePlace,
		}
		reservations = append(reservations, reservation)

		// Print
		if len(noPrint) == 0 || !noPrint[0] {
			fmt.Println("ID Réservation:", idReservation)
			fmt.Println("Horaire Début:", horaireDebut)
			fmt.Println("Horaire Fin:", horaireFin)

			if err != nil {
				Log.Error("Impossible de récupérer l'état de la réservation")
				fmt.Println("ID Etat:", idEtat)
			} else {
				fmt.Println("Etat : ", etatResult[0]["nom_etat"])
			}

			fmt.Println("Nom Salle :", salleName)
			fmt.Println("Place Salle :", sallePlace)
		}

	}

	if len(noPrint) == 0 || !noPrint[0] {
		Println("------------------------------")
	}

	return reservations
}

//
// ------------------------------------------------------------------------------------------------ //
//

func getDateAndHour() (time.Time, time.Time) {

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

	return departureDate, departureTime
}

//
// ------------------------------------------------------------------------------------------------ //
//

func menuReserv() {
	for {
		Println("\n-----------------------------------------------------\nMenu Réservation\n-----------------------------------------------------")
		Println("1. Lister les reservations\n2. Lister les reservations par salles\n3. Lister les reservations par date\n4. Créer une réservation\n5. Annuler une réservation\n6. Mettre à jour une reservation\n7. Exporter toutes les réservations en json\n8. Importer des réservations depuis un fichier json\n9. Menu Principal\nChoisissez une option :")
		_, err := fmt.Scanln(&optionReserv)
		if err != nil {
			Println("Erreur de saisie. Veuillez saisir un numéro valide.")
			continue
		}
		if optionReserv < 1 || optionReserv > 9 {
			Println("Option invalide. Veuillez choisir une option entre 1 et 7.")
			continue
		}
		break
	}
}

func retourMenuReserv() int64 {
	var choix int
	Println("\n-------------Retour-------------")
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
