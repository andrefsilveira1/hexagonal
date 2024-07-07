package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/andrefsilveira1/hexagonal/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type Server struct {
	Service application.ProductServiceInterface
}

func CreateNewServer() *Server {
	return &Server{}
}

func (s Server) serve() {

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log:", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
