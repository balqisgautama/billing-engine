package dtoin

type CreateLoanRequest struct {
	UserID int64   `json:"user_id" validate:"required,min=1"`
	Amount float64 `json:"amount" validate:"required,min=1"`
}
