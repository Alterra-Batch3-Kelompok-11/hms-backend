package religionRepository

import (
	"gorm.io/gorm"
	"hms-backend/models"
)

type ReligionRepository interface {
	GetAll() ([]models.Religion, error)
	GetById(id uint) (models.Religion, error)
}

type religionRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *religionRepository {
	return &religionRepository{db}
}

func (rep *religionRepository) GetAll() ([]models.Religion, error) {
	var religions []models.Religion

	if err := rep.db.Model([]models.Religion{}).Where("ID != 1").Find(&religions).Error; err != nil {
		return religions, err
	}

	return religions, nil
}
func (rep *religionRepository) GetById(id uint) (models.Religion, error) {
	religion := models.Religion{}

	if err := rep.db.Model(models.Religion{}).Where("ID = ?", id).First(&religion).Error; err != nil {
		return religion, err
	}

	return religion, nil
}
