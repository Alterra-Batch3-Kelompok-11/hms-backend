package roleRepository

import (
	"gorm.io/gorm"
	"hms-backend/models"
)

type RoleRepository interface {
	GetAll() ([]models.Role, error)
	GetById(id uint) (models.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *roleRepository {
	return &roleRepository{db}
}

func (rep *roleRepository) GetAll() ([]models.Role, error) {
	var roles []models.Role

	if err := rep.db.Model([]models.Role{}).Where("ID != 1").Find(&roles).Error; err != nil {
		return roles, err
	}

	return roles, nil
}
func (rep *roleRepository) GetById(id uint) (models.Role, error) {
	role := models.Role{}

	if err := rep.db.Model(models.Role{}).Where("ID = ?", id).First(&role).Error; err != nil {
		return role, err
	}

	return role, nil
}
