package dbmodels

import "time"

type Loan struct {
	ID                uint    `gorm:"primaryKey"`
	UserID            uint    `json:"user_id"`
	Amount            float64 `json:"amount"`
	WeeklyPayment     float64 `json:"weekly_payment"`
	OutstandingAmount float64 `json:"outstanding_amount"`
	IsDelinquent      bool    `json:"is_delinquent"`

	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`

	Billings []Billing `gorm:"foreignKey:LoanID"`
}
