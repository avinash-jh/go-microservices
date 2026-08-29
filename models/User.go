package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User defines the model and schema specifications for GORM.
type User struct {
	ID          string `json:"id" gorm:"type:char(36);primaryKey"`
	UserName    string `json:"user_name" gorm:"type:varchar(100);not null;unique"`
	Email       string `json:"email" gorm:"type:varchar(255);not null;unique"`
	PhoneNumber string `json:"phone_number" gorm:"type:varchar(20)"`
}

// TableName tells GORM the exact name of the table to create inside database 'db1'.
func (User) TableName() string {
	return "userdb"
}

// BeforeCreate is a GORM hook that generates a new UUID before inserting into MySQL
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return nil
}
