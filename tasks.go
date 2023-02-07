package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"
)

func CreateHttpServer(server *Server) {

	go func() {
		mux := http.NewServeMux()

		defaultHandler := func(w http.ResponseWriter, req *http.Request) {
			for _, resp := range server.HttpServer.Responses {
				if resp.Path == req.RequestURI {
					w.Write(resp.Body)
				}
			}
		}

		s := &http.Server{
			Addr:           ":" + strconv.Itoa(server.HttpServer.Port),
			Handler:        mux,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		for _, resp := range server.HttpServer.Responses {
			mux.HandleFunc(resp.Path, defaultHandler)
		}

		go func(s *http.Server) {
			<-server.Shutdown
			log.Println("Received shutdown signal....")
			if err := s.Shutdown(context.Background()); err != nil {
				// Error from closing listeners, or context timeout:
				log.Printf("HTTP server Shutdown: %v", err)
			}
			close(server.Shutdown)
		}(s)

		s.ListenAndServe()

	}()

}
