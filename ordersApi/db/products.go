package db

import (
	"fmt"
	"github.com/tsmoreland/go-web/ordersApi/models"
	"github.com/tsmoreland/go-web/ordersApi/utils"
	"sort"
	"sync"
)

type ProductDB struct {
	products sync.Map
}

// NewProducts creates a new empty product DB
func NewProducts() (*ProductDB, error) {
	p := &ProductDB{}
	// load start position
	products := map[string]models.Product{}
	if err := utils.ImportProducts(products); err != nil {
		return nil, err
	}

	return p, nil
}

// Exists checks whether a product with a give id exists
func (p *ProductDB) Exists(id string) error {

	if _, ok := p.products.Load(id); !ok {
		return fmt.Errorf("no product found for id %s", id)
	}

	return nil
}

// FindById returns a given product if one exists
func (p *ProductDB) FindById(id string) (models.Product, error) {
	prod, ok := p.products.Load(id)
	if !ok {
		return models.Product{}, fmt.Errorf("no product found for id %s", id)
	}
	return toProduct(prod), nil
}

// Upsert creates or updates a product in the orders DB
func (p *ProductDB) Upsert(prod models.Product) {
	p.products.Store(prod.ID, prod)
}

// FindAll returns all products in the system
func (p *ProductDB) FindAll() []models.Product {
	var allProducts []models.Product

	p.products.Range(func(key, value interface{}) bool {
		allProducts = append(allProducts, toProduct(value))
		return true
	})

	sort.Slice(allProducts, func(i, j int) bool {
		return allProducts[i].ID < allProducts[j].ID
	})
	return allProducts
}

func toProduct(maybeProduct any) models.Product {
	product, ok := maybeProduct.(models.Product)
	if !ok {
		panic(fmt.Errorf("error casting %v to Product", maybeProduct))
	}
	return product
}
