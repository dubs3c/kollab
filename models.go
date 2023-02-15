package main

import "github.com/google/uuid"

type HttpResponse struct {
	Path    string            `json:"Path"`
	Verb    string            `json:"Verb"`
	Headers map[string]string `json:"Headers,omitempty"`
	Body    []byte            `json:"Body"`
}

type HttpServer struct {
	Name      string         `json:"Name"`
	Port      int            `json:"Port"`
	Responses []HttpResponse `json:"Responses"`
}

type Server struct {
	Id         uuid.UUID
	HttpServer *HttpServer
	Shutdown   chan int
}

type Api struct {
	DefaultPaths *[]*HttpResponse
	Servers      *[]*Server
}
