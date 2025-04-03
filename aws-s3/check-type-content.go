package aws_s3

import (
	"mime"
	"net/http"
	"path/filepath"
)

// detect content type reads the first few bytes of the file to determine its type
func (c ClientS3) detectContentType(data []byte, fileName string) string {
	// Use http.DetectContentType on first 512 bytes
	contentType := http.DetectContentType(data)
	if contentType == contentTypeOctetStream {
		// Fallback to checking file extension
		ext := filepath.Ext(fileName)
		mimeType := mime.TypeByExtension(ext)
		if mimeType != "" {
			contentType = mimeType
		}
	}
	return contentType
}
