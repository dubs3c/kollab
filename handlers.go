package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// RespondWithError - Return an error
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJSON(w, code, map[string]string{"error": msg})
}

// RespondWithJSON - Respond with a json formatted string
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *Api) AddHttpServerHandler(w http.ResponseWriter, r *http.Request) {

	h := &HttpServer{}

	err := json.NewDecoder(r.Body).Decode(&h)
	if err != nil {
		log.Println(err)
		RespondWithError(w, 400, "Could not parse data")
		return
	}

	id := uuid.New()
	s := &Server{Id: id, HttpServer: h, Shutdown: make(chan int)}
	log.Printf("Created server '%s' with ID '%s'\n", s.HttpServer.Name, s.Id.String())

	// TODO: implement mutex
	*a.Servers = append(*a.Servers, s)

	CreateHttpServer(s)

	RespondWithJSON(w, 201, map[string]string{"data": "success"})

}

func (a *Api) AddDefaultHttpHandler(w http.ResponseWriter, r *http.Request) {
	data := &HttpResponse{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		RespondWithError(w, 400, "Could not parse data")
		return
	}

	if data.Path == "" {
		RespondWithError(w, 400, "Path can not be empty...")
		return
	}

	// TODO - improve lookup
	for _, p := range *a.DefaultPaths {
		if p.Path == data.Path {
			RespondWithError(w, 400, "Path already exists...")
			return
		}
	}

	// TODO: implement mutex
	*a.DefaultPaths = append(*a.DefaultPaths, data)

	RespondWithJSON(w, 201, map[string]string{"data": "success"})
}

func (a *Api) GetAllHttpServer(w http.ResponseWriter, r *http.Request) {

}

func (a *Api) GetHttpServer(w http.ResponseWriter, r *http.Request) {

}

func (a *Api) DeleteHttpServer(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "serverId")

	if id == "" {
		RespondWithError(w, 404, "Server ID not supplied")
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
	RespondWithError(w, 404, "Could find server")
}

func (a *Api) PutHttpServer(w http.ResponseWriter, r *http.Request) {

}

func (a *Api) AddTcpServer(w http.ResponseWriter, r *http.Request) {

}
