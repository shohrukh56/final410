// benchmark.go
package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func runReadBenchmarkConcurrent(b *testing.B, numCalls int) {
	fileSystem := NewFileSystem()
	filename := "benchmark_file.txt"
	content := "benchmark_content"
	fileSystem.WriteFile(filename, content)

	var totalTime time.Duration

	startTime := time.Now()

	var wg sync.WaitGroup
	wg.Add(numCalls)

	for i := 0; i < numCalls; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < b.N; j++ {
				_, err := fileSystem.ReadFile(filename)
				if err != nil {
					b.Fatalf("Error reading file: %s", err)
				}
			}
		}()
	}

	wg.Wait()

	totalTime = time.Since(startTime)

	fmt.Printf("numConcurrentCalls: %d\n", numCalls)
	fmt.Printf("Total Iterations: %d\n", b.N*numCalls)
	fmt.Printf("Total Time: %s\n", totalTime)
}

func runWriteBenchmarkConcurrent(b *testing.B, numCalls int) {
	fileSystem := NewFileSystem()
	filename := "benchmark_file.txt"
	content := "benchmark_content"

	var totalTime time.Duration

	startTime := time.Now()

	var wg sync.WaitGroup
	wg.Add(numCalls)

	for i := 0; i < numCalls; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < b.N; j++ {
				err := fileSystem.WriteFile(filename, content)
				if err != nil {
					b.Fatalf("Error writing file: %s", err)
				}
			}
		}()
	}

	wg.Wait()

	totalTime = time.Since(startTime)

	fmt.Printf("numConcurrentCalls: %d\n", numCalls)
	fmt.Printf("Total Iterations: %d\n", b.N*numCalls)
	fmt.Printf("Total Time: %s\n", totalTime)
}

func BenchmarkReadFileConcurrent5(b *testing.B) {
	runReadBenchmarkConcurrent(b, 5)
}

func BenchmarkReadFileConcurrent10(b *testing.B) {
	runReadBenchmarkConcurrent(b, 10)
}

func BenchmarkReadFileConcurrent20(b *testing.B) {
	runReadBenchmarkConcurrent(b, 20)
}

func BenchmarkWriteFileConcurrent5(b *testing.B) {
	runWriteBenchmarkConcurrent(b, 5)
}

func BenchmarkWriteFileConcurrent10(b *testing.B) {
	runWriteBenchmarkConcurrent(b, 10)
}

func BenchmarkWriteFileConcurrent20(b *testing.B) {
	runWriteBenchmarkConcurrent(b, 20)
}

// Export functions to get benchmark results
func RunReadBenchmarkConcurrent5() testing.BenchmarkResult {
	return testing.Benchmark(BenchmarkReadFileConcurrent5)
}

func RunReadBenchmarkConcurrent10() testing.BenchmarkResult {
	return testing.Benchmark(BenchmarkReadFileConcurrent10)
}

func RunReadBenchmarkConcurrent20() testing.BenchmarkResult {
	return testing.Benchmark(BenchmarkReadFileConcurrent20)
}

func RunWriteBenchmarkConcurrent5() testing.BenchmarkResult {
	return testing.Benchmark(BenchmarkWriteFileConcurrent5)
}

func RunWriteBenchmarkConcurrent10() testing.BenchmarkResult {
	return testing.Benchmark(BenchmarkWriteFileConcurrent10)
}

func RunWriteBenchmarkConcurrent20() testing.BenchmarkResult {
	return testing.Benchmark(BenchmarkWriteFileConcurrent20)
}
