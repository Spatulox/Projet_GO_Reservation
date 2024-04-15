package lechauve

import (
	. "Projet_GO_Reservation/pkg/bdd"
	. "Projet_GO_Reservation/pkg/const"
	. "Projet_GO_Reservation/pkg/log"
	. "Projet_GO_Reservation/pkg/models"
	"fmt"
	"strconv"
	"time"
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
			GetAllSalleDispo(nil, nil)
		case 3:
			GetSalleById(nil)
		case 4:
			CreateRoom(nil, nil)
		case 5:
			DeleteRoomByID(nil)
		case 6:
			Println("Retour menu principal")
			return
		}
		if retourMenuSalle() == 2 {
			return
		}
	}
}

//
// ------------------------------------------------------------------------------------------------ //
//

func GetAllSalle() []Salle {
	result, err := bdd.SelectDB(SALLES, []string{"id_salle", "nom", "place"}, nil, nil)
	if err != nil {
		Log.Error("Impossible de sélectionner dans la BDD : ", err)
		return nil
	}

	if result == nil || len(result) == 0 {
		Log.Error("Impossible de sélectionner les données")
		return nil
	}

	salleStruct := printSalle(result)

	return salleStruct
}

//
// ------------------------------------------------------------------------------------------------ //
//

func GetSalleById(salle *int) []Salle {

	var id int

	if salle == nil {

		result := GetAllSalle()

		var minIdSalle = result[0].IdSalle
		var maxIdSalle = result[len(result)-1].IdSalle

		var idSalle int
		fmt.Printf("Taper id de la salle que vous voulez entre %d et %d : ", minIdSalle, maxIdSalle)
		for {

			_, err := fmt.Scanln(&idSalle)

			if err != nil {
				Println("Erreur de saisie. Veuillez saisir un numéro entier valide : .")
				continue
			}

			if int64(idSalle) < minIdSalle || int64(idSalle) > maxIdSalle {
				fmt.Printf("Option invalide. Veuillez choisir une option entre %d et %d\n : ", minIdSalle, maxIdSalle)
				continue
			}

			f := false
			for _, m := range result {
				if m.IdSalle == int64(idSalle) {
					f = true
					break
				}
			}
			if f == false {
				Println("Cette salle n'existe pas")
				fmt.Printf("Taper id de la salle que vous voulez entre %d et %d : ", minIdSalle, maxIdSalle)
				continue
			}

			break
		}

		id = idSalle
	} else {
		id = *salle
	}

	condition := fmt.Sprintf("id_salle = %d", id)

	result, err := bdd.SelectDB(SALLES, []string{"id_salle", "nom", "place"}, nil, &condition)
	if err != nil {
		Log.Error("Impossible de sélectionner dans la BDD : ", err)
		return nil
	}

	if result == nil || len(result) == 0 {
		Log.Error("Impossible de sélectionner les données")
		return nil
	}

	salles := printSalle(result)

	return salles

}

//
// ------------------------------------------------------------------------------------------------ //
//

func CreateRoom(name *string, capacity *int) bool {

	if name == nil {
		Println("Taper le nom de la nouvelle salle")
		for {

			_, err := fmt.Scanln(name)

			if err != nil || *name == NullString {
				Println("Erreur de saisie. Veuillez saisir une nom valide : ")
				continue
			}

			break
		}
	}

	if capacity == nil {
		Println("Taper la capaciter de la nouvelle salle")
		for {
			_, err := fmt.Scanln(capacity)

			if err != nil || *capacity < 1 {
				Println("Erreur de saisie. Veuillez saisir une capacité valide et supérieur à 0 : ")
				continue
			}
			break
		}
	}

	columns := []string{"nom", "place"}
	values := []string{*name, strconv.Itoa(*capacity)}

	bdd.InsertDB("SALLES", columns, values)

	Log.Infos("Salle créée avec succès")

	return true
}

//
// ------------------------------------------------------------------------------------------------ //
//

func DeleteRoomByID(salle *int) bool {

	var id int

	if salle == nil {
		var err error

		for {
			Println("Taper id de la salle que vous voulez")
			id, err = fmt.Scanln(&id)
			if err != nil {
				break
			}
			fmt.Println("Erreur : Veuillez entrer un nombre entier.")
			_, _ = fmt.Scanln()
		}
	} else {
		id = *salle
	}

	if err := CheckId(id); err != nil {
		Log.Error("Erreur lors de la vérification de l'existence de la salle : ", err)
		return false
	}

	condition := fmt.Sprintf("id_salle = %d", id)
	bdd.DeleteDB("SALLES", &condition)
	Log.Infos("Salle supprimée avec succès")
	return true
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

//
// ------------------------------------------------------------------------------------------------ //
//

func GetAllSalleDispo(debutDateTime *string, endDateTime *string) []Salle {

	if debutDateTime == nil || endDateTime == nil {
		var debutDate, debutTime, finDate, finTime time.Time

		debutDate, debutTime = getDateAndHour()

		finDate, finTime = getDateAndHour()

		*debutDateTime = debutDate.Format("2006-01-02") + " " + debutTime.Format("15:04:00")
		*endDateTime = finDate.Format("2006-01-02") + " " + finTime.Format("15:04:00")
	} else {
		*debutDateTime += ":00"
		*endDateTime += ":00"
	}

	condition := "SALLES.id_salle NOT IN" +
		"(SELECT DISTINCT RESERVER.id_salle FROM RESERVER " +
		"INNER JOIN RESERVATIONS ON RESERVER.id_reservation = RESERVATIONS.id_reservation " +
		"WHERE (horaire_start BETWEEN '" + *debutDateTime + "' AND '" + *endDateTime + "'" +
		" OR horaire_end BETWEEN '" + *debutDateTime + "' AND '" + *endDateTime + "'))"

	result, err := bdd.SelectDB(SALLES, []string{"id_salle", "nom", "place"}, nil, &condition)
	if err != nil {
		Log.Error("Impossible de sélectionner dans la BDD : ", err)
		return nil
	}

	if result == nil || len(result) == 0 {
		Log.Error("Impossible de sélectionner les données")
		return nil
	}

	salleStruct := printSalle(result)

	return salleStruct
}

//
// ------------------------------------------------------------------------------------------------ //
//

func menuSalle() {
	for {
		Println("-----------------------------------------------------\nBienvenue dans le Menu Salle\n-----------------------------------------------------\n")
		Println("1.Lister les salles \n2.Lister de salles disponibles \n3.Selectioner une salles avec un id \n4.cree une salle \n5.supprimer une salle \n6.Retour menu principal\nChoisissez une option :")
		_, err := fmt.Scanln(&optionSalle)
		if err != nil {
			Println("Erreur de saisie. Veuillez saisir un numéro valide.")
			continue
		}
		if optionSalle < 1 || optionSalle > 6 {
			Println("Option invalide. Veuillez choisir une option entre 1 et 5.")
			continue
		}
		break
	}
}

//
// ------------------------------------------------------------------------------------------------ //
//

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

//
// ------------------------------------------------------------------------------------------------ //
//

func printSalle(result []map[string]interface{}, noPrint ...bool) []Salle {

	if len(noPrint) == 0 || !noPrint[0] {
		Println("------------------------------")
		Println("----- SALLES DISPONIBLES -----")
	}

	var salles []Salle

	for _, salle := range result {
		if len(noPrint) == 0 || !noPrint[0] {
			println("------------------------------")
		}
		id_salle := salle["id_salle"].(int64)
		nom := salle["nom"].(string)
		place := salle["place"].(int64)

		if len(noPrint) == 0 || !noPrint[0] {
			fmt.Println("ID salle:", id_salle)
		}
		fmt.Println("Nom:", nom)
		fmt.Println("Place:", place)

		s := Salle{
			IdSalle:    id_salle,
			NomSalle:   nom,
			PlaceSalle: place,
		}
		salles = append(salles, s)
	}

	if len(noPrint) == 0 || !noPrint[0] {
		println("------------------------------")
	}

	return salles
}
