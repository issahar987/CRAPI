package main

import (
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

	ViewAllOrders(token, config.TargetURL, 100)
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
