package main

import (
	"docker-proxy/src/core"
	"docker-proxy/src/middleware"
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
	r.GET("/containers", apiv1.ContainerList)
	r.GET("/images", apiv1.ImageList)
	r.POST("/start", apiv1.ContainerStart)
	r.POST("/stop", apiv1.ContainerStop)

	handle := middleware.LoggingMiddleware(r.Mux)

	fmt.Println("Server start success on " + config.Port)
	if err := http.ListenAndServe(config.Port, handle); err != nil {
		panic(err)
	}
}
