package routeur

import (
	"ProjetYmmersion2/controller"
	"fmt"
	"net/http"
	"os"
)

func InitServ() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/perso", controller.Perso)
	http.HandleFunc("/modifsuppr", controller.ModifSuppr)
	http.HandleFunc("/ajout", controller.Ajout)
	http.HandleFunc("/submit", controller.SubmitHandler)

	rootDoc, _ := os.Getwd()
	FileServ := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", FileServ))

	fmt.Println("Listening at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
