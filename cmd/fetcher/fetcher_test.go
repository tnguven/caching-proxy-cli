package fetcher_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tnguven/caching-proxy-cli/cmd/fetcher"
)

func TestFetchURL_Success(t *testing.T) {
	expectedBody := "Hello, World!"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(expectedBody))
	}))
	defer server.Close()

	body, err := fetcher.FetchURL(server.URL)
	assert.NoError(t, err)
	assert.Equal(t, expectedBody, string(body))
}

func TestFetchURL_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}))
	defer server.Close()

	_, err := fetcher.FetchURL(server.URL)
	assert.Error(t, err)
}

func TestFetchURL_InvalidURL(t *testing.T) {
	_, err := fetcher.FetchURL("http://invalid-url")
	assert.Error(t, err)
}
