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
		fmt.Println("erreur")
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
	http.Redirect(w, r, "/perso", http.StatusMovedPermanently)
}
