package main

import (
	"final410/plotter"
	"fmt"
	"sync"
)

// FileSystem represents a simple file system structure
type FileSystem struct {
	data  map[string]string
	mutex sync.Mutex
}

// NewFileSystem creates a new instance of FileSystem
func NewFileSystem() *FileSystem {
	return &FileSystem{
		data: make(map[string]string),
	}
}

// ReadFile simulates reading a file from the file system
func (fs *FileSystem) ReadFile(filename string) (string, error) {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()

	content, exists := fs.data[filename]
	if !exists {
		return "", fmt.Errorf("file not found: %s", filename)
	}

	return content, nil
}

// WriteFile simulates writing a file to the file system
func (fs *FileSystem) WriteFile(filename, content string) error {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()

	fs.data[filename] = content
	return nil
}

func main() {

	// Create a new file system
	fileSystem := NewFileSystem()

	// Number of concurrent operations
	numOperations := 5

	// Wait group for synchronization
	var wg sync.WaitGroup

	// Concurrent read operations
	for i := 0; i < numOperations; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			filename := fmt.Sprintf("file%d.txt", index)
			content, err := fileSystem.ReadFile(filename)
			if err != nil {
				fmt.Printf("Error reading %s: %s\n", filename, err)
			} else {
				fmt.Printf("Read content from %s: %s\n", filename, content)
			}
		}(i)
	}

	// Concurrent write operations
	for i := 0; i < numOperations; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			filename := fmt.Sprintf("file%d.txt", index)
			content := fmt.Sprintf("Content%d", index)
			err := fileSystem.WriteFile(filename, content)
			if err != nil {
				fmt.Printf("Error writing to %s: %s\n", filename, err)
			} else {
				fmt.Printf("Write content to %s: %s\n", filename, content)
			}
		}(i)
	}
	readResults := RunReadBenchmark()
	readResults.String()
	writeResults := RunWriteBenchmark()
	plotter.CreatePlot("File Read Benchmark Plot", readResults, "read_plot.png")
	plotter.CreatePlot("File Write Benchmark Plot", writeResults, "write_plot.png")

	// Wait for all goroutines to finish
	wg.Wait()
}
