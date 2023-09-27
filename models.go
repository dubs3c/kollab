package main

import (
	"database/sql"

	"github.com/google/uuid"
)

var SQL_INIT string = `
CREATE TABLE IF NOT EXISTS paths (
	id 			INTEGER PRIMARY KEY AUTOINCREMENT,
	uuid 		TEXT UNIQUE,
	url 	  	TEXT UNIQUE,
	verb	  	TEXT,
	headers   	JSON,
	body 		TEXT,
	fk_server	INTEGER,
	FOREIGN KEY(fk_server) REFERENCES servers(id)
);

CREATE TABLE IF NOT EXISTS servers (
	id		INTEGER PRIMARY KEY AUTOINCREMENT,
	uuid	TEXT UNIQUE,
	name 	TEXT UNIQUE,
	type	TEXT,
	port	INTEGER UNIQUE
);

CREATE TABLE IF NOT EXISTS paths_events (
	id  	INTEGER PRIMARY KEY AUTOINCREMENT,
	log		JSON,
	fk_path	INTEGER,
	FOREIGN KEY(fk_path) REFERENCES paths(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS server_events (
	id  		INTEGER PRIMARY KEY AUTOINCREMENT,
	log			JSON,
	fk_server	INTEGER,
	FOREIGN KEY(fk_server) REFERENCES servers(id) ON DELETE CASCADE
);

PRAGMA foreign_keys = ON;
`

type Path struct {
	Id      int       `json:"-" `
	UUID    uuid.UUID `json:"Id,omitempty"`
	Path    string    `json:"Path"`
	Verb    string    `json:"Verb"`
	Headers []string  `json:"Headers,omitempty"`
	Body    []byte    `json:"Body"`
}

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
	Name       string `json:"Name"`
	Type       string `json:"Type"`
	Port       int    `json:"Port"`
	HttpServer *HttpServer
	Shutdown   chan int
}

type Api struct {
	DB           *sql.DB
	DefaultPaths *[]*PathResponse
	Servers      *[]*Server
}

type EventPath struct {
	Path           string
	RequestHeaders map[string][]string
	IP             string
	UUID           uuid.UUID `json:"Id"`
}
