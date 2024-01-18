package controller

import (
	"ProjetYmmersion2/data"
	"ProjetYmmersion2/temps"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var ListAventurier []data.Aventurier

func Perso(w http.ResponseWriter, r *http.Request) {
	dataFile, errFile := os.ReadFile("./data/fichier.json")
	if errFile != nil {
		fmt.Fprintf(w, "erreur lecture fichier fichier.json")
		return
	}

	errDecode := json.Unmarshal(dataFile, &ListAventurier)
	if errDecode != nil {
		fmt.Fprintf(w, "erreur unmarshal fichier.json")
		return
	}

	temps.Temp.ExecuteTemplate(w, "perso", ListAventurier)
}
