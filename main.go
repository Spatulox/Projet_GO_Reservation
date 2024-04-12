package main

import (
	// Import functions
	. "Projet_GO_Reservation/src"
	"fmt"
)

func main() {
	var bdd Db

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
}
