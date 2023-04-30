package entity_test

import (
	"testing"

	"github.com/gabrielborel/microservice-go/internal/order/entity"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId_WhenCreatingANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order, err := entity.NewOrder("", 100.0, 10.0)
	assert.Nil(t, order)
	if assert.Error(t, err) {
		assert.Equal(t, "invalid id", err.Error())
	}
}

func TestGivenAnEmptyPrice_WhenCreatingANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order, err := entity.NewOrder("123", 0, 10.0)
	assert.Nil(t, order)
	if assert.Error(t, err) {
		assert.Equal(t, "invalid price", err.Error())
	}
}

func TestGivenAnEmptyTax_WhenCreatingANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order, err := entity.NewOrder("123", 100.0, 0)
	assert.Nil(t, order)
	if assert.Error(t, err) {
		assert.Equal(t, "invalid tax", err.Error())
	}
}

func TestGivenValidParams_WhenCreatingANewOrder_ThenShouldReceiverANewOrderWithAllParams(t *testing.T) {
	order, err := entity.NewOrder("123", 100.0, 10.0)
	assert.NoError(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 100.0, order.Price)
	assert.Equal(t, 10.0, order.Tax)
}

func TestGivenValidParams_WhenCalculatingFinalPrice_ThenShouldCalculateFinalPriceAndSetItOnFinalPriceProperty(t *testing.T) {
	order, err := entity.NewOrder("123", 100.0, 10.0)
	assert.NoError(t, err)
	err = order.CalculateFinalPrice()
	assert.NoError(t, err)
	assert.Equal(t, 110.0, order.FinalPrice)
}
