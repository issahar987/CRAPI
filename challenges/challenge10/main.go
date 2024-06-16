package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	configurator "github.com/tomek-skrond/crapiconfigurator"
)

const (
	CONV_PARAMS = "-v codec h2"
	VIDEO_NAME  = "nicevideo.mp4"
)

func main() {
	pwd, _ := os.Getwd()
	config, err := configurator.GetConfig(pwd + "/config.json")
	if err != nil {
		log.Fatalln(err)
	}
	// to check video after you can use also:
	// curl -X GET -H "Authorization: Bearer $valid_token" http://crapi.bobaklabs.com:8888/identity/api/v2/user/videos/$videoId > file.out
	// get jwt token
	token := configurator.GetJWTToken(config.LoginURL, config.Email, config.Password)
	if token == "" {
		log.Fatalln("token empty")
	}

	idFlag := flag.Int("id", 1, "ID of the video that is to be edited")

	flag.Parse()

	if *idFlag <= 0 {
		log.Fatalln("wrong video id (should be over 0)")
	}

	// fmt.Println(token)
	UnauthorizedVideoParameterEdition(config.TargetURL, token, *idFlag)
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
