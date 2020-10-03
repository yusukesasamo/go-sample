package model

import (
	"time"
)

// Item model
type Item struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Name      string    `json:"name"`
	Price     uint      `json:"price"`
	StockFlg  uint      `json:"stock_flg"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
