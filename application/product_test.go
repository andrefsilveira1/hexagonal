package application_test

import (
	"testing"

	"github.com/andrefsilveira1/hexagonal/application"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "André"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

}

func TestProduct_Enable_With_Error(t *testing.T) {
	product := application.Product{}
	product.Name = "André"
	product.Status = application.DISABLED
	product.Price = 0
	err := product.Enable()
	require.Equal(t, "the price must be greater than 0", err.Error())

}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "André"
	product.Status = application.DISABLED
	product.Price = 0
	err := product.Enable()
	require.Equal(t, "the price must be greater than 0", err.Error())

	err = product.Disable()
	require.Nil(t, err)
}

func TestProduct_Disable_With_Error(t *testing.T) {
	product := application.Product{}
	product.Name = "André"
	product.Status = application.DISABLED
	product.Price = 10
	err := product.Disable()
	require.Equal(t, "the price must be equal to 0", err.Error())
}

func Test_Is_Valid(t *testing.T) {
	product := application.Product{
		Id:     uuid.NewString(),
		Name:   "André",
		Status: application.DISABLED,
		Price:  10.0,
	}

	_, err := product.IsValid()
	require.Nil(t, err)

}

func Test_Is_Valid_With_Error(t *testing.T) {
	product := application.Product{
		Id:     uuid.NewString(),
		Name:   "André",
		Status: "INVALID",
		Price:  10.0,
	}

	_, err := product.IsValid()
	require.Equal(t, "status undefined", err.Error())

}
