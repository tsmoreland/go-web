package repo

import (
	"fmt"
	"github.com/tsmoreland/go-web/ordersApi/db"
	"github.com/tsmoreland/go-web/ordersApi/models"
	"log"
	"math"
)

// repo holds all the dependencies required for repo operations
type repo struct {
	products *db.ProductDB
	orders   *db.OrderDB
	incoming chan models.Order
	done     chan struct{}
}

// Repo is the interface we expose to outside packages
type Repo interface {
	CreateOrder(item models.Item) (*models.Order, error)
	GetAllProducts() []models.Product
	GetOrder(id string) (models.Order, error)
	Close()
}

// New creates a new Order repo with the correct database dependencies
func New() (Repo, error) {
	p, err := db.NewProducts()
	if err != nil {
		return nil, err
	}
	o := repo{
		products: p,
		orders:   db.NewOrders(),
		incoming: make(chan models.Order),
		done:     make(chan struct{}),
	}

	go o.processOrders()

	return &o, nil
}

// GetAllProducts returns all products in the system
func (r *repo) GetAllProducts() []models.Product {
	return r.products.FindAll()
}

// GetOrder returns the given order if one exists
func (r *repo) GetOrder(id string) (models.Order, error) {
	return r.orders.FindById(id)
}

// CreateOrder creates a new order for the given item
func (r *repo) CreateOrder(item models.Item) (*models.Order, error) {
	if err := r.validateItem(item); err != nil {
		return nil, err
	}
	order := models.NewOrder(item)
	r.orders.Upsert(order)

	select {
	case r.incoming <- order:
		r.orders.Upsert(order)
		return &order, nil
	case <-r.done:
		return nil, fmt.Errorf("orders app is closed, try again later")
	}
}

// validateItem runs validations on a given order
func (r *repo) validateItem(item models.Item) error {
	if item.Amount < 1 {
		return fmt.Errorf("order amount must be at least 1:got %d", item.Amount)
	}
	if err := r.products.Exists(item.ProductID); err != nil {
		return fmt.Errorf("product %s does not exist", item.ProductID)
	}
	return nil
}

func (r *repo) processOrders() {
	for {
		select {
		case order := <-r.incoming:
			r.processOrder(&order)
			r.orders.Upsert(order)
			fmt.Printf("Processing order %s completed\n", order.ID)
		case <-r.done:
			log.Println("Order processing stopped.")
			return
		}
	}
}

// processOrder is an internal method which completes or rejects an order
func (r *repo) processOrder(order *models.Order) {
	item := order.Item
	product, err := r.products.FindById(item.ProductID)
	if err != nil {
		order.Status = models.OrderStatusRejected
		order.Error = err.Error()
		return
	}
	if product.Stock < item.Amount {
		order.Status = models.OrderStatusRejected
		order.Error = fmt.Sprintf("not enough stock for product %s:got %d, want %d", item.ProductID, product.Stock, item.Amount)
		return
	}
	remainingStock := product.Stock - item.Amount
	product.Stock = remainingStock
	r.products.Upsert(product)

	total := math.Round(float64(order.Item.Amount)*product.Price*100) / 100
	order.Total = &total
	order.Complete()
}

// Close closes incoming orders allowing existing orders to complete but no new ones to be accepted.
func (r *repo) Close() {
	close(r.done)
}
