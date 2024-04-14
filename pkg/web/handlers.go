package web

import (
	. "Projet_GO_Reservation/pkg/const"
	. "Projet_GO_Reservation/pkg/reservation"
	"fmt"
	"net/http"
	"strconv"
)

func EnableHandlers() {

	http.HandleFunc(RouteIndex, IndexHandler)
	http.HandleFunc(RouteIndexReservation, ReservationHandler)
	http.HandleFunc(RouteListReservationRoom, ListByIdReservationHandler)
	http.HandleFunc(RouteListReservationDate, ListByDateReservationHandler)

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
		templates.ExecuteTemplate(w, "menu.html", nil)
	}
}

//
// ------------------------------------------------------------------------------------------------ //
//

func ReservationHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		result := ListReservations(nil)
		if result == nil {
			Log.Error("Data are null for unknown reason :/")
			templates.ExecuteTemplate(w, "reservations.html", nil)
		}
		templates.ExecuteTemplate(w, "reservations.html", result)

	}
}

func ListByIdReservationHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("idRoom")

	if r.Method == http.MethodGet {

		if idStr == "" {
			// Si aucun ID n'est fourni, redirigez vers la page de liste des réservations
			http.Redirect(w, r, "/reservations", http.StatusSeeOther)
			return
		}

		idRoom, err := strconv.Atoi(idStr)
		if err != nil {
			// Gestion de l'erreur de conversion en entier
			http.Error(w, "ID de salle invalide", http.StatusBadRequest)
			return
		}

		fmt.Println(idRoom)
		result := ListReservationsByRoom(&idRoom)

		if result == nil {
			Log.Error("No result")
			templates.ExecuteTemplate(w, "reservations.html", nil)
		}
		templates.ExecuteTemplate(w, "reservations.html", result)

	}
}

func ListByDateReservationHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("idDate")

	if r.Method == http.MethodGet {
		result := ListReservationsByDate(&idStr)
		if result == nil {
			Log.Error("Data are null for unknown reason :/")
			templates.ExecuteTemplate(w, "reservations.html", nil)
		}
		templates.ExecuteTemplate(w, "reservations.html", result)

	}
}

//
// ------------------------------------------------------------------------------------------------ //
//

/*func CreateReservationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templates.ExecuteTemplate(w, "reservations.html", nil)
	} else if r.Method == http.MethodPost {
		horaireStart := r.FormValue("horaire_start")
		horaireEnd := r.FormValue("horaire_end")
		salle := r.FormValue("salle")

		result := CreateReservation(horaireStart, horaireEnd, salle)
		if result == false {
			err := fmt.Errorf("An error occured")
			templates.ExecuteTemplate(w, "reservations.html", struct{ Error string }{err.Error()})
			return
		}

		http.Redirect(w, r, RouteIndex, http.StatusSeeOther)
	}
}*/

//
// ------------------------------------------------------------------------------------------------ //
//

/*
func SallesHandler(w http.ResponseWriter, r *http.Request) {
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
