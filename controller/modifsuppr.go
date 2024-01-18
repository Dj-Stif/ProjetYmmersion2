package controller

import (
	"ProjetYmmersion2/temps"
	"net/http"
)

func ModifSuppr(w http.ResponseWriter, r *http.Request) {
	temps.Temp.ExecuteTemplate(w, "modifsuppr", nil)
}
