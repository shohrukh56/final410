package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sync"
)

// FileSystem represents a simple file system structure
type FileSystem struct {
	data  map[string]fileEntry
	mutex sync.Mutex
}

type fileEntry struct {
	content  string
	checksum string
}

// NewFileSystem creates a new instance of FileSystem
func NewFileSystem() *FileSystem {
	return &FileSystem{
		data: make(map[string]fileEntry),
	}
}

// ReadFile simulates reading a file from the file system
func (fs *FileSystem) ReadFile(filename string) (string, error) {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()

	entry, exists := fs.data[filename]
	if !exists {
		return "", fmt.Errorf("file not found: %s", filename)
	}

	// Verify checksum before returning content
	if !verifyChecksum(entry.content, entry.checksum) {
		return "", fmt.Errorf("checksum verification failed for file: %s", filename)
	}

	return entry.content, nil
}

// WriteFile simulates writing a file to the file system
func (fs *FileSystem) WriteFile(filename, content string) error {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()

	// Generate checksum for the content
	checksum := generateChecksum(content)

	// Store content and checksum in the file system
	fs.data[filename] = fileEntry{content: content, checksum: checksum}
	return nil
}

// generateChecksum generates an MD5 checksum for the given content
func generateChecksum(content string) string {
	hasher := md5.New()
	hasher.Write([]byte(content))
	return hex.EncodeToString(hasher.Sum(nil))
}

// verifyChecksum verifies the content against the given checksum
func verifyChecksum(content, checksum string) bool {
	return generateChecksum(content) == checksum
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

	// Wait for all goroutines to finish
	wg.Wait()
}
