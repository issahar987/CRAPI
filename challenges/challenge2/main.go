package main

import (
	"log"
	"os"
)

func main() {
	// read config
	pwd, _ := os.Getwd()
	config, err := GetConfig(pwd + "/config.json")
	if err != nil {
		log.Fatalln(err)
	}

	// get jwt token
	token := GetJWTToken(config.LoginURL, config.Email, config.Password)
	if token == "" {
		log.Fatalln("token empty")
	}

	// check N report IDs
	GetAllReportsByID(50, token, config.TargetURL)
}
