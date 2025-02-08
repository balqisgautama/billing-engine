package dbmodels

import "time"

type Billing struct {
	ID      uint      `gorm:"primaryKey"`
	LoanID  uint      `json:"loan_id"`
	DueDate time.Time `json:"due_date"`
	Amount  float64   `json:"amount"`
	Week    int       `json:"week"`
	IsPaid  bool      `json:"is_paid"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
