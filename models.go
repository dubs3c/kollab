package main

import (
	"github.com/google/uuid"
)

type PathResponse struct {
	Id      uuid.UUID `json:"Id,omitempty"`
	Path    string    `json:"Path"`
	Verb    string    `json:"Verb"`
	Headers []string  `json:"Headers,omitempty"`
	Body    []byte    `json:"Body"`
}

type PathRequest struct {
	Id      string   `json:"Id,omitempty"`
	Path    string   `json:"Path"`
	Verb    string   `json:"Verb"`
	Headers []string `json:"Headers,omitempty"`
	Body    []byte   `json:"Body"`
}

type HttpServer struct {
	Name      string         `json:"Name"`
	Port      int            `json:"Port"`
	Responses []PathResponse `json:"Responses"`
}

type Server struct {
	Id         uuid.UUID
	HttpServer *HttpServer
	Shutdown   chan int
}

type Api struct {
	DefaultPaths *[]*PathResponse
	Servers      *[]*Server
}
