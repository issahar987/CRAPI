package main

// import (
// 	"log"
// 	"os"
// )

// func main() {
// 	// read config
// 	pwd, _ := os.Getwd()
// 	config, err := GetConfig(pwd + "/config.json")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	// get jwt token
// 	token := GetJWTToken(config.LoginURL, config.Email, config.Password)
// 	if token == "" {
// 		log.Fatalln("token empty")
// 	}

// 	UploadVideo("vidio.mp4", token, config.TargetURL)
// }
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

	UploadVideo("videos/vidio.mp4", token, url)
	UploadVideo("videos/laoganma.mp4", token, url)
}
