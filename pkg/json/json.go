package json

import (
	. "Projet_GO_Reservation/pkg/const"
	. "Projet_GO_Reservation/pkg/log"
	"encoding/json"
	"os"
)

func DataToJson(data []map[string]interface{}) bool {

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
	// Lire le fichier JSON
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

	// Afficher les données
	return data
}
