package models

import "math"

type Statistics struct {
	CompletedOrders int     `json:"completedOrders"`
	RejectedOrders  int     `json:"rejectedOrders"`
	Revenue         float64 `json:"revenue"`
	ReversedOrders  int     `json:"reversedOrders"`
}

// Combine adds numbers from two statistics objects
func Combine(this, that Statistics) Statistics {

	return Statistics{
		CompletedOrders: this.CompletedOrders + that.CompletedOrders,
		RejectedOrders:  this.RejectedOrders + that.RejectedOrders,
		Revenue:         math.Round((this.Revenue+that.Revenue)*100) / 100,
	}
}
