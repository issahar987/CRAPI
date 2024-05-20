package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	configurator "github.com/tomek-skrond/crapiconfigurator"
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

func GetPayload(token string, ids []string) {
	for _, id := range ids {
		url := fmt.Sprintf("http://crapi.bobaklabs.com:8888/identity/api/v2/vehicle/%s/location", id)

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

	var posts []Post
	if err := json.Unmarshal(body, &posts); err != nil {
		log.Fatalln(err)
	}
	var ids []string
	for _, p := range posts {
		vehicleId := p.Author.VehicleID
		if vehicleId != "" {
			ids = append(ids, vehicleId)
		}
		// fmt.Println(p.AuthorID)
	}

	return ids
}
