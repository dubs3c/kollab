package main

import "github.com/google/uuid"

type HttpResponse struct {
	Path    string            `json:"path"`
	Verb    string            `json:"verb"`
	Headers map[string]string `json:"headers,omitempty"`
	Body    []byte            `json:"body"`
}

type HttpServer struct {
	Name      string         `json:"name"`
	Port      int            `json:"port"`
	Responses []HttpResponse `json:"responses"`
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
