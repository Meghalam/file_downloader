package main

import (
	"file_downloader/data_store"
	"file_downloader/file_reader"
	def "file_downloader/internal/definitions"
	"file_downloader/utility"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) < 2 {
		logrus.Fatal("Input file path is missing, Usage: go run main.go <filepath>")
		os.Exit(1)
	}
	inputFilepath := os.Args[1]
	urlAggregatorChan := make(chan string, def.MaxWorkers)
	downloadResultChan := make(chan def.UrlDownloadResult, def.MaxWorkers)
	var fileReaderWg, downloaderWg, dataStoreWg sync.WaitGroup
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logrus.Info("Starting File Downloader pipeline")
	fileType := utility.GetFileType(inputFilepath)
	var fileReader file_reader.FileReader
	// todo: add switch case for other file types when they are supported
	switch fileType {
	case def.CsvFileType:
		fileReader = file_reader.NewCsvFileReader()
	default:
		logrus.Fatalf("Error while getting filetype: Unknown Filetype")
	}
	// Read file
	fileReaderWg.Add(1)
	go fileReader.ReadFile(inputFilepath, urlAggregatorChan, &fileReaderWg)

	//data storage
	var dataStore data_store.Datastore
	dataStore = data_store.NewDiskStore(def.ResultPath)
	dataStoreWg.Add(1)
	go dataStore.StoreData(downloadResultChan, &dataStoreWg)

	//download Urls
	sem := make(chan struct{}, def.MaxWorkers)
	for url := range urlAggregatorChan {
		downloaderWg.Add(1)
		sem <- struct{}{}
		go func(url string) {
			defer downloaderWg.Done()
			downloadFromUrl(url, downloadResultChan)
			<-sem
		}(url)
	}

	downloaderWg.Wait()
	logrus.Info("Done waiting")
	close(downloadResultChan)
	dataStoreWg.Wait()
	time.Sleep(100 * time.Millisecond)
	logrus.Info("Processing completed")

}

func downloadFromUrl(url string, downloadResultChan chan<- def.UrlDownloadResult) {
	logrus.Infof("Processing URL: %s", url)
	start := time.Now()
	// response, err := http.Get("https://" + url)
	response, err := http.Get(url)
	if err != nil {
		logrus.Errorf("failed to retrieve from url %s, Error: %v", url, err)
		downloadResultChan <- def.UrlDownloadResult{Url: url, Err: err}
		return
	}
	defer response.Body.Close()
	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		logrus.Warnf("Error while reading response bosy for url %s, Error: %v", url, err)
		downloadResultChan <- def.UrlDownloadResult{Url: url, Err: err}
		return
	}
	logrus.Infof("Successfully downloaded from url %s in %v", url, time.Since(start))
	downloadResultChan <- def.UrlDownloadResult{Url: url, Content: respBody}
}
