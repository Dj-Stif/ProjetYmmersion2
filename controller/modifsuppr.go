package controller

import (
	"ProjetYmmersion2/data"
	"ProjetYmmersion2/temps"
	"fmt"
	"net/http"
	"strconv"
)

var Id int

func Modif(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Println("erreur modif")
	}
	for _, t := range ListAventurier {
		if t.Id == Id {
			data.Player = t
			break
		}
	}
	Id = id

	temps.Temp.ExecuteTemplate(w, "modif", data.Player)
}

func Suppr(w http.ResponseWriter, r *http.Request) {
	ListAventurier, err = ReadJson()
	if err != nil {
		fmt.Println("erreur suppr", err)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Println("erreur modif")
	}
	for i, c := range ListAventurier {
		if c.Id == id {
			ListAventurier = append(ListAventurier[:i], ListAventurier[i+1:]...)
		}
	}
	EditJson(ListAventurier)

	http.Redirect(w, r, "/perso", http.StatusMovedPermanently)
}
