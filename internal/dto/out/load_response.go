package dtoout

type LoanResponse struct {
	ID                uint    `json:"id"`
	UserID            int64   `json:"user_id"`
	Amount            float64 `json:"amount"`
	WeeklyPayment     float64 `json:"weekly_payment"`
	OutstandingAmount float64 `json:"outstanding_amount"`
	IsDelinquent      bool    `json:"is_delinquent"`
}
