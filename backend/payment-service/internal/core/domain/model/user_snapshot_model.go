package model

import "time"

type UserSnapshot struct {
	ID        int64      `gorm:"primaryKey" json:"id"`
	UserID    int64      `gorm:"not null;uniqueIndex" json:"user_id"`
	Name      string     `gorm:"type:varchar(100);not null" json:"name"`
	Email     string     `gorm:"type:varchar(100);not null" json:"email"`
	Phone     string     `gorm:"type:varchar(20);not null" json:"phone"`
	Address   string     `gorm:"type:text;not null" json:"address"`
	Photo     string     `gorm:"type:text" json:"photo"`
	RoleID    string     `gorm:"type:varchar(50);not null" json:"role_id"`
	RoleName  string     `gorm:"type:varchar(50);not null" json:"role_name"`
	Lat       string     `gorm:"type:varchar(20)" json:"lat"`
	Lng       string     `gorm:"type:varchar(20)" json:"lng"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	LastUsed  *time.Time `json:"last_used"`
}
