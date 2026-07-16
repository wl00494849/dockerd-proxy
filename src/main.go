package main

import (
	"docker-proxy/src/core"
	"docker-proxy/src/routers"
	apiv1 "docker-proxy/src/routers/api_v1"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	config := core.LoadConfig()
	r := routers.New()

	r.GET("/ping", apiv1.Ping)
	r.GET("/list", apiv1.ContainerList)
	r.POST("/start", apiv1.ContainerStart)
	r.POST("/stop", apiv1.ContainerStop)

	fmt.Println("Server start success on " + config.Port)
	if err := http.ListenAndServe(config.Port, r); err != nil {
		panic(err)
	}
}
