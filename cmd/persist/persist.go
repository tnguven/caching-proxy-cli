package persist

import (
	"fmt"
	"os"
	"path/filepath"
)

const cacheDir = "./cache"

func SaveToCache(hash string, data []byte) error {
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return err
	}

	cachePath := filepath.Join(cacheDir, hash)

	return os.WriteFile(cachePath, data, 0644)
}

func GetCachedResponse(hash string) ([]byte, bool) {
	cachePath := filepath.Join(cacheDir, hash)
	if _, err := os.Stat(cachePath); os.IsNotExist(err) {
		return nil, false
	}

	data, err := os.ReadFile(cachePath)
	if err != nil {
		return nil, false
	}

	return data, true
}

func ClearCache() error {
	if err := os.RemoveAll(cacheDir); err != nil {
		fmt.Println("Error clearing cache:", err)
		return err
	}

	fmt.Println("Cache cleared.")

	return nil
}
