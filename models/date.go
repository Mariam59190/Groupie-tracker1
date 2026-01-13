package models

// Date représente les dates de concerts d'un artiste
type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}
