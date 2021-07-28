package drinks

import (

	"github.com/go-playground/validator/v10"

	"github.com/ioartigiano/ioartigiano-be/internal/entity"
	log "github.com/sirupsen/logrus"
)

type Service interface {
	GetDrinks(drinks entity.Drink) ([]entity.Drink, error)

	Validate(data interface{}) error

}

type service struct {
	repo      Repository
	validator *validator.Validate
}

// NewService creates a new authentication service.
func NewService(repo Repository, validator *validator.Validate) Service {
	return service{repo, validator}
}

func (s service) GetDrinks(drinkse entity.Drink) ([]entity.Drink, error) {

	drinks, err := s.repo.GetDrinks(drinkse)
	if err != nil {
		log.Error("GetDrinks: ", err)
		return []entity.Drink{}, err
	}

	return drinks, nil
}

func (s service) Validate(data interface{}) error {
	err := s.validator.Struct(data)

	if err != nil {
		log.Error(err)
	}
	return nil
}
