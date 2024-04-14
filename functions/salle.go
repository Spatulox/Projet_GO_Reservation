package lechauve

import (
	. "Projet_GO_Reservation/bdd"
	. "Projet_GO_Reservation/const"
	. "Projet_GO_Reservation/log"
	"fmt"
	"strconv"
)

var optionSalle int
var bdd Db

func MenuSalle() {
	for {
		menuSalle()
		switch optionSalle {
		case 1:
			GetAllSalle()
		case 2:
			GetSalleById()
		case 3:
			CreateRoom()
		case 4:
			DeleteRoomByID()
		case 5:
			Println("Retour menu principal")
			return
		}
		if retourMenuSalle() == 2 {
			return
		}
	}
}

func GetAllSalle() []map[string]interface{} {
	result, err := bdd.SelectDB(SALLES, []string{"id_salle", "nom", "place"}, nil, nil, true)
	if err != nil {
		Log.Error("Impossible de sélectionner dans la BDD : ", err)
		return nil
	}

	if result == nil || len(result) == 0 {
		Log.Error("Impossible de sélectionner les données")
		return nil
	}

	Println("------------------------------")
	Println("----------- SALLES -----------")
	for _, salle := range result {
		Println("------------------------------")
		id_salle := salle["id_salle"]
		nom := salle["nom"]
		place := salle["place"]

		fmt.Println("ID salle:", id_salle)
		fmt.Println("Nom:", nom)
		fmt.Println("Place:", place)
	}
	Println("------------------------------")

	return result
}

func GetSalleById() {

	fmt.Println("Taper id de la salle que vous voulez")
	id := 0
	fmt.Scanln(&id)
	condition := fmt.Sprintf("id_salle = %d", id)

	result, err := bdd.SelectDB(SALLES, []string{"nom", "place"}, nil, &condition, true)
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

	bdd.InsertDB("SALLES", columns, values, true)

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
	result, err := bdd.SelectDB("SALLES", []string{"id_salle"}, nil, &condition)
	if err != nil {
		return fmt.Errorf("Erreur lors de la vérification de l'existence de la salle : %v", err)
	}

	if len(result) == 0 {
		return fmt.Errorf("La salle avec l'ID %d n'existe pas", id)
	}

	return nil
}

func menuSalle() {
	for {
		Println("-----------------------------------------------------\nBienvenue dans le Menu Salle\n-----------------------------------------------------\n")
		Println("1.Lister les salles \n2.Selectioner une salles avec un id \n3.cree une salle \n4.supprimer une salle \n5.Retour menu principal\nChoisissez une option :")
		_, err := fmt.Scanln(&optionSalle)
		if err != nil {
			Println("Erreur de saisie. Veuillez saisir un numéro valide.")
			continue
		}
		if optionSalle < 1 || optionSalle > 5 {
			Println("Option invalide. Veuillez choisir une option entre 1 et 5.")
			continue
		}
		break
	}
}

func retourMenuSalle() int {
	var choix int
	Println("1. Retourner au menu Salle\n2. Retourner au menu Principal\nChoisissez une option :")
	fmt.Scanln(&choix)
	switch choix {
	case 1:
	// Rien à faire ici, le programme reviendra automatiquement à la boucle principale
	case 2:
		Println("Retour au menu principal!\n\n")
		return 2
	default:
		Println("Option invalide, veuillez réessayer.")
		retourMenuSalle()
	}
	return 1
}
