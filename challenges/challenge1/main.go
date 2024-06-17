package main

import (
	"fmt"
	"log"

	configurator "github.com/tomek-skrond/crapiconfigurator/v2"
)

func main() {

	// pwd, err := os.Getwd()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	config, err := configurator.GetConfig("../challenge-automation/config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	loginurl := fmt.Sprintf("%s%s", config.Hostname, config.LoginURL)
	token := configurator.GetJWTToken(loginurl, config.Email, config.Password)

	url := fmt.Sprintf("%s%s", config.Hostname, config.TargetURL)

	ids := GetVehicleIDs(url, token)

	fmt.Println("Retrieved vehicle location data:")
	GetPayload(url, token, ids)
}
