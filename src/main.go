package main

import (
	"docker-proxy/src/core"
	"docker-proxy/src/routers"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	config := core.GetConfig()
	r := routers.New()

	r.GET("/ping", func(w http.ResponseWriter, r *http.Request) {

		resp := map[string]string{
			"OK":      "true",
			"Message": "pong",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
		fmt.Println("pong")
	})

	fmt.Println("Server start success on " + config.Port)
	if err := http.ListenAndServe(config.Port, r); err != nil {
		panic(err)
	}
}
