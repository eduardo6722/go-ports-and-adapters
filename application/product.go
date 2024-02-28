package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
	Status string  `valid:"required"`
}

type IProduct interface {
	GetID() string
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetName() string
	GetStatus() string
	GetPrice() float64
}

type ProductReader interface {
	Get(id string) (IProduct, error)
}

type ProductWriter interface {
	Save(product IProduct) (IProduct, error)
}

type IProductService interface {
	Get(id string) (IProduct, error)
	Create(name string, price float64) (IProduct, error)
	Enable(product IProduct) (IProduct, error)
	Disable(product IProduct) (IProduct, error)
}

type IProductPersistence interface {
	ProductReader
	ProductWriter
}

type Foo struct {
	Bar string
}

func NewProduct() *Product {
	return &Product{
		ID:     uuid.NewV4().String(),
		Status: DISABLED,
	}
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
		return false, nil
	}
	if p.Status != DISABLED && p.Status != ENABLED {
		return false, errors.New("the status must be either disabled or enabled")
	}
	if p.Price < 0 {
		return false, errors.New("the price must be greater than or equal to zero")
	}
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("the price must be greater than zero")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}
	return errors.New("the price must be zero")
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
