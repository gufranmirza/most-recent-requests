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

var lastTransactionsDuration int64 = 10 // it is the timeframe in seconds to reterive last transactions
var transactionsPerSecond = 100000      // Number of transactions happening in a second
var timeSeriesDatbase = make(map[int64]Transaction, lastTransactionsDuration)

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
				PUT(float64(rand.Intn(5.00-1.00)+1.00), time.Now().UTC().Unix())
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
	data, ok := timeSeriesDatbase[idx]
	if !ok || data.Timestamp < timestamp {
		timeSeriesDatbase[idx] = Transaction{
			Timestamp:    timestamp,
			Transactions: 1,
			Amount:       amount,
		}
	} else {
		timeSeriesDatbase[idx] = Transaction{
			Timestamp:    data.Timestamp,
			Amount:       (data.Amount + amount),
			Transactions: (data.Transactions + 1),
		}
	}
	fmt.Printf("Added Transaction Amount=$%v Timestamp=%v, Index=%v, Database=%v\n", amount, timestamp, idx, timeSeriesDatbase)
}

// GET returns the average and sum of transactions in specified time frame
func GET() (sum float64, avg float64, transactions int) {
	for _, item := range timeSeriesDatbase {
		sum += item.Amount
		transactions += item.Transactions
	}

	return sum, sum / float64(transactions), transactions
}
