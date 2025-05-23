package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey:ProductID" json:"id"`
	Name      Name      `gorm:"embedded" json:"name"`
	Address   Address   `gorm:"embedded" json:"address"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedBy string    `json:"updatedBy"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedBy string    `json:"deletedBy"`
	DeletedAt time.Time `json:"deletedAt"`
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
