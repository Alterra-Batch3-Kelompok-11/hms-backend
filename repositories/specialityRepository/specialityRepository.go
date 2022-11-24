package specialityRepository

import (
	"gorm.io/gorm"
	"hms-backend/models"
)

type SpecialityRepository interface {
	GetAll() ([]models.Speciality, error)
	GetById(id uint) (models.Speciality, error)
	Create(payload models.Speciality) (models.Speciality, error)
	Update(id uint, payload models.Speciality) (models.Speciality, error)
	Delete(id uint) error
}

type specialityRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *specialityRepository {
	return &specialityRepository{db}
}

func (rep *specialityRepository) GetAll() ([]models.Speciality, error) {
	var specialities []models.Speciality

	if err := rep.db.Model([]models.Speciality{}).Find(&specialities).Error; err != nil {
		return specialities, err
	}

	return specialities, nil
}
func (rep *specialityRepository) GetById(id uint) (models.Speciality, error) {
	speciality := models.Speciality{}

	if err := rep.db.Model(models.Speciality{}).Where("ID = ?", id).First(&speciality).Error; err != nil {
		return speciality, err
	}

	return speciality, nil
}
func (rep *specialityRepository) Create(payload models.Speciality) (models.Speciality, error) {
	err := rep.db.Create(&payload).Error
	return payload, err
}
func (rep *specialityRepository) Update(id uint, payload models.Speciality) (models.Speciality, error) {
	err := rep.db.Model(models.Speciality{}).Where("ID = ?", id).Updates(&payload).Error
	return payload, err
}
func (rep *specialityRepository) Delete(id uint) error {
	var speciality models.Speciality
	return rep.db.Delete(&speciality, id).Error
}
