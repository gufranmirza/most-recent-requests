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

// mike@mike:~/go/src/github.com/gufranmirza/most-recent-requests/thread_safe_solution$ go test -bench=BenchmarkPUTUnixSecond -run=1000000000
// goos: linux
// goarch: amd64
// pkg: github.com/gufranmirza/most-recent-requests/thread_safe_solution
// BenchmarkPUTUnixSecond-4         6536997               182 ns/op
// PASS
// ok      github.com/gufranmirza/most-recent-requests/thread_safe_solution        1.382s

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

// mike@mike:~/go/src/github.com/gufranmirza/most-recent-requests/thread_safe_solution$  go test -bench=BenchmarkPUTUnixNano -run=1000000000
// goos: linux
// goarch: amd64
// pkg: github.com/gufranmirza/most-recent-requests/thread_safe_solution
// BenchmarkPUTUnixNano-4           6042093               192 ns/op
// PASS
// ok      github.com/gufranmirza/most-recent-requests/thread_safe_solution        1.371s

func BenchmarkGET(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// var lastTransactionsDuration int64 = 30
		// var transactionsPerSecond = 1000000
		// var timeSeriesDatbase = make(map[int64]Transaction, lastTransactionsDuration)
		GET()
	}
}

// mike@mike:~/go/src/github.com/gufranmirza/most-recent-requests/thread_safe_solution$ go test -bench=. -run=10000000X
// goos: linux
// goarch: amd64
// pkg: github.com/gufranmirza/most-recent-requests/thread_safe_solution
// BenchmarkPUTUnixSecond-4         6410214               181 ns/op
// BenchmarkPUTUnixNano-4           6154660               194 ns/op
// BenchmarkGET-4                    536581              2417 ns/op
// PASS
// ok      github.com/gufranmirza/most-recent-requests/thread_safe_solution        4.071s
