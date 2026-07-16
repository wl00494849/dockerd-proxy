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

func ContainerList(w http.ResponseWriter, r *http.Request) {
	lst := dockerd.NewDockerdCli().Containers()
	httpx.NewWirter(w).JSON(200, lst)
}

func ContainerStart(w http.ResponseWriter, r *http.Request) {
	imageId := r.PostFormValue("image_id")
	err := dockerd.NewDockerdCli().Start(imageId)
	if err != nil {
		panic(err)
	}

	httpx.NewWirter(w).JSON(200, "")
}

func ContainerStop(w http.ResponseWriter, r *http.Request) {

}
