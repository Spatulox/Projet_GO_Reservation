package json

import (
	. "Projet_GO_Reservation/pkg/const"
	. "Projet_GO_Reservation/pkg/log"
	. "Projet_GO_Reservation/pkg/models"
	"encoding/json"
	"fmt"
	"os"
)

// DataToJson Get the data with a []Reservation Structure (may change to export rooms)
// Export it into a human-readable json file called data.json into the root folder of the project
func DataToJson(data []Reservation) bool {

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		Log.Error("Erreur lors de la conversion en JSON:", err)
		return false
	}

	Println(string(jsonData))

	err = os.WriteFile("data.json", jsonData, 0644)
	if err != nil {
		Log.Error("Erreur lors de l'écriture du fichier:", err)
		return false
	}

	Log.Infos("Fichier JSON enregistré avec succès.")

	return true
}

// JsonToData Get the data path
// Transform the json file into a []map[string]interface{}
func JsonToData(data []map[string]interface{}) bool {

	for _, d := range data {
		fmt.Println(d)
	}

	// For all reservation inside the file

	// Need to check if the room still exist
	// If still exist, is it the same ? (Number of Place for the most part)
	// Then : Need to check if the room is available at the date/hour of the reservation

	// Need to check if the reservation already exist (ID)
	// Yes => Check if it's the same for all the point (No upload so)
	//		The same =>		No upload
	//		Not the same =>	Create a new reservation (ID useless)
	// No => Upload of the reservation

	// End of For
	return true
}
