package definitions

type UrlDownloadResult struct {
	Url     string
	Content []byte
	Err     error
}

const (
	MaxWorkers = 50
	ResultPath = "/tmp/"
)

const (
	TextFileType  = "Text File"
	CsvFileType   = "CSV File"
	JsonFileType  = "JSON File"
	XmlFileType   = "XML File"
	JpegFileType  = "JPEG Image"
	PngFileType   = "PNG Image"
	GifFileType   = "GIF Image"
	PdfFileType   = "PDF Document"
	WordFileType  = "Word Document"
	ExcelFileType = "Excel Spreadsheet"
)
