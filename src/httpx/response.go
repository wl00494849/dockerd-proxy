package httpx

import (
	"encoding/json"
	"net/http"
)

type Wirter struct {
	rw http.ResponseWriter
}

func NewWirter(w http.ResponseWriter) *Wirter {
	return &Wirter{
		rw: w,
	}
}

func (w *Wirter) JSON(status int, data any) {
	w.rw.Header().Set("Content-Type", "application/json")
	w.rw.WriteHeader(status)

	err := json.NewEncoder(w.rw).Encode(data)
	if err != nil {
		panic(err)
	}
}
