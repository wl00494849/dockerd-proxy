package httpx

import (
	"encoding/json"
	"net/http"
)

type Writer struct {
	rw http.ResponseWriter
}

func NewWriter(w http.ResponseWriter) *Writer {
	return &Writer{
		rw: w,
	}
}

func (w *Writer) JSON(status int, data any) {
	w.rw.Header().Set("Content-Type", "application/json")
	w.rw.WriteHeader(status)

	err := json.NewEncoder(w.rw).Encode(data)
	if err != nil {
		panic(err)
	}
}
