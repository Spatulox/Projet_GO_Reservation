package json

import (
	. "Projet_GO_Reservation/pkg/const"
	. "Projet_GO_Reservation/pkg/log"
	. "Projet_GO_Reservation/pkg/models"
	"encoding/json"
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
func JsonToData(path *string) []map[string]interface{} {

	var err error
	var jsonData []byte

	if path == nil {
		Log.Error("Vous devez spécifier un chemin pour le fichier à importer")
		Log.Debug("En cours de création")
		return nil
	} else {
		jsonData, err = os.ReadFile(*path)
	}

	if err != nil {
		Log.Error("Erreur lors de la lecture du fichier:", err)
		return nil
	}

	// Transformer le JSON en slice de maps
	var data []map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		Log.Error("Erreur lors de la conversion en Go:", err)
		return nil
	}

	return data
}
