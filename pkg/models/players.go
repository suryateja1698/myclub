package models

type Player struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Country   string `json:"country"`
	Position  string `json:"position"`
	Age       int    `json:"age"`
}
