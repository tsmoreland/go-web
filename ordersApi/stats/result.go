package stats

import (
	"github.com/tsmoreland/go-web/ordersApi/models"
	"sync"
)

type result struct {
	latest models.Statistics
	lock   sync.Mutex
}
type Result interface {
	Get() models.Statistics
	Combine(stats models.Statistics)
}

// Get returns the latest statistics
func (r *result) Get() models.Statistics {
	r.lock.Lock()
	defer r.lock.Unlock()
	return r.latest
}

// Combine updates the latest statistics by adding stats
func (r *result) Combine(stats models.Statistics) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.latest = models.Combine(r.latest, stats)
}
