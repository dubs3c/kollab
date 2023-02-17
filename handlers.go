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

func (a *Api) AddPath(w http.ResponseWriter, r *http.Request) {

	data := &PathRequest{}
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

	pr := &PathResponse{
		Id:      uuid.New(),
		Path:    data.Path,
		Verb:    data.Verb,
		Headers: data.Headers,
		Body:    data.Body,
	}
	// TODO: implement mutex
	*a.DefaultPaths = append(*a.DefaultPaths, pr)

	RespondWithJSON(w, 201, pr)
}

func (a *Api) GetAllPaths(w http.ResponseWriter, r *http.Request) {
	data := []PathResponse{}

	for _, p := range *a.DefaultPaths {
		data = append(data, *p)
	}

	RespondWithJSON(w, 200, data)
}

func (a *Api) GetPath(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "pathId")
	uid, err := uuid.Parse(id)
	if err != nil {
		RespondWithError(w, 404, "ID not recognized")
		return
	}

	if id == "" {
		RespondWithError(w, 404, "Server ID not supplied")
		return
	}

	data := &PathResponse{}

	for _, p := range *a.DefaultPaths {
		if p.Id == uid {
			data = p
			break
		}
	}

	if (data == &PathResponse{}) {
		RespondWithError(w, 404, "Path ID not found")
	} else {
		RespondWithJSON(w, 200, data)
	}

}

func (a *Api) UpdatePath(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "pathId")

	if id == "" {
		RespondWithError(w, 404, "Server ID not supplied")
		return
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		RespondWithError(w, 404, "ID not recognized")
		return
	}

	data := &PathResponse{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		RespondWithError(w, 400, "Could not parse data")
		return
	}

	if id != data.Id.String() {
		RespondWithError(w, 400, "UUID from Path does not match UUID from object")
		return
	}

	found := false
	for _, p := range *a.DefaultPaths {
		if p.Id == uid {
			*p = *data
			found = true
			break
		}
	}

	if !found {
		RespondWithError(w, 404, "Path ID not found")
	} else {
		RespondWithJSON(w, 200, data)
	}
}

func (a *Api) DeletePath(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "pathId")

	if id == "" {
		RespondWithError(w, 404, "Server ID not supplied")
		return
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		RespondWithError(w, 404, "ID not recognized")
		return
	}

	data := &PathResponse{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		RespondWithError(w, 400, "Could not parse data")
		return
	}

	if id != data.Id.String() {
		RespondWithError(w, 400, "UUID from Path does not match UUID from object")
		return
	}

	// yea it works stfu
	found := false
	lol := []*PathResponse{}
	for _, p := range *a.DefaultPaths {
		if p.Id == uid {
			found = true
			continue
		} else {
			lol = append(lol, p)
		}
	}

	if found {
		*a.DefaultPaths = lol
	}

	if !found {
		RespondWithError(w, 404, "Path ID not found")
	} else {
		RespondWithJSON(w, 202, map[string]string{"status": "success"})
	}
}

func (a *Api) GetAllHttpServer(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, 400, map[string]string{"error": "not implemented yet"})
}

func (a *Api) GetHttpServer(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, 400, map[string]string{"error": "not implemented yet"})
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
