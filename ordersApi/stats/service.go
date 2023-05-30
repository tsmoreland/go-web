package stats

import (
	"github.com/tsmoreland/go-web/ordersApi/models"
	"log"
	"math/rand"
	"time"
)

const WorkerCount = 3

type statsService struct {
	result    *result
	processed <-chan models.Order
	done      <-chan struct{}
	pubStats  chan models.Statistics
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
		pubStats:  make(chan models.Statistics, WorkerCount),
	}
	for i := 0; i < WorkerCount; i++ {
		go s.processStats()
	}
	go s.reconcile()
	return &s
}

func (s *statsService) processStats() {
	log.Println("Starting status service")
	for {
		select {
		case order := <-s.processed:
			pubStats := s.processOrder(order)
			s.pubStats <- pubStats
		case <-s.done:
			log.Println("Stats processing stopped")
		}
	}
}

func (s *statsService) reconcile() {
	for {
		select {
		case pubStats := <-s.pubStats:
			s.result.Combine(pubStats)
		case <-s.done:
			log.Println("Reconcile stopped")
		}
	}
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
