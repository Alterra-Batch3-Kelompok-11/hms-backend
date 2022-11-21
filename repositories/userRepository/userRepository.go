package userRepository

import (
	"gorm.io/gorm"
	"hms-backend/models"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetById(id uint) (models.User, error)
	GetByUsernamePassword(username, password string) (models.User, error)
	GetByUsername(username string) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(id uint, user models.User) (models.User, error)
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (rep *userRepository) GetAll() ([]models.User, error) {
	var users []models.User

	if err := rep.db.Model([]models.User{}).Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
func (rep *userRepository) GetById(id uint) (models.User, error) {
	user := models.User{}

	if err := rep.db.Model(models.User{}).Where("ID = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
func (rep *userRepository) GetByUsernamePassword(username, password string) (models.User, error) {
	user := models.User{}

	if err := rep.db.Model(models.User{}).Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
func (rep *userRepository) GetByUsername(username string) (models.User, error) {
	user := models.User{}

	if err := rep.db.Model(models.User{}).Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
func (rep *userRepository) Create(user models.User) (models.User, error) {
	err := rep.db.Create(&user).Error
	return user, err
}
func (rep *userRepository) Update(id uint, user models.User) (models.User, error) {
	err := rep.db.Model(models.User{}).Where("ID = ?", id).Updates(&user).Error
	return user, err
}
func (rep *userRepository) Delete(id uint) error {
	var user models.User
	return rep.db.Delete(&user, id).Error
}
