package models

type Deck struct {
	ID     int32  `json:"id"`
	Name   string `json:"name"`
	UserID int32  `json:"user_id"`
	Cards  []Card `json:"cards"`
}
