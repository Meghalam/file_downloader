package data_store

import (
	"crypto/rand"
	def "file_downloader/internal/definitions"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/sirupsen/logrus"
)

type DiskStore struct {
	AbsoluteFilePath string
}

func NewDiskStore(absFilePath string) *DiskStore {
	ds := new(DiskStore)
	ds.AbsoluteFilePath = absFilePath
	return ds
}

// stores file to disk in the given absolute path
func (ds *DiskStore) StoreData(downloadResultChan <-chan def.UrlDownloadResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for result := range downloadResultChan {
		if result.Err != nil {
			continue
		}
		fileName := randomFileName()
		filePath := filepath.Join(ds.AbsoluteFilePath, fileName+".txt")
		err := os.WriteFile(filePath, result.Content, 0644)
		if err != nil {
			logrus.Errorf("Error while writing to filepath %s, Error: %v", filePath, err)
		} else {
			logrus.Infof("File saved successfully, %s", filePath)
		}
	}
}

// creates random file names
func randomFileName() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
