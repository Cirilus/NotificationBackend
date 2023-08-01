package models

import "github.com/google/uuid"

type Account struct {
	Id       uuid.UUID `json:"id,omitempty"`
	Telegram string    `yaml:"telegram,omitempty"`
	Email    string    `yaml:"email,omitempty"`
}
