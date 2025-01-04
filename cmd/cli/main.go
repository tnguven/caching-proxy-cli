package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	"github.com/tnguven/caching-proxy-cli/cmd/fetcher"
	"github.com/tnguven/caching-proxy-cli/cmd/persist"
)

func hashURL(url string) string {
	h := sha256.New()
	h.Write([]byte(url))
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	origin := flag.String("origin", "", "URL of the server to forward requests to (required)")
	port := flag.String("port", "", "port to listen on (optional)")
	clearCache := flag.Bool("clear-cache", false, "Clear the cache (optional)")

	flag.Parse()

	var err error

	if *clearCache {
		if err := persist.ClearCache(); err != nil {
			fmt.Println("Error clearing cache:", err)
			os.Exit(1)
		}
	}

	if *origin == "" {
		fmt.Println("Error: --origin is required.")
		os.Exit(1)
	}

	requestURL := *origin + *port
	hash := hashURL(requestURL)

	if cachedResponse, found := persist.GetCachedResponse(hash); found {
		fmt.Println("X-Cache: HIT")
		fmt.Println(string(cachedResponse))
		return
	}

	response, err := fetcher.FetchURL(requestURL)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	if err := persist.SaveToCache(hash, response); err != nil {
		fmt.Println("Error saving to cache:", err)
	}

	fmt.Println("X-Cache: MISS")
	fmt.Println(string(response))
}
