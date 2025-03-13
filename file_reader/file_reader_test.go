package file_reader

import (
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCsvFileReader(t *testing.T) {
	// Create a temporary input file
	inputFile, err := os.CreateTemp("/tmp/", "input.csv")
	assert.NoError(t, err)
	defer os.Remove(inputFile.Name())

	// Write test URLs to the input file
	_, err = inputFile.WriteString("urls\nhttp://example.com\nhttp://example.org\n")
	assert.NoError(t, err)
	inputFile.Close()

	// Test CsvFileReader
	urlAggregatorChan := make(chan string, 2)
	var wg sync.WaitGroup
	wg.Add(1)
	reader := NewCsvFileReader()
	go reader.ReadFile(inputFile.Name(), urlAggregatorChan, &wg)

	// Verify URLs
	urls := []string{}
	for url := range urlAggregatorChan {
		urls = append(urls, url)
	}

	assert.ElementsMatch(t, []string{"http://example.com", "http://example.org"}, urls)
}
