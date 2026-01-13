package models

// Relation fait le lien entre artistes, dates et lieux
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
