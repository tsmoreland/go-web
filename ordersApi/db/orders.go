package db

import (
	"fmt"
	"github.com/tsmoreland/go-web/ordersApi/models"
	"sync"
)

type OrderDB struct {
	placedOrders sync.Map
}

// NewOrders creates a new empty order service
func NewOrders() *OrderDB {
	return &OrderDB{}
}

// FindById order for a given id, if one exists
func (o *OrderDB) FindById(id string) (models.Order, error) {
	order, ok := o.placedOrders.Load(id)
	if !ok {
		return models.Order{}, fmt.Errorf("no order found for %s order id", id)
	}

	return toOrder(order), nil
}

// Upsert creates or updates an order in the orders DB
func (o *OrderDB) Upsert(order models.Order) {
	o.placedOrders.Store(order.ID, order)
}

func toOrder(maybeOrder any) models.Order {
	order, ok := maybeOrder.(models.Order)
	if !ok {
		panic(fmt.Errorf("error casting %v to order", maybeOrder))
	}
	return order
}
