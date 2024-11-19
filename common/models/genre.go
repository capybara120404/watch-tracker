package models

type Genre struct {
	ID    uint   `json:"id"`
	Title string `json:"title,omitempty"`
}
