package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Transaction represents a transaction made by a user
type Transaction struct {
	Timestamp    int64   // timestamp in unix seconds
	Amount       float64 // transaction amount in $
	Transactions int     // Number of transactions contributed to total Amount
}

// TransactionsAnalytics represents the transaction analytics like last active transactions in a timeframe
type TransactionsAnalytics struct {
	TotalAmount   float64 // Total transaction amount  in given time frame
	AverageAmount float64 // Averate transaction amount
	Transactions  int     // Total transactions occurred in given time frame
}

var lastTransactionsDuration int64 = 10 // it is the timeframe in seconds to reterive last transactions
var transactionsPerSecond = 100000      // Number of transactions happening in a second
var timeSeriesDatbase = NewMap(lastTransactionsDuration)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-c
		fmt.Println("Ctrl+C pressed in Terminal. Terminating Program")
		cancel()
	}()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				// Concurrent writes to the database by multiple go-routine
				go func() {
					PUT(float64(rand.Intn(5.00-1.00)+1.00), time.Now().UTC().Unix())
				}()
				time.Sleep(time.Second / time.Duration(transactionsPerSecond))
			}
		}
	}(ctx)

	go func(ctx context.Context) {
		ticker := time.NewTicker(time.Second * time.Duration(lastTransactionsDuration))
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				sum, avg, transactions := GET()
				fmt.Printf("Last Transactions in Last %v Seconds Transactions=%v Sum=%v Average=%v\n", lastTransactionsDuration, transactions, sum, avg)
			}
		}
	}(ctx)

	<-ctx.Done()
}

// PUT Adds a transaction reccord into database
func PUT(amount float64, timestamp int64) {
	idx := timestamp % lastTransactionsDuration

	data, ok := timeSeriesDatbase.Get(idx)
	if !ok || data.Timestamp < timestamp { // also drop requests which is delayed and not contribute to time-series
		timeSeriesDatbase.Set(idx, &Transaction{
			Timestamp:    timestamp,
			Transactions: 1,
			Amount:       amount,
		})
	} else {
		timeSeriesDatbase.Set(idx, &Transaction{
			Timestamp:    data.Timestamp,
			Amount:       (data.Amount + amount),
			Transactions: (data.Transactions + 1),
		})
	}

	// writing excessive data on console will hurt the performance, so removed database state
	fmt.Printf("Added Transaction Amount=$%v Timestamp=%v, Index=%v \n", amount, timestamp, idx)
}

// GET returns the average and sum of transactions in specified time frame
func GET() (sum float64, avg float64, transactions int) {
	analytics := <-timeSeriesDatbase.Analytics()
	return analytics.TotalAmount, analytics.AverageAmount, analytics.Transactions
}
