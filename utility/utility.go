package utility

import (
	"path/filepath"
	"strings"
)

// getFileType returns the file type based on the file extension
func GetFileType(filePath string) string {
	ext := strings.ToLower(filepath.Ext(filePath)) // Extract and normalize extension

	// Define common file types based on extensions
	fileTypes := map[string]string{
		".txt":  "Text File",
		".csv":  "CSV File",
		".json": "JSON File",
		".xml":  "XML File",
		".jpg":  "JPEG Image",
		".png":  "PNG Image",
		".gif":  "GIF Image",
		".pdf":  "PDF Document",
		".doc":  "Word Document",
		".docx": "Word Document",
		".xls":  "Excel Spreadsheet",
		".xlsx": "Excel Spreadsheet",
	}

	if fileType, found := fileTypes[ext]; found {
		return fileType
	}
	return "Unknown File Type"
}
