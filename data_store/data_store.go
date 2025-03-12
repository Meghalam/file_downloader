package data_store

import (
	def "file_downloader/internal/definitions"
	"sync"
)

type Datastore interface {
	StoreData(downloadResultChan <-chan def.UrlDownloadResult, wg *sync.WaitGroup)
}
