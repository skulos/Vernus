// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// func main() {
// 	// Define the JSON payload
// 	payload := map[string]interface{}{
// 		"name":    "example",
// 		"version": "v1.0.0",
// 	}

// 	// Convert the payload to JSON
// 	payloadBytes, err := json.Marshal(payload)
// 	if err != nil {
// 		fmt.Println("Error marshaling JSON:", err)
// 		return
// 	}

// 	// Create a request with the JSON payload
// 	req, err := http.NewRequest("POST", "http://localhost:8080/register", bytes.NewBuffer(payloadBytes))
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return
// 	}

// 	// Set the Content-Type header to indicate JSON
// 	req.Header.Set("Content-Type", "application/json")

// 	// Create an HTTP client and send the request
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Error sending request:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	// Check the response status code
// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Println("Unexpected response status code:", resp.StatusCode)
// 		return
// 	}

// 	// Read the response body
// 	respBody, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response body:", err)
// 		return
// 	}

//		fmt.Println("Response body:", string(respBody))
//		// fmt.Println()
//	}
package client

import (
	"Vernus/artefacts"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func SendNewRelease(endpointURL string, release artefacts.ReleaseArtifact) error {
	// Convert release data to JSON
	jsonData, err := json.Marshal(release)
	if err != nil {
		return err
	}

	// Send POST request to the endpoint
	resp, err := http.Post(endpointURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return err // Or handle the error based on the desired behavior
	}

	// Request successful
	log.Println("New release sent successfully!")
	return nil
}
