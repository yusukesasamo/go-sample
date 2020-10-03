package model

import (
	"time"
)

// UserPurchaseHistory model
type UserPurchaseHistory struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	ItemID    uint      `json:"item_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
