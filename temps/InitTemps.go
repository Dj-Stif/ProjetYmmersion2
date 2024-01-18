package temps

import (
	"fmt"
	"html/template"
	"os"
)

var Temp *template.Template

func InitTemps() {
	temps, errTemp := template.ParseGlob("./temps/*.html")
	if errTemp != nil {
		fmt.Printf("Erreur de Templates : %v", errTemp.Error())
		os.Exit(1)
	}
	Temp = temps
}
