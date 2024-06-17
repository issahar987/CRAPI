package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

type Challenge struct {
	Path    string `yaml:"path"`
	Command string `yaml:"command"`
	Label   string `yaml:"label"`
	ID      int    `yaml:"id"`
}

type Config struct {
	Challenges map[string]Challenge `yaml:"challenges"`
}

func main() {
	configPathFlag := flag.String("config", "./config.yaml", "Task configuration file")

	flag.Parse()

	if _, err := os.Stat(*configPathFlag); err != nil {
		log.Fatalln("Configuration file does not exist - exiting!")
	}

	challenges := ReadChallenges(*configPathFlag)

	ExecuteChallenges(challenges)
}

func ReadChallenges(configPath string) []Challenge {
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Unmarshal the YAML into the struct
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Create a slice of challenges
	challenges := make([]Challenge, 0, len(config.Challenges))
	for _, challenge := range config.Challenges {
		challenges = append(challenges, challenge)
	}

	// Sort the challenges by ID
	sort.Slice(challenges, func(i, j int) bool {
		return challenges[i].ID < challenges[j].ID
	})

	return challenges
}

func ExecuteChallenges(challenges []Challenge) {
	for _, c := range challenges {
		fmt.Println("-----------------------------------------------------------------------")
		fmt.Println("|", c.Label, "|")
		fmt.Println("-----------------------------------------------------------------------")

		fullCommand := fmt.Sprintf("%s", c.Command)

		splitCmd := strings.Split(fullCommand, " ")
		if err := os.Chdir(c.Path); err != nil {
			log.Fatalln(err)
		}

		cmd := exec.Command(splitCmd[0], splitCmd[1:]...)

		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(output))
	}
}
