package drinks

import (
	"github.com/ioartigiano/ioartigiano-be/internal/entity"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	GetDrinks(drinks entity.Drink) ([]entity.Drink, error)

}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new album repository
func NewRepository(db *gorm.DB) Repository {
	return repository{db}
}

func (r repository) GetDrinks(drinks entity.Drink) ([]entity.Drink, error) {
	var product []entity.Drink
	err := r.db.Preload("ProductAssets").Where("id = ?", "id").First(&drinks).Error
	return product, err
}