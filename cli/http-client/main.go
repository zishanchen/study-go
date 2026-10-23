package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func fetch(url string) error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("User-Agent", "study-go-http-client/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}

	defer resp.Body.Close()

	fmt.Println("status:", resp.Status)
	fmt.Println("status code:", resp.StatusCode)
	fmt.Println("content-type:", resp.Header.Get("Content-Type"))

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response error: %w", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("HTTP error: %s", resp.Status)
	}

	fmt.Println("\n--- Body Preview ---")

	limit := 300
	if len(body) < limit {
		limit = len(body)
	}

	fmt.Println(string(body[:limit]))

	return nil
}

func main() {
	url := "https://example.com"

	if len(os.Args) > 1 {
		url = os.Args[1]
	}

	fmt.Println("GET:", url)

	err := fetch(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nRequest completed successfully.")
}
