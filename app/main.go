package main

import (
	"fmt"
)

var choix int

func main() {
	menu()

}

func menu() {

	fmt.Println("Bienvenue dans le Service de Réservation en Ligne\n-----------------------------------------------------\n")
	fmt.Println("1. Lister les salles disponibles\n2. Créer une réservation\n3. Annuler une réservation\n4. Visualiser les réservations\n5. Quitter\nChoisissez une option :")
	fmt.Scanln(&choix)
}
