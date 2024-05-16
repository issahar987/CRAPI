package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetPayload(token string, ids []string) {
	for _, id := range ids {
		url := fmt.Sprintf("http://crapi.bobaklabs.com:8888/identity/api/v2/vehicle/%s/location", id)

		req, _ := http.NewRequest("GET", url, nil)
		req = ConfigureRequest(req, token)

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

	req = ConfigureRequest(req, token)

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
	}

	return ids
}
