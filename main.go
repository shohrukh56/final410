package main

import (
	"crypto/md5"
	"fmt"
	"sync"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// FileSystem represents a simple file system structure
type FileSystem struct {
	data  map[string]fileData
	mutex sync.Mutex
}

type fileData struct {
	content  string
	checksum string
}

// NewFileSystem creates a new instance of FileSystem
func NewFileSystem() *FileSystem {
	return &FileSystem{
		data: make(map[string]fileData),
	}
}

// ReadFile simulates reading a file from the file system
func (fs *FileSystem) ReadFile(filename string, limiter chan struct{}) (string, error) {
	// Acquire the semaphore
	limiter <- struct{}{}
	defer func() {
		// Release the semaphore
		<-limiter
	}()

	fs.mutex.Lock()
	defer fs.mutex.Unlock()

	file, exists := fs.data[filename]
	if !exists {
		return "", fmt.Errorf("file not found: %s", filename)
	}

	// Verify checksum
	calculatedChecksum := calculateChecksum(file.content)
	if calculatedChecksum != file.checksum {
		return "", fmt.Errorf("checksum verification failed for file: %s", filename)
	}

	return file.content, nil
}

// WriteFile simulates writing a file to the file system with checksum
func (fs *FileSystem) WriteFile(filename, content string, limiter chan struct{}) {
	// Acquire the semaphore
	limiter <- struct{}{}
	defer func() {
		// Release the semaphore
		<-limiter
	}()

	fs.mutex.Lock()
	defer fs.mutex.Unlock()

	// Generate checksum for the content
	checksum := calculateChecksum(content)

	// Store content and checksum in the data map
	fs.data[filename] = fileData{
		content:  content,
		checksum: checksum,
	}
}

// calculateChecksum calculates the MD5 checksum for the given content
func calculateChecksum(content string) string {
	hash := md5.New()
	hash.Write([]byte(content))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// benchmarkWriteFile benchmarks the WriteFile operation with concurrent goroutines
func benchmarkWriteFile(fileSystem *FileSystem, content string, iterations int, maxNumGoroutines int) time.Duration {
	startTime := time.Now()

	limiter := make(chan struct{}, maxNumGoroutines)
	var wg sync.WaitGroup

	// Concurrent write operations
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func(filename, content string) {
			defer wg.Done()
			fileSystem.WriteFile(filename, content, limiter)
		}(fmt.Sprintf("file%d.txt", i), content)
	}

	// Wait for all write goroutines to finish
	wg.Wait()

	endTime := time.Now()

	// Calculate the total duration for the benchmarkWriteFile function
	totalDuration := endTime.Sub(startTime)

	return totalDuration
}

// benchmarkReadFile benchmarks the ReadFile operation with concurrent goroutines
func benchmarkReadFile(fileSystem *FileSystem, iterations int, maxNumGoroutines int) time.Duration {
	startTime := time.Now()

	limiter := make(chan struct{}, maxNumGoroutines)
	var wg sync.WaitGroup

	// Prepare data for benchmark
	for i := 0; i < iterations; i++ {
		filename := fmt.Sprintf("file%d.txt", i)
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			_, err := fileSystem.ReadFile(filename, limiter)
			if err != nil {
				fmt.Printf("Error reading %s: %s\n", filename, err)
			}
		}(filename)
	}

	// Wait for all read goroutines to finish
	wg.Wait()

	endTime := time.Now()

	// Calculate the total duration for the benchmarkReadFile function
	totalDuration := endTime.Sub(startTime)

	return totalDuration
}

func main() {
	iterationsList := []int{1000000}
	maxNumGoroutinesList := []int{2, 5, 20, 100} // Set the desired number of goroutines

	for _, iterations := range iterationsList {
		for _, maxNumGoroutines := range maxNumGoroutinesList {
			fmt.Printf("iterations: %v\n", iterations)
			fmt.Printf("max concurrent goroutines: %v\n", maxNumGoroutines)
			fileSystem := NewFileSystem()

			// Larger content for the file
			content := "BenchmarkContentBenchmarkContentBenchmarkContentBenchmarkContentBenchmarkContent"

			// Benchmark WriteFile operation
			writeDuration := benchmarkWriteFile(fileSystem, content, iterations, maxNumGoroutines)
			fmt.Printf("WriteFile duration: %v\n", writeDuration)

			// Benchmark ReadFile operation
			readDuration := benchmarkReadFile(fileSystem, iterations, maxNumGoroutines)
			fmt.Printf("ReadFile duration: %v\n", readDuration)
			PrintCPUUsage()
		}
		fmt.Println("----------------------------------")
	}
}

func PrintCPUUsage() {
	percentages, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Printf("Error getting CPU usage: %v\n", err)
		return
	}

	fmt.Printf("CPU Usage: %.2f%%\n\n", percentages[0])
}
