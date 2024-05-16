package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GetConfig(path string) (*Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config *Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return config, nil
}

func GetJWTToken(loginurl, email, password string) string {

	jsonData, err := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Send a POST request to the login endpoint
	resp, err := http.Post(loginurl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	defer resp.Body.Close()

	// Decode the response body into TokenResponse
	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	return tokenResp.Token
}

func ConfigureRequest(req *http.Request, token string) *http.Request {
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Thunder Client (https://www.thunderclient.com)")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	return req
}
