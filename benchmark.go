// benchmark.go
package main

import "testing"

func BenchmarkReadFile(b *testing.B) {
	fileSystem := NewFileSystem()
	filename := "benchmark_file.txt"
	content := "benchmark_content"
	fileSystem.WriteFile(filename, content)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := fileSystem.ReadFile(filename)
		if err != nil {
			b.Fatalf("Error reading file: %s", err)
		}
	}
	b.ReportMetric(float64(b.N), "iterations")
}

func BenchmarkWriteFile(b *testing.B) {
	fileSystem := NewFileSystem()
	filename := "benchmark_file.txt"
	content := "benchmark_content"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := fileSystem.WriteFile(filename, content)
		if err != nil {
			b.Fatalf("Error writing file: %s", err)
		}
	}
	b.ReportMetric(float64(b.N), "iterations")
}

// Export functions to get benchmark results
func RunReadBenchmark() testing.BenchmarkResult {
	b := testing.Benchmark(BenchmarkReadFile)
	return b
}

func RunWriteBenchmark() testing.BenchmarkResult {
	b := testing.Benchmark(BenchmarkWriteFile)
	return b
}
