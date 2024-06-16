package main

import (
	"fmt"
	"log"
	"os"

	configurator "github.com/tomek-skrond/crapiconfigurator"
)

func main() {

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	config, err := configurator.GetConfig(pwd + "/config.json")
	if err != nil {
		log.Fatalln(err)
	}

	token := configurator.GetJWTToken(config.LoginURL, config.Email, config.Password)

	url := config.TargetURL

	ids := GetVehicleIDs(url, token)

	fmt.Println("Retrieved vehicle location data:")
	GetPayload(url, token, ids)
}
