package model

import "time"

type User struct {
	BaseModel
	Name        string    `gorm:"type:varchar(255);not null;index:idx_name"`
	Password    string    `gorm:"type:varchar(64);not null"`
	Email       string    `gorm:"type:varchar(255);not null;default:'';uniqueIndex:uniq_idx_email"`
	Phone       string    `gorm:"type:varchar(20);not null;default:''"`
	Avatar      string    `gorm:"type:varchar(255);not null;default:''"`
	LastLoginAt time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
}
