package apiv1

import (
	"docker-proxy/src/dockerd"
	"docker-proxy/src/httpx"
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseMsg struct {
	IsSuccess bool   `json:"is_sucess"`
	Message   string `json:"message"`
}

type ContainerRequest struct {
	ImageID  string `json:"image_id"`
	RepoTags string `json:"tags"`
}

func ContainerList(w http.ResponseWriter, r *http.Request) {
	lst := dockerd.NewDockerdCli().Containers()
	httpx.NewWirter(w).JSON(200, lst)
}

func ImageList(w http.ResponseWriter, r *http.Request) {
	lst := dockerd.NewDockerdCli().Images()
	httpx.NewWirter(w).JSON(200, lst)
}

func ContainerStart(w http.ResponseWriter, r *http.Request) {
	var param ContainerRequest
	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	fmt.Printf("Image ID:%s \n", param.ImageID)
	err := dockerd.NewDockerdCli().Start(param.ImageID)
	if err != nil {
		panic(err)
	}

	httpx.NewWirter(w).JSON(200, &ResponseMsg{
		IsSuccess: true,
		Message:   fmt.Sprintf("Success Start:%s", param.RepoTags),
	})
}

func ContainerStop(w http.ResponseWriter, r *http.Request) {
	var param ContainerRequest
	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	fmt.Printf("Image ID:%s \n", param.ImageID)
	err := dockerd.NewDockerdCli().Stop(param.ImageID)
	if err != nil {
		panic(err)
	}

	httpx.NewWirter(w).JSON(200, &ResponseMsg{
		IsSuccess: true,
		Message:   fmt.Sprintf("Success stop:%s", param.RepoTags),
	})
}
