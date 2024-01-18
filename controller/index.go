package controller

import (
	"ProjetYmmersion2/temps"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	temps.Temp.ExecuteTemplate(w, "Index", nil)

}
