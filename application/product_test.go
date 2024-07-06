package application_test

import (
	"testing"

	"github.com/andrefsilveira1/hexagonal/application"
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
