package application

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProductEnable(t *testing.T) {
	product := NewProduct()
	product.Name = "test"
	product.Price = 10
	product.Status = DISABLED

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0

	err = product.Enable()
	require.Equal(t, "the price must be greater than zero", err.Error())
}

func TestProductDisable(t *testing.T) {
	product := NewProduct()
	product.Name = "test"
	product.Price = 10
	product.Status = ENABLED

	err := product.Disable()
	require.Equal(t, "the price must be zero", err.Error())

	product.Price = 0

	err = product.Disable()
	require.Nil(t, err)
}

func TestIsValid(t *testing.T) {
	product := NewProduct()
	product.Name = "test"
	product.Price = 10
	product.Status = DISABLED

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "invalid"
	_, err = product.IsValid()
	require.Equal(t, "the status must be either disabled or enabled", err.Error())

	product.Status = ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -1
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater than or equal to zero", err.Error())
}

func TestGetName(t *testing.T) {
	product := NewProduct()
	product.Name = "test"
	product.Price = 10
	require.Equal(t, "test", product.GetName())
}

func TestGetStatus(t *testing.T) {
	product := NewProduct()
	product.Name = "test"
	require.Equal(t, DISABLED, product.GetStatus())
}

func TestGetPrice(t *testing.T) {
	product := NewProduct()
	product.Price = 10
	require.Equal(t, float64(10), product.GetPrice())
}
