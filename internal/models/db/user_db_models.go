package dbmodels

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `json:"username"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`

	Loans []Loan `gorm:"foreignKey:UserID"`
}
