package controller

import (
	"ProjetYmmersion2/data"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var Aventurier []data.Aventurier

func GetDataFromJson() {
	data, err := os.ReadFile("data/data.json") //ouverture et lecture du json
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	json.Unmarshal(data, &Aventurier) //passage en json vers la struct
}

//fonction qui récupere le json

func SubmitHandler(w http.ResponseWriter, r *http.Request) {

	ageForm, ageErr := strconv.Atoi(r.FormValue("age"))
	if ageErr != nil {
		fmt.Println("erreur")
	}

	var handForm bool

	if r.FormValue("hand") == "right" {
		handForm = true
	}

	dataForm := data.Aventurier{r.FormValue("lastname"), r.FormValue("firstname"), ageForm, handForm}

	dataFile, errFile := os.ReadFile("./data/fichier.json")
	if errFile != nil {
		fmt.Fprintf(w, "erreur lecture fichier fichier.json")
		return
	}

	var dataAventurier []data.Aventurier

	errDecode := json.Unmarshal(dataFile, &dataAventurier)
	if errDecode != nil {
		fmt.Fprintf(w, "erreur unmarshal fichier.json")
		return
	}

	dataAventurier = append(dataAventurier, dataForm)

	dataEncode, errEncode := json.Marshal(dataAventurier)
	if errEncode != nil {
		fmt.Fprintf(w, "erreur marshal fichier.json")
		return
	}

	errWrite := os.WriteFile("./data/fichier.json", dataEncode, 0644)
	if errWrite != nil {
		fmt.Fprintf(w, "erreur ecriture fichier.json")
		return
	}

	http.Redirect(w, r, "/ajout", http.StatusMovedPermanently)

}

func ReadJson() ([]data.Aventurier, error) {
	jsonFile, err := os.ReadFile("Chemin vers le JSON")
	if err != nil {
		fmt.Println("Error reading", err.Error())
	}

	var jsonData []data.Aventurier
	err = json.Unmarshal(jsonFile, &jsonData)
	return jsonData, err
}

func EditJson(ModifiedArticle []data.Aventurier) {

	modifiedJson, errMarshal := json.Marshal(ModifiedArticle)
	if errMarshal != nil {
		fmt.Println("Error encodage", errMarshal.Error())
		return
	}

	err := os.WriteFile("Chemin vers le JSON", modifiedJson, 0644)
	if err != nil {
		fmt.Println("Erreur lors de la modification du fichier fichier.json", err)
		return
	}
}
