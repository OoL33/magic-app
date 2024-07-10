package database

import (
	"context"
	"magic-app/internal/database/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDecksForUser(ctx context.Context, db *pgxpool.Pool, userID int32) ([]models.Deck, error) {
	rows, err := db.Query(ctx, "SELECT id, user_id, name FROM decks WHERE user_id=$1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var decks []models.Deck
	for rows.Next() {
		var deck models.Deck
		err := rows.Scan(&deck.ID, &deck.UserID, &deck.Name)
		if err != nil {
			return nil, err
		}
		decks = append(decks, deck)
	}
	return decks, nil
}
