package data_store

import (
	def "file_downloader/internal/definitions"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiskStore(t *testing.T) {
	// Create a temporary result path
	resultPath, err := os.MkdirTemp("", "results")
	assert.NoError(t, err)
	defer os.RemoveAll(resultPath)

	// Test DiskStore
	store := NewDiskStore(resultPath)
	downloadResultChan := make(chan def.UrlDownloadResult, 2)
	var wg sync.WaitGroup
	wg.Add(1)
	go store.StoreData(downloadResultChan, &wg)

	// Send test data to the channel
	downloadResultChan <- def.UrlDownloadResult{Url: "http://example.com", Content: []byte("Example content")}
	close(downloadResultChan)
	wg.Wait()

	// Verify stored data
	files, err := os.ReadDir(resultPath)
	assert.NoError(t, err)
	assert.Len(t, files, 1)

	content, err := os.ReadFile(resultPath + "/" + files[0].Name())
	assert.NoError(t, err)
	assert.Equal(t, "Example content", string(content))
}
