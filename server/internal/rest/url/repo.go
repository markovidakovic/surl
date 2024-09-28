package url

import (
	"github.com/markovidakovic/surl/server/internal/db/models"
	"gorm.io/gorm"
)

type Repository interface {
	Find() ([]models.Url, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

func (r *repository) Find() ([]models.Url, error) {
	var urls []models.Url
	result := r.db.Find(&urls)
	if result.Error != nil {
		return nil, result.Error
	}
	return urls, nil
}
