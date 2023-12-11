// filesystem_benchmark_test.go
package main

import (
	"fmt"
	"testing"
)

func TestFileSystemBenchmarkReadFile(b *testing.T) {
	benchmarkResult5 := RunReadBenchmarkConcurrent5()
	printAverageTime(5, benchmarkResult5)

	benchmarkResult10 := RunReadBenchmarkConcurrent10()
	printAverageTime(10, benchmarkResult10)

	benchmarkResult20 := RunReadBenchmarkConcurrent20()
	printAverageTime(20, benchmarkResult20)
}

func TestFileSystemBenchmarkWriteFile(b *testing.T) {
	benchmarkResult5 := RunWriteBenchmarkConcurrent5()
	printAverageTime(10, benchmarkResult5)

	benchmarkResult10 := RunWriteBenchmarkConcurrent10()
	printAverageTime(10, benchmarkResult10)

	benchmarkResult20 := RunWriteBenchmarkConcurrent20()
	printAverageTime(10, benchmarkResult20)

	// Add assertions or validations if needed
	fmt.Println(benchmarkResult5, benchmarkResult10, benchmarkResult20)
}

func printAverageTime(numConcurrentCalls int, result testing.BenchmarkResult) {
	averageTimePerOp := result.T.Seconds() / float64(result.N)

	fmt.Printf("numConcurrentCalls: %d\n", numConcurrentCalls)
	fmt.Printf("Total Iterations: %d\n", result.N)
	fmt.Printf("Total Time: %s\n", result.T)
	fmt.Printf("Average Time per Operation: %.12f seconds\n", averageTimePerOp)
	fmt.Println()
}
