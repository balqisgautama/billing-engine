package dtoout

type PaymentResponse struct {
	ID         uint    `json:"id"`
	LoanID     uint    `json:"loan_id"`
	Amount     float64 `json:"amount"`
	WeekNumber int     `json:"week_number"`
	CreatedAt  string  `json:"created_at"` // Use string for formatted date
}
