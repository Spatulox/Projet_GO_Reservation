package _const

import (
	. "Projet_GO_Reservation/pkg/log"
)

// Constante pour les strings
const (
	NullString = ""
	NullInt    = 0
)

// Constante pour les tables SQL
const (
	ETAT         = "ETATS"
	RESERVATIONS = "RESERVATIONS"
	RESERVER     = "RESERVER"
	SALLES       = "SALLES"
)

// Instanciation d'une variable Log pour éviter de la faire dans les autres fichiers.
var Log LogHelper
