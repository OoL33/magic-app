package database

import (
	"context"
	"magic-app/internal/database/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetCardsForUser(ctx context.Context, db *pgxpool.Pool, userID int32) ([]models.Card, error) {
	query := `
	SELECT c.id, c.name, c.colors, c.image_urls
	FROM cards c
	JOIN user_cards uc ON c.id = uc.card_id
	WHERE uc.user_id = $1`

	rows, err := db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []models.Card
	for rows.Next() {
		var card models.Card
		err := rows.Scan(&card.ID, &card.Name, &card.Colors, &card.ImageUrls)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}
