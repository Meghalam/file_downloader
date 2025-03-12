package file_reader

import "sync"

// interface added to support all types of file in future when required. e.g - word, csv, excel etc
type FileReader interface {
	ReadFile(filePath string, urlAggregator chan<- string, wg *sync.WaitGroup)
}
