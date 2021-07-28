package ingredients

import (
	log "github.com/sirupsen/logrus"

	"github.com/go-playground/validator/v10"
	"github.com/ioartigiano/ioartigiano-be/internal/entity"
)

type Service interface {
	GetIngredients() ([]entity.Ingredient, error)

	Validate(data interface{}) error
}

type service struct {
	repo        Repository
	validator   *validator.Validate
}

// NewService creates a new authentication service.
func NewService(repo Repository, validator *validator.Validate) Service {
	return service{repo, validator}
}

func (s service) GetIngredients() ([]entity.Ingredient, error) {

	ingredients, err := s.repo.GetAllIngredients()
	if err != nil {
		log.Error("GetIngredients: ", err)
		return []entity.Ingredient{}, err
	}

	return ingredients, nil
}

func (s service) Validate(data interface{}) error {
	err := s.validator.Struct(data)

	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
