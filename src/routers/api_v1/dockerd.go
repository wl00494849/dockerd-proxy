package apiv1

import (
	"docker-proxy/src/dockerd"
	"docker-proxy/src/httpx"
	"fmt"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {

	resp := map[string]string{
		"status":  http.StatusText(200),
		"message": "pong",
	}

	httpx.NewWirter(w).JSON(200, resp)
	fmt.Println("pong")
}

func DockerList(w http.ResponseWriter, r *http.Request) {
	lst := dockerd.NewDockerdCli().ContainerList()
	httpx.NewWirter(w).JSON(200, lst)
	fmt.Println(lst)
}
