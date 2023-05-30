package stats

import (
	"github.com/tsmoreland/go-web/ordersApi/models"
	"log"
	"math/rand"
	"time"
)

type statsService struct {
	result    *result
	processed <-chan models.Order
	done      <-chan struct{}
}

func (s *statsService) GetStats() models.Statistics {
	return s.result.Get()
}

type Service interface {
	GetStats() models.Statistics
}

func New(processed <-chan models.Order, done chan struct{}) Service {
	s := statsService{
		result:    &result{},
		processed: processed,
		done:      done,
	}
	go s.processStats()
	return &s
}

func (s *statsService) processStats() {
	log.Println("Starting status service")
	for {
		select {
		case order := <-s.processed:
			pubStats := s.processOrder(order)
			s.reconcile(pubStats)
		case <-s.done:
			log.Println("Stats processing stopped")
		}
	}
}

func (s *statsService) reconcile(pubStats models.Statistics) {
	s.result.Combine(pubStats)
}

func (s *statsService) processOrder(order models.Order) models.Statistics {
	randomSleep()
	if order.Status == models.OrderStatusCompleted {
		return models.Statistics{
			CompletedOrders: 0,
			Revenue:         *order.Total,
		}
	}
	return models.Statistics{
		RejectedOrders: 0,
	}
}

func randomSleep() {
	rand.Seed(time.Now().UnixNano()) // TODO: replace deprecated funcition, though it'll likely be gone before we need to
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}
