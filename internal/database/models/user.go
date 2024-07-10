package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"unique"`
	Password string `json:"password"`
	Decks    []Deck `json:"decks"`
	Cards    []Card `json:"cards"`
}

func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
