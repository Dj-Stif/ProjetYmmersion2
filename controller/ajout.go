package controller

import (
	"ProjetYmmersion2/temps"
	"net/http"
)

func Ajout(w http.ResponseWriter, r *http.Request) {
	temps.Temp.ExecuteTemplate(w, "ajout", nil)
}
