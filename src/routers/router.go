package routers

import "net/http"

type Router struct {
	Mux *http.ServeMux
}

func New() *Router {
	return &Router{
		Mux: http.NewServeMux(),
	}
}

func (r *Router) GET(pattern string, handler http.HandlerFunc) {
	r.Mux.HandleFunc("GET "+pattern, handler)
}

func (r *Router) POST(pattern string, handler http.HandlerFunc) {
	r.Mux.HandleFunc("POST "+pattern, handler)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.Mux.ServeHTTP(w, req)
}
