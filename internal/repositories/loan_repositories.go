package repositories

import (
	dbmodels "billing-engine/internal/models/db"

	"gorm.io/gorm"
)

type LoanRepository struct {
	DB *gorm.DB
}

func NewLoanRepository(db *gorm.DB) *LoanRepository {
	return &LoanRepository{
		DB: db,
	}
}

func (repo *LoanRepository) CreateLoan(loan *dbmodels.Loan) error {
	return repo.DB.Create(loan).Error
}

func (repo *LoanRepository) GetLoanByID(id uint) (*dbmodels.Loan, error) {
	var loan dbmodels.Loan
	err := repo.DB.First(&loan, id).Error
	return &loan, err
}

func (repo *LoanRepository) UpdateLoan(loan *dbmodels.Loan) error {
	return repo.DB.Save(loan).Error
}

func (repo *LoanRepository) GetAllLoans() ([]dbmodels.Loan, error) {
	var loans []dbmodels.Loan
	err := repo.DB.Find(&loans).Error
	return loans, err
}

func (repo *LoanRepository) GetLoansByUserID(userID uint) ([]dbmodels.Loan, error) {
	var loans []dbmodels.Loan
	err := repo.DB.Where("user_id = ?", userID).Find(&loans).Error
	return loans, err
}

func (repo *LoanRepository) GetLoansByUserIDIsDelinquent(userID uint) ([]dbmodels.Loan, error) {
	var loans []dbmodels.Loan
	err := repo.DB.Where("user_id = ? AND is_delinquent = ?", userID, true).Find(&loans).Error
	return loans, err
}
