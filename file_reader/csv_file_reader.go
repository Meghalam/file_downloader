package file_reader

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

type CsvFileReader struct{}

func NewCsvFileReader() *CsvFileReader {
	return new(CsvFileReader)
}

func (cfr *CsvFileReader) ReadFile(filePath string, urlAggregator chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(filePath)
	if err != nil {
		logrus.Fatalf("Error while opening file %s", filePath)
	}
	defer file.Close()
	csvReader := csv.NewReader(bufio.NewReader(file))
	//skipping the header
	_, err = csvReader.Read()
	if err != nil {
		logrus.Errorf("Error while reading csv header %v", err)
	}
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			// exit on reaching end of file
			break
		}
		if err != nil {
			logrus.Errorf("Error while reading csv record:, %v", err)
		}
		urlAggregator <- record[0]
	}
	close(urlAggregator)
}
