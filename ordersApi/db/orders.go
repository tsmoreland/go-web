package db

import (
	"fmt"
	"github.com/tsmoreland/go-web/ordersApi/models"
)

type OrderDB struct {
	placedOrders map[string]models.Order
}

// NewOrders creates a new empty order service
func NewOrders() *OrderDB {
	return &OrderDB{
		placedOrders: make(map[string]models.Order),
	}
}

// FindById order for a given id, if one exists
func (o *OrderDB) FindById(id string) (models.Order, error) {
	order, ok := o.placedOrders[id]
	if !ok {
		return models.Order{}, fmt.Errorf("no order found for %s order id", id)
	}

	return order, nil
}

// Upsert creates or updates an order in the orders DB
func (o *OrderDB) Upsert(order models.Order) {
	o.placedOrders[order.ID] = order
}
