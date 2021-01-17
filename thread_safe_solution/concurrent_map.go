package main

import (
	"sync"
)

// Map type that can be safely shared between
// goroutines that require read/write access to a map
type Map struct {
	sync.RWMutex
	items map[int64]Transaction
}

func NewMap(size int64) *Map {
	return &Map{
		items: make(map[int64]Transaction, size),
	}
}

// Sets a key in a concurrent map
func (m *Map) Set(key int64, value *Transaction) {
	m.Lock()
	defer m.Unlock()

	m.items[key] = *value
}

// Gets a key from a concurrent map
func (m *Map) Get(key int64) (Transaction, bool) {
	m.Lock()
	defer m.Unlock()

	value, ok := m.items[key]
	return value, ok
}

func (m *Map) Analytics() <-chan TransactionsAnalytics {
	c := make(chan TransactionsAnalytics, len(m.items))

	go func() {
		m.Lock()
		defer m.Unlock()
		sum := 0.0
		transactions := 0

		for _, v := range m.items {
			sum += v.Amount
			transactions += v.Transactions
		}

		c <- TransactionsAnalytics{
			TotalAmount: sum, 
			AverageAmount: (sum/float64(transactions)), 
			Transactions: transactions ,
		}
		close(c)

	}()

	return c
}