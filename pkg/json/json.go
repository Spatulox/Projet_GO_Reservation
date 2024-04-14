package json

import (
	. "Projet_GO_Reservation/pkg/const"
	. "Projet_GO_Reservation/pkg/log"
	"encoding/json"
	"os"
)

func Json(data []map[string]interface{}) bool {
	/*data := []map[string]interface{}{
		{"name": "John Doe", "age": 30, "email": "john.doe@example.com"},
		{"name": "Jane Smith", "age": 25, "email": "jane.smith@example.com"},
	}*/

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
