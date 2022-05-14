package dto

type AreaDTO struct {
	From CoordinateDTO `json:"from"`
	To   CoordinateDTO `json:"to"`
}
