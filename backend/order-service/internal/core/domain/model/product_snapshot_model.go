package model

import "time"

type ProductSnapshot struct {
	ID           int64      `json:"id"`
	Name         string     `json:"name"`
	Stock        int        `json:"stock"`
	Image        string     `json:"image"`
	RegulerPrice int64      `json:"reguler_price"`
	SalePrice    int64      `json:"sale_price"`
	Unit         string     `gorm:"column:unit;default:'gram'"`
	Weight       int        `gorm:"column:weight;default:0"`
	CreatedAt    time.Time  `json:"created_at"`
	LastUsed     *time.Time `json:"last_used"`
}
