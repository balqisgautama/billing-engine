package dtoin

type MakeBillingRequest struct {
	LoanID int64   `json:"loan_id" validate:"required,min=1"`
	Amount float64 `json:"amount" validate:"required,min=1"`
}
