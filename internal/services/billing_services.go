package services

import (
	dtoin "billing-engine/internal/dto/in"
	dbmodels "billing-engine/internal/models/db"
	"billing-engine/internal/repositories"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// BillingService struct
type BillingService struct {
	loanRepo    *repositories.LoanRepository
	billingRepo *repositories.BillingRepository
}

// NewBillingService creates a new instance of BillingService
func NewBillingService(
	loanRepo *repositories.LoanRepository,
	billingRepo *repositories.BillingRepository,
) *BillingService {
	return &BillingService{
		loanRepo:    loanRepo,
		billingRepo: billingRepo,
	}
}

// CreateLoan handles the creation of a new loan
func (service *BillingService) CreateBilling(c *gin.Context) {
	var billing dtoin.MakeBillingRequest
	if err := c.ShouldBindJSON(&billing); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the request
	if err := dtoin.Validate(&billing); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	billingModel := &dbmodels.Billing{
		LoanID:    uint(billing.LoanID),
		Amount:    billing.Amount,
		DueDate:   time.Now().AddDate(0, 0, 30),
		Week:      1,
		CreatedAt: time.Now(),
		IsPaid:    false,
	}

	// Create the loan
	err := service.billingRepo.CreateBilling(billingModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, billing)
}
