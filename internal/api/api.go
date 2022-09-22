package api

import (
	"net/http"
)

func HttpGet(url string, params map[string]string) ([]byte, error) {
	if len(params) != 0 {
		url += "?"
		for key, value := range params {
			url += key + "=" + value + "&"
		}
	}
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body := switchContentEncoding(resp)
	return body, nil
}

func HttpGetWithHeader(headers map[string]string, url string, params map[string]string) ([]byte, error) {
	if len(params) != 0 {
		url += "?"
		for key, value := range params {
			url += key + "=" + value + "&"
		}
	}
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body := switchContentEncoding(resp)
	return body, nil
}
