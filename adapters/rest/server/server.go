package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/eduardo6722/go-hexagonal/adapters/rest/handlers"
	"github.com/eduardo6722/go-hexagonal/application"
	"github.com/gorilla/mux"
)

type WebServer struct {
	Service application.IProductService
}

func NewWebServer(service *application.ProductService) *WebServer {
	return &WebServer{
		Service: service,
	}
}

func (s *WebServer) Serve() {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handlers.MakeProductHandlers(r, n, s.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("error: %v", err.Error())
	}
}
