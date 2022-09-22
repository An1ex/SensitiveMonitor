package api

import (
	"compress/flate"
	"compress/gzip"
	"io"
	"net/http"

	"github.com/andybalholm/brotli"
)

func switchContentEncoding(resp *http.Response) []byte {
	var bodyReader io.Reader
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		bodyReader, _ = gzip.NewReader(resp.Body)
	case "deflate":
		bodyReader = flate.NewReader(resp.Body)
	case "br":
		bodyReader = brotli.NewReader(resp.Body)
	default:
		bodyReader = resp.Body
	}
	body, _ := io.ReadAll(bodyReader)
	return body
}
