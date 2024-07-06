package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float32
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	Id     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float32 `valid: "float"`
	Status string  `valid:"required"`
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("status undefined")
	}

	if p.Price < 0 {
		return false, errors.New("the price must be greater or equal to zero")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}

	return errors.New("the price must be greater than 0")

}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}

	return errors.New("the price must be equal to 0")
}

func (p *Product) GetId() string {
	return p.Id
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}
func (p *Product) GetPrice() float32 {
	return p.Price
}
