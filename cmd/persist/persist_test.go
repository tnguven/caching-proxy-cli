package persist_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tnguven/caching-proxy-cli/cmd/persist"
)

func TestSaveToCache_Success(t *testing.T) {
	hash := "testhash"
	data := []byte("testdata")

	err := persist.SaveToCache(hash, data)
	assert.NoError(t, err)

	cachePath := filepath.Join("./cache", hash)
	defer os.Remove(cachePath)

	savedData, err := os.ReadFile(cachePath)
	assert.NoError(t, err)
	assert.Equal(t, data, savedData)
}

func TestGetCachedResponse_Found(t *testing.T) {
	hash := "testhash"
	data := []byte("testdata")

	cachePath := filepath.Join("./cache", hash)
	os.MkdirAll("./cache", 0755)
	os.WriteFile(cachePath, data, 0644)
	defer os.Remove(cachePath)

	cachedData, found := persist.GetCachedResponse(hash)
	assert.True(t, found)
	assert.Equal(t, data, cachedData)
}

func TestGetCachedResponse_NotFound(t *testing.T) {
	hash := "nonexistenthash"

	cachedData, found := persist.GetCachedResponse(hash)
	assert.False(t, found)
	assert.Nil(t, cachedData)
}

func TestClearCache_Success(t *testing.T) {
	os.MkdirAll("./cache", 0755)
	defer os.RemoveAll("./cache")

	err := persist.ClearCache()
	assert.NoError(t, err)

	_, err = os.Stat("./cache")
	assert.True(t, os.IsNotExist(err))
}
