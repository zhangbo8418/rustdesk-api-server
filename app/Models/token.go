package Models

import (
	"time"
)

type Token struct {
	ID             uint64 `json:"id" gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	UserID         uint64 `json:"user_id"`
	User           User
	ClientID       string `json:"client_id" gorm:"index;type:varchar(16)"`
	AccessToken    string `json:"access_token" gorm:"uniqueIndex"`
	LoginTokenType string `json:"login_token_type" gorm:"index;type:varchar(50)"`
}
