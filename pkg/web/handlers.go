// -*- coding: utf-8 -*-

package web

import (
	. "Projet_GO_Reservation/pkg/const"
	. "Projet_GO_Reservation/pkg/models"
	. "Projet_GO_Reservation/pkg/reservation"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func EnableHandlers() {

	// Create the directory with static file like CSS and JS
	staticDir := http.Dir("templates/src")
	staticHandler := http.FileServer(staticDir)
	http.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	http.HandleFunc(RouteIndex, IndexHandler)
	http.HandleFunc(RouteIndexReservation, ReservationHandler)
	http.HandleFunc(RouteListReservation, ListByRoomDateIdReservationHandler)
	http.HandleFunc(RouteCreateReservation, CreateReservationHandler)
	http.HandleFunc(RouteCancelReservation, CancelReservationHandler)
	http.HandleFunc(RouteUpdateReservation, UpdateReservationHanlder)

	http.HandleFunc(RouteGetAllRoolAvailable, GetAllRoomAvailHandler)

	Log.Infos("Handlers Enabled")

	Log.Infos("Starting server on port " + PORT)
	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		Log.Error("Failed to start server: ", err)
		return
	}
	Log.Infos("Server stopped on port " + PORT)

}

//
// ------------------------------------------------------------------------------------------------ //
//

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		currentTime := time.Now().Format("2006-01-02 15:04")
		templates.ExecuteTemplate(w, "menu.html", map[string]interface{}{
			"now": currentTime,
		})
	}

	/*	if r.Method == http.MethodGet {
		templates.ExecuteTemplate(w, "menu.html", nil)
	}*/
}

//
// ------------------------------------------------------------------------------------------------ //
//

func ReservationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if r.Method == http.MethodGet {
		result := ListReservations(nil)
		if result == nil {
			Log.Error("Data are null for unknown reason :/")
			var msg = "Impossible to retrieve data"
			//http.Redirect(w, r, "/reservation?message="+msg, http.StatusSeeOther)
			// Exécuter le template avec l'URL et le message
			templates.ExecuteTemplate(w, "reservations.html", map[string]interface{}{
				"message": msg,
				"result":  nil,
			})
			return
		}
		templates.ExecuteTemplate(w, "reservations.html", map[string]interface{}{
			"result":  result,
			"message": nil,
		})

	}
}

func ListByRoomDateIdReservationHandler(w http.ResponseWriter, r *http.Request) {

	roomStr := r.URL.Query().Get("idRoom")

	dateStr := r.URL.Query().Get("idDate")

	idStr := r.URL.Query().Get("idReserv")

	if r.Method == http.MethodGet {

		if roomStr == NullString && dateStr == NullString && idStr == NullString {
			// Si aucun ID n'est fourni, redirigez vers la page de liste des réservations
			var msg = "Vous ne pouvez pas acceder à cette page sans spécifier un truc :/"
			http.Redirect(w, r, "/reservation?message="+msg, http.StatusSeeOther)
			return
		}

		var result []Reservation

		if roomStr != NullString {
			idRoom, err := strconv.Atoi(roomStr)
			if err != nil {
				// Gestion de l'erreur de conversion en entier
				http.Error(w, "ID de salle invalide", http.StatusBadRequest)
				return
			}

			Log.Infos("Listing des réservations par Salles")
			result = ListReservationsByRoom(&idRoom)
		}

		if dateStr != NullString {
			Log.Infos("Listing des réservations par Date")
			dateStr = strings.Replace(dateStr, "T", " ", 1)
			dateStr = dateStr + ":00"
			result = ListReservationsByDate(&dateStr)
		}

		if idStr != NullString {
			Log.Infos("Listing des réservations par ID (reservation)")
			var tmp = "id_reservation=" + idStr
			result = ListReservations(&tmp)
			// It have a special pages yes
			if result != nil {
				templates.ExecuteTemplate(w, "soloReservation.html", result)
			} else {
				var msg = "Impossible de trouver cette réservation"
				http.Redirect(w, r, "/reservation?message="+msg, http.StatusSeeOther)
			}
			return
		}

		if result == nil {
			Log.Error("No result")
			var msg = "Impossible te retrieve data"
			http.Redirect(w, r, "/reservation?message="+msg, http.StatusSeeOther)
			return
		}

		//templates.ExecuteTemplate(w, "reservations.html", result)
		templates.ExecuteTemplate(w, "reservations.html", map[string]interface{}{
			"message": nil,
			"result":  result,
		})

	}
}

//
// ------------------------------------------------------------------------------------------------ //
//

func CreateReservationHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if r.Method == http.MethodGet {
		templates.ExecuteTemplate(w, "creerReservations.html", map[string]interface{}{
			"message": nil,
			"result":  nil,
		})
		//templates.ExecuteTemplate(w, "creerReservations.html", nil)
	} else if r.Method == http.MethodPost {
		horaireStartDate := r.FormValue("horaire_start_date")
		horaireStartTime := r.FormValue("horaire_start_time") + ":00"
		horaireEndDate := r.FormValue("horaire_end_date")
		horaireEndTime := r.FormValue("horaire_end_time") + ":00"
		salle := r.FormValue("id_salle")

		salleInt64, err := strconv.ParseInt(salle, 10, 64)

		resultSalle := GetAllSalle()
		leBool := false
		for _, m := range resultSalle {
			if m.IdSalle == salleInt64 {
				leBool = true
				break
			}
		}

		if leBool == false {
			var msg = "Cette ID de salle n'existe pas :/"
			Log.Error(msg)
			http.Redirect(w, r, "/reservation/create?message="+msg, http.StatusSeeOther)
			return
		}

		if err != nil {
			var msg = "Erreur dans le format de la date/heure de début"
			Log.Error(msg)
			http.Redirect(w, r, "/reservation/create?message="+msg, http.StatusSeeOther)
			return
		}

		horaireStartDateTime, err := time.Parse("2006-01-02 15:04:05", horaireStartDate+" "+horaireStartTime)
		if err != nil {
			var msg = "Erreur dans le format de la date/heure de début"
			Log.Error(msg)
			http.Redirect(w, r, "/reservation/create?message="+msg, http.StatusSeeOther)
			return
		}

		horaireEndDateTime, err := time.Parse("2006-01-02 15:04:05", horaireEndDate+" "+horaireEndTime)
		if err != nil {
			var msg = "Erreur dans le format de la date/heure de début"
			Log.Error(msg)
			http.Redirect(w, r, "/reservation/create?message="+msg, http.StatusSeeOther)
			return
		}

		today := time.Now()

		if horaireStartDateTime.Before(today) || horaireStartDateTime.Equal(today) {
			var msg = "Impossible de créer une réservation avant aujourd'hui ou aujourd'hui"
			Log.Error(msg)
			http.Redirect(w, r, "/reservation/create?message="+msg, http.StatusSeeOther)
			return
		}

		if horaireStartDateTime.After(horaireEndDateTime) {
			var msg = "La fin de la réservation doit être après le début de celle-ci"
			Log.Error(msg)
			http.Redirect(w, r, "/reservation/create?message="+msg, http.StatusSeeOther)
			return
			//La fin est avant le début ??
		}

		if horaireStartDateTime.Equal(horaireEndDateTime) {
			var msg = "Vous ne pouvez pas faire des réservation de moins de 1 minute"
			Log.Error(msg)
			http.Redirect(w, r, "/reservation/create?message="+msg, http.StatusSeeOther)
			return
			// La fin doit être différente du début
		}

		horaireStartSeconds := horaireStartDate + " " + horaireStartTime
		horaireEndSeconds := horaireEndDate + " " + horaireEndTime

		result := CreateReservation(&salleInt64, &horaireStartSeconds, &horaireEndSeconds)
		if result == false {
			var msg = "An error occured"
			Log.Error(msg)
			http.Redirect(w, r, "/reservation?message="+msg, http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, RouteIndexReservation, http.StatusSeeOther)
	}
}

//
// ------------------------------------------------------------------------------------------------ //
//

func CancelReservationHandler(w http.ResponseWriter, r *http.Request) {

	idReserv := r.URL.Query().Get("idReserv")
	var tmp = "id_reservation=" + idReserv

	result := ListReservations(&tmp)
	if len(result) == 0 {
		Log.Error("Aucune reservation contient cet ID")
		http.Error(w, "Aucune réservation trouvée pour l'ID de salle "+idReserv, http.StatusBadRequest)
		return
	}

	newChoix, err := strconv.Atoi(idReserv)

	if err != nil {
		Log.Error("Impossible to convert the id from string to int")
		http.Error(w, "Impossible to convert the id from string to int : "+idReserv, http.StatusBadRequest)
		return
	}

	CancelReservation(newChoix)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Réservation annulée avec succès"))
	Log.Infos("Reservation annulée avec succès !")
	return
}

//
// ------------------------------------------------------------------------------------------------ //
//

func UpdateReservationHanlder(w http.ResponseWriter, r *http.Request) {
	leString := r.URL.Query().Get("idReserv")

	Log.Debug(leString)

	idReserv := strings.Split(leString, "?")[0]
	newEtat := strings.Split(leString, "?etat=")[1]

	Log.Debug(idReserv)
	Log.Debug(newEtat)
	//return

	if idReserv == "nil" || newEtat == "nil" {
		var msg = "Vous ne pouvez pas mettre à jour une réservation avec le même statut :/"
		Log.Error(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	var err error
	var newReserv, newState int

	newReserv, err = strconv.Atoi(idReserv)
	if err != nil {
		var msg = "Impossible de transformer l'ID reserv string en int : " + idReserv
		Log.Error(msg, err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	newState, err = strconv.Atoi(newEtat)
	if err != nil {
		var msg = "Impossible de transformer l'ID etat string en int : " + newEtat
		Log.Error(msg, err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	UpdateReservation(&newState, &newReserv)
	//return
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Réservation annulée avec succès"))
	Log.Infos("Reservation annulée avec succès !")
	//http.Redirect(w, r, "", http.StatusSeeOther)
	return
}

//
// ------------------------------------------------------------------------------------------------ //
//

func GetAllRoomAvailHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		//templates.ExecuteTemplate(w, "sallesDispo.html", nil)
		// Affichage de toutes les salles disponible selon un horaire, etc... (OU PAS, on peut le faire aussi dans la page des salles de bases)
	} else if r.Method == http.MethodPost {

		var params struct {
			StartDateTime string `json:"startDateTime"`
			EndDateTime   string `json:"endDateTime"`
		}
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			var msg = "Erreur lors de la lecture des paramètres"
			Log.Error(msg)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}
		/*
			fmt.Println(params.StartDateTime)
			fmt.Println(params.EndDateTime)
			return*/

		sallesAvail := GetAllSalleDispo(&params.StartDateTime, &params.EndDateTime)
		if sallesAvail == nil {
			var msg = "Impossible de récupérer les salles disponibles"
			Log.Error(msg)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}

		jsonData, err := json.Marshal(sallesAvail)
		if err != nil {
			var msg = "Erreur lors de la conversion en JSON, impossible d'avoir les salles disponibles"
			http.Error(w, msg, http.StatusInternalServerError)
			Log.Error(msg, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)

		Log.Infos("Envoie des salles disponibles avec succès !")
	}

}

/*
func SallesHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if r.Method == http.MethodGet {
		templates.ExecuteTemplate(w, "salles.html", nil)
	} else if r.Method == http.MethodPost {
		idReservation := r.FormValue("id_reservation")

		err := CancelReservation(idReservation)
		if err != nil {
			templates.ExecuteTemplate(w, "salles.html", struct{ Error string }{err.Error()})
			return
		}

		// Rediriger vers la page d'accueil
		http.Redirect(w, r, routeIndex, http.StatusSeeOther)
	}
}*/
