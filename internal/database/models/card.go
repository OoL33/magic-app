package models

import (
	"github.com/jackc/pgtype"
)

type Card struct {
	ID        int32            `json:"id"`
	Name      string           `json:"name"`
	Colors    pgtype.TextArray `json:"colors"`
	ImageUrls string           `json:"image_urls"`
	DeckID    int32            `json:"deck_id"`
	UserID    int32            `json:"user_id"`
}
