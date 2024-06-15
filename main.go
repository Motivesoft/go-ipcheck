package main

import (
	"fmt"
	"io"
	"net/http"
)

// Entry point of the program.
//
// It retrieves the external/public IP address by sending a GET request to the URL "https://ident.me".
// It creates a new HTTP client and sends the request using the client.
// It reads the response body and prints it to the console.
//
// No parameters.
// No return value, but will print the error result if the call fails.
func main() {
	// URL to retrieve
	url := "https://ident.me"

	// Create a new HTTP client
	client := &http.Client{}

	// Send a GET request to the URL
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the response body
	fmt.Println(string(body))
}
