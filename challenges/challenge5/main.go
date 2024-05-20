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

	url := config.TargetURL

	UploadVideo("videos/laoganma.mp4", token, url)
	UploadVideo("videos/vidio.mp4", token, url)
}
