package main

import (
	"fmt"
	"log"

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

	fmt.Println("Your account balance before buying product:", GetMoney(config.Hostname, token))

	product, err := BuyProduct(url, token)
	if err != nil {
		log.Fatalln(err)

	}

	fmt.Println("Your account balance after buying product:", GetMoney(config.Hostname, token))
	ReturnProduct(config.Hostname, *product, token)
	fmt.Println("Your account balance after returning product:", GetMoney(config.Hostname, token))

}
