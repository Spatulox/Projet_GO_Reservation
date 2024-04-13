package lechauve

import (
	// Import functions

	. "Projet_GO_Reservation/src"
	"fmt"
)

var bdd Db

func GetAllSalle() {
	result, err := bdd.SelectDB(SALLES, []string{"id_salle", "nom", "place"}, nil, true)
	if err != nil {
		Log.Error("Impossible de sélectionner dans la BDD : ", err)
		return
	}

	if result == nil || len(result) == 0 {
		Log.Error("Impossible de sélectionner les données")
		return
	}

	for _, salle := range result {
		id_salle := salle["id_salle"]
		nom := salle["nom"]
		place := salle["place"]

		fmt.Println("ID salle:", id_salle)
		fmt.Println("Nom:", nom)
		fmt.Println("Place:", place)
	}
}
