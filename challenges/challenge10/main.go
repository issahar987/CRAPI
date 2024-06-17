package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	configurator "github.com/tomek-skrond/crapiconfigurator/v2"
)

const (
	CONV_PARAMS = "-v codec h21234"
	VIDEO_NAME  = "nicevideo.mp4"
)

func main() {
	config, err := configurator.GetConfig("../challenge-automation/config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	loginurl := fmt.Sprintf("%s%s", config.Hostname, config.LoginURL)
	token := configurator.GetJWTToken(loginurl, config.Email, config.Password)

	url := fmt.Sprintf("%s%s", config.Hostname, config.TargetURL)

	idFlag := flag.Int("id", 1, "ID of the video that is to be edited")

	flag.Parse()

	if *idFlag <= 0 {
		log.Fatalln("wrong video id (should be over 0)")
	}

	// fmt.Println(token)
	UnauthorizedVideoParameterEdition(url, token, *idFlag)
}

func PrintBefore(method, url, token string) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req = configurator.ConfigureRequest(req, token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := configurator.CustomHttpClient().Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	respBody, err := configurator.ReadBody(resp)
	if err != nil {
		log.Fatalln(err)
	}

	modifiedJSON, err := ExcludeParameter(respBody)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Video parameters before edit: ")
	fmt.Println(string(modifiedJSON))
}
func UnauthorizedVideoParameterEdition(targeturl, token string, id int) {
	client := configurator.CustomHttpClient()

	url := fmt.Sprintf("%s/%d", targeturl, id)
	payload := map[string]interface{}{
		"videoName":         VIDEO_NAME,
		"conversion_params": CONV_PARAMS,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln(err)
	}

	PrintBefore("GET", url, token)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalln(err)
	}

	req = configurator.ConfigureRequest(req, token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	respBody, err := configurator.ReadBody(resp)
	if err != nil {
		log.Fatalln(err)
	}

	modifiedJSON, err := ExcludeParameter(respBody)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Video parameters after edit: ")
	fmt.Println(string(modifiedJSON))
}

func ExcludeParameter(respBody []byte) ([]byte, error) {

	var result map[string]interface{}

	if err := json.Unmarshal(respBody, &result); err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
		return nil, err
	}

	delete(result, "profileVideo")

	modifiedJSON, err := json.Marshal(result)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
		return nil, err
	}

	return modifiedJSON, nil
}
