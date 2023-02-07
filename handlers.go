package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (a *Api) AddHttpServerHandler(w http.ResponseWriter, r *http.Request) {

	/*
		need to add paths to Server struct here. To avoid deadlocks/racing,
		lock access to struct before updating. A mutex needs to be added to the struct
	*/

	h := &HttpServer{}

	err := json.NewDecoder(r.Body).Decode(&h)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}

	id := uuid.New()
	s := &Server{Id: id, HttpServer: h, Shutdown: make(chan int)}
	log.Printf("Created server '%s' with ID '%s'\n", s.HttpServer.Name, s.Id.String())

	// TODO: implement mutex
	*a.Servers = append(*a.Servers, s)

	CreateHttpServer(s)

	w.WriteHeader(201)

}

func (a *Api) AddDefaultHttpHandler(w http.ResponseWriter, r *http.Request) {
	data := &HttpResponse{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}

	// TODO: implement mutex
	*a.DefaultPaths = append(*a.DefaultPaths, data)

	w.WriteHeader(201)
}

func (a *Api) GetAllHttpServer(w http.ResponseWriter, r *http.Request) {

}

func (a *Api) GetHttpServer(w http.ResponseWriter, r *http.Request) {

}

func (a *Api) DeleteHttpServer(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "serverId")

	if id == "" {
		w.WriteHeader(404)
		return
	}

	log.Printf("Deleting HTTP server with id '%s'\n", id)

	for _, server := range *a.Servers {
		if server.Id.String() == id {
			log.Printf("Shutting down http server '%s'\n", server.HttpServer.Name)
			server.Shutdown <- 1
			w.WriteHeader(200)
			return
		}
	}
	w.WriteHeader(404)
}

func (a *Api) PutHttpServer(w http.ResponseWriter, r *http.Request) {

}

func (a *Api) AddTcpServer(w http.ResponseWriter, r *http.Request) {

}
