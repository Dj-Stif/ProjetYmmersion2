package data

type Aventurier struct {
	LastName  string `json:"lastname"`
	FirstName string `json:"firstname"`
	Age       int    `json:"age"`
	Hand      bool   `json:"hand"`
	Id        int    `json:"id"`
}

var Player Aventurier
