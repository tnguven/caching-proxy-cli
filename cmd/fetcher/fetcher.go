package fetcher

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func FetchURL(requestURL string) ([]byte, error) {
	response, err := http.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching URL: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching URL: received status code %d", response.StatusCode)
	}

	var body bytes.Buffer
	if _, err := io.Copy(&body, response.Body); err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return body.Bytes(), nil
}
