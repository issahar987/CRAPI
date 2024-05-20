package main

import (
	"fmt"
	"log"
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

	fmt.Println("Your account balance before buying product:", GetMoney(token))

	product, err := BuyProduct(token)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Your account balance after buying product:", GetMoney(token))
	ReturnProduct(*product, token)
	fmt.Println("Your account balance after returning product:", GetMoney(token))

}
