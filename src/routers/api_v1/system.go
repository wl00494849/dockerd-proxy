package apiv1

import (
	"docker-proxy/src/httpx"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	httpx.NewWirter(w).JSON(200, &ResponseMsg{
		IsSuccess: true,
		Message:   "pong",
	})
}
