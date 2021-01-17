package main

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkPUTUnixSecond(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// var lastTransactionsDuration int64 = 30
		// var transactionsPerSecond = 1000000
		// var timeSeriesDatbase = make(map[int64]Transaction, lastTransactionsDuration)
		// test exits in ~2 seconds
		// timeSeriesDatbase is ~ 1% filled || Hiting only ~1% slots in DB
		PUT(float64(rand.Intn(100.00-1.00)+1.00), time.Now().UTC().Unix())
	}
}

// mike@mike:~/go/src/github.com/gufranmirza/most-recent-requests/thread_unsafe_solution$ go test -bench=BenchmarkPUTUnixSecond -run=1000000000
// goos: linux
// goarch: amd64
// pkg: github.com/gufranmirza/most-recent-requests/thread_unsafe_solution
// BenchmarkPUTUnixSecond-4         9494624               123 ns/op
// PASS
// ok      github.com/gufranmirza/most-recent-requests/thread_unsafe_solution      1.298s

func BenchmarkPUTUnixNano(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// var lastTransactionsDuration int64 = 30
		// var transactionsPerSecond = 1000000
		// var timeSeriesDatbase = make(map[int64]Transaction, lastTransactionsDuration)
		// test exits in ~2 seconds
		// timeSeriesDatbase is ~ 40% filled - Using UnixNano for filling || Hitting ~40% slots in DB
		PUT(float64(rand.Intn(100.00-1.00)+1.00), time.Now().UTC().UnixNano())
	}
}

// mike@mike:~/go/src/github.com/gufranmirza/most-recent-requests/thread_unsafe_solution$ go test -bench=BenchmarkPUTUnixNano -run=1000000000
// goos: linux
// goarch: amd64
// pkg: github.com/gufranmirza/most-recent-requests/thread_unsafe_solution
// BenchmarkPUTUnixNano-4           7889335               148 ns/op
// PASS
// ok      github.com/gufranmirza/most-recent-requests/thread_unsafe_solution      1.332s

func BenchmarkGET(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// var lastTransactionsDuration int64 = 30
		// var transactionsPerSecond = 1000000
		// var timeSeriesDatbase = make(map[int64]Transaction, lastTransactionsDuration)
		GET()
	}
}

// mike@mike:~/go/src/github.com/gufranmirza/most-recent-requests/thread_unsafe_solution$ go test -bench=. -run=10000000X
// goos: linux
// goarch: amd64
// pkg: github.com/gufranmirza/most-recent-requests/thread_unsafe_solution
// BenchmarkPUTUnixSecond-4         9789592               124 ns/op
// BenchmarkPUTUnixNano-4           7863362               148 ns/op
// BenchmarkGET-4                   2588995               461 ns/op
// PASS
// ok      github.com/gufranmirza/most-recent-requests/thread_unsafe_solution      4.344s
