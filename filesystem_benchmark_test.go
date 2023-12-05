// filesystem_benchmark_test.go
package main

import (
	"fmt"
	"testing"
)

func TestFileSystemBenchmarkReadFile(b *testing.T) {
	benchmarkResult := RunReadBenchmark()
	// Add assertions or validations if needed
	fmt.Println(benchmarkResult)
}

func TestFileSystemBenchmarkWriteFile(b *testing.T) {
	benchmarkResult := RunWriteBenchmark()
	// Add assertions or validations if needed
	fmt.Println(benchmarkResult)
}
