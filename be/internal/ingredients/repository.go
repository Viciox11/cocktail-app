package ingredients

import (
	"github.com/ioartigiano/ioartigiano-be/internal/entity"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	GetAllIngredients() ([]entity.Ingredient, error)
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new album repository
func NewRepository(db *gorm.DB) Repository {
	return repository{db}
}

func (r repository) GetAllIngredients() ([]entity.Ingredient, error) {
	var ingredients []entity.Ingredient
	err := r.db.Find(&ingredients).Error
	return ingredients, err
}