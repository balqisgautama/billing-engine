package repositories

import (
	dbmodels "billing-engine/internal/models/db"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (repo *UserRepository) CreateUser(User *dbmodels.User) error {
	return repo.DB.Create(User).Error
}

func (repo *UserRepository) GetUserByID(id uint) (*dbmodels.User, error) {
	var User dbmodels.User
	err := repo.DB.First(&User, id).Error
	return &User, err
}

func (repo *UserRepository) GetUsersByLoadID(id uint) ([]dbmodels.User, error) {
	var Users []dbmodels.User
	err := repo.DB.Where("loan_id = ?", id).Find(&Users).Error
	return Users, err
}

func (repo *UserRepository) UpdateUser(User *dbmodels.User) error {
	return repo.DB.Save(User).Error
}

func (repo *UserRepository) GetAllUsers() ([]dbmodels.User, error) {
	var Users []dbmodels.User
	err := repo.DB.Find(&Users).Error
	return Users, err
}
