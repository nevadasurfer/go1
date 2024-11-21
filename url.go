// Go program checking the response status code
// Go program checking the response status code
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Specify the input file containing URLs
	inputFile := "urls.txt"

	// Open the file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Process each URL
	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())
		if url == "" {
			continue // Skip empty lines
		}

		// Check the URL
		checkURL(url)
	}

	// Check for errors while reading the file
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}

func checkURL(url string) {
	// Ensure the URL has a scheme (http or https)
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	// Make a GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching URL %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	// Print the status code
	fmt.Printf("URL: %s, Status Code: %d\n", url, resp.StatusCode)
}

