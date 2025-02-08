package repositories

import (
	dbmodels "billing-engine/internal/models/db"

	"gorm.io/gorm"
)

type BillingRepository struct {
	DB *gorm.DB
}

func NewBillingRepository(db *gorm.DB) *BillingRepository {
	return &BillingRepository{
		DB: db,
	}
}

func (repo *BillingRepository) CreateBilling(Billing *dbmodels.Billing) error {
	return repo.DB.Create(Billing).Error
}

func (repo *BillingRepository) GetBillingByID(id uint) (*dbmodels.Billing, error) {
	var Billing dbmodels.Billing
	err := repo.DB.First(&Billing, id).Error
	return &Billing, err
}

func (repo *BillingRepository) GetBillingsByLoadID(id uint) ([]dbmodels.Billing, error) {
	var Billings []dbmodels.Billing
	err := repo.DB.Where("loan_id = ?", id).Find(&Billings).Error
	return Billings, err
}

func (repo *BillingRepository) UpdateBilling(Billing *dbmodels.Billing) error {
	return repo.DB.Save(Billing).Error
}

func (repo *BillingRepository) GetAllBillings() ([]dbmodels.Billing, error) {
	var Billings []dbmodels.Billing
	err := repo.DB.Find(&Billings).Error
	return Billings, err
}
