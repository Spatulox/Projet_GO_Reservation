package lechauve

import (
	// Import functions

	. "Projet_GO_Reservation/src"
	"fmt"
	"strconv"
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

func GetSalleById() {

	fmt.Println("Taper id de la salle que vous voulez")
	id := 0
	fmt.Scanln(&id)
	condition := fmt.Sprintf("id_salle = %d", id)

	result, err := bdd.SelectDB(SALLES, []string{"nom", "place"}, &condition, true)
	if err != nil {
		Log.Error("Impossible de sélectionner dans la BDD : ", err)
		return
	}

	if result == nil || len(result) == 0 {
		Log.Error("Impossible de sélectionner les données")
		return
	}

	firstMap := result[0]
	id_salle := firstMap["id_salle"]
	nom := firstMap["nom"]
	place := firstMap["place"]

	fmt.Println("ID salle:", id_salle)
	fmt.Println("Nom:", nom)
	fmt.Println("Place:", place)

}

func CreateRoom() {
	name := ""
	capacity := 0
	fmt.Println("Taper le nom de la nouvelle salle")
	fmt.Scanln(&name)
	fmt.Println("Taper la capaciter de la nouvelle salle")
	fmt.Scanln(&capacity)

	columns := []string{"nom", "place"}
	values := []string{name, strconv.Itoa(capacity)}

	bdd.InsertDB("SALLES", columns, values, nil, true)

	Log.Infos("Salle créée avec succès")
}

func DeleteRoomByID() {
	fmt.Println("Taper id de la salle que vous voulez")
	id := 0
	fmt.Scanln(&id)

	if err := CheckId(id); err != nil {
		Log.Error("Erreur lors de la vérification de l'existence de la salle : ", err)
		return
	}

	condition := fmt.Sprintf("id_salle = %d", id)
	bdd.DeleteDB("SALLES", &condition, true)
	Log.Infos("Salle supprimée avec succès")
	return
}

func CheckId(id int) error {
	condition := fmt.Sprintf("id_salle = %d", id)
	result, err := bdd.SelectDB("SALLES", []string{"id_salle"}, &condition)
	if err != nil {
		return fmt.Errorf("Erreur lors de la vérification de l'existence de la salle : %v", err)
	}

	if len(result) == 0 {
		return fmt.Errorf("La salle avec l'ID %d n'existe pas", id)
	}

	return nil
}
