package main

import (
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

	ViewAllOrders(token, url, 20)
}

func ViewAllOrders(token, target_url string, howMany int) {

	for i := 1; i <= howMany; i++ {
		x := fmt.Sprintf("%s/%d", target_url, i)
		fmt.Println(x)
		request, err := http.NewRequest("GET", fmt.Sprintf("%s/%d", target_url, i), nil)
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
		fmt.Println(string(responseBody))
	}

}
