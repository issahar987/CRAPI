package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	config, err := GetConfig(pwd + "/config.json")
	if err != nil {
		log.Fatalln(err)
	}

	token := GetJWTToken(config.LoginURL, config.Email, config.Password)

	url := config.TargetURL

	ids := GetVehicleIDs(url, token)

	fmt.Println("Retrieved vehicle location data:")
	GetPayload(token, ids)
}
