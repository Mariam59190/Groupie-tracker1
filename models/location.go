package models

// Location représente les lieux de concerts d'un artiste
type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}
