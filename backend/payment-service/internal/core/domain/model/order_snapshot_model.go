package model

import "time"

type OrderSnapshot struct {
	ID              int64      `gorm:"primaryKey" json:"id"`
	OrderID         int64      `gorm:"not null;uniqueIndex" json:"order_id"`
	OrderCode       string     `gorm:"type:varchar(100);not null" json:"order_code"`
	OrderDatetime   string     `gorm:"type:varchar(50);not null" json:"order_datetime"`
	Status          string     `gorm:"type:varchar(50);not null" json:"order_status"`
	PaymentMethod   string     `gorm:"type:varchar(50);not null" json:"payment_method"`
	ShippingFee     int64      `gorm:"not null" json:"shipping_fee"`
	ShippingType    string     `gorm:"type:varchar(50);not null" json:"shipping_type"`
	Remarks         string     `gorm:"type:text" json:"remarks"`
	TotalAmount     int64      `gorm:"not null" json:"total_amount"`
	CustomerName    string     `gorm:"type:varchar(100);not null" json:"customer_name"`
	CustomerPhone   string     `gorm:"type:varchar(20);not null" json:"customer_phone"`
	CustomerAddress string     `gorm:"type:text;not null" json:"customer_address"`
	CustomerEmail   string     `gorm:"type:varchar(100);not null" json:"customer_email"`
	CustomerID      int64      `gorm:"not null" json:"customer_id"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	LastUsed        *time.Time `json:"last_used"`
}
