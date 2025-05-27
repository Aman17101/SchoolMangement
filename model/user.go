package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primarykey" json:"id"`
	Active    bool      `json:"active" example:"true"`
	CreatedBy string    `json:"created_by" binding:"required"  example:"aman4"`
	Email     string    `json:"email" binding:"required" gorm:"unique;not null" example:"aman@1"`
	Password  string    `json:"password" binding:"required" gorm:"not null"  example:"password13"`
	CreatedAt time.Time `json:"created_at" `
	UpdatedBy string    `json:"updated_by" `
	UpdatedAt time.Time `json:"updated_at" `
	DeletedBy string    `json:"deleted_by" `
	DeletedAt  time.Time `json:"deleted_at" `
}

type Name struct {
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
}

type Address struct {
	Lane     string `json:"lane"`
	Village  string `json:"village"`
	City     string `json:"city"`
	District string `json:"district"`
	Pincode  int    `json:"pincode"`
	State    string `json:"state"`
}
