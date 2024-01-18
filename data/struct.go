package data

type Aventurier struct {
	LastName  string `json:"lastname"`
	FirstName string `json:"firstname"`
	Age       int    `json:"age"`
	Hand      bool   `json:"hand"`
}

var Player Aventurier
