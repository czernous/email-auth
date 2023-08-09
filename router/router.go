package router

import (
	"email-auth/handler"
	"net/http"
	"os"
)

type Router struct {
	*http.ServeMux
}

func (r *Router) HandleRoute(method string, path string, handler http.HandlerFunc) {
	r.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != method {

			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if req.Header.Get("authApiKey") != os.Getenv("AUTH_API_KEY") {
			http.Error(w, "Wrong authApiKey", http.StatusUnauthorized)
			return
		}

		handler(w, req)

	})

}

func SetupRoutes() http.Handler {
	mux := &Router{
		ServeMux: http.NewServeMux(),
	}
	// TODO: register routes and handlers here, add respective handlers to /handler/routename.go
	//mux.HandleRoute("GET", "/hello", http.HandlerFunc(handleHello))
	mux.HandleRoute("GET", "/token", http.HandlerFunc(handler.HandleToken))
	return mux
}
