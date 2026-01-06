package benchmark

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	chi "github.com/creativeyann17/GolangAPIsTest/internal/chi"
	echo "github.com/creativeyann17/GolangAPIsTest/internal/echo"
	fasthttp "github.com/creativeyann17/GolangAPIsTest/internal/fasthttp"
	fiber "github.com/creativeyann17/GolangAPIsTest/internal/fiber"
	gin "github.com/creativeyann17/GolangAPIsTest/internal/gin"
	hertz "github.com/creativeyann17/GolangAPIsTest/internal/hertz"
	httprouter "github.com/creativeyann17/GolangAPIsTest/internal/httprouter"
)

const (
	port        = ":8080"
	numRequests = 10000
	warmupReqs  = 100
)

type BenchResult struct {
	Framework   string
	MeanTime    time.Duration
	MinTime     time.Duration
	MaxTime     time.Duration
	TotalTime   time.Duration
	SuccessRate float64
}

func TestAllFrameworks(t *testing.T) {
	frameworks := []string{"Fiber", "Gin", "Echo", "Chi", "HttpRouter", "Hertz", "FastHTTP"}
	results := make([]BenchResult, 0, len(frameworks))

	for _, fw := range frameworks {
		t.Run(fw, func(t *testing.T) {
			result := benchmarkFramework(t, fw)
			results = append(results, result)
		})
	}

	// Print summary
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("BENCHMARK SUMMARY")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("%-15s | %-12s | %-12s | %-12s | %-12s\n",
		"Framework", "Mean Time", "Min Time", "Max Time", "Success Rate")
	fmt.Println(strings.Repeat("-", 80))

	for _, r := range results {
		fmt.Printf("%-15s | %-12s | %-12s | %-12s | %.2f%%\n",
			r.Framework,
			r.MeanTime.Round(time.Microsecond),
			r.MinTime.Round(time.Microsecond),
			r.MaxTime.Round(time.Microsecond),
			r.SuccessRate)
	}
	fmt.Println(strings.Repeat("=", 80))
}

type ServerInterface interface {
	Stop() error
}

func benchmarkFramework(t *testing.T, framework string) BenchResult {
	// Start server
	log.Printf("Starting %s server...", framework)
	var server ServerInterface
	switch framework {
	case "Fiber":
		server = fiber.Start(port)
	case "Gin":
		server = gin.Start(port)
	case "Echo":
		server = echo.Start(port)
	case "Chi":
		server = chi.Start(port)
	case "HttpRouter":
		server = httprouter.Start(port)
	case "Hertz":
		server = hertz.Start(port)
	case "FastHTTP":
		server = fasthttp.Start(port)
	}
	defer server.Stop()

	// Wait for server to be ready
	waitForServer(t)

	// Warmup
	log.Printf("Warming up %s (%d requests)...", framework, warmupReqs)
	for i := 0; i < warmupReqs; i++ {
		makeRequest()
	}

	// Benchmark
	log.Printf("Benchmarking %s (%d requests)...", framework, numRequests)
	times := make([]time.Duration, 0, numRequests)
	successCount := 0

	for i := 0; i < numRequests; i++ {
		start := time.Now()
		if makeRequest() {
			elapsed := time.Since(start)
			times = append(times, elapsed)
			successCount++
		}
	}

	// Calculate stats
	var totalTime time.Duration
	minTime := times[0]
	maxTime := times[0]

	for _, t := range times {
		totalTime += t
		if t < minTime {
			minTime = t
		}
		if t > maxTime {
			maxTime = t
		}
	}

	meanTime := totalTime / time.Duration(len(times))
	successRate := (float64(successCount) / float64(numRequests)) * 100

	result := BenchResult{
		Framework:   framework,
		MeanTime:    meanTime,
		MinTime:     minTime,
		MaxTime:     maxTime,
		TotalTime:   totalTime,
		SuccessRate: successRate,
	}

	log.Printf("%s - Mean: %s, Min: %s, Max: %s, Success: %.2f%%",
		framework,
		meanTime.Round(time.Microsecond),
		minTime.Round(time.Microsecond),
		maxTime.Round(time.Microsecond),
		successRate)

	return result
}

func makeRequest() bool {
	resp, err := http.Get("http://localhost" + port + "/hello")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	io.ReadAll(resp.Body)
	return resp.StatusCode == http.StatusOK
}

func waitForServer(t *testing.T) {
	for i := 0; i < 50; i++ {
		resp, err := http.Get("http://localhost" + port + "/hello")
		if err == nil {
			resp.Body.Close()
			time.Sleep(50 * time.Millisecond)
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
	t.Fatal("Server failed to start")
}
