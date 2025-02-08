// services/Loan_service.go
package services

import (
	dtoin "billing-engine/internal/dto/in"
	dtoout "billing-engine/internal/dto/out"
	dbmodels "billing-engine/internal/models/db"
	"billing-engine/internal/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// LoanService struct
type LoanService struct {
	loanRepo *repositories.LoanRepository
	userRepo *repositories.UserRepository
}

// NewLoanService creates a new instance of LoanService
func NewLoanService(
	loanRepo *repositories.LoanRepository,
	userRepo *repositories.UserRepository,
) *LoanService {
	return &LoanService{
		loanRepo: loanRepo,
		userRepo: userRepo,
	}
}

// CreateLoan handles the creation of a new loan
func (service *LoanService) CreateLoan(c *gin.Context) {
	var request dtoin.CreateLoanRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the request
	if err := dtoin.Validate(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := service.userRepo.GetUserByID(uint(request.UserID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loans, err := service.loanRepo.GetLoansByUserIDIsDelinquent(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(loans) > 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please pay your bills first."})
		return
	}

	loan := dbmodels.Loan{
		UserID:            user.ID,
		Amount:            request.Amount,
		WeeklyPayment:     request.Amount * 1.1 / 50, // 10% interest
		OutstandingAmount: request.Amount * 1.1,
	}

	if err := service.loanRepo.CreateLoan(&loan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create loan"})
		return
	}

	response := dtoout.LoanResponse{
		ID:                loan.ID,
		UserID:            request.UserID,
		Amount:            loan.Amount,
		WeeklyPayment:     loan.WeeklyPayment,
		OutstandingAmount: loan.OutstandingAmount,
		IsDelinquent:      loan.IsDelinquent,
	}

	c.JSON(http.StatusCreated, response)
}

// MakePayment handles the payment for a loan
// func (service *LoanService) MakePayment(c *gin.Context) {
// 	var request dtoin.MakePaymentRequest
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	loanIDStr := c.Param("id")
// 	loanID, err := strconv.ParseUint(loanIDStr, 10, 32)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
// 		return
// 	}
// 	loan, err := service.loanRepo.GetLoanByID(uint(loanID))
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
// 		return
// 	}

// 	temp := loan.OutstandingAmount - request.Amount
// 	if temp < 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Payment amount exceeds outstanding amount"})
// 		return
// 	}

// 	if request.Amount != loan.WeeklyPayment || request.Amount <= 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment amount"})
// 		return
// 	}

// 	loan.OutstandingAmount -= request.Amount
// 	loan.UpdatedAt = time.Now()
// 	if err := service.loanRepo.UpdateLoan(loan); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update loan"})
// 		return
// 	}

// 	weekNumber, _ := time.Now().ISOWeek() // Get the current week number
// 	payment := dbmodels.Billing{
// 		LoanID:    loan.ID,
// 		Amount:    request.Amount,
// 		Week:      weekNumber,
// 		CreatedAt: time.Now(),
// 	}
// 	service.loanRepo.DB.Create(&payment)

// 	response := dtoout.LoanResponse{
// 		ID:                loan.ID,
// 		Amount:            loan.Amount,
// 		WeeklyPayment:     loan.WeeklyPayment,
// 		OutstandingAmount: loan.OutstandingAmount,
// 		IsDelinquent:      loan.IsDelinquent,
// 	}

// 	c.JSON(http.StatusOK, response)
// }

// GetOutstanding retrieves the outstanding amount for a loan
func (service *LoanService) GetOutstanding(c *gin.Context) {
	loanIDStr := c.Param("id")
	loanID, err := strconv.ParseUint(loanIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
		return
	}
	loan, err := service.loanRepo.GetLoanByID(uint(loanID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"outstanding_amount": loan.OutstandingAmount})
}

// IsDelinquent checks if a loan is delinquent
func (service *LoanService) IsDelinquent(c *gin.Context) {
	loanIDStr := c.Param("id")
	loanID, err := strconv.ParseUint(loanIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
		return
	}
	loan, err := service.loanRepo.GetLoanByID(uint(loanID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	missedPayments := 0
	for _, payment := range loan.Billings {
		if payment.Amount == 0 {
			missedPayments++
		} else {
			missedPayments = 0
		}
		if missedPayments >= 2 {
			loan.IsDelinquent = true
			break
		}
	}
	service.loanRepo.UpdateLoan(loan)

	c.JSON(http.StatusOK, gin.H{"is_delinquent": loan.IsDelinquent})
}
