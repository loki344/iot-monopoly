package dto

type CardDTO struct {
	Title     string `json:"title"`
	Text      string `json:"text"`
	Confirmed bool   `json:"confirmed"`
}
