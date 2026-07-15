package core

import (
	"os"
)

type config struct {
	DockerHost string
	Port       string
}

func GetConfig() config {
	return config{
		DockerHost: os.Getenv("DOCKER_HOST"),
		Port:       os.Getenv("PORT"),
	}
}
