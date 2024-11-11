package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	name := "Product A"
	price := 1000

	product, err := NewProduct(name, price)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, name, product.Name)
	assert.Equal(t, price, product.Price)
	assert.NotEmpty(t, product.ID)

}

func TestNewProductWhenNameIsEmpty(t *testing.T) {
	price := 1000

	product, err := NewProduct("", price)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestNewProductWhenPriceIsZero(t *testing.T) {
	name := "Product A"

	product, err := NewProduct(name, 0)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestNewProductWhenPriceIsNegative(t *testing.T) {
	name := "Product A"

	product, err := NewProduct(name, -1000)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestNewProductWhenItIsValid(t *testing.T) {
	name := "Product A"
	price := 1000

	product, _ := NewProduct(name, price)

	err := product.Validate()
	assert.Nil(t, err)
}
