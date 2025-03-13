package main

import (
	def "file_downloader/internal/definitions"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	// Create a temporary input file
	inputFile, err := os.CreateTemp("/tmp/", "input-*.csv")
	assert.NoError(t, err)
	defer os.Remove(inputFile.Name())

	// Write test URLs to the input file
	_, err = inputFile.WriteString("urls\nhttp://example.com/test1\nhttp://example.com/test2\n")
	assert.NoError(t, err)
	inputFile.Close()

	// Redirect log output to avoid cluttering test output
	logrus.SetOutput(io.Discard)

	// Run the main function with the test input file
	os.Args = []string{"cmd", inputFile.Name()}
	main()
}

func TestDownloadFromUrl(t *testing.T) {
	// Mock HTTP server
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	go server.ListenAndServe()
	defer server.Close()

	// Test downloadFromUrl function
	url := "http://localhost:8080"
	downloadResultChan := make(chan def.UrlDownloadResult, 1)
	downloadFromUrl(url, downloadResultChan)
	result := <-downloadResultChan

	assert.NoError(t, result.Err)
	assert.Equal(t, "Hello, world!", string(result.Content))
}
